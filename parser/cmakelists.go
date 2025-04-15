
//line ragel/cmakelists.go.rl:1
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


//line parser/cmakelists.go:35
const parserequesturi_start int = 12
const parserequesturi_first_final int = 12
const parserequesturi_error int = 0

const parserequesturi_en_main int = 12


//line parser/cmakelists.go:43
	{
	cs = parserequesturi_start
	}

//line parser/cmakelists.go:48
	{
	if p == pe {
		goto _test_eof
	}
	switch cs {
	case 12:
		goto st_case_12
	case 1:
		goto st_case_1
	case 2:
		goto st_case_2
	case 0:
		goto st_case_0
	case 3:
		goto st_case_3
	case 4:
		goto st_case_4
	case 5:
		goto st_case_5
	case 6:
		goto st_case_6
	case 7:
		goto st_case_7
	case 8:
		goto st_case_8
	case 9:
		goto st_case_9
	case 13:
		goto st_case_13
	case 10:
		goto st_case_10
	case 11:
		goto st_case_11
	case 14:
		goto st_case_14
	}
	goto st_out
tr25:
//line ragel/cmakelists.rl:10

    fmt.Printf("%#v\n", tmp_args)
    commands = append(commands, Command{tmp_args[0], tmp_args[1:]})
    tmp_args = tmp_args[:0]

	goto st12
	st12:
		if p++; p == pe {
			goto _test_eof12
		}
	st_case_12:
//line parser/cmakelists.go:99
		switch data[p] {
		case 32:
			goto st12
		case 34:
			goto tr22
		case 35:
			goto st14
		}
		switch {
		case data[p] > 13:
			if 40 <= data[p] && data[p] <= 41 {
				goto st0
			}
		case data[p] >= 9:
			goto st12
		}
		goto tr20
tr20:
//line ragel/cmakelists.rl:4
 pb = p 
	goto st1
tr24:
//line ragel/cmakelists.rl:10

    fmt.Printf("%#v\n", tmp_args)
    commands = append(commands, Command{tmp_args[0], tmp_args[1:]})
    tmp_args = tmp_args[:0]

//line ragel/cmakelists.rl:4
 pb = p 
	goto st1
	st1:
		if p++; p == pe {
			goto _test_eof1
		}
	st_case_1:
//line parser/cmakelists.go:136
		switch data[p] {
		case 32:
			goto tr1
		case 40:
			goto tr3
		case 41:
			goto st0
		}
		switch {
		case data[p] > 13:
			if 34 <= data[p] && data[p] <= 35 {
				goto st0
			}
		case data[p] >= 9:
			goto tr1
		}
		goto st1
tr1:
//line ragel/cmakelists.rl:6

    tmp_args = append(tmp_args, data[pb:p])

	goto st2
	st2:
		if p++; p == pe {
			goto _test_eof2
		}
	st_case_2:
//line parser/cmakelists.go:165
		switch data[p] {
		case 32:
			goto st2
		case 40:
			goto st3
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st2
		}
		goto st0
st_case_0:
	st0:
		cs = 0
		goto _out
tr3:
//line ragel/cmakelists.rl:6

    tmp_args = append(tmp_args, data[pb:p])

	goto st3
	st3:
		if p++; p == pe {
			goto _test_eof3
		}
	st_case_3:
//line parser/cmakelists.go:191
		switch data[p] {
		case 32:
			goto st3
		case 34:
			goto tr7
		case 35:
			goto st0
		case 40:
			goto st0
		case 41:
			goto st13
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st3
		}
		goto tr6
tr6:
//line ragel/cmakelists.rl:4
 pb = p 
	goto st4
	st4:
		if p++; p == pe {
			goto _test_eof4
		}
	st_case_4:
//line parser/cmakelists.go:217
		switch data[p] {
		case 32:
			goto tr10
		case 34:
			goto st0
		case 35:
			goto tr11
		case 40:
			goto st0
		case 41:
			goto tr12
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto tr3
		}
		goto st4
tr10:
//line ragel/cmakelists.rl:6

    tmp_args = append(tmp_args, data[pb:p])

	goto st5
	st5:
		if p++; p == pe {
			goto _test_eof5
		}
	st_case_5:
//line parser/cmakelists.go:245
		switch data[p] {
		case 32:
			goto st5
		case 34:
			goto tr7
		case 35:
			goto st8
		case 40:
			goto st0
		case 41:
			goto st13
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st3
		}
		goto tr6
tr7:
//line ragel/cmakelists.rl:4
 pb = p 
	goto st6
	st6:
		if p++; p == pe {
			goto _test_eof6
		}
	st_case_6:
//line parser/cmakelists.go:271
		if data[p] == 34 {
			goto st7
		}
		goto st6
	st7:
		if p++; p == pe {
			goto _test_eof7
		}
	st_case_7:
		switch data[p] {
		case 32:
			goto tr10
		case 35:
			goto tr11
		case 41:
			goto tr12
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto tr3
		}
		goto st0
tr11:
//line ragel/cmakelists.rl:6

    tmp_args = append(tmp_args, data[pb:p])

	goto st8
	st8:
		if p++; p == pe {
			goto _test_eof8
		}
	st_case_8:
//line parser/cmakelists.go:304
		switch data[p] {
		case 10:
			goto st9
		case 13:
			goto st9
		}
		goto st8
	st9:
		if p++; p == pe {
			goto _test_eof9
		}
	st_case_9:
		switch data[p] {
		case 32:
			goto st5
		case 35:
			goto st8
		case 41:
			goto st13
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st3
		}
		goto st0
tr12:
//line ragel/cmakelists.rl:6

    tmp_args = append(tmp_args, data[pb:p])

	goto st13
	st13:
		if p++; p == pe {
			goto _test_eof13
		}
	st_case_13:
//line parser/cmakelists.go:340
		switch data[p] {
		case 32:
			goto tr25
		case 34:
			goto tr26
		case 35:
			goto tr27
		}
		switch {
		case data[p] > 13:
			if 40 <= data[p] && data[p] <= 41 {
				goto st0
			}
		case data[p] >= 9:
			goto tr25
		}
		goto tr24
tr22:
//line ragel/cmakelists.rl:4
 pb = p 
	goto st10
tr26:
//line ragel/cmakelists.rl:10

    fmt.Printf("%#v\n", tmp_args)
    commands = append(commands, Command{tmp_args[0], tmp_args[1:]})
    tmp_args = tmp_args[:0]

//line ragel/cmakelists.rl:4
 pb = p 
	goto st10
	st10:
		if p++; p == pe {
			goto _test_eof10
		}
	st_case_10:
//line parser/cmakelists.go:377
		if data[p] == 34 {
			goto st11
		}
		goto st10
	st11:
		if p++; p == pe {
			goto _test_eof11
		}
	st_case_11:
		switch data[p] {
		case 32:
			goto tr1
		case 40:
			goto tr3
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto tr1
		}
		goto st0
tr27:
//line ragel/cmakelists.rl:10

    fmt.Printf("%#v\n", tmp_args)
    commands = append(commands, Command{tmp_args[0], tmp_args[1:]})
    tmp_args = tmp_args[:0]

	goto st14
	st14:
		if p++; p == pe {
			goto _test_eof14
		}
	st_case_14:
//line parser/cmakelists.go:410
		switch data[p] {
		case 10:
			goto st12
		case 13:
			goto st12
		}
		goto st14
	st_out:
	_test_eof12: cs = 12; goto _test_eof
	_test_eof1: cs = 1; goto _test_eof
	_test_eof2: cs = 2; goto _test_eof
	_test_eof3: cs = 3; goto _test_eof
	_test_eof4: cs = 4; goto _test_eof
	_test_eof5: cs = 5; goto _test_eof
	_test_eof6: cs = 6; goto _test_eof
	_test_eof7: cs = 7; goto _test_eof
	_test_eof8: cs = 8; goto _test_eof
	_test_eof9: cs = 9; goto _test_eof
	_test_eof13: cs = 13; goto _test_eof
	_test_eof10: cs = 10; goto _test_eof
	_test_eof11: cs = 11; goto _test_eof
	_test_eof14: cs = 14; goto _test_eof

	_test_eof: {}
	if p == eof {
		switch cs {
		case 13:
//line ragel/cmakelists.rl:10

    fmt.Printf("%#v\n", tmp_args)
    commands = append(commands, Command{tmp_args[0], tmp_args[1:]})
    tmp_args = tmp_args[:0]

//line parser/cmakelists.go:444
		}
	}

	_out: {}
	}

//line ragel/cmakelists.go.rl:41


    if p != eof {  // input date not fully consumed
        err = errors.New("input data not fully consumed");
        return commands, err
    }

    if cs < parserequesturi_first_final {
        err = fmt.Errorf("urlquery parse error: %s %d", data, cs)
    }
    return commands, err
}
