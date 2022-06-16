package main

import (
	"fmt"
	"os"
	"trest/internal"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("USAGE: trest [cli/icli/tui]")
		return
	}

	if os.Args[1] == "cli" {
		internal.StartCli()
	}

	if os.Args[1] == "icli" {
		internal.StartIcli()
	}

	if os.Args[1] == "tui" {
		fmt.Println("not implemented yet")
	}
}
