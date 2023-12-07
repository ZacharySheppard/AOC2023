package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Triplet struct {
	destination int
	start       int
	count       int
}

func parseInput(filename string) []string {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		panic("Invalid File")
	}
	return strings.Split(string(content), "\r\n\r\n")
}

func parseTriplet(line string) Triplet {
	strings := strings.Split(line, " ")
	numbers := []int{}
	if len(strings) != 3 {
		fmt.Println(strings)
		panic("invalid triplet")
	}

	for _, number := range strings {
		convertedNumber, err := strconv.Atoi(number)
		if err != nil {
			panic("Read bad line")
		}
		numbers = append(numbers, convertedNumber)
	}

	return Triplet{numbers[0], numbers[1], numbers[2]}
}

func maxInt(a int, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}

func makeSeedMap(triplets []Triplet, max int) []int {
	seedMap := make([]int, max)
	for _, triplet := range triplets {
		for i := 0; i < triplet.count; i++ {
			seedMap[triplet.start+i] = triplet.destination + i
		}
	}

	for indx, num := range seedMap {
		if num == 0 {
			seedMap[indx] = indx
		}
	}

	return seedMap
}

func parseMap(block string) []int {

	lines := strings.Split(block, "\r\n")
	triplets := []Triplet{}

	max := 0
	for _, line := range lines[1:] {
		triplet := parseTriplet(line)
		triplets = append(triplets, triplet)
		max = maxInt(triplet.count+triplet.start, max)
	}
	return makeSeedMap(triplets, max)
}

func main() {
	input := parseInput("input.txt")
	fmt.Println(parseMap(input[1]))

}
