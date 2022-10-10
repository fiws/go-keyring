package main

import (
	"fmt"

	"github.com/zalando/go-keyring"
)

func main() {

	test(1024 * 1024)
}

func test(length int) {
	serviceName := "go-keyring-test"

	longString := "(\"\n💣\n\t{test}^💣💣@#';DROP TABLE"
	for a := 0; a <= length; a++ {
		longString += "a"
	}
	longString += "z"

	err := keyring.Set(serviceName, "demo", longString)

	fmt.Println("err", err)

	// read value again
	value, err := keyring.Get(serviceName, "demo")
	if err != nil {
		panic(err)
	}

	if value == longString {
		fmt.Printf("✅ Test with %d characters passed", length)
	} else {
		fmt.Printf("❌ Test with %d characters failed, saved value was different", length)
	}

	keyring.Delete(serviceName, "demo")
}
