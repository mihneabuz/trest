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
			name: "get",
			exec: func(args []string) {
				fmt.Printf("GET : [%s]\n", args[0])
				res := Get(args)
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
		s := strings.Split(scanner.Text(), " ")

		for _, cmd := range commands {
			if s[0] == cmd.name {
				cmd.exec(s[1:])
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
