package main

import "testing"

func TestFindSmallestArea_ReturnsNumIfLandWAreSmallestArea(test *testing.T) {
	input := []int{2, 3, 4}

	output := FindSmallestArea(input)

	if output != 6 {
		test.Errorf("Expected: 6, Received: %d", output)
	}

	input = []int{1, 1, 10}

	output = FindSmallestArea(input)

	if output != 1 {
		test.Errorf("Expected: 1, Received: %d", output)
	}
}

func TestFindSmallestArea_ReturnsNumIfWandHAreSmallestArea(test *testing.T) {
	input := []int{4, 3, 2}

	output := FindSmallestArea(input)

	if output != 6 {
		test.Errorf("Expected: 6, Received: %d", output)
	}
}

func TestFindSmallestArea_ReturnsNumIfLandHAreSmallestArea(test *testing.T) {
	input := []int{1, 4, 2}

	output := FindSmallestArea(input)

	if output != 2 {
		test.Errorf("Expected: 2, Received: %d", output)
	}
}

func TestFormatDimensionsIntoSlice_ReturnsSliceWithOneIntSlice(test *testing.T) {
	input := "2x3x4"

	output := FormatDimensionsIntoSlice(input)

	if output[0][0] != 2 {
		test.Errorf("Expected: 2, Received: %d", output[0][0])
	}

	if output[0][1] != 3 {
		test.Errorf("Expected: 3, Received: %d", output[0][0])
	}

	if output[0][2] != 4 {
		test.Errorf("Expected: 4, Received: %d", output[0][0])
	}

	input = "1x1x10"

	output = FormatDimensionsIntoSlice(input)

	if output[0][0] != 1 {
		test.Errorf("Expected: 1, Received: %d", output[0][0])
	}

	if output[0][1] != 1 {
		test.Errorf("Expected: 1, Received: %d", output[0][0])
	}

	if output[0][2] != 10 {
		test.Errorf("Expected: 10, Received: %d", output[0][0])
	}
}

func TestFormatDimensionsIntoSlice_ReturnsSliceWithTwoIntSlices(test *testing.T) {
	input := "2x3x4\n1x1x10"

	output := FormatDimensionsIntoSlice(input)

	if output[0][0] != 2 {
		test.Errorf("Expected: 2, Received: %d", output[0][0])
	}

	if output[0][1] != 3 {
		test.Errorf("Expected: 3, Received: %d", output[0][0])
	}

	if output[0][2] != 4 {
		test.Errorf("Expected: 4, Received: %d", output[0][0])
	}

	if output[1][0] != 1 {
		test.Errorf("Expected: 1, Received: %d", output[0][0])
	}

	if output[1][1] != 1 {
		test.Errorf("Expected: 1, Received: %d", output[0][0])
	}

	if output[1][2] != 10 {
		test.Errorf("Expected: 10, Received: %d", output[0][0])
	}
}

func TestCalcSurfaceArea_ReturnsSurfaceAreaTotal(test *testing.T) {
	input := []int{2, 3, 4}

	output := CalcSurfaceArea(input)

	if output != 52 {
		test.Errorf("Expected: 52, Received: %d", output)
	}

	input = []int{1, 1, 10}

	output = CalcSurfaceArea(input)

	if output != 42 {
		test.Errorf("Expected: 42, Received: %d", output)
	}
}

func TestPart1_ReturnsSqFtForSingleDimensionSet(test *testing.T) {
	rawInput := "2x3x4"
	formattedInput := FormatDimensionsIntoSlice(rawInput)

	output := Part1(formattedInput)

	if output != 58 {
		test.Errorf("Expected: 58, Received: %d", output)
	}

	rawInput = "1x1x10"
	formattedInput = FormatDimensionsIntoSlice(rawInput)

	output = Part1(formattedInput)

	if output != 43 {
		test.Errorf("Expected: 43, Received: %d", output)
	}
}

func TestPart1_ReturnsSqFtForMultipleDimensionSets(test *testing.T) {
	rawInput := "2x3x4\n1x1x10"
	formattedInput := FormatDimensionsIntoSlice(rawInput)

	output := Part1(formattedInput)

	if output != 101 {
		test.Errorf("Expected: 101, Received: %d", output)
	}
}