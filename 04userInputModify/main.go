package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("Welcome to my program, what is your age? ")
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Welcome to ,whats your age hottie?")
	input, _ := reader.ReadString('\n')
	age, err := strconv.ParseInt(strings.TrimSpace(input), 0, 64)
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Printf("Success: %d\n", age)
		fmt.Printf("You will be %d years old at the end of 2023\n", age+1)
	}
}
