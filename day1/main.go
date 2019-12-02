package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

var TotalFuel = 0

func calc(num string) int {
	i, err := strconv.Atoi(num)
	if err != nil {
		return 0
		log.Fatal(err)
	}
	i = i / 3
	return i - 2
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		TotalFuel += calc(scanner.Text())
	}

	fmt.Println(TotalFuel)

	// TotalFuel = 100756 // TODO: Remove
    // original := TotalFuel

	// fuelLeft := TotalFuel
	// for fuelLeft > 0 {
	// 	fuelLeft = calc(strconv.Itoa(fuelLeft)) // convert to string to use same method
	// 	// fmt.Println(fuelLeft)
	// 	// fmt.Println(TotalFuel)
	// 	if fuelLeft > 0 {
	// 		TotalFuel += fuelLeft
	// 	}
	// }

	// fmt.Println(TotalFuel) // Result for part 2

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
