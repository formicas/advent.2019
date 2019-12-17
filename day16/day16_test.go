package day16

import (
	"testing"
)

func Test_buildPatterns(t *testing.T) {
	patterns := buildPatterns(5)
	expectedPatterns := [][]int{
		[]int{0, 1, 0, -1},
		[]int{0, 0, 1, 1, 0, 0, -1, -1},
		[]int{0, 0, 0, 1, 1, 1, 0, 0, 0, -1, -1, -1},
		[]int{0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0, -1, -1, -1, -1},
		[]int{0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0, -1, -1, -1, -1, -1},
	}

	if len(patterns) != len(expectedPatterns) {
		t.Errorf("Patterns length is %d, Expected %d", len(patterns), len(expectedPatterns))
	}

	for i, p := range patterns {
		if len(p) != len(expectedPatterns[i]) {
			t.Errorf("Pattern %d length is %d, expected %d", i, len(p), len(expectedPatterns[i]))
		}
		for j, v := range p {
			if v != expectedPatterns[i][j] {
				t.Errorf("(%d,%d) is %d, expected %d", i, j, v, expectedPatterns[i][j])
			}
		}
	}
}

func Test_fft(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8}
	expectedOutputs := [][]int{
		[]int{4, 8, 2, 2, 6, 1, 5, 8},
		[]int{3, 4, 0, 4, 0, 4, 3, 8},
		[]int{0, 3, 4, 1, 5, 5, 1, 8},
		[]int{0, 1, 0, 2, 9, 4, 9, 8},
	}
	patterns := buildPatterns(len(input))

	for _, expectedOutput := range expectedOutputs {
		input = fft(input, &patterns)
		for i, v := range input {
			if v != expectedOutput[i] {
				t.Errorf("Got %v, Expected %v", input, expectedOutput)
				break
			}

		}
	}

}

func Test_fft100(t *testing.T) {
	tests := []struct {
		input    string
		expected []int
	}{
		{
			input:    "80871224585914546619083218645595",
			expected: []int{2, 4, 1, 7, 6, 1, 7, 6},
		},
		{
			input:    "19617804207202209144916044189917",
			expected: []int{7, 3, 7, 4, 5, 4, 1, 8},
		},
		{
			input:    "69317163492948606335995924319873",
			expected: []int{5, 2, 4, 3, 2, 1, 3, 3},
		},
	}

	for i, test := range tests {
		sequence := buildSequence(test.input)
		patterns := buildPatterns(len(sequence))
		for i := 0; i < 100; i++ {
			sequence = fft(sequence, &patterns)
		}

		for j, expected := range test.expected {
			if sequence[j] != expected {
				t.Errorf("Test %d: got %v, expected %v", i, sequence, test.expected)
				break
			}
		}
	}
}

func Test_fastFft(t *testing.T) {
	tests := []struct {
		input    string
		offset   int
		expected string
	}{
		{
			input:    "03036732577212944063491565474664",
			offset:   303673,
			expected: "84462026",
		},
		{
			input:    "02935109699940807407585447034323",
			offset:   293510,
			expected: "78725270",
		},
		{
			input:    "03081770884921959731165446850517",
			offset:   308177,
			expected: "53553731",
		},
	}

	for i, test := range tests {
		output := fastFft(test.input, test.offset, 100)
		if output != test.expected {
			t.Errorf("Test %d: got %v, expected %v", i, output, test.expected)
		}
	}
}
