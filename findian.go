package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func regards() {
	fmt.Println("**********************************")
	fmt.Println("Hi enter a the characters i, a, n ")
	fmt.Println("**********************************")
}

func goodbye() {
	fmt.Println("**********************************")
	fmt.Println("Thank you for used this program")
	fmt.Println("**********************************")
}

func main() {

	regards()

	fmt.Println("Please enter your character")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	check := scanner.Text()
	if strings.HasPrefix(check, "i") && strings.Contains(check, "a") && strings.HasSuffix(check, "n") {

		fmt.Println("Found")

	} else {
		fmt.Println("No found")
	}
	goodbye()
}
