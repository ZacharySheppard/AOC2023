package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Part struct {
	xRange []int
	y      int
	value  int
}

type Symbol struct {
	value rune
	x     int
	y     int
}

func isNumber(r rune) bool {
	convertedValue := r - '0'
	if convertedValue < 0 {
		return false
	}
	if convertedValue > 9 {
		return false
	}

	return true
}

func isSymbol(r rune) bool {
	if r == '.' {
		return false
	}

	return !isNumber(r)
}

func parsePartNumbers(line string, y int) []Part {
	parts := []Part{}
	for x := 0; x < len(line)-1; x++ {
		current := rune(line[x])
		if !isNumber(current) {
			continue
		}
		valueString := ""
		start := x
		for isNumber(current) && x < len(line) {
			valueString += string(line[x])
			x++
			if x > len(line)-1 {
				break
			}
			current = rune(line[x])
		}
		stop := x - 1
		value, err := strconv.Atoi(valueString)
		if err != nil {
			panic("bad value")
		}
		part := Part{[]int{start, stop}, y, value}
		parts = append(parts, part)
	}

	return parts
}

func parseSymbols(line string, y int) []Symbol {
	locations := []Symbol{}

	for x, char := range line {
		if isSymbol(char) {
			locations = append(locations, Symbol{char, x, y})
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

func distance(a int, b int) int {
	if a-b > 0 {
		return a - b
	} else {
		return b - a
	}
}

func makeRange(a int, b int) []int {
	rng := []int{}
	for i := a; i <= b; i++ {
		rng = append(rng, i)
	}
	return rng
}

func isAdjacent(part Part, symbol Symbol) bool {
	if distance(symbol.y, part.y) > 1 {
		return false
	}
	start := part.xRange[0]
	stop := part.xRange[1]
	anyCloseEnough := false
	for _, position := range makeRange(start, stop) {
		if distance(position, symbol.x) <= 1 {
			anyCloseEnough = true
			break
		}
	}
	return anyCloseEnough

}

func partOne(parts []Part, symbols []Symbol) int {
	sum := 0
	for _, part := range parts {
		for _, symbol := range symbols {

			if isAdjacent(part, symbol) {
				sum += part.value
				break
			}
		}
	}
	return sum
}

func partTwo(parts []Part, symbols []Symbol) int {
	sum := 0
	for _, symbol := range symbols {
		if symbol.value != '*' {
			continue
		}
		adjacent := []Part{}
		for _, part := range parts {
			if isAdjacent(part, symbol) {
				adjacent = append(adjacent, part)
			}
		}

		if len(adjacent) == 2 {
			sum += adjacent[0].value * adjacent[1].value
		}

	}
	return sum
}

func main() {
	symbols := []Symbol{}
	parts := []Part{}
	input := parseInput("input.txt")
	for y, line := range input {
		symbols = append(symbols, parseSymbols(line, y)...)
		parts = append(parts, parsePartNumbers(line, y)...)
	}
	fmt.Println(partOne(parts, symbols))
	fmt.Println(partTwo(parts, symbols))
}
