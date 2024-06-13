package main

import (
    "bufio"
    "fmt"
    "os"
    "os/exec"
    "path/filepath"
    "strings"
)

type Shell struct {
    Commands map[string] func(* Shell, string)
    Reader * bufio.Reader
    Cwd string
}

func NewShell() * Shell {
    cwd, _ := os.Getwd()
    return &Shell {
        Commands: make(map[string] func(* Shell, string)),
        Reader: bufio.NewReader(os.Stdin),
        Cwd: cwd,
    }
}

func main() {
    myShell := NewShell()
    myShell.Commands["echo"] = echo
    myShell.Commands["exit"] = exit
    myShell.Commands["type"] = typeCommand
    myShell.Commands["pwd"] = printWorkingDir
    myShell.Commands["cd"] = changeDirectory

    for {
        fmt.Fprint(os.Stdout, "$ ")
        cmd, err := myShell.Reader.ReadString('\n')
        if err != nil {
            fmt.Fprintf(os.Stderr, "Error reading command: %v\n", err)
            continue
        }
        cmd = strings.TrimSpace(cmd)

        args := strings.SplitN(cmd, " ", 2)
        command := args[0]
        arg := ""
        if len(args) > 1 {
            arg = args[1]
        }

        if cmdFunc, ok := myShell.Commands[command];
        ok {
            cmdFunc(myShell, arg)
        } else {
            executeExternalCommand(command, arg)
        }
    }
}

func echo(shell *Shell, statement string) {
    fmt.Fprintln(os.Stdout, statement)
}

func exit(shell *Shell, command string) {
    os.Exit(0)
}

func typeCommand(shell *Shell, command string) {
    if _, ok := shell.Commands[command];
    ok {
        fmt.Printf("%s is a shell builtin\n", command)
        return
    }

    if path := findExecutable(command);
    path != "" {
        fmt.Printf("%s is %s\n", command, path)
    } else {
        fmt.Printf("%s: not found\n", command)
    }
}

func findExecutable(command string) string {
    pathDirs := strings.Split(os.Getenv("PATH"), ":")
    for _,
    dir := range pathDirs {
        execPath := filepath.Join(dir, command)
        if _,
        err := os.Stat(execPath);err == nil {
            return execPath
        }
    }
    return ""
}

func executeExternalCommand(command string, arg string) {
	cmd := exec.Command(command, strings.Fields(arg)...)
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    cmd.Stdin = os.Stdin

    err := cmd.Run()
    if err != nil {
        fmt.Printf("%s: command not found\n", command)
    }
}

func printWorkingDir(shell *Shell, command string) {
    fmt.Println(shell.Cwd)
}

func changeDirectory(shell *Shell, newPath string) {
    if newPath == "" {
        return
    }
    if newPath == "~" {
        newPath = os.Getenv("HOME")
    }
    err := os.Chdir(newPath)
    if err != nil {
        fmt.Printf("cd: %s: No such file or directory\n", newPath)
    }
    path, _ := os.Getwd()
    shell.Cwd = path
}