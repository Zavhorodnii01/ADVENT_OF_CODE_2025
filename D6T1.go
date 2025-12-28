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

	freshSet := make(map[int]bool)

	for _, r := range ranges {
		index := strings.Index(r, "-")

		start, _ := strconv.Atoi(r[:index])
		end, _ := strconv.Atoi(r[index+1:])

		for i := start; i <= end; i++ {
			freshSet[i] = true
		}
	}

	fmt.Println(len(freshSet))
}
