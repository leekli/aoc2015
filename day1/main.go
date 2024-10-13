package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	// Data setup
	rawInput, err := readFileToString("input.txt")

	if err != nil {
		os.Exit(-1)
	}

	formattedInput := RemoveNonValidChars(rawInput)

	// Part 1
	p1Output := Day1Part1(formattedInput)

	fmt.Printf("Day 1, Part 1 Output: %d\n", p1Output)

	// Part 2
	p2Output := Day1Part2(formattedInput)

	fmt.Printf("Day 1, Part 2 Output: %d\n", p2Output)
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

func Day1Part2(input string) int {
	// Should return the number [int] which represents the position of the first character that causes Santa to enter the basement (floor -1)

	if len(input) == 0 {
		return 0
	}

	floorNum := 0
	charPosition := 0

	splitInput := strings.Split(input, "")

	for i := 0; i < len(splitInput); i++ {
		if (splitInput[i] == "(") {
			floorNum++
		}

		if (splitInput[i] == ")") {
			floorNum--
		}

		if (floorNum == -1) {
			charPosition = i + 1
			break
		}
	}

	return charPosition
}