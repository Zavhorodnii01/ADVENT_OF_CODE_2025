//go:build ignore

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Range struct {
	start int
	end   int
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var ranges []Range
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}

		index := strings.Index(line, "-")
		start, _ := strconv.Atoi(line[:index])
		end, _ := strconv.Atoi(line[index+1:])

		ranges = append(ranges, Range{start, end})
	}

	// sort ranges by start
	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].start < ranges[j].start
	})

	total := 0
	currentStart := ranges[0].start
	currentEnd := ranges[0].end

	for _, r := range ranges[1:] {
		if r.start <= currentEnd+1 {
			if r.end > currentEnd {
				currentEnd = r.end
			}
		} else {
			total += currentEnd - currentStart + 1
			currentStart = r.start
			currentEnd = r.end
		}
	}

	total += currentEnd - currentStart + 1

	fmt.Println(total)
}
