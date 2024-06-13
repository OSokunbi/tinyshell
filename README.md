## tinyshell

tinyshell is a simple shell that can run commands and execute scripts. it is written in Go in order to learn more about the language and how shells work.

its pretty lightweight, its written in aroudn 125 lines of code. it was pretty fun to make and i learned a lot about how shells work. 

### structure
its architecture is pretty simple, i have a couple of types for the shell, including
- `Cwd` - the current working directory
- `Reader` - a reader that reads input from the user
- `Commands` - a map of commands that the shell can run

Commands is a map of strings to functions, where the string is the name of the command and the function is the function that the command runs.
- each function takes in a string as an argument as well as a pointer to the current shell session

i have a type called shell that holds the state of the shell, and a function. the shell has a prompt that is displayed to the user, and the user can enter commands. the shell reads the input, parses it, and then executes it. the shell can run commands and execute scripts.

### builtins
the shell has a couple of builtins, including
- `cd` - change directory
- `exit` - exit the shell
- `pwd` - print the current working directory
- `echo` - print a string to the console

that's about it, feel free to use this or modify and make additions to it as you see fit. i hope you enjoy it!

# ʕ˙Ⱉ˙ʔ rawr!