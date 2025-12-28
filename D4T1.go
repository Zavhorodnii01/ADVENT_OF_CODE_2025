//go:build ignore

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var grid [][]rune

	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, []rune(line))
	}
	sum := 0

	for i := 0; i < len(grid); i++ {
		count := 0
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '@' {
				if i != 0 && (grid[i-1][j] == '@') {
					count++
				}
				if i != 0 && j != 0 && grid[i-1][j-1] == '@' {
					count++
				}
				if i != 0 && j != len(grid[i])-1 && grid[i-1][j+1] == '@' {
					count++
				}
				if i != len(grid)-1 && grid[i+1][j] == '@' {
					count++
				}
				if i != len(grid)-1 && j != 0 && grid[i+1][j-1] == '@' {
					count++
				}
				if i != len(grid)-1 && j != len(grid[i])-1 && grid[i+1][j+1] == '@' {
					count++
				}
				if j != len(grid[i])-1 && grid[i][j+1] == '@' {
					count++
				}
				if j != 0 && grid[i][j-1] == '@' {
					count++
				}

				if count < 4 {
					sum++
				}
				count = 0
			}

		}
	}
	fmt.Println(sum)
}
