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

func TestHashBeingsWithFiveZeros_ReturnsFalseForStringLessThanFiveCharsLong(test *testing.T) {
	input := "abcd"

	output := HashBeingsWithFiveZeros(input)

	if output != false {
		test.Errorf("Expected: false, Received: %t", output)
	}
}

func TestHashBeingsWithFiveZeros_ReturnsFalseForCodeWhichDoesNotStartWithFiveZeros(test *testing.T) {
	input := "df6f58808ebfd3e609c234cf2283a989"

	output := HashBeingsWithFiveZeros(input)

	if output != false {
		test.Errorf("Expected: false, Received: %t", output)
	}

	input = "0000dkf7bx"

	output = HashBeingsWithFiveZeros(input)

	if output != false {
		test.Errorf("Expected: false, Received: %t", output)
	}
}

func TestHashBeingsWithFiveZeros_ReturnsTrueForCodeWhichDoesStartWithFiveZeros(test *testing.T) {
	input := "00000df08ebfd3e609c234cf2283a989"

	output := HashBeingsWithFiveZeros(input)

	if output != true {
		test.Errorf("Expected: true, Received: %t", output)
	}
}

func TestPartOne_ReturnsValidHashWithFiveZeros(test *testing.T) {
	input := "abcdef"

	output := Part1(input)

	if output != 609043 {
		test.Errorf("Expected: 609043, Received: %d", output)
	}

	input = "pqrstuv"

	output = Part1(input)

	if output != 1048970 {
		test.Errorf("Expected: 1048970, Received: %d", output)
	}
}