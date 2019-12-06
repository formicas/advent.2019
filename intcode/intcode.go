package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	program := []int{3, 225, 1, 225, 6, 6, 1100, 1, 238, 225, 104, 0, 1102, 91, 92, 225, 1102, 85, 13, 225, 1, 47, 17, 224, 101, -176, 224, 224, 4, 224, 1002, 223, 8, 223, 1001, 224, 7, 224, 1, 223, 224, 223, 1102, 79, 43, 225, 1102, 91, 79, 225, 1101, 94, 61, 225, 1002, 99, 42, 224, 1001, 224, -1890, 224, 4, 224, 1002, 223, 8, 223, 1001, 224, 6, 224, 1, 224, 223, 223, 102, 77, 52, 224, 1001, 224, -4697, 224, 4, 224, 102, 8, 223, 223, 1001, 224, 7, 224, 1, 224, 223, 223, 1101, 45, 47, 225, 1001, 43, 93, 224, 1001, 224, -172, 224, 4, 224, 102, 8, 223, 223, 1001, 224, 1, 224, 1, 224, 223, 223, 1102, 53, 88, 225, 1101, 64, 75, 225, 2, 14, 129, 224, 101, -5888, 224, 224, 4, 224, 102, 8, 223, 223, 101, 6, 224, 224, 1, 223, 224, 223, 101, 60, 126, 224, 101, -148, 224, 224, 4, 224, 1002, 223, 8, 223, 1001, 224, 2, 224, 1, 224, 223, 223, 1102, 82, 56, 224, 1001, 224, -4592, 224, 4, 224, 1002, 223, 8, 223, 101, 4, 224, 224, 1, 224, 223, 223, 1101, 22, 82, 224, 1001, 224, -104, 224, 4, 224, 1002, 223, 8, 223, 101, 4, 224, 224, 1, 223, 224, 223, 4, 223, 99, 0, 0, 0, 677, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1105, 0, 99999, 1105, 227, 247, 1105, 1, 99999, 1005, 227, 99999, 1005, 0, 256, 1105, 1, 99999, 1106, 227, 99999, 1106, 0, 265, 1105, 1, 99999, 1006, 0, 99999, 1006, 227, 274, 1105, 1, 99999, 1105, 1, 280, 1105, 1, 99999, 1, 225, 225, 225, 1101, 294, 0, 0, 105, 1, 0, 1105, 1, 99999, 1106, 0, 300, 1105, 1, 99999, 1, 225, 225, 225, 1101, 314, 0, 0, 106, 0, 0, 1105, 1, 99999, 8, 226, 677, 224, 102, 2, 223, 223, 1005, 224, 329, 1001, 223, 1, 223, 1007, 226, 226, 224, 1002, 223, 2, 223, 1006, 224, 344, 101, 1, 223, 223, 108, 226, 226, 224, 1002, 223, 2, 223, 1006, 224, 359, 1001, 223, 1, 223, 107, 226, 677, 224, 102, 2, 223, 223, 1006, 224, 374, 101, 1, 223, 223, 8, 677, 677, 224, 102, 2, 223, 223, 1006, 224, 389, 1001, 223, 1, 223, 1008, 226, 677, 224, 1002, 223, 2, 223, 1006, 224, 404, 101, 1, 223, 223, 7, 677, 677, 224, 1002, 223, 2, 223, 1005, 224, 419, 101, 1, 223, 223, 1108, 226, 677, 224, 1002, 223, 2, 223, 1005, 224, 434, 101, 1, 223, 223, 1108, 226, 226, 224, 102, 2, 223, 223, 1005, 224, 449, 1001, 223, 1, 223, 107, 226, 226, 224, 102, 2, 223, 223, 1005, 224, 464, 101, 1, 223, 223, 1007, 677, 677, 224, 102, 2, 223, 223, 1006, 224, 479, 101, 1, 223, 223, 1007, 226, 677, 224, 102, 2, 223, 223, 1005, 224, 494, 1001, 223, 1, 223, 1008, 226, 226, 224, 1002, 223, 2, 223, 1005, 224, 509, 1001, 223, 1, 223, 1108, 677, 226, 224, 1002, 223, 2, 223, 1006, 224, 524, 1001, 223, 1, 223, 108, 677, 677, 224, 1002, 223, 2, 223, 1005, 224, 539, 101, 1, 223, 223, 108, 226, 677, 224, 1002, 223, 2, 223, 1005, 224, 554, 101, 1, 223, 223, 1008, 677, 677, 224, 1002, 223, 2, 223, 1006, 224, 569, 1001, 223, 1, 223, 1107, 677, 677, 224, 102, 2, 223, 223, 1005, 224, 584, 1001, 223, 1, 223, 7, 677, 226, 224, 102, 2, 223, 223, 1005, 224, 599, 1001, 223, 1, 223, 8, 677, 226, 224, 1002, 223, 2, 223, 1005, 224, 614, 1001, 223, 1, 223, 7, 226, 677, 224, 1002, 223, 2, 223, 1006, 224, 629, 101, 1, 223, 223, 1107, 677, 226, 224, 1002, 223, 2, 223, 1005, 224, 644, 1001, 223, 1, 223, 1107, 226, 677, 224, 102, 2, 223, 223, 1006, 224, 659, 1001, 223, 1, 223, 107, 677, 677, 224, 1002, 223, 2, 223, 1005, 224, 674, 101, 1, 223, 223, 4, 223, 99, 226}
	runIntCode(program)
}

