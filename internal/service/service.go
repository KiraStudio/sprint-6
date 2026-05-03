package service

import (
	"fmt"
	morse "sprint-6/pkg/morse"
	"strings"
)

func isMorse(input string) (bool, error) {
	text := strings.TrimSpace(input)
	if input == "" {
		return false, fmt.Errorf("input contains empty string")
	}

	text = strings.ReplaceAll(text, "-", "")
	text = strings.ReplaceAll(text, ".", "")
	text = strings.ReplaceAll(text, " ", "")

	if text == "" {
		return true, nil
	}

	return false, nil
}

func UniversalConverter(input string) (string, error) {
	innerMorse, err := isMorse(input)
	if err != nil {
		return "", err
	}

	result := ""

	if innerMorse {
		result = morse.ToText(input)
	} else {
		result = morse.ToMorse(input)
	}

	return result, nil
}
