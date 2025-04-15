// DO NOT EDIT!!! GENERETED BY ragel AS:
//  ragel -Z -G2 -e -o ragel_urlparse.go urlparse.go.rl
// it might be helpful to generate the FSM graph in debug:
//   ragel -Vp urlparse.go.rl -o urlparse.dot
//   dot urlparse.dot -Tpng -o urlparse.png

package parser

import (
    "fmt"
    "errors"
)

type Command struct {
    Name string
    Args []string
}

func (cmd Command) String() string {
    return fmt.Sprintf("%s(%s)",cmd.Name, cmd.Args)
}

func ParseCMakelists(data string) ([]Command, error) {
    cs, p, pe := 0, 0, len(data)
    eof := pe
    pb := p
    var err error
    commands := make([]Command, 0, 10)
    tmp_args := make([]string, 0, 10)

%%{
machine parserequesturi;

include common "cmakelists.rl";

main := cmakelists;

write data;
write init;
write exec;
}%%

    if p != eof {  // input date not fully consumed
        err = errors.New("input data not fully consumed");
        return commands, err
    }

    if cs < parserequesturi_first_final {
        err = fmt.Errorf("urlquery parse error: %s %d", data, cs)
    }
    return commands, err
}
