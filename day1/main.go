package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func calc(num int) int {
	num = num / 3
	return num - 2
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var Total = 0
	var Total2 = 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())

		if err != nil {
			break
		}

		Total += calc(num)
		fuelLeft := num

		for {
			fuelLeft = calc(fuelLeft) // convert to string to use same method
			if fuelLeft < 0 {
				break
			}
			Total2 += fuelLeft
		}
	}

	fmt.Printf("Part 1: %d\n", Total) // Result of Part 1
	fmt.Printf("Part 2: %d", Total2)  // Result of Part2

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
