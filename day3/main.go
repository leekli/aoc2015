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

	moveInput := RemoveNonValidChars(rawInput)

	// Part 1
	housesGrid := BuildHousesGrid(2501, 2501)
	
	p1Output := Part1(moveInput, housesGrid)

	fmt.Printf("Day 3, Part 1 Output: %d\n", p1Output)

	// Part 2
	housesGrid = BuildHousesGrid(2501, 2501)
	
	p2Output := Part2(moveInput, housesGrid)

	fmt.Printf("Day 3, Part 2 Output: %d\n", p2Output)
}

func Part1(moves string, grid [][]int) int {
	// Turn string into array
	movesArray := strings.Split(moves, "")

	// If moves are empty, return straight away
	if (len(movesArray) == 0) {
		return 0
	}

	// Have a total presents var 
	totalPresentsDelivered := 0

	// Set the starting point of the grid (find middle position function)
	startingRow, startingCol := FindMiddleStartingPoint(grid)

	// Have a variable to track current position
	// 	- Set current position var to starting position
	currentPos := []int{startingRow, startingCol}

	// 	- Mark this starting location to 1 as visited
	grid[currentPos[0]][currentPos[1]] = 1

	//	- Increase total presents var to +1 for starting position present delivered
	totalPresentsDelivered += 1

	// Start looping through list of moves
	for i := 0; i < len(movesArray); i++ {
		// 	Log where we currently are
		currentRow := currentPos[0]
		currentCol := currentPos[1]

		// Move current position to the next move position depending if going north, south, east, west...

		// 	- Moving north ^
		if (movesArray[i] == "^") {
			currentPos = []int{currentRow - 1, currentCol}
		}

		// 	- Moving south v
		if (movesArray[i] == "v") {
			currentPos = []int{currentRow + 1, currentCol}
		}

		// 	- Moving east >
		if (movesArray[i] == ">") {
			currentPos = []int{currentRow, currentCol + 1}
		}

		// 	- Moving west <
		if (movesArray[i] == "<") {
			currentPos = []int{currentRow , currentCol - 1}
		}

		//	- Check if new, current house location has been visited already or not (true/false)
		nodeVisited := HasHouseBeenVisited(grid, currentPos)

		// 	- If false (house not visited) then increase total presents delivered by 1 and mark house as visited
		if (!nodeVisited) {
			totalPresentsDelivered += 1

			grid[currentPos[0]][currentPos[1]] = 1
		}
	}

	// Return the total presents delivered number
	return totalPresentsDelivered
}

func Part2(moves string, grid [][]int) int {
	// Turn string into array
	movesArray := strings.Split(moves, "")

	// If moves are empty, return straight away
	if (len(movesArray) == 0) {
		return 0
	}

	// Have a total presents var 
	totalPresentsDelivered := 0

	// Start both at the same location (FindMiddleStartingPoint)
	// - Both Santas to hold the same position for the start
	startingRow, startingCol := FindMiddleStartingPoint(grid)
	santa1CurrentPos := []int{startingRow, startingCol}
	santa2CurrentPos := []int{startingRow, startingCol}

	// 	- Mark this starting location to 1 as visited
	grid[startingRow][startingCol] = 1

	//	- Increase total presents var to +1 for starting position present delivered by both Santa's
	totalPresentsDelivered += 1

	// A var to track which santa's turn it is next (Start with 'Santa 1' before loop starts) - 1 for santa 1, 2 for santa 2
	var santaToGoNext int = 1

	// Start going through moves list
	for i := 0; i < len(movesArray); i++ {
		// Work out which Santa's turn it is (using the var above)
		var currentSanta = santaToGoNext

		// 	Log where we currently are & update depending on which Santa we are
		var currentRow, currentCol int

		if currentSanta == 1 {
			currentRow, currentCol = santa1CurrentPos[0], santa1CurrentPos[1]
		} else {
			currentRow, currentCol = santa2CurrentPos[0], santa2CurrentPos[1]		
		}

		// Move current position to the next move position depending if going north, south, east, west...

		// 	- Moving north ^
		if (movesArray[i] == "^") {
			newPos := []int{currentRow - 1, currentCol}

			if currentSanta == 1 {
				santa1CurrentPos = newPos
			} else {
				santa2CurrentPos = newPos
			}
		}

		// 	- Moving south v
		if (movesArray[i] == "v") {
			newPos := []int{currentRow + 1, currentCol}

			if currentSanta == 1 {
				santa1CurrentPos = newPos
			} else {
				santa2CurrentPos = newPos
			}
		}

		// 	- Moving east >
		if (movesArray[i] == ">") {
			newPos := []int{currentRow, currentCol + 1}

			if currentSanta == 1 {
				santa1CurrentPos = newPos
			} else {
				santa2CurrentPos = newPos
			}
		}

		// 	- Moving west <
		if (movesArray[i] == "<") {
			newPos := []int{currentRow , currentCol - 1}

			if currentSanta == 1 {
				santa1CurrentPos = newPos
			} else {
				santa2CurrentPos = newPos
			}
		}

		//	- Check if new, current house location has been visited already or not (true/false)
		var nodeVisited bool

		if currentSanta == 1 {
			nodeVisited = HasHouseBeenVisited(grid, santa1CurrentPos)
		} else {
			nodeVisited = HasHouseBeenVisited(grid, santa2CurrentPos)
		}

		// 	- If false (house not visited) then increase total presents delivered by 1 and mark house as visited
		if (!nodeVisited) {
			totalPresentsDelivered += 1

			if currentSanta == 1 {
				grid[santa1CurrentPos[0]][santa1CurrentPos[1]] = 1
			} else {
				grid[santa2CurrentPos[0]][santa2CurrentPos[1]] = 1
			}
		}

		// Switch the Santa ready for the next move
		if currentSanta == 1 {
			santaToGoNext = 2
		} else {
			santaToGoNext = 1
		}
	}

	// Return the total presents delivered number
	return totalPresentsDelivered
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
		if char == '^' || char == 'v' || char == '>' || char == '<' {
			amendedOutput.WriteRune(char)
		} else if !unicode.IsSpace(char) {
			// Ignore everything else
		}
	}

	return amendedOutput.String()
}

func BuildHousesGrid(rows int, cols int) [][]int {
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

func FindMiddleStartingPoint(grid [][]int) (int, int) {
	// Assume all grids are odd numbers, could/should make it work for evens eventually 
	// Also assumes an identical rows & cols grid
	
	middleRow := len(grid) / 2
	middleCol := len(grid[0]) / 2

	return middleRow, middleCol
}

func HasHouseBeenVisited(grid [][]int, coOrds []int) bool {
	coOrdValue := grid[coOrds[0]][coOrds[1]]

	if (coOrdValue == 1) {
		return true
	}

	return false
}