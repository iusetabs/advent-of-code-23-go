package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

func getSymbolIndices(line string) []int {
	symbolIndices := []int{}
	for pos, charAsRune := range line {
		if charAsRune == 46 {
			continue
		} else if charAsRune < 48 || charAsRune > 57 {
			symbolIndices = append(symbolIndices, pos)
		}
	}
	return symbolIndices
}

func symbolCloseToIndex(symbolIndices []int, startingIndx int, numberAsString string) bool {
	var i = startingIndx - 1
	for i < startingIndx+len(numberAsString)+1 {
		if slices.Contains(symbolIndices, i) {
			return true
		}
		i += 1
	}
	return false
}

func main() {
	file, err := os.Open("./day3/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	numberPattern := regexp.MustCompile("[0-9]+")
	var total = 0
	var previousLine = ""

parseLines:
	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) == "" {
			break parseLines
		} else {
			symbolIndices := getSymbolIndices(line)
			previousSymbolIndices := getSymbolIndices(previousLine)
			numbers := numberPattern.FindAllString(line, -1)
			numbersIndices := numberPattern.FindAllStringIndex(line, -1)

			previousNumbers := numberPattern.FindAllString(previousLine, -1)
			previousNumbersIndices := numberPattern.FindAllStringIndex(previousLine, -1)

			for i, numberAsString := range numbers {
				startingIndx := numbersIndices[i][0]
				if symbolCloseToIndex(symbolIndices, startingIndx, numberAsString) || symbolCloseToIndex(previousSymbolIndices, startingIndx, numberAsString) {
					number, _ := strconv.Atoi(numberAsString)
					println(number)
					total += number
				}
			}
			for i, previousNumberAsString := range previousNumbers {
				startingIndx := previousNumbersIndices[i][0]
				if symbolCloseToIndex(symbolIndices, startingIndx, previousNumberAsString) {
					number, _ := strconv.Atoi(previousNumberAsString)
					println(number)
					total += number
				}
			}

			previousLine = line
		}
	}
	fmt.Printf("%s %v\n", "p1 =", total)
}
