package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

func DirExists(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return info.IsDir()
}

func ParseGitRepoURL(repoURL string) (string, string, error) {
	fmt.Printf("parseurl %s ...\n", repoURL)

	var parsedURL *url.URL
	var err error
	if strings.HasPrefix(repoURL, "git@") {
		parts := strings.SplitN(repoURL, ":", 2)
		if len(parts) != 2 {
			return "", "", fmt.Errorf("invalid git SSH URL: %s", repoURL)
		}
		host := parts[0][4:]
		path := strings.TrimSuffix(parts[1], ".git")
		return host, path, nil
	}

	parsedURL, err = url.Parse(repoURL)
	if err != nil {
		return "", "", err
	}
	path := strings.TrimSuffix(parsedURL.Path, ".git")
	return parsedURL.Host, strings.TrimPrefix(path, "/"), nil
}

func GetGitRepoFile(host, path, filename string) (string, error) {
	if host != "github.com" {
		return "", fmt.Errorf("support repo from github.com only")
	}

	// url := fmt.Sprintf("https://raw.githubusercontent.com/%s/%s/%s/%s", owner, repo, branch, filePath)
	host = "raw.githubusercontent.com"

	file_from_master := fmt.Sprintf("https://%s/%s/%s/%s", host, path, master_branch, filename)
	fmt.Printf("GET %s\n", file_from_master)
	resp, err := http.Get(file_from_master)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound { // not fond in `master`` branch, try `main` branch
		file_from_main := fmt.Sprintf("https://%s/%s/%s/%s", host, path, main_branch, filename)
		fmt.Printf("GET %s\n", file_from_main)
		resp, err = http.Get(file_from_main)
		if err != nil {
			return "", err
		}
		defer resp.Body.Close()
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func CopyFile(from, to string) error {
	fmt.Printf("copy %s %s ...\n", from, to)
	data, err := os.ReadFile(from)
	if err != nil {
		return err
	}
	return os.WriteFile(to, data, os.ModePerm)
}

func matchOptionNamesInCMakeLists(filePath string) ([]string, error) {
	pattern := `option\(\s*([a-zA-Z_][a-zA-Z0-9_]*)\s+`
	fmt.Printf("grep %s %s ...\n", pattern, filePath)
	// 编译正则表达式，用于匹配 option 命令中的 name
	re, err := regexp.Compile(pattern)
	if err != nil {
		return nil, err
	}

	// 打开文件
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var optionNames []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		matches := re.FindStringSubmatch(line)
		if len(matches) > 1 {
			optionNames = append(optionNames, matches[1])
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return optionNames, nil
}

func ListCMakeOptions(dir string) (map[string]string, error) {
	fmt.Printf("cd %s ...\n", dir)
	if err := os.Chdir(dir); err != nil {
		return nil, err
	}

	opts, err := matchOptionNamesInCMakeLists(cmakefile)
	if err != nil {
		return nil, err
	}

	allopts := map[string]string{}
	for _, opt := range opts {
		allopts[opt] = ""
	}

	{ // check if cmake exist
		cmd := exec.Command("cmake")
		if err := cmd.Run(); err != nil {
			fmt.Printf("[ERROR] %s\n", err)
			return nil, err
		}
	}
	fmt.Printf("exec `cmake -L` ...\n")
	cmd := exec.Command("cmake", "-L")
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Run() // `cmake -L` may error for many reasons, ignore it. we just need options

	// ROARING_BUILD_LTO:BOOL=OFF
	re, err := regexp.Compile(`([a-zA-Z_][a-zA-Z0-9_]*):([a-zA-Z_][a-zA-Z0-9_]*)=([a-zA-Z_][a-zA-Z0-9_]*)`)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(bytes.NewReader(out.Bytes()))
	for scanner.Scan() {
		line := scanner.Text()
		matches := re.FindStringSubmatch(line)
		if len(matches) == 4 {
			name := matches[1]
			value := matches[3]
			if _, ok := allopts[name]; ok && matches[2] == "BOOL" {
				allopts[name] = value
			}
		}
	}

	return allopts, nil
}

const (
	master_branch = "master"
	main_branch   = "main"
	cmakefile     = "CMakeLists.txt"
)

// cmake_add_deps git-repoo
// cmake_add_deps dir
func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Printf("usage: cmake_add_deps [git-repo] or cmake_add_deps [dir]\n")
		os.Exit(1)
	}

	dir := args[1]
	tmpdir, err := os.MkdirTemp("", "cmake_add_deps-*")
	if err != nil {
		panic(err)
	}
	defer os.RemoveAll(tmpdir)

	tmpfile := tmpdir + "/" + cmakefile
	isdir := DirExists(dir)
	if DirExists(dir) {
		if err = CopyFile(dir+"/"+cmakefile, tmpfile); err != nil {
			fmt.Printf("[ERROR] copy %s %s\n", dir+"/"+cmakefile, tmpfile)
			os.Exit(1)
		}
	} else {
		host, path, err := ParseGitRepoURL(dir)
		if err != nil {
			fmt.Printf("[ERROR] invalid git repo: %s\n", dir)
			os.Exit(1)
		}

		data, err := GetGitRepoFile(host, path, cmakefile)
		if err != nil {
			fmt.Printf("[ERROR] %s\n", err)
			os.Exit(1)
		}

		fmt.Printf("write %s to %s ...\n", cmakefile, tmpfile)

		if err = os.WriteFile(tmpfile, []byte(data), os.ModePerm); err != nil {
			fmt.Printf("[ERROR] write %s %s\n", tmpfile, err)
			os.Exit(1)
		}
	}

	opt_default_values, err := ListCMakeOptions(tmpdir)
	if err != nil {
		fmt.Printf("[ERROR] %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("analyse success for cmake module %s\nimport and config it as follwing:\n\n", dir)

	for key, value := range opt_default_values {
		fmt.Printf("#set(%s %s)\n", key, value)
	}
	if isdir {
		fmt.Printf("#add_subdirectory(%s)\n\n", dir)
	} else {
		fmt.Printf("#add_subdirectory(%s)\n\n", "path/to/repo")
		fmt.Printf("#you may get the repo using:\n git clone %s %s\n", dir, "path/to/repo")
		fmt.Printf("#or add as submodule using:\n git submodule add %s %s\n", dir, "path/to/repo")
	}
}
