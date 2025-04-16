# cmaketools

## cmake_add_deps

analyse CMakeLists.txt options for target module and tell you how to config it.

target module can be github repo or localpath.

### install

`go install github.com/longqimin/cmaketools/cmake_add_deps@latest`

### usage

```
cmake_add_deps
usage: cmake_add_deps [git-repo] or cmake_add_deps [dir]
```

#### add dependency module from github repo

`cmake_add_deps https://github.com/RoaringBitmap/CRoaring.git`

output:

```
analyse success for cmake module https://github.com/RoaringBitmap/CRoaring.git
import and config it as follwing:

#set(ENABLE_ROARING_TESTS ON)
#set(ROARING_BUILD_C_TESTS_AS_CPP OFF)
#set(ROARING_SANITIZE_UNDEFINED OFF)
#set(ROARING_USE_CPM ON)
#set(ROARING_DISABLE_NEON OFF)
#set(ROARING_EXCEPTIONS ON)
#set(ROARING_DISABLE_X64 OFF)
#set(ROARING_DISABLE_AVX OFF)
#set(ROARING_DISABLE_AVX512 OFF)
#set(ROARING_BUILD_STATIC ON)
#set(ROARING_BUILD_C_AS_CPP OFF)
#set(ROARING_LINK_STATIC OFF)
#set(ROARING_BUILD_LTO OFF)
#set(ROARING_SANITIZE OFF)
#set(ROARING_SANITIZE_THREADS OFF)
#set(ENABLE_ROARING_MICROBENCHMARKS OFF)
#add_subdirectory(path/to/repo)

#you may get the repo using:
 git clone https://github.com/RoaringBitmap/CRoaring.git path/to/repo
#or add as submodule using:
 git submodule add https://github.com/RoaringBitmap/CRoaring.git path/to/repo
```

#### add dependency module from path

`cmake_add_deps thirdparty/CRoaring/`

output:

```
analyse success for cmake module thirdparty/CRoaring/
import and config it as follwing:

#set(ROARING_DISABLE_AVX OFF)
#set(ROARING_DISABLE_AVX512 OFF)
#set(ROARING_SANITIZE_THREADS OFF)
#set(ROARING_USE_CPM ON)
#set(ROARING_BUILD_STATIC ON)
#set(ROARING_BUILD_LTO OFF)
#set(ROARING_BUILD_C_AS_CPP OFF)
#set(ROARING_SANITIZE OFF)
#set(ENABLE_ROARING_TESTS ON)
#set(ROARING_EXCEPTIONS ON)
#set(ROARING_DISABLE_X64 OFF)
#set(ROARING_DISABLE_NEON OFF)
#set(ROARING_SANITIZE_UNDEFINED OFF)
#set(ROARING_LINK_STATIC OFF)
#set(ROARING_BUILD_C_TESTS_AS_CPP OFF)
#set(ENABLE_ROARING_MICROBENCHMARKS OFF)
#add_subdirectory(thirdparty/CRoaring/)
```
