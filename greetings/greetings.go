package greetings

import (
	"fmt"
	"github.com/google/uuid"
)

// Hello returns a greeting for the named person.
func Hello(name string) string {
	// Return a greeting that embeds the name in a message.
	var id = uuid.NewString()
	message := fmt.Sprintf("Hi, %v. Welcome! - %v", name, id)
	return message
}
