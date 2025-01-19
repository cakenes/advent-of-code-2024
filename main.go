package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"sort"
	"strings"
)

func main() {
	day1()
	day2()
	day3()
	day4()
}

func read(filename string) []string {
	file, err := os.Open(filename)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

func day1() {
	lines := read("day1.txt")

	var first []int
	var second []int

	final := 0

	for _, line := range lines {
		numbers := strings.Fields(line)
		if len(numbers) >= 2 {
			var tmpfirst, tmpsecond int
			fmt.Sscanf(numbers[0], "%d", &tmpfirst)
			fmt.Sscanf(numbers[1], "%d", &tmpsecond)
			first = append(first, tmpfirst)
			second = append(second, tmpsecond)
		}
	}

	sort.Ints(first)
	sort.Ints(second)

	for i := 0; i < len(first) && i < len(second); i++ {
		final += int(math.Abs(float64(first[i] - second[i])))
	}

	fmt.Println("Day 1:", final)
}

func day2() {
	lines := read("day2.txt")

	var lineNumbers [][]int

	final := 0

	for _, line := range lines {
		var numbers []int
		fields := strings.Fields(line)
		for _, field := range fields {
			var tmp int
			fmt.Sscanf(field, "%d", &tmp)
			numbers = append(numbers, tmp)
		}
		lineNumbers = append(lineNumbers, numbers)
	}

	for _, numbers := range lineNumbers {
		increasing := false
		decreasing := false

		for i := 1; i < len(numbers); i++ {
			diff := numbers[i] - numbers[i-1]
			if diff != 1 && diff != 2 && diff != 3 {
				increasing = true
			}
			if diff != -1 && diff != -2 && diff != -3 {
				decreasing = true
			}
		}

		if (increasing || decreasing) && !(increasing && decreasing) {
			final++
		}
	}

	fmt.Println("Day 2:", final)
}

func day3() {
	lines := read("day3.txt")

	multiPattern := `mul\(\d+,\d+\)`
	multiRe := regexp.MustCompile(multiPattern)
	numberPattern := `\d+`
	numberRe := regexp.MustCompile(numberPattern)

	final := 0

	for _, line := range lines {
		matches := multiRe.FindAllString(line, -1)
		for _, match := range matches {
			numbers := numberRe.FindAllString(match, -1)
			var first, second int
			fmt.Sscanf(numbers[0], "%d", &first)
			fmt.Sscanf(numbers[1], "%d", &second)
			final += first * second
		}
	}

	fmt.Println("Day 3:", final)
}

func day4() {
	lines := read("day4.txt")

	xmasPattern := `XMAS`
	xmasRe := regexp.MustCompile(xmasPattern)
	samxPattern := `SAMX`
	samxRe := regexp.MustCompile(samxPattern)

	final := 0

	// Horizontal
	for _, row := range lines {
		for i := 0; i <= len(row)-4; i++ { // Check all possible substrings to handle overlapping
			substr := row[i : i+4]
			if xmasRe.MatchString(substr) || samxRe.MatchString(substr) {
				final++
			}
		}
	}

	// Vertical
	for col := 0; col < len(lines[0]); col++ {
		for row := 0; row <= len(lines)-4; row++ { // Check all possible substrings to handle overlapping
			substr := ""
			for k := 0; k < 4; k++ {
				substr += string(lines[row+k][col])
			}
			if xmasRe.MatchString(substr) || samxRe.MatchString(substr) {
				final++
			}
		}
	}

	// Diagonal
	for i := 0; i <= len(lines)-4; i++ {
		for j := 0; j <= len(lines[0])-4; j++ { // Check all possible substrings to handle overlapping
			substr := ""
			for k := 0; k < 4; k++ {
				substr += string(lines[i+k][j+k])
			}
			if xmasRe.MatchString(substr) || samxRe.MatchString(substr) {
				final++
			}

			substr = ""
			for k := 0; k < 4; k++ {
				substr += string(lines[i+k][j+3-k])
			}
			if xmasRe.MatchString(substr) || samxRe.MatchString(substr) {
				final++
			}
		}
	}

	fmt.Println("Day 4:", final)
}
