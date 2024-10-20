package main

import "testing"

func TestParseInput_ReturnsEmptyStringForEmptyInput(test *testing.T) {
	input := ""

	output := ParseInput(input)

	if output[0] != "" {
		test.Errorf("Expected: '', Received: %s", output)
	}
}

func TestParseInput_ReturnsSingleElementSliceForStringOfOneLine(test *testing.T) {
	input := "aeiouaeiouaeiou"

	output := ParseInput(input)

	if output[0] != "aeiouaeiouaeiou" {
		test.Errorf("Expected: 'aeiouaeiouaeiou', Received: %s", output)
	}
}

func TestParseInput_ReturnsMultipleElementSliceForStringOfMultipleLines(test *testing.T) {
	input := "aeiouaeiouaeiou\nugknbfddgicrmopn"

	output := ParseInput(input)

	if output[0] != "aeiouaeiouaeiou" {
		test.Errorf("Expected: 'aeiouaeiouaeiou', Received: %s", output)
	}

	if output[1] != "ugknbfddgicrmopn" {
		test.Errorf("Expected: 'ugknbfddgicrmopn', Received: %s", output)
	}

	input = "jchzalrnumimnmhp\nhaegwjzuvuyypxyu\ndvszwmarrgswjxmb\naeiouaeiouaeiou\nugknbfddgicrmopn"

	output = ParseInput(input)

	if output[2] != "dvszwmarrgswjxmb" {
		test.Errorf("Expected: 'dvszwmarrgswjxmb', Received: %s", output)
	}

	if output[4] != "ugknbfddgicrmopn" {
		test.Errorf("Expected: 'ugknbfddgicrmopn', Received: %s", output)
	}
}

func TestContainsAtLeast3Vowels_ReturnsFalseForEmptyInput(test *testing.T) {
	input := ""

	output := ContainsAtLeast3Vowels(input)

	if output != false {
		test.Errorf("Expected: 'false', Received: %t", output)
	}
}

func TestContainsAtLeast3Vowels_ReturnsFalseIfInputIsLessThan3Chars(test *testing.T) {
	input := "a"

	output := ContainsAtLeast3Vowels(input)

	if output != false {
		test.Errorf("Expected: 'false', Received: %t", output)
	}

	input = "ae"

	output = ContainsAtLeast3Vowels(input)

	if output != false {
		test.Errorf("Expected: 'false', Received: %t", output)
	}
}

func TestContainsAtLeast3Vowels_ReturnsFalseIfLessThan3Vowels(test *testing.T) {
	input := "zbg"

	output := ContainsAtLeast3Vowels(input)

	if output != false {
		test.Errorf("Expected: 'false', Received: %t", output)
	}

	input = "aev"

	output = ContainsAtLeast3Vowels(input)

	if output != false {
		test.Errorf("Expected: 'false', Received: %t", output)
	}

	input = "isklunbcq"

	output = ContainsAtLeast3Vowels(input)

	if output != false {
		test.Errorf("Expected: 'false', Received: %t", output)
	}
}

func TestContainsAtLeast3Vowels_ReturnsTrueIf3OrMoreVowels(test *testing.T) {
	input := "aei"

	output := ContainsAtLeast3Vowels(input)

	if output != true {
		test.Errorf("Expected: 'true', Received: %t", output)
	}

	input = "aeiou"

	output = ContainsAtLeast3Vowels(input)

	if output != true {
		test.Errorf("Expected: 'true', Received: %t", output)
	}

	input = "xazegov"

	output = ContainsAtLeast3Vowels(input)

	if output != true {
		test.Errorf("Expected: 'true', Received: %t", output)
	}

	input = "ugknbfddgicrmopn"

	output = ContainsAtLeast3Vowels(input)

	if output != true {
		test.Errorf("Expected: 'true', Received: %t", output)
	}
}

func TestContainsAtLeastOneLetterTwice_ReturnsFalseForEmptyInput(test *testing.T) {
	input := ""

	output := ContainsAtLeastOneLetterTwice(input)

	if output != false {
		test.Errorf("Expected: 'false', Received: %t", output)
	}
}

func TestContainsAtLeastOneLetterTwice_ReturnsFalseIfInputIsLessThan2Chars(test *testing.T) {
	input := "a"

	output := ContainsAtLeastOneLetterTwice(input)

	if output != false {
		test.Errorf("Expected: 'false', Received: %t", output)
	}
}

func TestContainsAtLeastOneLetterTwice_ReturnsFalseForTwoLettersNotIdentical(test *testing.T) {
	input := "ab"

	output := ContainsAtLeastOneLetterTwice(input)

	if output != false {
		test.Errorf("Expected: 'false', Received: %t", output)
	}

	input = "izkl"

	output = ContainsAtLeastOneLetterTwice(input)

	if output != false {
		test.Errorf("Expected: 'false', Received: %t", output)
	}

	input = "lalalalalalalalala"

	output = ContainsAtLeastOneLetterTwice(input)

	if output != false {
		test.Errorf("Expected: 'false', Received: %t", output)
	}
}

