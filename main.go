package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	
	"github.com/unknown3472/golang_test.git/golang_test/git"
)

func main() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	filePath := filepath.Join(dir, "userinfo.txt")
	file, err := os.Create(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	username, err := git.GetUserName()
	if err != nil {
		log.Fatal(err)
	}
	_, err = file.WriteString(fmt.Sprintf("Username: %s\n", username))
	if err != nil {
		log.Fatal(err)
	}
	email, err := git.GetUserEmail()
	if err != nil {
		log.Fatal(err)
	}
	_, err = file.WriteString(fmt.Sprintf("Email: %s\n", email))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("User information has been written to userinfo.txt successfully.")
}
