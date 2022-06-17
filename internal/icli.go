package internal

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"golang.org/x/term"
)

const prompt = ">>> "

func printLine(buffer []byte) {
	fmt.Printf("\r\033[K\r")
	fmt.Printf("%s%s", prompt, buffer)
}

type icommand struct {
	name string
	usage string
	exec func([]string) string
}

func getIcommands() []icommand {
	return []icommand {
		{
			name: "get",
			usage: "get [url]",
			exec: func(args []string) string {
				if len(args) < 1 {
					return "bad cmd"
				}
				return Get(args)
			},
		},
		{
			name: "post",
			usage: "post [url] [body]",
			exec: func(args []string) string {
				if len(args) < 2 {
					return "bad cmd"
				}
				return Post(args)
			},
		},
		{
			name: "post",
			usage: "delete [url]",
			exec: func(args []string) string {
				if len(args) < 1 {
					return "bad cmd"
				}
				return Delete(args)
			},
		},
	}
}

func StartIcli() {
	fmt.Println("trest - test your http endpoints")

	commands := getIcommands()

	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		panic(err)
	}
	defer term.Restore(int(os.Stdin.Fd()), oldState)

	reader := bufio.NewReader(os.Stdin)
	var buffer []byte
	var history [][]byte
	history_idx := -1
	for {
		printLine(buffer)
		c, _ := reader.ReadByte()

		// check for arrows
		if c == 27 {
			c, _ := reader.ReadByte()
			if c == 91 {
				c, _ := reader.ReadByte()

				// up arrow
				if c == 65 {
					if history_idx == -1 {
						history_idx = len(history) - 1

						if history_idx >= 0 {
							buffer = history[history_idx]
						}
					} else {
						if history_idx > 1 {
							history_idx -= 1

							buffer = history[history_idx]
						}
					}

				// down arrow
				} else if c == 66 {
					if history_idx != -1 && history_idx < len(history) {
						history_idx += 1
						if history_idx == len(history) {
							buffer = buffer[:0]
						} else {
							buffer = history[history_idx]
						}
					}
				}
			}
			continue
		}

		// check for enter/return
		if c == '\r' || c == '\n' {
			fmt.Println()

			start := 0
			for _, b := range buffer {
				if b != ' ' {
					break
				}
				start += 1
			}
			cmd := string(buffer[start:])

			history = append(history, append([]byte{}, buffer...))
			history_idx = -1

			if cmd == "exit" || cmd == "quit" {
				break
			}

			split := strings.Split(cmd, " ")
			for _, cmd := range commands {
				if split[0] == cmd.name {
					res := cmd.exec(split[1:])
					for _, line := range strings.Split(res, "\n") {
						if len(line) > 0 {
							fmt.Printf("\r\033[K\r%s\n", line)
						}
					}
					break
				}
			}

			buffer = buffer[:0]

			// check for backspace
		} else if c == 127 && len(buffer) > 0 {
			buffer = buffer[0 : len(buffer) - 1]

		} else {
			buffer = append(buffer, c)
		}

		// break on ctrl-c
		if c == 3 {
			break
		}
	}
}
