package main

import (
	"errors"
	"fmt"
)

func Signup(name, password string) error {

	if name == "" {
		return errors.New("name is required")
	}

	if password == "" {
		return errors.New("password is required")
	}

	return nil
}

func main() {

	err := Signup("Sabari", "")

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Signup Successful")
}
func Login(name, password string) error {

	if name == "" || password == "" {
		return errors.New("name and password are required")
	}

	if name == "admin" && password == "admin" {
		return nil
	}

	return errors.New("invalid username or password")
}