package main

import (
    "fmt"
    "bufio"
    "os"
    "strings"
    "trest/internal"
)

type command struct {
    name string
    exec func([]string)
}

func getCommands() []command {
    commands := []command {
        {
            name: "post",
            exec: func (args []string) {
                fmt.Printf("POST: [%s] [type: %s] [body: %s]\n", args[0], args[1], args[2])
                internal.Post(args)
            },
        },
        {
            name: "get",
            exec: func (args []string) {
                fmt.Printf("GET : [%s]\n", args[0])
                res := internal.Get(args)
                fmt.Println(res)
            },
        },
        {
            name: "exit",
            exec:  func (_ []string) { os.Exit(0) },
        },
    }

    return commands
}

func main()  {
    commands := getCommands()

    prompt := func() { fmt.Print(">>> ") }
    scanner := bufio.NewScanner(os.Stdin)
    prompt()
    for scanner.Scan() {
        s := strings.Split(scanner.Text(), " ")

        for _, cmd := range commands {
            if s[0] == cmd.name {
                cmd.exec(s[1:])
                break
            }
        }

        prompt()
    }

    if err := scanner.Err(); err != nil {
        fmt.Fprintln(os.Stderr, err)
        os.Exit(1)
    }
}