package main

import (
	"bufio"
	"fmt"
	"github.com/awesomebfm/minepanel/pkg/auth"
	"github.com/awesomebfm/minepanel/pkg/database"
	"log"
	"os"
	"strings"
	"syscall"
	"time"

	"golang.org/x/term"
)

func main() {
	// Connect to the database
	db, err := database.NewDatabase(os.Getenv("POSTGRES_URL"))
	if err != nil {
		log.Fatalf("error connecting to database: %q", err.Error())
	}
	defer db.Close()

	// Create an auth object
	ath := auth.NewAuth(
		&auth.Params{
			Memory:      64 * 1024,
			Iterations:  3,
			Parallelism: 2,
			SaltLength:  16,
			KeyLength:   32,
		},
		90*24*time.Hour,
	)

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

	// Create user
	var user auth.User
	user.Username = username
	user.CreatedAt = time.Now()
	user.LastLogin = time.Now()

	// Hash password
	hashedPassword, err := ath.HashPassword(password)
	if err != nil {
		log.Fatalf("error hashing password: %q", err.Error())
		return
	}
	user.HashedPassword = hashedPassword

	// Store user
	err = db.PersistUser(&user)
	if err != nil {
		log.Fatalf("error storing user: %q", err.Error())
		return
	}

	fmt.Println("SUCCESS! A user now exists with these credentials")
}
