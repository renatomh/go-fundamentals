package main

import (
	"fmt"
	"intro-tests/addresses"
)

// We can run all the package tests with the following command
// go test ./... --coverprofile cover.txt
// We can read the coverage report with the commands below (shows the functions being tested)
// go tool cover --func=cover.txt
// To show as HTML page
// go tool cover --html=cover.txt

func main() {
	addressType := addresses.AddressType("P. Sherman 42, Wallaby Way")
	fmt.Println(addressType)
}
