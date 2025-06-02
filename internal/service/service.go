package service

import (
	"errors"
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func Convert(input string) (string, error) {
	if len(input) == 0 {
		return "", errors.New("empty string")
	}

	isMorseCode := true
	for _, char := range input {
		if !strings.ContainsAny(string(char), ".- ") {
			isMorseCode = false
			break
		}
	}

	var result string

	if isMorseCode {
		result = morse.ToText(input)
	} else {
		result = morse.ToMorse(input) 
	}

	return result, nil
}
