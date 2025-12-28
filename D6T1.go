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

	var line1, line2, line3, line4, line5 []string
	i := 0
	for scanner.Scan() {
		line := scanner.Text()
		if i == 0 {
			line1 = strings.Fields(line)
		}
		if i == 1 {
			line2 = strings.Fields(line)
		}
		if i == 2 {
			line3 = strings.Fields(line)
		}
		if i == 3 {
			line4 = strings.Fields(line)
		}
		if i == 4 {
			line5 = strings.Fields(line)
		}
		i++
	}

	sum := 0

	for j := 0; j < len(line1); j++ {
		a, _ := strconv.Atoi(line1[j])
		b, _ := strconv.Atoi(line2[j])
		c, _ := strconv.Atoi(line3[j])
		d, _ := strconv.Atoi(line4[j])

		if line5[j] == "+" {
			sum += a + b + c + d
		} else {
			sum += a * b * c * d
		}
	}

	fmt.Println(sum)
}
