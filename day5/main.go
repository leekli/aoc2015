package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
	"unicode"
)

func main() {
	// Data setup
	rawInput, err := readFileToString("input.txt")

	if err != nil {
		os.Exit(-1)
	}

	parsedInput := ParseInput(rawInput)

	// Part 1
	p1Output := Part1(parsedInput)

	fmt.Printf("Day 5, Part 1 Output: %d\n", p1Output)

	// Part 2
	p2Output := Part2(parsedInput)

	fmt.Printf("Day 5, Part 2 Output: %d\n", p2Output)
}

func Part1(input []string) int {
	niceStringsTotal := 0

	for i := 0; i < len(input); i++ {
		if IsANiceSting(input[i]) {
			niceStringsTotal++
		}
	}

	return niceStringsTotal
}

func Part2(input []string) int {
	niceStringsTotal := 0

	for i := 0; i < len(input); i++ {
		if IsANiceString_Part2(input[i]) {
			niceStringsTotal++
		}
	}

	return niceStringsTotal
}

func readFileToString(filePath string) (string, error) {
	content, err := os.ReadFile(filePath)

	if err != nil {
		return "", err
	}
	
	return string(content), nil
}

func ParseInput(rawData string) []string {
	if len(rawData) == 0 {
		return []string{""}
	}

	var parsedInput []string

	if !strings.Contains(rawData, "\n") {
		return []string{rawData}
	}

	if strings.Contains(rawData, "\n") {
		parsedInput = strings.Split(rawData, "\n")
	}

	return parsedInput
}

func ContainsAtLeast3Vowels(input string) bool {
	if (len(input) < 3) {
		return false
	}
	
	contains3VowelsResult := false

	vowelsList := []string{"a", "e", "i", "o", "u"}

	vowelCount := 0

	for _, char := range input {
        if slices.Contains(vowelsList, string(char)) {
            vowelCount++
        }
    }

	if vowelCount >= 3 {
		contains3VowelsResult = true
	}

	return contains3VowelsResult
}

func ContainsAtLeastOneLetterTwice(input string) bool {
	if (len(input) < 2) {
		return false
	}

	contains1LetterTwice := false

	for i := 0; i < len(input) - 1; i++ {
		if input[i] == input[i + 1] {
			contains1LetterTwice = true
			break
		}
	}

	return contains1LetterTwice
}

func ContainsDisallowedStrings(input string) bool {
	if (len(input) < 2) {
		return false
	}

	containsDisallowedStrings := false

	disallowedList := []string{"ab", "cd", "pq", "xy"}

	for i := 0; i < len(disallowedList); i++ {
		if strings.Contains(input, disallowedList[i]) {
			containsDisallowedStrings = true
			break
		}
	}

	return containsDisallowedStrings
}

func Contains2LettersAtLeastTwice(input string) bool {
	if (len(input) < 4) {
		return false
	}

	contains2LettersAtLeastTwiceNoOverlaps := false

	var pairsMap = make(map[string]int)

	for i := 0; i < len(input) - 1; i++ {
		firstLetter := string(input[i])
		secondLetter := string(input[i + 1])
		currentPair := firstLetter + secondLetter

		_, keyExists := pairsMap[currentPair]

		if keyExists {
			pairsMap[currentPair] += 1
		} else {
			pairsMap[currentPair] = 1
		}

		if i + 2 < len(input) {
			overLappingLetter := string(input[i + 2])

			if firstLetter == secondLetter && firstLetter == overLappingLetter {
				contains2LettersAtLeastTwiceNoOverlaps = false

				break
			}
		}
	}

	for _, total := range pairsMap {
		if total >= 2 {
			contains2LettersAtLeastTwiceNoOverlaps = true

			break
		}
	}

	return contains2LettersAtLeastTwiceNoOverlaps
}

func ContainsRepeatingLetterWithOneLetterBetween(input string) bool {
	if (len(input) < 3) {
		return false
	}

	contains2LettersWith1LetterBetween := false

	for i := 0; i < len(input) - 2; i++ {
		if input[i] == input[i + 2] {
			if unicode.IsLetter(rune(input[i + 1])) {
				contains2LettersWith1LetterBetween = true
				break			
			}
		}
	}

	return contains2LettersWith1LetterBetween
}

func IsANiceSting(input string) bool {
	niceString := false

	hasCorrectVowels := ContainsAtLeast3Vowels(input)
	hasDoubleLetters := ContainsAtLeastOneLetterTwice(input)
	hasDisAllowedStrings := ContainsDisallowedStrings(input)

	if hasCorrectVowels && hasDoubleLetters && !hasDisAllowedStrings {
		niceString = true
	}

	return niceString
}

func IsANiceString_Part2(input string) bool {
	niceString := false

	has2LettersAtLeastTwice := Contains2LettersAtLeastTwice(input)

	has2LettersWith1Between := ContainsRepeatingLetterWithOneLetterBetween(input)

	if has2LettersAtLeastTwice && has2LettersWith1Between {
		niceString = true
	}

	return niceString
}