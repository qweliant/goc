package main

import (
	"fmt"
	"os"
	"os/user"
	"goc/repl"
)

func main() {
	user, err := user.Current() // user.Current() returns a pointer to a user struct
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is the go c programming language or goc for short! Try some goc out\n", user.Username)
	fmt.Printf("Feel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
