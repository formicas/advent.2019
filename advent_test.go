package main

import (
	"testing"

	"advent.2019/intcode"
)

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
