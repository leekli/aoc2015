package main

import (
	"testing"
)

func TestParseInput_ReturnsEmptyArrayForEmptyInputString(test *testing.T) {
	input := ""

	output := ParseInput(input)
	outputLength := len(output)

	if outputLength != 0 {
		test.Errorf("Expected: 0, Received: %d", outputLength)
	}
}

func TestParseInput_ReturnsInstructionsForToggleSingleCoOrd(test *testing.T) {
	input := "toggle 0,0 through 999,999"

	output := ParseInput(input)

	if output[0][0] != "toggle" {
		test.Errorf("Expected: 'toggle', Received: %s", output[0][0])
	}

	if output[0][1] != "0,0" {
		test.Errorf("Expected: '0,0', Received: %s", output[0][1])
	}

	if output[0][2] != "999,999" {
		test.Errorf("Expected: '999,999', Received: %s", output[0][1])
	}
}

func TestParseInput_ReturnsInstructionsForTurnOnOrOffSingleCoOrd(test *testing.T) {
	input := "turn on 0,0 through 999,999"

	output := ParseInput(input)

	if output[0][0] != "on" {
		test.Errorf("Expected: 'on', Received: %s", output[0][0])
	}

	if output[0][1] != "0,0" {
		test.Errorf("Expected: '0,0', Received: %s", output[0][1])
	}

	if output[0][2] != "999,999" {
		test.Errorf("Expected: '999,999', Received: %s", output[0][1])
	}

	input = "turn off 0,0 through 999,999"

	output = ParseInput(input)

	if output[0][0] != "off" {
		test.Errorf("Expected: 'off', Received: %s", output[0][0])
	}

	if output[0][1] != "0,0" {
		test.Errorf("Expected: '0,0', Received: %s", output[0][1])
	}

	if output[0][2] != "999,999" {
		test.Errorf("Expected: '999,999', Received: %s", output[0][1])
	}
}

func TestParseInput_ReturnsInstructionsForMultiLineInstructions(test *testing.T) {
	input := "turn on 0,0 through 999,999\ntoggle 0,0 through 999,0\nturn off 499,499 through 500,500"

	output := ParseInput(input)

	if output[0][0] != "on" {
		test.Errorf("Expected: 'on', Received: %s", output[0])
	}

	if output[0][2] != "999,999" {
		test.Errorf("Expected: '999,999', Received: %s", output[0])
	}

	if output[1][0] != "toggle" {
		test.Errorf("Expected: 'toggle', Received: %s", output[0])
	}

	if output[1][1] != "0,0" {
		test.Errorf("Expected: '0,0', Received: %s", output[0])
	}

	if output[2][0] != "off" {
		test.Errorf("Expected: 'off', Received: %s", output[0])
	}

	if output[2][1] != "499,499" {
		test.Errorf("Expected: '499,499', Received: %s", output[0])
	}

	if output[2][2] != "500,500" {
		test.Errorf("Expected: '500,500', Received: %s", output[0])
	}
}

func TestBuildHouseGrid_ReturnsCorrectSizeGrid(test *testing.T) {
	rows := 21
	cols := 21

	output := BuildHouseGrid(rows, cols)

	if len(output) != rows {
		test.Errorf("Expected: 21, Received: %d", len(output))
	}

	rows = 101
	cols = 101

	output = BuildHouseGrid(rows, cols)

	if len(output) != rows {
		test.Errorf("Expected: 101, Received: %d", len(output))
	}
}

