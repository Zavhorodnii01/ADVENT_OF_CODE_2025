//go:build ignore

package main

import (
	"bufio"
	"fmt"
	"os"
)

type Pos struct {
	i, j int
}

func remove(grid [][]rune) int {
	var toRemove []Pos

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {

			if grid[i][j] == '@' {
				count := 0

				if i != 0 && grid[i-1][j] == '@' {
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
					toRemove = append(toRemove, Pos{i, j})
				}
			}
		}
	}

	for _, p := range toRemove {
		grid[p.i][p.j] = '.'
	}

	return len(toRemove)
}

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

	for {
		removed := remove(grid)
		if removed == 0 {
			break
		}
		sum += removed
	}

	fmt.Println(sum)
}
