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

		max1 := 0
		max2 := 0
		max1Position := 0
		max2Position := 0

		for i := 0; i < len(numbers); i++ {
			val, _ := strconv.Atoi(numbers[i])
			if max1 < val {
				max1 = val
				max1Position = i
			}
		}

		for i := 0; i < len(numbers); i++ {
			val, _ := strconv.Atoi(numbers[i])
			if max2 <= val && i != max1Position {
				max2 = val
				max2Position = i
			}
		}

		if max2 != 9 {
			max2 := 0
			for i := max1Position + 1; i < len(numbers); i++ {
				val, _ := strconv.Atoi(numbers[i])
				if max2 < val {
					max2 = val
					max2Position = i
				}
			}
		}

		largest := ""
		if max1Position < max2Position {
			largest = numbers[max1Position] + numbers[max2Position]
		} else {
			largest = numbers[max2Position] + numbers[max1Position]
		}

		largestInt, _ := strconv.Atoi(largest)
		sum += largestInt
		fmt.Println(largestInt)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Println(sum)
}