func TestBuildHouseGrid_ReturnsGridAllSetToZero(test *testing.T) {
	rows := 21
	cols := 21

	output := BuildHouseGrid(rows, cols)

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

func TestChangeBulbState_ChangesOffToOn(test *testing.T) {
	rows := 5
	cols := 5

	grid := BuildHouseGrid(rows, cols)

	instruction := "on"
	rowToChange := 3
	colToChnage := 2
	
	ChangeBulbState(grid, instruction, rowToChange, colToChnage)

	if grid[3][2] != 1 {
		test.Errorf("Expected: 1, Received: %d", grid[3][2])
	}
}

func TestChangeBulbState_ChangesOnToOff(test *testing.T) {
	rows := 5
	cols := 5

	grid := BuildHouseGrid(rows, cols)

	grid[1][4] = 1

	instruction := "off"
	rowToChange := 1
	colToChnage := 4
	
	ChangeBulbState(grid, instruction, rowToChange, colToChnage)

	if grid[1][4] != 0 {
		test.Errorf("Expected: 0, Received: %d", grid[1][4])
	}
}

func TestChangeBulbState_TogglesBulb(test *testing.T) {
	rows := 5
	cols := 5

	grid := BuildHouseGrid(rows, cols)

	grid[2][1] = 1

	instruction := "toggle"
	rowToChange := 2
	colToChnage := 1
	
	ChangeBulbState(grid, instruction, rowToChange, colToChnage)

	if grid[2][1] != 0 {
		test.Errorf("Expected: 0, Received: %d", grid[2][1])
	}

	instruction = "toggle"
	rowToChange = 4
	colToChnage = 3
	
	ChangeBulbState(grid, instruction, rowToChange, colToChnage)

	if grid[4][3] != 1 {
		test.Errorf("Expected: 1, Received: %d", grid[4][3])
	}
}

func TestExecuteInstruction_Part1_CompletesForSingleInstructionSingleBulb(test *testing.T) {
	rows := 5
	cols := 5

	grid := BuildHouseGrid(rows, cols)

	testInstruction := []string{"on", "0,0", "0,0"}

	ExecuteInstruction_Part1(grid, testInstruction)

	if grid[0][0] != 1 {
		test.Errorf("Expected: 1, Received: %d", grid[0][0])
	}

	testInstruction = []string{"on", "2,2", "2,2"}

	ExecuteInstruction_Part1(grid, testInstruction)

	if grid[2][2] != 1 {
		test.Errorf("Expected: 1, Received: %d", grid[2][2])
	}

	testInstruction = []string{"off", "0,0", "0,0"}

	ExecuteInstruction_Part1(grid, testInstruction)

	if grid[0][0] != 0 {
		test.Errorf("Expected: 0, Received: %d", grid[0][0])
	}

	testInstruction = []string{"toggle", "3,3", "3,3"}

	ExecuteInstruction_Part1(grid, testInstruction)

	if grid[3][3] != 1 {
		test.Errorf("Expected: 1, Received: %d", grid[3][3])
	}
}

func TestExecuteInstruction_Part1_CompletesForInstructionsForBulbRange(test *testing.T) {
	rows := 1000
	cols := 1000

	grid := BuildHouseGrid(rows, cols)

	testInstruction := []string{"on", "0,0", "999,999"}

	ExecuteInstruction_Part1(grid, testInstruction)

	if grid[0][0] != 1 {
		test.Errorf("Expected: 1, Received: %d", grid[0][0])
	}

	if grid[10][16] != 1 {
		test.Errorf("Expected: 1, Received: %d", grid[10][16])
	}

	if grid[345][234] != 1 {
		test.Errorf("Expected: 1, Received: %d", grid[2][2])
	}

	if grid[999][999] != 1 {
		test.Errorf("Expected: 1, Received: %d", grid[3][3])
	}

	grid = BuildHouseGrid(rows, cols)

	testInstruction = []string{"toggle", "0,0", "999,0"}

	ExecuteInstruction_Part1(grid, testInstruction)

	if grid[0][0] != 1 {
		test.Errorf("Expected: 1, Received: %d", grid[0][0])
	}

	if grid[222][0] != 1 {
		test.Errorf("Expected: 1, Received: %d", grid[222][0])
	}

	if grid[999][0] != 1 {
		test.Errorf("Expected: 1, Received: %d", grid[999][0])
	}

	if grid[345][234] != 0 {
		test.Errorf("Expected: 0, Received: %d", grid[345][234])
	}

	grid = BuildHouseGrid(rows, cols)

	testInstruction = []string{"off", "499,4999", "500,500"}

	grid[567][890] = 1

	ExecuteInstruction_Part1(grid, testInstruction)

	if grid[499][499] != 0 {
		test.Errorf("Expected: 0, Received: %d", grid[499][499])
	}

	if grid[499][500] != 0 {
		test.Errorf("Expected: 0, Received: %d", grid[499][500])
	}

	if grid[500][499] != 0 {
		test.Errorf("Expected: 0, Received: %d", grid[500][499])
	}

	if grid[500][500] != 0 {
		test.Errorf("Expected: 0, Received: %d", grid[500][500])
	}

	if grid[234][999] != 0 {
		test.Errorf("Expected: 0, Received: %d", grid[234][999])
	}

	if grid[567][890] != 1 {
		test.Errorf("Expected: 1, Received: %d", grid[567][890])
	}
}

func TestCountBulbsSwitchedOn_ReturnsCorrectCountForInstructions(test *testing.T) {
	rows := 1000
	cols := 1000

	grid := BuildHouseGrid(rows, cols)

	testInstruction := []string{"on", "0,0", "999,999"}

	ExecuteInstruction_Part1(grid, testInstruction)

	bulbCount := CountBulbsSwitchedOn(grid)

	if bulbCount != 1000000 {
		test.Errorf("Expected: 1000000, Received: %d", bulbCount)
	}

	grid = BuildHouseGrid(rows, cols)

	testInstruction = []string{"toggle", "0,0", "999,0"}

	ExecuteInstruction_Part1(grid, testInstruction)

	bulbCount = CountBulbsSwitchedOn(grid)

	if bulbCount != 1000 {
		test.Errorf("Expected: 1000, Received: %d", bulbCount)
	}

	grid = BuildHouseGrid(rows, cols)

	testInstruction = []string{"on", "499,499", "500,500"}

	ExecuteInstruction_Part1(grid, testInstruction)

	bulbCount = CountBulbsSwitchedOn(grid)

	if bulbCount != 4 {
		test.Errorf("Expected: 4, Received: %d", bulbCount)
	}
}