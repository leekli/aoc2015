package main

import "testing"

func TestRemoveNonValidChars_ReturnsEmptyStringForEmptyInput(test *testing.T) {
	input := ""

	output := RemoveNonValidChars(input)

	if output != "" {
		test.Errorf("Expected: '', Received: %s", output)
	}
}

func TestRemoveNonValidChars_ReturnsEmptyStringForNonValidChars(test *testing.T) {
	input := "^"

	output := RemoveNonValidChars(input)

	if output != "" {
		test.Errorf("Expected: '', Received: %s", output)
	}

	input = "£$@^%$"

	output = RemoveNonValidChars(input)

	if output != "" {
		test.Errorf("Expected: '', Received: %s", output)
	}
}

func TestRemoveNonValidChars_ReturnsStringContainingOnlyValidCharsForSingleInput(test *testing.T) {
	input := "("

	output := RemoveNonValidChars(input)

	if output != "(" {
		test.Errorf("Expected: '(', Received: %s", output)
	}

	input = ")"

	output = RemoveNonValidChars(input)

	if output != ")" {
		test.Errorf("Expected: ')', Received: %s", output)
	}
}

func TestRemoveNonValidChars_ReturnsStringContainingOnlyValidCharsForMultipleValidInput(test *testing.T) {
	input := "()()()()"

	output := RemoveNonValidChars(input)

	if output != "()()()()" {
		test.Errorf("Expected: '()()()()', Received: %s", output)
	}

	input = "))))))))"

	output = RemoveNonValidChars(input)

	if output != "))))))))" {
		test.Errorf("Expected: '))))))))', Received: %s", output)
	}
}

func TestRemoveNonValidChars_ReturnsStringContainingOnlyValidCharsForMixedInput(test *testing.T) {
	input := ")(*^*&^$%^$£@%$^£$   &*^))(((&%%))"

	output := RemoveNonValidChars(input)

	if output != ")())((())" {
		test.Errorf("Expected: ')())((())', Received: %s", output)
	}
}

func TestPartOne_ReturnsZeroForNoMovement(test *testing.T) {
	input := ""

	output := Day1Part1(input)

	if output != 0 {
		test.Errorf("Expected: 0, Received: %d", output)
	}
}

func TestPartOne_ReturnsNumberForOneMovement(test *testing.T) {
	input := "("

	output := Day1Part1(input)

	if output != 1 {
		test.Errorf("Expected: 1, Received: %d", output)
	}

	input = ")"

	output = Day1Part1(input)

	if output != -1 {
		test.Errorf("Expected: -1, Received: %d", output)
	}
}

func TestPartOne_ReturnsNumberForMultipleSameMovements(test *testing.T) {
	input := "((("

	output := Day1Part1(input)

	if output != 3 {
		test.Errorf("Expected: 3, Received: %d", output)
	}

	input = ")))))"

	output = Day1Part1(input)

	if output != -5 {
		test.Errorf("Expected: -5, Received: %d", output)
	}
}

func TestPartOne_ReturnsNumberForMultipleDifferentMovementsBackToZeroFloor(test *testing.T) {
	input := "(())"

	output := Day1Part1(input)

	if output != 0 {
		test.Errorf("Expected: 0, Received: %d", output)
	}

	input = "()()"

	output = Day1Part1(input)

	if output != 0 {
		test.Errorf("Expected: 0, Received: %d", output)
	}
}

func TestPartOne_ReturnsNumberForMultipleDifferentMovementsToAPositiveFloor(test *testing.T) {
	input := "(()(()("

	output := Day1Part1(input)

	if output != 3 {
		test.Errorf("Expected: 3, Received: %d", output)
	}

	input = "))((((("

	output = Day1Part1(input)

	if output != 3 {
		test.Errorf("Expected: 3, Received: %d", output)
	}
}

func TestPartOne_ReturnsNumberForMultipleDifferentMovementsToANegativeFloor(test *testing.T) {
	input := "())"

	output := Day1Part1(input)

	if output != -1 {
		test.Errorf("Expected: -1, Received: %d", output)
	}

	input = "))("

	output = Day1Part1(input)

	if output != -1 {
		test.Errorf("Expected: -1, Received: %d", output)
	}

	input = ")())())"

	output = Day1Part1(input)

	if output != -3 {
		test.Errorf("Expected: -3, Received: %d", output)
	}
}

func TestPartTwo_ReturnsZeroIfInputEmpty(test *testing.T) {
	input := ""

	output := Day1Part2(input)

	if output != 0 {
		test.Errorf("Expected: 0, Received: %d", output)
	}
}

func TestPartTwo_ReturnsZeroIfSingleCharNeverEntersBasement(test *testing.T) {
	input := "("

	output := Day1Part2(input)

	if output != 0 {
		test.Errorf("Expected: 0, Received: %d", output)
	}
}

func TestPartTwo_ReturnsPositionForSingleCharEnteringBasement(test *testing.T) {
	input := ")"

	output := Day1Part2(input)

	if output != 1 {
		test.Errorf("Expected: 1, Received: %d", output)
	}
}

func TestPartTwo_ReturnsZeroIfMultipleCharNeverEntersBasement(test *testing.T) {
	input := "(((((("

	output := Day1Part2(input)

	if output != 0 {
		test.Errorf("Expected: 0, Received: %d", output)
	}
}

func TestPartTwo_ReturnsPositionForMultiCharEnteringBasement(test *testing.T) {
	input := "())"

	output := Day1Part2(input)

	if output != 3 {
		test.Errorf("Expected: 3, Received: %d", output)
	}

	input = "()())"

	output = Day1Part2(input)

	if output != 5 {
		test.Errorf("Expected: 5, Received: %d", output)
	}
}