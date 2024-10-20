package main

import "testing"

func TestCreateFullSecretKey_ReturnsEmptyStringForEmptyInput(test *testing.T) {
	inputStr := ""
	inputNum := 0

	output := CreateFullSecretKey(inputStr, inputNum)

	if output != "" {
		test.Errorf("Expected: '', Received: %s", output)
	}
}

func TestCreateFullSecretKey_ReturnsNewStringWithSingleCharAndSingleNum(test *testing.T) {
	inputStr := "a"
	inputNum := 1

	output := CreateFullSecretKey(inputStr, inputNum)

	if output != "a1" {
		test.Errorf("Expected: 'a1', Received: %s", output)
	}

	inputStr = "a"
	inputNum = 9

	output = CreateFullSecretKey(inputStr, inputNum)

	if output != "a9" {
		test.Errorf("Expected: 'a9', Received: %s", output)
	}
}

func TestCreateFullSecretKey_ReturnsNewStringWithMultipleCharsAndNums(test *testing.T) {
	inputStr := "abcdef"
	inputNum := 7538

	output := CreateFullSecretKey(inputStr, inputNum)

	if output != "abcdef7538" {
		test.Errorf("Expected: 'abcdef7538', Received: %s", output)
	}
}

func TestGetMD5Hash_ReturnsEmptyStringWhenGivenEmptyInput(test *testing.T) {
	input := ""

	output := GetMD5Hash(input)

	if output != "" {
		test.Errorf("Expected: '', Received: %s", output)
	}
}

func TestGetMD5Hash_ReturnsTheMD5HashCode(test *testing.T) {
	input := "a"

	output := GetMD5Hash(input)

	if output != "0cc175b9c0f1b6a831c399e269772661" {
		test.Errorf("Expected: '0cc175b9c0f1b6a831c399e269772661', Received: %s", output)
	}

	input = "abcde12345"

	output = GetMD5Hash(input)

	if output != "df6f58808ebfd3e609c234cf2283a989" {
		test.Errorf("Expected: 'df6f58808ebfd3e609c234cf2283a989', Received: %s", output)
	}
}

func TestHashBeingsWithNumOfZeros_ReturnsFalseForStringLessThanFiveCharsLong(test *testing.T) {
	input := "abcd"
	inputNums := 5

	output := HashBeingsWithNumOfZeros(input, inputNums)

	if output != false {
		test.Errorf("Expected: false, Received: %t", output)
	}
}

func TestHashBeingsWithNumOfZeros_ReturnsFalseForCodeWhichDoesNotStartWithFiveZeros(test *testing.T) {
	input := "df6f58808ebfd3e609c234cf2283a989"
	inputNums := 5

	output := HashBeingsWithNumOfZeros(input, inputNums)

	if output != false {
		test.Errorf("Expected: false, Received: %t", output)
	}

	input = "0000dkf7bx"
	inputNums = 5

	output = HashBeingsWithNumOfZeros(input, inputNums)

	if output != false {
		test.Errorf("Expected: false, Received: %t", output)
	}
}

func TestHashBeingsWithNumOfZeros_ReturnsTrueForCodeWhichDoesStartWithFiveZeros(test *testing.T) {
	input := "00000df08ebfd3e609c234cf2283a989"
	inputNums := 5

	output := HashBeingsWithNumOfZeros(input, inputNums)

	if output != true {
		test.Errorf("Expected: true, Received: %t", output)
	}
}

func TestHashBeingsWithNumOfZeros_ReturnsFalseWhenSixZerosAsInput(test *testing.T) {
	input := "00000df08ebfd3e609c234cf2283a989"
	inputNums := 6

	output := HashBeingsWithNumOfZeros(input, inputNums)

	if output != false {
		test.Errorf("Expected: false, Received: %t", output)
	}
}

func TestHashBeingsWithNumOfZeros_ReturnsTrueWhenSixZerosAsInput(test *testing.T) {
	input := "000000df08ebfd3e609c234cf2283a989"
	inputNums := 6

	output := HashBeingsWithNumOfZeros(input, inputNums)

	if output != true {
		test.Errorf("Expected: true, Received: %t", output)
	}
}

func TestSolve_ReturnsValidHashWithFiveZeros(test *testing.T) {
	input := "abcdef"
	inputNum := 5

	output := Solve(input, inputNum)

	if output != 609043 {
		test.Errorf("Expected: 609043, Received: %d", output)
	}

	input = "pqrstuv"
	inputNum = 5

	output = Solve(input, inputNum)

	if output != 1048970 {
		test.Errorf("Expected: 1048970, Received: %d", output)
	}

	input = "iwrupvqb"
	inputNum = 5

	output = Solve(input, inputNum)

	if output != 346386 {
		test.Errorf("Expected: 346386, Received: %d", output)
	}
}

func TestSolve_ReturnsValidHashWithSixZeros(test *testing.T) {
	input := "iwrupvqb"
	inputNum := 6

	output := Solve(input, inputNum)

	if output != 9958218 {
		test.Errorf("Expected: 9958218, Received: %d", output)
	}
}