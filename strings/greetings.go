package greetings

import (
	"errors"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
	if name == "" {
		return name, errors.New("empty name")
	}
	message := "Hello " + name
	return message, nil
}
