//go:build ignore

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	position := 50
	password := 0
	maxNumber := 100

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) < 2 {
			continue
		}
		dir := line[0]
		steps, _ := strconv.Atoi(line[1:])

		if dir == 'R' {
			for i := 0; i < steps; i++ {
				position = (position + 1) % maxNumber
				if position == 0 {
					password++
				}
			}
		} else if dir == 'L' {
			for i := 0; i < steps; i++ {
				position = (position - 1 + maxNumber) % maxNumber
				if position == 0 {
					password++
				}
			}
		}
	}

	fmt.Println("Password (method 0x434C49434B):", password)
}
