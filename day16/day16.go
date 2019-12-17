package day16

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

//DoTheThing that day 16 is supposed to do
func DoTheThing() {
	input := "59762770781817719190459920638916297932099919336473880209100837309955133944776196290131062991588533604012789279722697427213158651963842941000227675363260513283349562674004015593737518754413236241876959840076372395821627451178924619604778486903040621916904575053141824939525904676911285446889682089563075562644813747239285344522507666595561570229575009121663303510763018855038153974091471626380638098740818102085542924937522595303462725145620673366476987473519905565346502431123825798174326899538349747404781623195253709212209882530131864820645274994127388201990754296051264021264496618531890752446146462088574426473998601145665542134964041254919435635"
	sequence := buildSequence(input)
	patterns := buildPatterns(len(sequence))
	for i := 0; i < 100; i++ {
		sequence = fft(sequence, &patterns)
	}

	for i := 0; i < 8; i++ {
		fmt.Print(sequence[i])
	}
	fmt.Println()

	fmt.Println(fastFft(input, 5976277, 100))
}

func buildSequence(input string) []int {
	runes := []rune(input)
	sequence := make([]int, len(runes))
	for i, r := range runes {
		sequence[i], _ = strconv.Atoi(string(r))
	}
	return sequence
}

func buildPatterns(n int) [][]int {
	patterns := make([][]int, n)
	patterns[0] = []int{0, 1, 0, -1}
	basePattern := patterns[0]
	for i := 1; i < n; i++ {
		pattern := make([]int, len(basePattern)*(i+1))
		for j, v := range basePattern {
			for k := 0; k <= i; k++ {
				pattern[j*(i+1)+k] = v
			}
		}
		patterns[i] = pattern
	}

	return patterns
}

func fft(input []int, patternsPointer *[][]int) []int {
	patterns := *patternsPointer
	output := make([]int, len(input))
	for i := range input {
		output[i] = applyPattern(input, &patterns[i])
	}

	return output
}

func applyPattern(input []int, patternPointer *[]int) int {
	pattern := *patternPointer
	i := 1
	iv := 0

	for _, v := range input {
		iv += pattern[i%len(pattern)] * v
		i++
	}

	return int(math.Abs(float64(iv % 10)))
}

func fastFft(input string, offset int, iterations int) string {
	//all these examples, including the actual test, only care about output in the second half of the array
	//that means for each iteration, each digit is just the sum of the digits after it

	runes := []rune(input)
	totalLength := len(runes) * 10000
	truncatedLength := totalLength - offset
	output := make([]int, truncatedLength)
	runedex := len(runes) - 1
	//build phase0
	for i := truncatedLength - 1; i >= 0; i-- {
		output[i], _ = strconv.Atoi(string(runes[runedex]))
		runedex--
		if runedex < 0 {
			runedex = len(runes) - 1
		}
	}
	for i := 0; i < iterations; i++ {
		//don't need to do the last item - it never changes
		newOutput := make([]int, truncatedLength)
		for j := truncatedLength - 1; j >= 0; j-- {
			if j == truncatedLength-1 {
				newOutput[j] = output[j]
			} else {
				newOutput[j] = (output[j] + newOutput[j+1]) % 10
			}
		}
		output = newOutput
	}

	var str strings.Builder
	for i := 0; i < 8; i++ {
		str.WriteString(strconv.Itoa(output[i]))
	}

	return str.String()
}
