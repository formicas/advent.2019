package day7

import (
	"fmt"
	"strconv"

	"advent.2019/intcode"
)

//DoTheThing that day 7 is supposed to do
func DoTheThing() {
	program := []int{3, 8, 1001, 8, 10, 8, 105, 1, 0, 0, 21, 38, 63, 88, 97, 118, 199, 280, 361, 442, 99999, 3, 9, 1002, 9, 3, 9, 101, 2, 9, 9, 1002, 9, 4, 9, 4, 9, 99, 3, 9, 101, 3, 9, 9, 102, 5, 9, 9, 101, 3, 9, 9, 1002, 9, 3, 9, 101, 3, 9, 9, 4, 9, 99, 3, 9, 1002, 9, 2, 9, 1001, 9, 3, 9, 102, 3, 9, 9, 101, 2, 9, 9, 1002, 9, 4, 9, 4, 9, 99, 3, 9, 102, 2, 9, 9, 4, 9, 99, 3, 9, 102, 4, 9, 9, 101, 5, 9, 9, 102, 2, 9, 9, 101, 5, 9, 9, 4, 9, 99, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 99, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 99, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 99, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 99, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 99}
	values := []int{5, 6, 7, 8, 9}
	maxValue := 0
	heapPermutation(program, values, 5, &maxValue, true)
}

func heapPermutation(program []int, values []int, size int, maxValue *int, feedback bool) {
	if size == 1 {
		value := doTheThing(program, values, feedback)
		if value > *maxValue {
			fmt.Printf("%v outputs %d\n", values, value)
			*maxValue = value
		}
	} else {
		for i := 0; i < size; i++ {
			heapPermutation(program, values, size-1, maxValue, feedback)

			if size%2 == 1 {
				temp := values[0]
				values[0] = values[size-1]
				values[size-1] = temp
			} else {
				temp := values[i]
				values[i] = values[size-1]
				values[size-1] = temp
			}
		}
	}
}

func doTheThing(program []int, inputs []int, feedback bool) int {
	programs := make([]intcode.Program, 5)

	for i := 0; i < 5; i++ {
		code := make([]int, len(program))
		copy(code, program[:])
		var input chan int
		if i == 0 {
			input = make(chan int, 2)
		} else {
			input = programs[i-1].Output
		}
		output := make(chan int, 2)
		programs[i] = intcode.Program{Code: code, Input: input, Output: output, Name: strconv.Itoa(i)}
	}

	//pass in the initial values
	for i := range programs {
		programs[i].Input <- inputs[i]
	}

	if feedback {
		programs[4].Output = programs[0].Input
		programs[4].Halted = make(chan int)
	}

	for _, p := range programs {
		go p.Run()
	}

	programs[0].Input <- 0
	if !feedback {
		result := <-programs[4].Output
		return result
	}
	//fmt.Println("Waiting for 4 to halt")
	<-programs[4].Halted
	//fmt.Println("It has halted")
	result := <-programs[4].Output
	return result

}
