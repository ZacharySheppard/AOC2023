package main

import (
	"fmt"
	"math"
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

func findMinOf(array []int) int {
	min := math.MaxInt
	for _, val := range array {
		if val < min {
			min = val
		}
	}

	return min
}

func isInRange(triplet Triplet, seed int) bool {
	if triplet.start <= seed && seed <= triplet.start+triplet.count {
		return true
	} else {
		return false
	}
}

func findMapPosition(triplet Triplet, seed int) int {
	for i := 0; i <= triplet.count; i++ {
		if triplet.start+i == seed {
			return triplet.destination + i
		}
	}
	fmt.Println(seed, " with ", triplet)
	panic("always call isInRange before calling this function")
}

func mapTo(triplets []Triplet, seed int) int {
	for _, triplet := range triplets {
		if isInRange(triplet, seed) {
			return findMapPosition(triplet, seed)
		}
	}
	return seed
}

func parseMap(block string) []Triplet {

	lines := strings.Split(block, "\r\n")
	triplets := []Triplet{}
	for _, line := range lines[1:] {
		triplet := parseTriplet(line)
		triplets = append(triplets, triplet)
	}
	return triplets
}

func parseInputNumbers(line string) []int {
	input := strings.Split(line, ":")
	if len(input) < 2 {
		panic("Got bad input expected :")
	}
	seeds := []int{}
	numbers := strings.Split(strings.TrimSpace(input[1]), " ")
	for _, number := range numbers {
		value, err := strconv.Atoi(number)
		if err != nil {
			fmt.Println("with seeds as", seeds)
			panic("Got bad value in seeds string")
		}
		seeds = append(seeds, value)
	}
	return seeds
}

func parseSeeds(line string) []int {
	seeds := parseInputNumbers(line)
	return seeds
}

func makeRange(min int, count int) []int {
	numbers := []int{}
	for i := min; i < min+count; i++ {
		numbers = append(numbers, i)
	}
	return numbers
}

func parseSeedRanges(line string) []int {
	seeds := []int{}
	numbers := parseInputNumbers(line)
	if len(numbers)%2 != 0 {
		panic("Bad input scheme")
	}
	for i := 0; i < len(numbers); i += 2 {
		seeds = append(seeds, makeRange(numbers[i], numbers[i+1])...)
	}
	return seeds
}

func containsOverlap(source Triplet, destination Triplet) bool {
	if source.destination+source.count < destination.start {
		return false
	}

	if source.destination+source.count > destination.start+destination.count {
		return false
	}

	return true
}

func splitIntervals(source Triplet, destination Triplet) []Triplet {
	
	for
}

func flattenTriplets(source []Triplet, destination []Triplet) []Triplet {

	output = []Triplet{}
	for _, triplet := range source {

	}

	return []Triplet{}
}

func partOne(input []string) int {

	seeds := parseSeeds(input[0])
	blocks := input[1:]
	for _, block := range blocks {
		triplets := parseMap(block)
		for index, seed := range seeds {
			seeds[index] = mapTo(triplets, seed)
		}
	}
	return findMinOf(seeds)
}

func partTwo(input []string) int {
	seeds := parseSeedRanges(input[0])
	blocks := input[1:]
	triplets := []Triplet{}
	for _, block := range blocks {
		triplet := parseMap(block)
		triplets = flattenTriplets(triplet, triplets)
	}

	for index, seed := range seeds {
		seeds[index] = mapTo(triplets, seed)
	}
	return findMinOf(seeds)

}

func main() {
	input := parseInput("input.txt")
	fmt.Println(partOne(input))
	fmt.Println(partTwo(input))

}
