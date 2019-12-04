package main

import (
    "encoding/csv"
    "fmt"
    "os"
    "strconv"
)

func getInput() ([]string, []string) {
	file, err := os.Open("input2.txt")
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
    fmt.Println(firstC)
    fmt.Println(secondC)
    fmt.Println(intersectionPoints(firstC, secondC))
}

type point struct {
    x int
    y int
}

func coordinates(list []string) []point {
    var coordinates []point
    current := point{x: 0, y: 0}
    for _, i := range list {
        direction := string(i[0])
        distance, _ := strconv.Atoi(i[1:len(i)])
        // TODO: Need to add intermediate points as well for intersection logic
        // For every movement a coordinate should be created
        if direction == "R" {
            current.x += distance
        } else if direction == "L" {
            current.x -= distance
        } else if direction == "U" {
            current.y += distance
        } else if direction == "D" {
            current.y -= distance
        } else {
            panic("hit else??")
        }

        coordinates = append(coordinates, current)
    }

    return coordinates
}
