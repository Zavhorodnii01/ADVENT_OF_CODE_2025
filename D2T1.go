//go:build ignore

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isDoubleSequence(n int) bool {
	s := strconv.Itoa(n)
	l := len(s)
	if l%2 != 0 {
		return false
	}
	half := l / 2
	return s[:half] == s[half:]
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var allRanges []string

	for scanner.Scan() {
		line := scanner.Text()
		elements := strings.Split(line, ",")
		allRanges = append(allRanges, elements...)
	}

	sum := 0

	for _, element := range allRanges {
		index := strings.Index(element, "-")
		if index == -1 {
			continue
		}

		curNum, _ := strconv.Atoi(element[:index])
		endNum, _ := strconv.Atoi(element[index+1:])

		for n := curNum; n <= endNum; n++ {
			if isDoubleSequence(n) {
				sum += n
			}
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Println(sum)
}
