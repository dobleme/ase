package main

import (
	"fmt"
	"os"
	"os/user"

	"ase/internal/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hi %s! Welcome to Ase lang!\n", user.Username)
	repl.Start(os.Stdin, os.Stdout)
}
