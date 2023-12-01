package main

import (
	"os"
	"strings"
)

type Number uint8

const (
	One Number = iota
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
)

func convertSubStringToNumber(s string) int {
	switch s {
	case "zero":
		return 0
	case "one":
		return 1
	case "two":
		return 2
	case "three":
		return 3
	case "four":
		return 4
	case "five":
		return 5
	case "six":
		return 6
	case "seven":
		return 7
	case "eight":
		return 8
	case "nine":
		return 9
	default:
		return -1
	}
}

func searchForNumbers(line string) {
	length := len(line)

}

func convertEndsTo2DigitNumber(numbers []int) int {
	length := len(numbers)
	if length < 1 {
		panic("bad input set")
	}
	return 10*numbers[0] + numbers[length-1]
}

func getNumbersFromLine(line string) []int {
	digits := make([]int, 0)

	for _, character := range line {
		value := int(character - '0')
		if 0 <= value && value <= 9 {
			digits = append(digits, value)
			continue
		}

	}
	return digits
}

func getInputAsStringArray(filepath string) []string {

	text, err := os.ReadFile(filepath)
	if err != nil {
		println("couldn't read file")
		return []string{}
	}

	return strings.Split(string(text), "\n")
}

func main() {

	sum := 0
	for _, line := range getInputAsStringArray("input1.txt") {
		numbers := getNumbersFromLine(line)
		sum += convertEndsTo2DigitNumber(numbers)
	}
	println(sum)

}
