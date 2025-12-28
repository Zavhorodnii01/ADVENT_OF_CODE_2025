//go:build ignore

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	dial := 50
	countAtZero := 0

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) < 2 {
			continue
		}

		dir := line[0] // 'L' or 'R'
		dist, err := strconv.Atoi(line[1:])
		if err != nil {
			log.Fatal(err)
		}

		if dir == 'L' {
			dial = (dial - dist) % 100
			if dial < 0 {
				dial += 100
			}
		} else if dir == 'R' {
			dial = (dial + dist) % 100
		} else {
			log.Fatalf("Unknown direction: %c", dir)
		}

		if dial == 0 {
			countAtZero++
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("The password (number of times dial points at 0) is: %d\n", countAtZero)
}