func runIntCode(intcode []int) {
	pointer := 0
	for {
		instruction := intcode[pointer]

		opCode := instruction % 100
		instruction = instruction / 100

		p1Mode := instruction % 10
		instruction = instruction / 10

		p2Mode := instruction % 10
		instruction = instruction / 10

		p3Mode := instruction % 10

		switch opCode {
		case 1:
			slice := intcode[pointer : pointer+4]
			p1 := getValue(intcode, p1Mode, slice[1])
			p2 := getValue(intcode, p2Mode, slice[2])

			intcode[slice[3]] = p1 + p2

			pointer = pointer + 4
		case 2:
			slice := intcode[pointer : pointer+4]
			p1 := getValue(intcode, p1Mode, slice[1])
			p2 := getValue(intcode, p2Mode, slice[2])

			intcode[slice[3]] = p1 * p2

			pointer = pointer + 4
		case 3:
			slice := intcode[pointer : pointer+2]
			dest := slice[1]
			reader := bufio.NewReader(os.Stdin)
			fmt.Println("Enter text:")
			input, _ := reader.ReadString('\n')
			asInt, err := strconv.Atoi(strings.TrimSpace(input))
			if err != nil {
				fmt.Println(err)
				os.Exit(2)
			}
			intcode[dest] = asInt

			pointer = pointer + 2
		case 4:
			slice := intcode[pointer : pointer+2]
			value := getValue(intcode, p1Mode, slice[1])
			fmt.Printf("Output: %d\n", value)

			pointer = pointer + 2
		case 5:
			slice := intcode[pointer : pointer+3]
			value := getValue(intcode, p1Mode, slice[1])
			newPointer := getValue(intcode, p2Mode, slice[2])

			if value != 0 {
				pointer = newPointer
			} else {
				pointer = pointer + 3
			}
		case 6:
			slice := intcode[pointer : pointer+3]
			value := getValue(intcode, p1Mode, slice[1])
			newPointer := getValue(intcode, p2Mode, slice[2])
			if value == 0 {
				pointer = newPointer
			} else {
				pointer = pointer + 3
			}
		case 7:
			slice := intcode[pointer : pointer+4]
			value := 0
			if getValue(intcode, p1Mode, slice[1]) < getValue(intcode, p2Mode, slice[2]) {
				value = 1
			}

			intcode[slice[3]] = value

			pointer = pointer + 4
		case 8:
			slice := intcode[pointer : pointer+4]
			value := 0
			if getValue(intcode, p1Mode, slice[1]) == getValue(intcode, p2Mode, slice[2]) {
				value = 1
			}

			intcode[slice[3]] = value

			pointer = pointer + 4
		case 99:
			return
		}
	}
}

func getValue(intcode []int, mode int, value int) int {
	switch mode {
	case 1:
		return value
	default:
		return intcode[value]
	}
}
