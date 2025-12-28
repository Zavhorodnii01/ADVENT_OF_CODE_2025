//go:build ignore

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var ranges []string
	for scanner.Scan() {
		line := scanner.Text()
		ranges = append(ranges, line)
	}

	file2, err2 := os.Open("input2.txt")
	if err2 != nil {
		panic(err2)
	}
	defer file2.Close()

	scanner2 := bufio.NewScanner(file2)

	var numbers []int
	for scanner2.Scan() {
		line := scanner2.Text()
		n, _ := strconv.Atoi(line)
		numbers = append(numbers, n)
	}

	count := 0

	for _, number := range numbers {
		isInRange := false

		for _, range1 := range ranges {
			index := strings.Index(range1, "-")

			start, _ := strconv.Atoi(range1[:index])
			end, _ := strconv.Atoi(range1[index+1:])

			if number >= start && number <= end {
				isInRange = true
				break
			}
		}

		if isInRange {
			count++
		}
	}

	fmt.Println(count)
}
