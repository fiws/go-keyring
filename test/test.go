package main

import (
	"fmt"

	"github.com/zalando/go-keyring"
)

func main() {

	serviceName := "go-keyring-test"

	longString := "(\"\n💣\n\t{test}^💣💣@#';DROP TABLE"
	for a := 0; a <= 2518; a++ {
		longString += "a"
	}
	longString += "z"

	length := len(longString)

	fmt.Println("Trying to add password with:")
	fmt.Println("Password length:", length)
	err := keyring.Set(serviceName, "demo", longString)

	fmt.Println("err", err)

	// read value again
	value, err := keyring.Get(serviceName, "demo")
	if err != nil {
		panic(err)
	}

	if value == longString {
		fmt.Println("✅ Set password looks good")
	} else {
		fmt.Println("OH NO! Retrieved password is not the same as stored:")
		fmt.Println("Value stored:", value)
	}
}
