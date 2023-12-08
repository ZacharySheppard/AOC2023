package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Hand struct {
	red   int
	green int
	blue  int
}

func tallyHand(line string) Hand {
	cubes := strings.Split(line, ",")
	hand := Hand{0, 0, 0}
	for _, cube := range cubes {
		formattedString := strings.Trim(cube, " ")
		colorCount := strings.Split(formattedString, " ")
		if len(colorCount) > 2 {
			panic(colorCount)
		}
		value, err := strconv.Atoi(colorCount[0])
		if err != nil {
			panic(colorCount[0])
		}

		switch colorCount[1] {
		case "red":
			hand.red += value
		case "blue":
			hand.blue += value
		case "green":
			hand.green += value
		default:
			panic("Bad input")
		}
	}
	return hand
}

func convertToHands(line string) []string {
	info := strings.Split(line, ":")
	if len(info) > 2 {
		panic("malformed input")
	}

	return strings.Split(info[1], ";")
}

func parseGameInput(filepath string, seperator string) []string {
	contents, err := os.ReadFile(filepath)
	if err != nil {
		fmt.Println("Couldn't Read file")
	}

	return strings.Split(string(contents), seperator)
}

func getGameNumber(line string) int {
	gameAndNumber := strings.Split(line, ":")[0]
	number := strings.Split(gameAndNumber, " ")[1]
	val, err := strconv.Atoi(number)
	if err != nil {
		fmt.Println("Somethign Went Wrong")
	}
	return val
}

func isValidHand(max Hand, current Hand) bool {
	if current.red > max.red {
		return false
	}
	if current.blue > max.blue {
		return false
	}
	if current.green > max.green {
		return false
	}
	return true
}

func partOne(input []string) int {
	total := 0
	maxRound := Hand{12, 13, 14}
	for _, round := range input {
		hands := convertToHands(round)
		gameNumber := getGameNumber(round)
		allHandsValid := true
		for _, hand := range hands {
			currentHand := tallyHand(hand)
			if !isValidHand(maxRound, currentHand) {
				allHandsValid = false
				break
			}
		}

		if allHandsValid {
			total += gameNumber
		}
	}
	return total
}

func partTwo(input []string) int {
	total := 0
	for _, round := range input {
		hands := convertToHands(round)
		smallestHand := Hand{0, 0, 0}
		for _, hand := range hands {
			currentHand := tallyHand(hand)
			smallestHand.red = max(currentHand.red, smallestHand.red)
			smallestHand.green = max(currentHand.green, smallestHand.green)
			smallestHand.blue = max(currentHand.blue, smallestHand.blue)
		}
		total += (smallestHand.red * smallestHand.green * smallestHand.blue)
	}
	return total
}

func main() {
	input := parseGameInput("input.txt", "\r\n")
	fmt.Println(partOne(input))
	fmt.Println(partTwo(input))

}
