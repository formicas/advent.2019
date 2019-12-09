package main

import (
	"testing"

	"advent.2019/intcode"
)

func Test_Day7_FirstExample(t *testing.T) {
	program := []int{3, 15, 3, 16, 1002, 16, 10, 16, 1, 16, 15, 15, 4, 15, 99, 0, 0}
	inputs := []int{4, 3, 2, 1, 0}

	result := doTheThing(program, inputs, false)

	if result != 43210 {
		t.Errorf("Thruster Output = %d; want 43210", result)
	}
}

func Test_Day7_SecondExample(t *testing.T) {
	program := []int{3, 23, 3, 24, 1002, 24, 10, 24, 1002, 23, -1, 23, 101, 5, 23, 23, 1, 24, 23, 23, 4, 23, 99, 0, 0}
	inputs := []int{0, 1, 2, 3, 4}

	result := doTheThing(program, inputs, false)

	if result != 54321 {
		t.Errorf("Thruster Output = %d, want 54321", result)
	}
}

func Test_Day7_Part2_FirstExample(t *testing.T) {
	program := []int{3, 26, 1001, 26, -4, 26, 3, 27, 1002, 27, 2, 27, 1, 27, 26, 27, 4, 27, 1001, 28, -1, 28, 1005, 28, 6, 99, 0, 0, 5}
	inputs := []int{9, 8, 7, 6, 5}
	result := doTheThing(program, inputs, true)

	if result != 139629729 {
		t.Errorf("Thruster Output = %d, want 139629729", result)
	}
}

func Test_Day7_Part2_SecondExample(t *testing.T) {
	program := []int{3, 52, 1001, 52, -5, 52, 3, 53, 1, 52, 56, 54, 1007, 54, 5, 55, 1005, 55, 26, 1001, 54, -5, 54, 1105, 1, 12, 1, 53, 54, 53, 1008, 54, 0, 55, 1001, 55, 1, 55, 2, 53, 55, 53, 4, 53, 1001, 56, -1, 56, 1005, 56, 6, 99, 0, 0, 0, 0, 10}
	inputs := []int{9, 7, 8, 5, 6}
	result := doTheThing(program, inputs, true)

	if result != 18216 {
		t.Errorf("Thruster Output = %d, want 18216", result)
	}
}

func Test_Day9_FirstExample(t *testing.T) {
	program := []int{109, 1, 204, -1, 1001, 100, 1, 100, 1008, 100, 16, 101, 1006, 101, 0, 99}
	outChannel := make(chan int, 50)

	p := intcode.Program{Code: program, Name: "test", Output: outChannel}

	p.Run()
	for i, v := range program {
		output := <-outChannel
		if output != v {
			t.Errorf("Item %d, expected %d, got %d", i, v, output)
		}
	}
}

func Test_Day9_SecondExample(t *testing.T) {
	program := []int{1102, 34915192, 34915192, 7, 4, 7, 99, 0}
	outChannel := make(chan int)

	p := intcode.Program{Code: program, Name: "test", Output: outChannel}
	go p.Run()
	output := <-outChannel
	if output != 1219070632396864 {
		t.Errorf("Expected 1219070632396864, was %d", output)
	}
}

func Test_Day9_ThirdExample(t *testing.T) {
	program := []int{104, 1125899906842624, 99}
	outChannel := make(chan int)

	p := intcode.Program{Code: program, Name: "test", Output: outChannel}
	go p.Run()
	output := <-outChannel
	if output != 1125899906842624 {
		t.Errorf("Expected 1125899906842624, was %d", output)
	}
}