func TestContainsAtLeastOneLetterTwice_ReturnsTrueForTwoLettersOnlyIdentical(test *testing.T) {
	input := "aa"

	output := ContainsAtLeastOneLetterTwice(input)

	if output != true {
		test.Errorf("Expected: 'true', Received: %t", output)
	}

	input = "zzz"

	output = ContainsAtLeastOneLetterTwice(input)

	if output != true {
		test.Errorf("Expected: 'true', Received: %t", output)
	}

	input = "bbbbbbbb"

	output = ContainsAtLeastOneLetterTwice(input)

	if output != true {
		test.Errorf("Expected: 'true', Received: %t", output)
	}

	input = "lalalalalalalalalalalalalalalalannlalalalala"

	output = ContainsAtLeastOneLetterTwice(input)

	if output != true {
		test.Errorf("Expected: 'true', Received: %t", output)
	}
}

func TestContainsDisallowedStrings_ReturnsFalseForEmptyInput(test *testing.T) {
	input := ""

	output := ContainsDisallowedStrings(input)

	if output != false {
		test.Errorf("Expected: 'false', Received: %t", output)
	}
}

func TestContainsDisallowedStrings_ReturnsFalseIfInputIsLessThan2Chars(test *testing.T) {
	input := "a"

	output := ContainsDisallowedStrings(input)

	if output != false {
		test.Errorf("Expected: 'false', Received: %t", output)
	}
}

func TestContainsDisallowedStrings_ReturnsTrueIfInputContainsDisallowedStrings(test *testing.T) {
	input := "ab"

	output := ContainsDisallowedStrings(input)

	if output != true {
		test.Errorf("Expected: 'true', Received: %t", output)
	}

	input = "cd"

	output = ContainsDisallowedStrings(input)

	if output != true {
		test.Errorf("Expected: 'true', Received: %t", output)
	}

	input = "pq"

	output = ContainsDisallowedStrings(input)

	if output != true {
		test.Errorf("Expected: 'true', Received: %t", output)
	}

	input = "xy"

	output = ContainsDisallowedStrings(input)

	if output != true {
		test.Errorf("Expected: 'true', Received: %t", output)
	}

	input = "abcdpqxy"

	output = ContainsDisallowedStrings(input)

	if output != true {
		test.Errorf("Expected: 'true', Received: %t", output)
	}

	input = "haegwjzuvuyypxyu"

	output = ContainsDisallowedStrings(input)

	if output != true {
		test.Errorf("Expected: 'true', Received: %t", output)
	}
}

func TestIsANiceSting_ReturnsFalseDueToLessThanThreeVowels(test *testing.T) {
	input := "dvszwmarrgswjxmb"

	output := IsANiceSting(input)

	if output != false {
		test.Errorf("Expected: 'false', Received: %t", output)
	}
}

func TestIsANiceSting_ReturnsFalseDueToNoDoubleLetter(test *testing.T) {
	input := "jchzalrnumimnmhp"

	output := IsANiceSting(input)

	if output != false {
		test.Errorf("Expected: 'false', Received: %t", output)
	}
}

func TestIsANiceSting_ReturnsFalseDueContainingDisallowedStrings(test *testing.T) {
	input := "haegwjzuvuyypxyu"

	output := IsANiceSting(input)

	if output != false {
		test.Errorf("Expected: 'false', Received: %t", output)
	}
}

func TestIsANiceSting_ReturnsTrueForANiceString(test *testing.T) {
	input := "ugknbfddgicrmopn"

	output := IsANiceSting(input)

	if output != true {
		test.Errorf("Expected: 'true', Received: %t", output)
	}

	input = "aaa"

	output = IsANiceSting(input)

	if output != true {
		test.Errorf("Expected: 'true', Received: %t", output)
	}
}

func TestPart1_ReturnsZeroAsContainsNoNiceString(test *testing.T) {
	input := []string{"jchzalrnumimnmhp"}

	output := Part1(input)

	if output != 0 {
		test.Errorf("Expected: '0', Received: %d", output)
	}	

	input = []string{"haegwjzuvuyypxyu"}

	output = Part1(input)

	if output != 0 {
		test.Errorf("Expected: '0', Received: %d", output)
	}
	
	input = []string{"dvszwmarrgswjxmb"}

	output = Part1(input)

	if output != 0 {
		test.Errorf("Expected: '0', Received: %d", output)
	}	

	input = []string{"jchzalrnumimnmhp", "haegwjzuvuyypxyu", "dvszwmarrgswjxmb"}

	output = Part1(input)

	if output != 0 {
		test.Errorf("Expected: '0', Received: %d", output)
	}
}

func TestPart1_ReturnsTotalOfNiceStrings(test *testing.T) {
	input := []string{"ugknbfddgicrmopn"}

	output := Part1(input)

	if output != 1 {
		test.Errorf("Expected: '1', Received: %d", output)
	}	

	input = []string{"ugknbfddgicrmopn", "aaa"}

	output = Part1(input)

	if output != 2 {
		test.Errorf("Expected: '2', Received: %d", output)
	}
}