package day19

import (
	"fmt"

	"advent.2019/intcode"
)

//DoTheThing that day 19 is supposed to do
func DoTheThing() {
	code := []int{109, 424, 203, 1, 21101, 11, 0, 0, 1105, 1, 282, 21102, 18, 1, 0, 1105, 1, 259, 2102, 1, 1, 221, 203, 1, 21102, 1, 31, 0, 1106, 0, 282, 21101, 0, 38, 0, 1106, 0, 259, 20102, 1, 23, 2, 21202, 1, 1, 3, 21101, 0, 1, 1, 21102, 57, 1, 0, 1105, 1, 303, 2101, 0, 1, 222, 20102, 1, 221, 3, 20101, 0, 221, 2, 21102, 259, 1, 1, 21101, 0, 80, 0, 1106, 0, 225, 21102, 135, 1, 2, 21101, 0, 91, 0, 1105, 1, 303, 2102, 1, 1, 223, 21001, 222, 0, 4, 21102, 259, 1, 3, 21102, 1, 225, 2, 21101, 0, 225, 1, 21101, 118, 0, 0, 1106, 0, 225, 20101, 0, 222, 3, 21101, 0, 12, 2, 21101, 0, 133, 0, 1106, 0, 303, 21202, 1, -1, 1, 22001, 223, 1, 1, 21102, 1, 148, 0, 1105, 1, 259, 1202, 1, 1, 223, 21002, 221, 1, 4, 20102, 1, 222, 3, 21101, 0, 17, 2, 1001, 132, -2, 224, 1002, 224, 2, 224, 1001, 224, 3, 224, 1002, 132, -1, 132, 1, 224, 132, 224, 21001, 224, 1, 1, 21102, 1, 195, 0, 105, 1, 109, 20207, 1, 223, 2, 21001, 23, 0, 1, 21101, 0, -1, 3, 21101, 214, 0, 0, 1105, 1, 303, 22101, 1, 1, 1, 204, 1, 99, 0, 0, 0, 0, 109, 5, 1202, -4, 1, 249, 21201, -3, 0, 1, 22102, 1, -2, 2, 22102, 1, -1, 3, 21102, 250, 1, 0, 1106, 0, 225, 21202, 1, 1, -4, 109, -5, 2106, 0, 0, 109, 3, 22107, 0, -2, -1, 21202, -1, 2, -1, 21201, -1, -1, -1, 22202, -1, -2, -2, 109, -3, 2105, 1, 0, 109, 3, 21207, -2, 0, -1, 1206, -1, 294, 104, 0, 99, 21201, -2, 0, -2, 109, -3, 2105, 1, 0, 109, 5, 22207, -3, -4, -1, 1206, -1, 346, 22201, -4, -3, -4, 21202, -3, -1, -1, 22201, -4, -1, 2, 21202, 2, -1, -1, 22201, -4, -1, 1, 22102, 1, -2, 3, 21101, 0, 343, 0, 1106, 0, 303, 1105, 1, 415, 22207, -2, -3, -1, 1206, -1, 387, 22201, -3, -2, -3, 21202, -2, -1, -1, 22201, -3, -1, 3, 21202, 3, -1, -1, 22201, -3, -1, 2, 22101, 0, -4, 1, 21102, 384, 1, 0, 1105, 1, 303, 1106, 0, 415, 21202, -4, -1, -4, 22201, -4, -3, -4, 22202, -3, -2, -2, 22202, -2, -4, -4, 22202, -3, -2, -3, 21202, -4, -1, -2, 22201, -3, -2, 1, 22102, 1, 1, -4, 109, -5, 2106, 0, 0}
	input := make(chan int)
	output := make(chan int)
	halted := make(chan int)
	p := intcode.Program{Code: code, Input: input, Output: output, Halted: halted}

	total := 0
	for y := 0; y < 50; y++ {
		for x := 0; x < 50; x++ {
			go p.Run()
			input <- x
			input <- y
			out := <-output
			total = total + out
			<-halted
		}
	}
	fmt.Printf("Total in 50x50: %d\n", total)

	width := 0
	y := 0
	dy := 100
	for width < 100 {
		width = 0
		for x := 0; x < y; x++ {
			go p.Run()
			input <- x
			input <- y
			width += <-output
			<-halted
		}
		y += dy
	}
	fmt.Printf("y: %d, width: %d\n", y, width)
	minX := 0
	for width >= 100 {
		width = 0
		minX = y
		for x := 0; x < y; x++ {
			go p.Run()
			input <- x
			input <- y
			out := <-output
			width += out
			<-halted
			if out != 0 && x < minX {
				minX = x
			}
		}
		y--
	}
	fmt.Printf("y: %d, minX: %d, width: %d\n", y, minX, width)

	minY := y

	topLeftX := minX
	topLeftY := minY
SuperSearch:
	for {
		topLeftX = minX - 1
		tempOut := 0
		for tempOut < 1 {
			topLeftX++
			go p.Run()
			input <- topLeftX
			input <- topLeftY
			tempOut = <-output
			<-halted
		}
	ShiftDown:
		for {
			area := 0
		ShiftRight:
			for x := topLeftX; x < topLeftX+100; x++ {
				for y := topLeftY; y < topLeftY+100; y++ {
					go p.Run()
					input <- x
					input <- y
					out := <-output
					<-halted
					area += out
					if out == 0 {
						if y == topLeftY {
							break ShiftDown
						}
						break ShiftRight
					}
				}
			}
			if area == 10000 {
				break SuperSearch
			}
			topLeftX++
		}
		topLeftY++
	}
	fmt.Printf("(%d,%d)\n", topLeftX, topLeftY)
}
