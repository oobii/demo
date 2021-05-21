package greetings2

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func init() {
rand.Seed(time.Now().UnixNano())
}

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}

	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}


func randomFormat() string {

	formats := []string {
		"Hi! %v",
		"Grat to see you, %v",
		"Hey, %v",
	}

	return formats[rand.Intn(len(formats))]
}