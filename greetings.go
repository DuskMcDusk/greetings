package greetings

import (
	"errors"
	"fmt"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("msg cannot be empty")
	}
	message := fmt.Sprintf("Hello %s!", name)
	return message, nil
}
