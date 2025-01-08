package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strings"
)

func main() {
	day1()
}

func day1() {
	file, err := os.Open("day1.txt")

	if err != nil {
		panic(err)
	}

	defer file.Close()

	var first []int
	var second []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		numbers := strings.Fields(line)
		if len(numbers) >= 2 {
			var tmpfirst, tmpsecond int
			fmt.Sscanf(numbers[0], "%d", &tmpfirst)
			fmt.Sscanf(numbers[1], "%d", &tmpsecond)
			first = append(first, tmpfirst)
			second = append(second, tmpsecond)
		}
	}

	final := 0

	sort.Ints(first)
	sort.Ints(second)

	for i := 0; i < len(first) && i < len(second); i++ {
		final += int(math.Abs(float64(first[i] - second[i])))
	}

	fmt.Println(final)
}
