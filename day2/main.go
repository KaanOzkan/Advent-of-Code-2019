package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
)

func prepare(list []int, i int, j int) []int {
	list[1] = i
	list[2] = j
	return list
}

func main() {
	original_list := getInput()
	list := original_list
	fmt.Printf("Input: %d\n\n", list)
	list = prepare(list, 12, 2)
	list = calculate(list)
	fmt.Println(list[0]) // Result of Part 1

	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
            original_list := getInput() // Go doesn't create copy on assigment so reset slice
			list := prepare(original_list, i, j)
            fmt.Println(original_list[0])
            list = calculate(list)
            if list[0] == 19690720 {
                fmt.Println(i) // Result of part 2
                fmt.Println(j) // Result of part 2
                return
            }
		}
	}

    fmt.Println("Failed")
}

func calculate(list []int) []int {
	i := 0 // index/iterator
    counter := 0
	for {
        counter += 1
        if counter > 150 {
            return list // Prevent infinite loop while debugging
        }
		num := list[i]
		if num == 99 {
			break
		}

		// Addition
		if num == 1 {
			// fmt.Println("adding")
			left := list[i+1]
			right := list[i+2]
			sum := list[left] + list[right]
			storageLoc := list[i+3]
			list[storageLoc] = sum
			i = i + 4
			// fmt.Println(list)
		} else if num == 2 {
			// fmt.Println("multiplying")
			left := list[i+1]
			right := list[i+2]
			mult := list[left] * list[right]
			storageLoc := list[i+3]
			list[storageLoc] = mult
			i = i + 4
			// fmt.Println(list)
		}
	}
	return list
}

func getInput() []int {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	r := csv.NewReader(file)
	var list []int
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			panic(err)
		}

		for _, i := range record {
			j, _ := strconv.Atoi(i)
			list = append(list, j)
		}
	}

	return list
}
