package main

import (
	"fmt"
)

func main() {

	input := []int{3, 8, 9, 1, 2, 5, 4, 6, 7}
	max := -1
	for i := 0; i < len(input); i++ {
		if max < input[i] {
			max = input[i]
		}
	}

	for turns := 0; turns < 10; turns++ {

		current := input[turns%len(input)]
		pickedUp := []int{}
		remaining := []int{}
		for y := 1; y < 4; y++ {
			pickedUp = append(pickedUp, input[(turns+y)%len(input)])
		}
		for y := 4; y < len(input); y++ {
			remaining = append(remaining, input[(turns+y)%len(input)])
		}
		remainingWithCurrent := append(remaining, current)

		counter := current - 1
		destination := -1

		for destination == -1 {
			if counter == 0 {
				counter = max
			}
			for _, value := range remainingWithCurrent {
				if value == counter {
					destination = value
				}
			}
			counter--
		}

		newCircle := []int{}
		newCircle = append(newCircle, current)
		newCircle = append(newCircle, destination)
		newCircle = append(newCircle, pickedUp...)
		newCircle = append(newCircle, remaining...)

		fmt.Println(newCircle)

	}

}
