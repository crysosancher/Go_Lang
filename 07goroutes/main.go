package main

import (
	"fmt"
	"time"
	// "time"
)

func sayHello() {
    for i := 0; i < 5; i++ {
        fmt.Println("Hello")
         time.Sleep(time.Millisecond * 500)
    }
}

func main() {
    go sayHello() // Start a new goroutine
    for i := 0; i < 5; i++ {
        fmt.Println("World")
         time.Sleep(time.Millisecond * 500)
    }
    // Without waiting, the program may exit before the goroutine finishes
    // Use time.Sleep to wait for the goroutine to complete or use sync.WaitGroup.
}
