package main

import (
	"bufio"
	"fmt"
	"os"
)
 func main()  {
	welcome :="Welcome to our hole"
	   fmt.Println( welcome)
		 reader:= bufio.NewReader(os.Stdin)
		 fmt.Println("Enter the rating for our pizza: ")
		 //comma ok || err ok Syntax
		 input, _ := reader.ReadString('\n')
		 fmt.Println("Thanks for rating, ", input)
 }