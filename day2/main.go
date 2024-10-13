package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	// Data setup
	rawInput, err := readFileToString("input.txt")

	if err != nil {
		os.Exit(-1)
	}

	formattedInput := FormatDimensionsIntoSlice(rawInput)

	// Part 1
	p1Output := Part1(formattedInput)
	fmt.Printf("Day 2, Part 1 Output: %d\n", p1Output)

	// Part 2
	p2Output := Part2(formattedInput)
	fmt.Printf("Day 2, Part 2 Output: %d\n", p2Output)
}

func readFileToString(filePath string) (string, error) {
	content, err := os.ReadFile(filePath)

	if err != nil {
		return "", err
	}
	
	return string(content), nil
}

func FormatDimensionsIntoSlice(input string) [][]int {
	splitStrings := strings.Split(input, "\n")

	rows := len(splitStrings) 
	
	dimensionsSlice := make([][]int, rows)

	for i := 0; i < len(splitStrings); i++ {
		splitString := strings.Split(splitStrings[i], "x")

		intsSlice := []int{}

		for j := 0; j < len(splitString); j++ {
			num, _ := strconv.Atoi(splitString[j])

			intsSlice = append(intsSlice, num)
		}

		for j := 0; j < rows; j++ {
			dimensionsSlice[i] = intsSlice
		}	
	}

	return dimensionsSlice
}

func Part1(input [][]int) int {
	totalSqFt := 0

	for i := 0; i < len(input); i++ {
		surfaceArea := CalcSurfaceArea(input[i])

		totalSqFt += surfaceArea

		smallestArea := FindSmallestArea(input[i])

		totalSqFt += smallestArea
	}

	return totalSqFt
}

func Part2(input [][]int) int {
	totalRibbonNeeded := 0

	for i := 0; i < len(input); i++ {
		ribbonForPresents := CalcRibbonRequiredForPresent(input[i])

		ribbonForBow := CalcRibbonRequiredForBow(input[i])

		totalRibbonNeeded += ribbonForPresents + ribbonForBow
	}

	return totalRibbonNeeded
}

func CalcSurfaceArea(dimensions []int) int {
	length := dimensions[0]
	width := dimensions[1]
	height := dimensions[2]

	calc1 := 2 * length * width
	calc2 := 2 * width * height
	calc3 := 2 * height * length

	total := calc1 + calc2 + calc3

	return total
}

func FindSmallestArea(dimensions []int) int {
	length := dimensions[0]
	width := dimensions[1]
	height := dimensions[2]

	area1 := length * width
	area2 := length * height
	area3 := width * height

	areasList := []int{area1, area2, area3}

	return slices.Min(areasList)
}

func CalcRibbonRequiredForPresent(dimensions []int) int {
	length := dimensions[0]
	width := dimensions[1]
	height := dimensions[2]

	perim1 := 2 * (length + width)
	perim2 := 2 * (length + height)
	perim3 := 2 * (width + height)

	perimsList := []int{perim1, perim2, perim3}

	return slices.Min(perimsList)
}

func CalcRibbonRequiredForBow(dimensions []int) int {
	length := dimensions[0]
	width := dimensions[1]
	height := dimensions[2]

	cubicTotal := length * width * height

	return cubicTotal
}