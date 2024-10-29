package main

import (
	"fmt"
	"monkey/repl"
	"os"
	"os/user"
)

func main() {
	current, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s!\n", current.Username)
	fmt.Println("Welcome to the Monkey programming language!")
	repl.Start(os.Stdin, os.Stdout)

}
