package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"
	"syscall"

	"github.com/Barrun-Kabir-Rifat/letseat/internal/database"
	"github.com/Barrun-Kabir-Rifat/letseat/internal/repositories"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/term"
)

func main() {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter admin username: ")
	username, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	username = strings.TrimSpace(username)

	if username == "" {
		panic("username cannot be empty")
	}

	fmt.Print("Enter admin password: ")
	password, err := term.ReadPassword(int(syscall.Stdin))
	fmt.Println()
	if err != nil {
		panic(err)
	}

	fmt.Print("Confirm admin password: ")
	confirmPassword, err := term.ReadPassword(int(syscall.Stdin))
	fmt.Println()
	if err != nil {
		panic(err)
	}

	if string(password) != string(confirmPassword) {
		panic("passwords do not match")
	}

	if len(password) < 8 {
		panic("password must contain at least 8 characters")
	}

	passwordHash, err := bcrypt.GenerateFromPassword(
		password,
		bcrypt.DefaultCost,
	)
	if err != nil {
		panic(err)
	}

	repo := repositories.NewAdminRepository(db)

	admin, err := repo.Create(
		context.Background(),
		username,
		string(passwordHash),
	)
	if err != nil {
		panic(err)
	}

	fmt.Printf(
		"Admin created successfully! ID: %d, Username: %s\n",
		admin.ID,
		admin.Username,
	)
}
