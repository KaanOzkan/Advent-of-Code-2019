package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

func getInput() ([]string, []string) {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	r := csv.NewReader(file)
	var list1 []string
	var list2 []string
	record, err := r.Read()

	if err != nil {
		panic(err)
	}

	for _, i := range record {
		list1 = append(list1, i)
	}

	record, err = r.Read()

	for _, i := range record {
		list2 = append(list2, i)
	}

	return list1, list2
}

func intersectionPoints(first []point, second []point) []point {
	var intersections []point
	for _, i := range first {
		for _, j := range second {
			if i.x == j.x && i.y == j.y {
				intersections = append(intersections, i)
			}
		}
	}
	return intersections
}

func main() {
	first, second := getInput()
	firstC := coordinates(first)
	secondC := coordinates(second)
	fmt.Println("Calculate intersections")
	fmt.Println(intersectionPoints(firstC, secondC))
}

type point struct {
	x int
	y int
}

func coordinates(list []string) []point {
	var coordinates []point
	coordinates = append(coordinates, point{x: 0, y: 0})
	for _, item := range list {
		direction := string(item[0])
		distance, _ := strconv.Atoi(item[1:len(item)])
		for i := 0; i < distance; i++ {
			prev := coordinates[len(coordinates)-1]
			current := prev
			if direction == "R" {
				current.x += 1
			} else if direction == "L" {
				current.x -= 1
			} else if direction == "U" {
				current.y += 1
			} else if direction == "D" {
				current.y -= 1
			} else {
				panic("invalid direction")
			}

			coordinates = append(coordinates, current)
		}
	}

	return coordinates
}
