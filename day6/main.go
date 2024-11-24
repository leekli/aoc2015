package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
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

	fmt.Printf("Day 6, Part 1 Output: %d\n", p1Output)

	// Part 2
	p2Output := Part2(parsedInput)

	fmt.Printf("Day 6, Part 2 Output: %d\n", p2Output)
}

func Part1(input [][]string) int {
	rows := 1000
	cols := 1000

	grid := BuildHouseGrid(rows, cols)

	for i := 0; i < len(input); i++ {
		currentInstruction := input[i]

		ExecuteInstruction_Part1(grid, currentInstruction)
	}

	bulbOnCount := CountBulbsSwitchedOn(grid)

	return bulbOnCount
}

func Part2(input [][]string) int {
	rows := 1000
	cols := 1000

	grid := BuildHouseGrid(rows, cols)

	for i := 0; i < len(input); i++ {
		currentInstruction := input[i]

		ExecuteInstruction_Part2(grid, currentInstruction)
	}

	bulbOnCount := CountBulbsSwitchedOn_Part2(grid)

	return bulbOnCount
}

func readFileToString(filePath string) (string, error) {
	content, err := os.ReadFile(filePath)

	if err != nil {
		return "", err
	}
	
	return string(content), nil
}

func ParseInput(rawInput string) [][]string {
	if len(rawInput) == 0 {
		return make([][]string, 0)
	}

	var numOfInstructions = len(strings.Split(rawInput, "\n"))

	var parsedInput = make([][]string, numOfInstructions)

	if !strings.Contains(rawInput, "\n") {
		splitInput := strings.Split(rawInput, " ")

		if len(splitInput) == 4 {
			parsedInput[0] = append(parsedInput[0], splitInput[0], splitInput[1], splitInput[3])
		}

		if len(splitInput) == 5 {
			parsedInput[0] = append(parsedInput[0], splitInput[1], splitInput[2], splitInput[4])
		}
	}

	if strings.Contains(rawInput, "\n") {
		instructionsSplit := strings.Split(rawInput, "\n")

		for i := 0; i < len(instructionsSplit); i++ {
			splitInstruction := strings.Split(instructionsSplit[i], " ")

			if len(splitInstruction) == 4 {
				parsedInput[i] = append(parsedInput[i], splitInstruction[0], splitInstruction[1], splitInstruction[3])
			}
	
			if len(splitInstruction) == 5 {
				parsedInput[i] = append(parsedInput[i], splitInstruction[1], splitInstruction[2], splitInstruction[4])
			}
		}
	}

	return parsedInput
}

func BuildHouseGrid(rows int, cols int) [][]int {
	houseGrid := make([][]int, rows)

	for i := range houseGrid {
		houseGrid[i] = make([]int, cols)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			houseGrid[i][j] = 0
		}
	}

	return houseGrid
}

func ChangeBulbState(grid [][]int, instruction string, row int, col int) {
	if instruction == "on" {
		grid[row][col] = 1
	}

	if instruction == "off" {
		grid[row][col] = 0
	}

	if instruction == "toggle" {
		if grid[row][col] == 0 {
			grid[row][col] = 1
		} else {
			grid[row][col] = 0
		}
	}
}

func ChangeBulbState_Part2(grid [][]int, instruction string, row int, col int) {
	if instruction == "on" {
		grid[row][col] += 1
	}

	if instruction == "off" {
		grid[row][col] -= 1

		if grid[row][col] < 0 {
			grid[row][col] = 0
		}
	}

	if instruction == "toggle" {
		grid[row][col] += 2
	}
}

func ExecuteInstruction_Part1(grid [][]int, instructions []string) {
	instructionCode := instructions[0]

	splitFrom := strings.Split(instructions[1], ",")
	splitTo := strings.Split(instructions[2], ",")

	startRow, _ := strconv.Atoi(splitFrom[0])
	finalRow, _ := strconv.Atoi(splitTo[0])
	startCol, _ := strconv.Atoi(splitFrom[1])
	finalCol, _ := strconv.Atoi(splitTo[1])

	for i := startRow; i <= finalRow; i++ {
		for j := startCol; j <= finalCol; j++ {
			ChangeBulbState(grid, instructionCode, i, j)
		}
	}
}

func ExecuteInstruction_Part2(grid [][]int, instructions []string) {
	instructionCode := instructions[0]

	splitFrom := strings.Split(instructions[1], ",")
	splitTo := strings.Split(instructions[2], ",")

	startRow, _ := strconv.Atoi(splitFrom[0])
	finalRow, _ := strconv.Atoi(splitTo[0])
	startCol, _ := strconv.Atoi(splitFrom[1])
	finalCol, _ := strconv.Atoi(splitTo[1])

	for i := startRow; i <= finalRow; i++ {
		for j := startCol; j <= finalCol; j++ {
			ChangeBulbState_Part2(grid, instructionCode, i, j)
		}
	}
}

func CountBulbsSwitchedOn(grid [][]int) int {
	bulbOnCount := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				bulbOnCount++
			}
		}
	}

	return bulbOnCount
}

func CountBulbsSwitchedOn_Part2(grid [][]int) int {
	bulbOnCount := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			bulbOnCount += grid[i][j]
		}
	}

	return bulbOnCount
}