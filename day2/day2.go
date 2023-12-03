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

func isGamePossible(line string) bool {
	maxRed, maxGreen, maxBlue := 12, 13, 14
	return isColorBelowMax("red", line, maxRed) && isColorBelowMax("blue", line, maxBlue) && isColorBelowMax("green", line, maxGreen)
}

func isColorBelowMax(color string, line string, max int) bool {
	colorPattern, _ := regexp.Compile(`(\d+) ` + color)
	number, _ := regexp.Compile(`\d+`)
	cubesWithColor := colorPattern.FindAllString(line, -1)
	for _, cubeWithColor := range cubesWithColor {
		n, _ := strconv.Atoi(number.FindString(cubeWithColor))
		if n > max {
			return false
		}
	}
	return true
}

func getMaxOfColor(color string, line string) int {
	var maxi = -1
	colorPattern, _ := regexp.Compile(`(\d+) ` + color)
	number, _ := regexp.Compile(`\d+`)
	cubesWithColor := colorPattern.FindAllString(line, -1)
	for _, cubeWithColor := range cubesWithColor {
		n, _ := strconv.Atoi(number.FindString(cubeWithColor))
		if maxi == -1 || n > maxi {
			maxi = n
		}
	}
	return maxi
}

func main() {
	file, err := os.Open("./day2/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var p1, p2 = 0, 0

parseCubes:
	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) == "" {
			break parseCubes
		} else {
			gameId, _ := strconv.Atoi(strings.Split(strings.Split(line, ":")[0], " ")[1])
			if isGamePossible(line) {
				p1 += gameId
			}
			power := getMaxOfColor("red", line) * getMaxOfColor("green", line) * getMaxOfColor("blue", line)
			p2 += power
		}
	}
	fmt.Printf("%s%v\n", "Sum of IDs: ", p1)
	fmt.Printf("%s%v\n", "Sum of the min powers is: ", p2)

}
