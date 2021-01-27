package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}
	
	// comment below  
	message := fmt.Sprintf(randomFormats(), name)
	// uncomment below to see the test failing
	// message := fmt.Sprintf(randomFormats())
	return message, nil
}


func Hellos(names []string) (map[string]string, error) {
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


func init() {
	rand.Seed(time.Now().UnixNano())
}


func randomFormats() string {
	formats := []string {
		"Hi, %v. Welcome",
		"Great to see you, %v!",
		"%v! Well met!",
	}

	return formats[rand.Intn(len(formats))]
}
