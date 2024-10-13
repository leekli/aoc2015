package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	// Part 1
	p1RawInput, err := readFileToString("p1.txt")

	if err != nil {
		os.Exit(-1)
	}

	p1InputNonParenthesis := RemoveNonValidChars(p1RawInput)
	p1OOutput := Day1Part1(p1InputNonParenthesis)

	fmt.Printf("Day 1, Part 1 Output: %d\n", p1OOutput)

	// Part 2
}

func readFileToString(filePath string) (string, error) {
	content, err := os.ReadFile(filePath)

	if err != nil {
		return "", err
	}
	
	return string(content), nil
}

func RemoveNonValidChars(input string) string {
	var amendedOutput strings.Builder

	for _, char := range input {
		if char == '(' || char == ')' {
			amendedOutput.WriteRune(char)
		} else if !unicode.IsSpace(char) {
			// Ignore everything else
		}
	}

	return amendedOutput.String()
}

func Day1Part1(input string) int {
	// Should return the number [int] floor Santa is on

	if len(input) == 0 {
		return 0
	}

	floorNum := 0

	splitInput := strings.Split(input, "")

	for i := 0; i < len(splitInput); i++ {
		if (splitInput[i] == "(") {
			floorNum++
		}

		if (splitInput[i] == ")") {
			floorNum--
		}
	}

	return floorNum
}