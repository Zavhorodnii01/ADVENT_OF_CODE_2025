////go:build ignore

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

	var grid [][]rune
	for scanner.Scan() {
		grid = append(grid, []rune(scanner.Text()))
	}

	rows := len(grid)
	cols := len(grid[0])

	sum := 0
	lastEmptyCol := -1

	for c := 0; c <= cols; c++ {
		isEmpty := true

		for r := 0; r < rows; r++ {
			if c < len(grid[r]) && grid[r][c] != ' ' {
				isEmpty = false
				break
			}
		}

		if isEmpty || c == cols {
			if lastEmptyCol+1 >= c {
				lastEmptyCol = c
			}

			opCol := -1

			for k := lastEmptyCol + 1; k < c; k++ {
				if k < len(grid[rows-1]) {
					if grid[rows-1][k] == '+' || grid[rows-1][k] == '*' {
						opCol = k
						break
					}
				}
			}

			if opCol == -1 {
				lastEmptyCol = c
			}

			var numbers []int

			for k := lastEmptyCol + 1; k < c; k++ {
				num := ""
				for r := 0; r < rows-1; r++ {
					if k < len(grid[r]) && grid[r][k] != ' ' {
						num += string(grid[r][k])
					}
				}
				if num != "" {
					n, _ := strconv.Atoi(num)
					numbers = append(numbers, n)
				}
			}

			cur := 0
			if grid[rows-1][opCol] == '+' {
				for _, n := range numbers {
					cur += n
				}
			} else {
				cur = 1
				for _, n := range numbers {
					cur *= n
				}
			}

			sum += cur
			lastEmptyCol = c
		}
	}

	fmt.Println(sum)
}
