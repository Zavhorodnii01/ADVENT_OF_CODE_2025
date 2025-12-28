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

	sum := 0
	for scanner.Scan() {
		line := scanner.Text()

		numbers := strings.Split(line, "")

		var allNeededNumbers []int
		max := 0
		maxPosition := -1

		for j := 0; j < 12; j++ {
			for i := maxPosition + 1; i < len(numbers)-(12-len(allNeededNumbers))+1; i++ {
				val, _ := strconv.Atoi(numbers[i])
				if max < val {
					max = val
					maxPosition = i
				}

			}

			allNeededNumbers = append(allNeededNumbers, max)
			max = 0
		}

		sumC := ""
		for _, num := range allNeededNumbers {
			sumC += strconv.Itoa(num)
		}

		sumN, _ := strconv.Atoi(sumC)

		sum += sumN

	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Println(sum)
}
