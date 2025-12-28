//go:build ignore

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isRepeatable(n string) bool {
	l := len(n)

	for splitter := 1; splitter <= l/2; splitter++ {
		if l%splitter != 0 {
			continue
		}

		var allSplits []string
		start := 0
		for start < l {
			end := start + splitter
			allSplits = append(allSplits, n[start:end])
			start = end
		}

		allSame := true
		for i := 1; i < len(allSplits); i++ {
			if allSplits[i] != allSplits[0] {
				allSame = false
				break
			}
		}

		if allSame {
			return true
		}
	}

	return false
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
			if isRepeatable(strconv.Itoa(n)) {
				sum += n
			}
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Println(sum)
}
