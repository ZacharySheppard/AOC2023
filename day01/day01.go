package main

import (
	"fmt"
	"os"
	"strings"
)

type DigitIndex struct {
	postion int
	value   int
}

func isDigit(r rune) bool {
	digit := r - '0'
	if 0 < digit && digit < 10 {
		return true
	} else {
		return false
	}
}

func hasDigitWord(line string) bool {
	words := []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	for _, word := range words {
		if strings.Contains(line, word) {
			return true
		}
	}
	return false
}

func convertDigitWordToDigit(line string) int {
	words := []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	for index, word := range words {
		position := strings.Index(strings.ToLower(line), word)
		if position >= 0 {
			return index
		}
	}
	panic("Couldn't find digit in word")
}

func getFirstDigitFromLine(line string) int {
	substring := ""
	for _, char := range line {
		if isDigit(char) {
			return int(char - '0')
		} else {
			substring += string(char)
			if len(substring) < 3 {
				continue
			}
			if hasDigitWord(substring) {
				return convertDigitWordToDigit(substring)
			}
		}
	}
	panic("no digit in line")
}

func getLastDigitFromLine(line string) int {
	substring := ""
	for i := len(line) - 1; i >= 0; i-- {
		char := rune(line[i])
		if isDigit(char) {
			return int(char - '0')
		} else {
			substring = string(char) + substring
			if len(substring) < 3 {
				continue
			}
			if hasDigitWord(substring) {
				return convertDigitWordToDigit(substring)
			}
		}
	}
	panic("no digit in line")
}

func parseInput(filepath string) string {
	contents, err := os.ReadFile(filepath)
	if err != nil {
		fmt.Println("Couldn't Read file")
	}

	return string(contents)
}

func main() {
	input := parseInput("input.txt")
	lines := strings.Split(input, "\n")

	sum := 0
	for _, line := range lines {
		first := getFirstDigitFromLine(line)
		last := getLastDigitFromLine(line)

		number := 10*first + last
		fmt.Println(number)
		sum += number
	}

	fmt.Println(sum)
}
