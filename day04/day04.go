package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseInput(filename string) []string {
	input, err := os.ReadFile(filename)
	if err != nil {
		panic("invalid file")
	}

	return strings.Split(string(input), "\r\n")
}

func stripGameNumber(line string) string {
	lines := strings.Split(line, ":")
	if len(lines) > 2 {
		panic("malformed line")
	}

	return lines[len(lines)-1]
}

func parseNumbers(stringOfNumbers string) []int {
	trimmedString := strings.TrimSpace(stringOfNumbers)
	numbersAsStrings := strings.Split(trimmedString, " ")
	numbers := []int{}
	for _, number := range numbersAsStrings {
		if number == "" {
			continue
		}
		trimmedNumber := strings.TrimSpace(number)
		value, err := strconv.Atoi(trimmedNumber)
		if err != nil {
			fmt.Println("malformed input", value)
			panic("")
		}
		numbers = append(numbers, value)
	}
	return numbers
}

func isWinner(number int, winningNumbers []int) bool {
	for _, winningNumber := range winningNumbers {
		if number == winningNumber {
			return true
		}
	}

	return false
}

func getNumberOfWinners(card []int, winningNumbers []int) int {
	numberOfWinners := 0
	for _, myNumber := range card {
		if isWinner(myNumber, winningNumbers) {
			numberOfWinners++
		}
	}
	return numberOfWinners
}

func Pow(base int, exponent int) int {
	if exponent == 0 {
		return 0
	}
	total := 1
	for i := 1; i < exponent; i++ {
		total = total * base
	}
	return total
}

func partOne(winningNumbers [][]int, cards [][]int) int {
	points := 0
	for i := 0; i < len(cards); i++ {
		wins := getNumberOfWinners(cards[i], winningNumbers[i])
		points += Pow(2, wins)
	}
	return points
}

func partTwo(winningNumbers [][]int, cards [][]int) int {
	total := 0
	cardsPerRound := [][][]int //why god why
	for i := 0; i < len(winningNumbers)-1; i++ {
		card := cardsPerRound[i]
		for {

		}
	}
}

func main() {

	input := parseInput("input.txt")
	cards := [][]int{}
	winningNumbers := [][]int{}
	for indx, line := range input {
		if line == "" {
			continue
		}
		strippedLine := stripGameNumber(line)
		splitLine := strings.Split(strippedLine, "|")
		if len(splitLine) != 2 {
			fmt.Println("line {} was bad", indx)
			return
		}
		winningNumbers = append(winningNumbers, parseNumbers(splitLine[0]))
		cards = append(cards, parseNumbers(splitLine[1]))
	}

	fmt.Println(partOne(winningNumbers, cards))
}
