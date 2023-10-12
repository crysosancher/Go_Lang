package main

import (
	"fmt"
)

func main() {
	// fmt.Println("Hello World")
	vibhu := Hand{"Large","Brown",true}
	fmt.Println(vibhu)
	fmt.Printf("vibhu's Hands are %v and %v\n",vibhu.Size,vibhu.color)
}
type Hand struct{
	Size string
	color string
	Available bool
}