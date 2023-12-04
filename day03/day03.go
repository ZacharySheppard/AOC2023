package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type PartNumber struct {
	start int
	stop  int
	value int
}

type Point struct {
	x int
	y int
}

func isSymbol(r rune) bool {
	symbols := []rune{'&', '*', '!', '@', '%', '#', '^'}
	for _, symbol := range symbols {
		if r == symbol {
			return true
		}
	}
	return false
}

func isNumber(r rune) bool {
	convertedValue := r - '0'
	if convertedValue > 9 {
		return false
	}

	if convertedValue < 0 {
		return false
	}

	return true
}

func parsePartNumbers(line string) []PartNumber {
	parts := []PartNumber{}

	for x := 0; x < len(line)-1; x++ {
		if isNumber(rune(line[x])) {
			part := PartNumber{x, 0, 0}
			valueString := ""
			for isNumber(rune(line[x])) && x < len(line)-1 {
				valueString += string(line[x])
				x++

			}
			part.stop = x - 1
			value, err := strconv.Atoi(valueString)
			if err != nil {
				panic("bad value")
			}
			part.value = value
			parts = append(parts, part)
		}
	}

	return parts
}

func getSymbolXCoord(line string) []int {
	locations := []int{}

	for x, char := range line {
		if isSymbol(char) {
			locations = append(locations, x)
		}
	}
	return locations
}

func parseInput(filename string) []string {
	input, err := os.ReadFile(filename)
	if err != nil {
		panic("invalid file")
	}

	return strings.Split(string(input), "\r\n")
}


func 

func main() {
	symbols := []Point{}
	parts := []PartNumber{}
	input := parseInput("input.txt")
	for y, line := range input {
		for _, x := range getSymbolXCoord(line) {
			symbols = append(symbols, Point{x, y})
			
		}
		incompleteParts := parsePartNumbers(line)
		for _, x 
		parts = append(parts, parsePartNumbers(line)...)
	}
	fmt.Println(parts)
	fmt.Println(symbols)
}
