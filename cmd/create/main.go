package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"syscall"
	"strings"

	"golang.org/x/term"
)

func main() {
	// Connect to the database
	// db, err := database.NewDatabase(os.Getenv("POSTGRES_URL"))
	// if err != nil {
	// 	log.Fatalf("error connecting to database: %q", err.Error())
	// }
	// defer db.Close()

	// Create the user
	fmt.Println(`[!] This tool will create the login credentials for a root superuser 
	only use this tool for generating your first user or if you have lost access 
	to the root account. For normal users please create them via the panel. [!]`)

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter username: ")
	username, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("error reading username: %q", err.Error())
		return
	}
	username = strings.TrimSpace(username)

	fmt.Print("Enter password: ")
	bytePassword, err := term.ReadPassword(int(syscall.Stdin))
	fmt.Println("")
	if err != nil {
		log.Fatalf("error reading password: %q", err.Error())
		return
	}
	password := string(bytePassword)
	password = strings.TrimSpace(password)

	fmt.Print("Repeat password: ")
	bytePasswordConfirm, err := term.ReadPassword(int(syscall.Stdin))
	fmt.Println("")
	if err != nil {
		log.Fatalf("error reading password confirmation: %q", err.Error())
		return
	}
	passwordConfirm := string(bytePasswordConfirm)
	passwordConfirm = strings.TrimSpace(passwordConfirm)

	if password != passwordConfirm {
		log.Fatalf("passwords do not match")
		return
	}

	// TODO: Actually create the user

}
