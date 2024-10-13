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
	input := "@"

	output := RemoveNonValidChars(input)

	if output != "" {
		test.Errorf("Expected: '', Received: %s", output)
	}

	input = "£$@%$"

	output = RemoveNonValidChars(input)

	if output != "" {
		test.Errorf("Expected: '', Received: %s", output)
	}
}

func TestRemoveNonValidChars_ReturnsCorrectStringForValidSingleChars(test *testing.T) {
	input := "^"

	output := RemoveNonValidChars(input)

	if output != "^" {
		test.Errorf("Expected: '^', Received: %s", output)
	}

	input = "v"

	output = RemoveNonValidChars(input)

	if output != "v" {
		test.Errorf("Expected: 'v', Received: %s", output)
	}

	input = ">"

	output = RemoveNonValidChars(input)

	if output != ">" {
		test.Errorf("Expected: '>', Received: %s", output)
	}

	input = "<"

	output = RemoveNonValidChars(input)

	if output != "<" {
		test.Errorf("Expected: '<', Received: %s", output)
	}
}

func TestRemoveNonValidChars_ReturnsCorrectStringForValidMixedChars(test *testing.T) {
	input := "£&(^)"

	output := RemoveNonValidChars(input)

	if output != "^" {
		test.Errorf("Expected: '^', Received: %s", output)
	}

	input = "£$^$*&^(&*<>)"

	output = RemoveNonValidChars(input)

	if output != "^^<>" {
		test.Errorf("Expected: '^^<>', Received: %s", output)
	}

	input = "!@<&*(>)$^hv)£x^"

	output = RemoveNonValidChars(input)

	if output != "<>^v^" {
		test.Errorf("Expected: '<>^v^', Received: %s", output)
	}

	input = "^^^<<<>>>vvv^^^><v^"

	output = RemoveNonValidChars(input)

	if output != "^^^<<<>>>vvv^^^><v^" {
		test.Errorf("Expected: '^^^<<<>>>vvv^^^><v^', Received: %s", output)
	}
}

func TestBuildHousesGrid_ReturnsCorrectSizeGrid(test *testing.T) {
	rows := 21
	cols := 21

	output := BuildHousesGrid(rows, cols)

	if len(output) != rows {
		test.Errorf("Expected: 21, Received: %d", len(output))
	}

	rows = 101
	cols = 101

	output = BuildHousesGrid(rows, cols)

	if len(output) != rows {
		test.Errorf("Expected: 101, Received: %d", len(output))
	}
}

func TestBuildHousesGrid_ReturnsGridAllSetToZero(test *testing.T) {
	rows := 21
	cols := 21

	output := BuildHousesGrid(rows, cols)

	if output[0][0] != 0 {
		test.Errorf("Expected: 0, Received: %d", len(output))
	}

	if output[4][9] != 0 {
		test.Errorf("Expected: 0, Received: %d", len(output))
	}

	if output[19][20] != 0 {
		test.Errorf("Expected: 0, Received: %d", len(output))
	}
}

func TestFindMiddleStartingPoint_ReturnsCorrectMiddlePointIndexes(test *testing.T) {
	rows := 21
	cols := 21
	testHouseGrid := BuildHousesGrid(rows, cols)

	middleRow, middleCol := FindMiddleStartingPoint(testHouseGrid)

	if middleRow != 10 && middleCol != 10 {
		test.Errorf("Expected: middleRow to be 10 and middleCol to be 10, Received: %d and %d", middleRow, middleCol)
	}

	rows = 201
	cols = 201
	testHouseGrid = BuildHousesGrid(rows, cols)

	middleRow, middleCol = FindMiddleStartingPoint(testHouseGrid)

	if middleRow != 100 || middleCol != 100 {
		test.Errorf("Expected: middleRow to be 100 and middleCol to be 10, Received: %d and %d", middleRow, middleCol)
	}
}

func TestHasHouseBeenVisited_ReturnsFalseForNotVisited(test *testing.T) {
	rows := 21
	cols := 21
	testHouseGrid := BuildHousesGrid(rows, cols)

	testHouseGrid[3][2] = 0

	testCoOrdsToCheck := []int{3, 2}

	output := HasHouseBeenVisited(testHouseGrid, testCoOrdsToCheck)

	if output != false {
		test.Errorf("Expected: false, Received: %t", output)
	}
}

func TestHasHouseBeenVisited_ReturnsTrueForVisited(test *testing.T) {
	rows := 21
	cols := 21
	testHouseGrid := BuildHousesGrid(rows, cols)

	testHouseGrid[6][9] = 1

	testCoOrdsToCheck := []int{6, 9}

	output := HasHouseBeenVisited(testHouseGrid, testCoOrdsToCheck)

	if output != true {
		test.Errorf("Expected: true, Received: %t", output)
	}
}

func TestPartOne_ReturnsZeroIfMovesAreEmpty(test *testing.T) {
	testMoves := ""

	rows := 6
	cols := 6
	testHouseGrid := BuildHousesGrid(rows, cols)

	output := Part1(testMoves, testHouseGrid)

	if output != 0 {
		test.Errorf("Expected: 0, Received: %d", output)
	}
}

func TestPartOne_ReturnsTotalForSingleMove(test *testing.T) {
	// Moving East
	testMoves := ">"

	rows := 6
	cols := 6
	testHouseGrid := BuildHousesGrid(rows, cols)

	output := Part1(testMoves, testHouseGrid)

	if output != 2 {
		test.Errorf("Expected: 2, Received: %d", output)
	}

	// Moving West
	testMoves = "<"

	rows = 6
	cols = 6
	testHouseGrid = BuildHousesGrid(rows, cols)

	output = Part1(testMoves, testHouseGrid)

	if output != 2 {
		test.Errorf("Expected: 2, Received: %d", output)
	}

	// Moving North
	testMoves = "^"

	rows = 6
	cols = 6
	testHouseGrid = BuildHousesGrid(rows, cols)

	output = Part1(testMoves, testHouseGrid)

	if output != 2 {
		test.Errorf("Expected: 2, Received: %d", output)
	}

	// Moving South
	testMoves = "v"

	rows = 6
	cols = 6
	testHouseGrid = BuildHousesGrid(rows, cols)

	output = Part1(testMoves, testHouseGrid)

	if output != 2 {
		test.Errorf("Expected: 2, Received: %d", output)
	}
}

func TestPartOne_ReturnsTotalForMultipleMovesAllNonVisited(test *testing.T) {
	testMoves := "^^>>^^<v"

	rows := 11
	cols := 11
	testHouseGrid := BuildHousesGrid(rows, cols)

	output := Part1(testMoves, testHouseGrid)

	if output != 9 {
		test.Errorf("Expected: 9, Received: %d", output)
	}
}

func TestPartOne_ReturnsTotalForMultipleMovesWithAlreadyVisited(test *testing.T) {
	testMoves := "^>v<"

	rows := 11
	cols := 11
	testHouseGrid := BuildHousesGrid(rows, cols)

	output := Part1(testMoves, testHouseGrid)

	if output != 4 {
		test.Errorf("Expected: 4, Received: %d", output)
	}

	testMoves = "^v^v^v^v^v"

	rows = 13
	cols = 13
	testHouseGrid = BuildHousesGrid(rows, cols)

	output = Part1(testMoves, testHouseGrid)

	if output != 2 {
		test.Errorf("Expected: 2, Received: %d", output)
	}
}