package internal

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type command struct {
	name string
	exec func([]string)
}

func getCommands() []command {
	commands := []command{
		{
			name: "get",
			exec: func(args []string) {
				fmt.Printf("GET : [%s]\n", args[0])
				res := Get(args)
				fmt.Println(res)
			},
		},
		{
			name: "post",
			exec: func(args []string) {
				if len(args) < 3 {
					args = append(args, "")
				}
				fmt.Printf("POST: [%s] [body: %s]\n", args[0], args[1])
				res := Post(args)
				fmt.Println(res)
			},
		},
		{
			name: "delete",
			exec: func(args []string) {
				if len(args) < 3 {
					args = append(args, "")
				}
				fmt.Printf("DEL : [%s] [body: %s]\n", args[0], args[1])
				res := Delete(args)
				fmt.Println(res)
			},
		},
		{
			name: "exit",
			exec: func(_ []string) { os.Exit(0) },
		},
	}

	return commands
}

func StartCli() {
	commands := getCommands()

	prompt := func() { fmt.Print(">>> ") }
	scanner := bufio.NewScanner(os.Stdin)
	prompt()
	for scanner.Scan() {
		s := scanner.Text()
		for len(s) > 0 && s[0] == ' ' {
			s = s[1:]
		}

		split := strings.Split(s, " ")

		for _, cmd := range commands {
			if split[0] == cmd.name {
				cmd.exec(split[1:])
				fmt.Println()
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
