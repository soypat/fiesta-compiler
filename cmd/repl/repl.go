package main

import (
	"fiesta-compiler/repl"
	"fmt"
	"log"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Hello %s! This is the Monkey programming language!\n\tDWTFYW (C) 2023\n", user.Username)
	repl.Start(os.Stdin, os.Stdout)
}
