package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func reverse(str string) (result string) {
	for _, v := range str {
		result = string(v) + result
	}
	return
}

func getFirstOrNil(a []string) string {
	if len(a) > 0 {
		return a[0]
	} else {
		return ""
	}
}

func indexOfNumber(number string, line string) int {
	numberAsRune := []rune(number)[0]
	for k, v := range []rune(line) {
		if numberAsRune == v {
			return k
		}
	}
	return -1
}

func indexOfWord(s, sep string, n int) int {
	idx := strings.Index(s[n:], sep)
	if idx > -1 {
		idx += n
	}
	return idx
}

func convertWordToNumAsString(word string) string {
	numbers := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten"}
	for k, v := range numbers {
		if word == v {
			return strconv.Itoa(k + 1)
		}
	}
	return ""
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var totalPartOne = 0
	var totalPartTwo = 0

parseCalibration:
	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) == "" {
			break parseCalibration
		} else {
			rePart1 := regexp.MustCompile("[0-9]")
			rePart2 := regexp.MustCompile("(one|two|three|four|five|six|seven|eight|nine)")
			rePart2Reversed := regexp.MustCompile("(enin|thgie|neves|xis|evif|ruof|eerht|owt|eno)")

			// Simple regex to get first number
			firstNumberAsDigit := rePart1.FindStringSubmatch(line)
			lastNumberAsDigit := rePart1.FindStringSubmatch(reverse(line))

			// Have to reverse the line and regex pattern to get the last number as a word
			firstNumberAsWord := getFirstOrNil(rePart2.FindStringSubmatch(line))
			lastNumberAsWord := reverse(getFirstOrNil(rePart2Reversed.FindStringSubmatch(reverse(line))))

			var firstNumberAsDigitIdx = -1
			var lastNumberAsDigitIdx = -1
			var firstNumberAsWordIdx = -1
			var lastNumberAsWordIdx = -1
			var firstNumberAsString = ""
			var lastNumberAsString = ""

			// Figure out the indices for the numbers as a singular digit and the numbers as a word
			if len(firstNumberAsDigit) > 0 {
				firstNumberAsDigitIdx = indexOfNumber(firstNumberAsDigit[0], line)
			}
			if len(lastNumberAsDigit) > 0 {
				lastNumberAsDigitIdx = len(line) - indexOfNumber(lastNumberAsDigit[0], reverse(line)) - 1
			}
			if len(firstNumberAsWord) > 0 {
				firstNumberAsWordIdx = indexOfWord(line, firstNumberAsWord, 0)
			}
			if len(lastNumberAsWord) > 0 {
				lastNumberAsWordIdx = len(line) - indexOfWord(reverse(line), reverse(lastNumberAsWord), 0) - len(lastNumberAsWord)
			}

			// Decide if the digit or the word is first and handle when there isn't a word or a number
			if firstNumberAsDigitIdx == -1 {
				firstNumberAsString = convertWordToNumAsString(firstNumberAsWord)
			} else if firstNumberAsWordIdx == -1 {
				firstNumberAsString = firstNumberAsDigit[0]
			} else if firstNumberAsWordIdx < firstNumberAsDigitIdx {
				firstNumberAsString = convertWordToNumAsString(firstNumberAsWord)
			} else {
				firstNumberAsString = firstNumberAsDigit[0]
			}

			// Decide if the digit or the word is last and handle when there isn't a word or a number
			if lastNumberAsDigitIdx == -1 {
				lastNumberAsString = convertWordToNumAsString(lastNumberAsWord)
			} else if lastNumberAsWordIdx == -1 {
				lastNumberAsString = lastNumberAsDigit[0]
			} else if lastNumberAsWordIdx > lastNumberAsDigitIdx {
				lastNumberAsString = convertWordToNumAsString(lastNumberAsWord)
			} else {
				lastNumberAsString = lastNumberAsDigit[0]
			}

			// Create the calibration values for both parts
			numberAsIntP1, err := strconv.Atoi(firstNumberAsDigit[0] + lastNumberAsDigit[0])
			numberAsIntP2, err := strconv.Atoi(firstNumberAsString + lastNumberAsString)
			if err != nil {
				panic(err)
			}

			totalPartOne += numberAsIntP1
			totalPartTwo += numberAsIntP2
		}
	}
	log.Println(fmt.Sprintf("%s%v", "part1 totalPartOne: ", totalPartOne))
	log.Println(fmt.Sprintf("%s%v", "totalPartTwo: ", totalPartTwo))

}
