package main

import (
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func main() {

	fmt.Println("# Golang Snippets")
	fmt.Println("## Hashed Password")

	password := "Pass-6942!"

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(hashedPassword))
}
