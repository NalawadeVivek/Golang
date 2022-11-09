package main

// Import the packages
import (
	"fmt"
	"math/rand"
	"time"
)

// main is the start point of execution
func main() {
	println("Hello", "World!")
	fmt.Println("Current Time ", time.Now())
	fmt.Println("Random Number", rand.Intn(9999))

}
