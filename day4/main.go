package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
)

func main() {
	// Part 1
	p1Input := "iwrupvqb"
	p1Output := Part1(p1Input)

	fmt.Printf("Day 4, Part 1 Output: %d\n", p1Output)
}

func Part1(input string) int {
	// Have a variable which is just the number, starts at 0 which will be used to increment the number (this is the value to return at the end)
	answer := 0

	// Have a variable for the ongoing for loop where the hash is not invalid, start at false (changes to true and stops the loop when a hash with five 0000s)
	validHashFound := false

	// For loop starts here
	for validHashFound != true {
	// -- For each iteration:
	//		-- Get the current number above
		currentNumToTry := answer
	//		-- Concatenate it to the given string (e.g. 'abcde + 1 (the num)') using .CreateFullSecretKey()
		fullSecretKeyToTry := CreateFullSecretKey(input, currentNumToTry)
	//		-- Get the MD5 hash using .GetMD5Hash()
		md5HashToTry := GetMD5Hash(fullSecretKeyToTry)
	//		-- Check if the hash begins with five zeros using .HashBeingsWithFiveZeros()
		hasBeginsWithFiveZeros := HashBeingsWithFiveZeros(md5HashToTry)

		if (hasBeginsWithFiveZeros) {
		//		-- If it does begin with 5 zeros, change the variable above to true to stop the loop
			validHashFound = true
		} else {
		//		-- If it does not begin with 5 zeros, increment the number variable above so it goes round again and tries the next number	
			answer++
		}
	}

	// Return the number tracked above
	return answer
}

func Part2() {}

func CreateFullSecretKey(secretKey string, numToAdd int) string {
	if (len(secretKey) == 0) {
		return ""
	}

	var fullSecretKey string

  	formattedNumAsStr := strconv.Itoa(numToAdd)

	fullSecretKey = secretKey + formattedNumAsStr

	return fullSecretKey
}

func GetMD5Hash(inputToConvert string) string {
	if (len(inputToConvert) == 0) {
		return ""
	}

	hash := md5.Sum([]byte(inputToConvert))

	return hex.EncodeToString(hash[:])
}

func HashBeingsWithFiveZeros(hashCode string) bool {
	if len(hashCode) >= 5 {

        if hashCode[:5] == "00000" {
            return true
        } else {
            return false
        }
    }

	return false
}