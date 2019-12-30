package day22

import (
	"fmt"
	"math/big"
	"regexp"
	"strconv"
	"strings"
)

//DoTheThing that day 22 is supposed to do
func DoTheThing() {
	instructions := `deal into new stack
deal with increment 25
cut -5919
deal with increment 56
deal into new stack
deal with increment 20
deal into new stack
deal with increment 53
cut 3262
deal with increment 63
cut 3298
deal into new stack
cut -4753
deal with increment 57
deal into new stack
cut 9882
deal with increment 42
deal into new stack
deal with increment 40
cut 2630
deal with increment 32
cut 1393
deal with increment 74
cut 2724
deal with increment 23
cut -3747
deal into new stack
cut 864
deal with increment 61
deal into new stack
cut -4200
deal with increment 72
cut -7634
deal with increment 32
deal into new stack
cut 6793
deal with increment 38
cut 7167
deal with increment 10
cut -9724
deal into new stack
cut 6047
deal with increment 37
cut 7947
deal with increment 63
deal into new stack
deal with increment 9
cut -9399
deal with increment 26
cut 1154
deal with increment 74
deal into new stack
cut 3670
deal with increment 45
cut 3109
deal with increment 64
cut -7956
deal with increment 39
deal into new stack
deal with increment 61
cut -9763
deal with increment 20
cut 4580
deal with increment 30
deal into new stack
deal with increment 62
deal into new stack
cut -997
deal with increment 54
cut -1085
deal into new stack
cut -9264
deal into new stack
deal with increment 11
cut 6041
deal with increment 9
deal into new stack
cut 5795
deal with increment 26
cut 5980
deal with increment 38
cut 1962
deal with increment 25
cut -565
deal with increment 45
cut 9490
deal with increment 21
cut -3936
deal with increment 64
deal into new stack
cut -7067
deal with increment 75
cut -3975
deal with increment 29
deal into new stack
cut -7770
deal into new stack
deal with increment 12
cut 8647
deal with increment 49`

	currentPosition := big.NewInt(2019)
	length := big.NewInt(10007)
	repeats := big.NewInt(1)
	a, b := big.NewInt(1), big.NewInt(0)
	for _, l := range strings.Split(instructions, "\n") {
		a, b = shuffle(a, b, length, l, false)
		// interim := calculate(a, b, currentPosition, length)
		// fmt.Println(interim)
	}

	a, b = repeat(a, b, length, repeats)
	currentPosition = calculate(a, b, currentPosition, length)
	fmt.Println(currentPosition)
	fmt.Println()
	//part 2 miracle brain fart:
	//finding the value in position x is equivalent to
	//apply the inverse shuffle and find the position of value x
	//inverse of deal into new stack is... deal into new stack
	//inverse of cut n is cut -n
	//inverse of deal in increment n is deal in modinv(n,L) -- this got me good
	//and go through the steps backwards

	currentPosition = big.NewInt(2020)
	length = big.NewInt(119315717514047)
	repeats = big.NewInt(101741582076661)
	lines := strings.Split(instructions, "\n")
	a, b = big.NewInt(1), big.NewInt(0)
	for i := len(lines) - 1; i >= 0; i-- {
		a, b = shuffle(a, b, length, lines[i], true)
		// interim := calculate(a, b, currentPosition, length)
		// fmt.Println(interim)
	}

	a, b = repeat(a, b, length, repeats)
	currentPosition = calculate(a, b, currentPosition, length)
	fmt.Println(currentPosition)
}

func shuffle(a, b, length *big.Int, instruction string, invert bool) (*big.Int, *big.Int) {
	dealRegex := regexp.MustCompile("deal with increment (\\d+)")
	cutRegex := regexp.MustCompile("cut (-?\\d+)")

	matches := dealRegex.FindStringSubmatch(instruction)
	if matches != nil {
		increment, _ := strconv.Atoi(matches[1])
		bigIncrement := big.NewInt(int64(increment))
		if invert {
			bigIncrement = bigIncrement.ModInverse(bigIncrement, length)
		}
		return a.Mul(a, bigIncrement).Mod(a, length), b.Mul(b, bigIncrement).Mod(b, length)
	}

	matches = cutRegex.FindStringSubmatch(instruction)
	if matches != nil {
		size, _ := strconv.Atoi(matches[1])
		bigSize := big.NewInt(int64(size))
		if invert {
			bigSize = bigSize.Neg(bigSize)
		}
		return a, b.Sub(b, bigSize).Mod(b, length)
	}

	return a.Neg(a), b.Neg(b).Sub(b, big.NewInt(1)).Mod(b, length)
}

func repeat(a, b, length, r *big.Int) (*big.Int, *big.Int) {
	switch {
	case r.Cmp(big.NewInt(0)) == 0:
		return big.NewInt(1), big.NewInt(0)
	case big.NewInt(0).Mod(r, big.NewInt(2)).Cmp(big.NewInt(0)) == 0:
		newA, newB, newR := big.NewInt(0), big.NewInt(0), big.NewInt(0)

		return repeat(newA.Mul(a, a).Mod(newA, length), newB.Mul(a, b).Add(newB, b).Mod(newB, length), length, newR.Div(r, big.NewInt(2)))
	default:
		newA, newB, newR := big.NewInt(0), big.NewInt(0), big.NewInt(0)
		c, d := repeat(a, b, length, newR.Sub(r, big.NewInt(1)))
		return newA.Mul(a, c).Mod(newA, length), newB.Mul(a, d).Add(newB, b).Mod(newB, length)
	}
}

func calculate(a, b, x, mod *big.Int) *big.Int {
	result := big.NewInt(0)
	result = result.Mul(a, x).Add(result, b).Mod(result, mod)
	if result.Cmp(big.NewInt(0)) < 0 {
		result = result.Add(result, mod)
	}
	return result
}
