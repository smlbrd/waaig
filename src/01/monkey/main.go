package main

import (
	"fmt"
	"go-interpreter/src/01/monkey/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Ook! Hello %s! This is the Monkey programming language!\n",
		user.Username)

	fmt.Printf("Feel free to type in some commands\n")

	repl.Start(os.Stdin, os.Stdout)
}
