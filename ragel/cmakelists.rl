%%{
machine common;

action mark_pb { pb = p }

action append_args {
    tmp_args = append(tmp_args, data[pb:p])
}

action add_command {
    fmt.Printf("%#v\n", tmp_args)
    commands = append(commands, Command{tmp_args[0], tmp_args[1:]})
    tmp_args = tmp_args[:0]
}

ws = space;
newline = '\n' | '\r';

string = ('"' ) (any - '"')* ('"') ;
comment = ' '* . '#' . (any - newline)* ;
commentline = comment . newline;

arg = ( (any - ws - '#' - '"' - '(' - ')')+ | string) >mark_pb %append_args ;
args = arg commentline* (ws+ arg commentline* )* ;

command = (arg ws* '(' ws* args? ws* ')') %add_command;
commands = (command | commentline) (ws* (command | commentline) )*;

cmakelists = ws* commands? ws* comment?;

}%%;