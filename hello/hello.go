package main

import (
	"fmt"
	"math/rand"

	"example.com/greetings"
)

func main() {
	var r = rand.Intn(10)
	message := greetings.Hello("Gladys")
	message2 := fmt.Sprintf("%v: %v", message, r)
	fmt.Println(message2)
}
