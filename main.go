package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	min, max := 1, 100
	// change the initial seed every round
	rand.Seed(time.Now().UnixNano())
	secretNumber := rand.Intn(max - min) + min
	fmt.Println("The secret number is: ", secretNumber)
}