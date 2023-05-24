package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/dobleme/ase-interpreter/internal/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hi %s! Welcome to Ase lang!\n", user.Username)
	repl.Start(os.Stdin, os.Stdout)
}
