package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {

	if name == "" {
		return "", errors.New("empty name")
	}

	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func Hellos(names []string) (map[string]string, error) {
	// In Go, you initialize a map with the following syntax: make(map[key-type]value-type).
	messages := make(map[string]string)

	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}

		messages[name] = message
	}

	return messages, nil

}

func randomFormat() string {
	// A slice of message formats

	formats := []string{
		"Hi, %v. welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	// return a randomly selected message format
	return formats[rand.Intn(len(formats))]

}
