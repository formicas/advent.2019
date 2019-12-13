package day12

import (
	"fmt"
	"math"
)

//DoTheThing that day 12 is supposed to do
func DoTheThing() {
	moons := []moon{
		{p: vector3{x: -6, y: 2, z: -9}},
		{p: vector3{x: 12, y: -14, z: -4}},
		{p: vector3{x: 9, y: 5, z: -6}},
		{p: vector3{x: -1, y: -4, z: 9}},
	}

	fmt.Println(getTotalEnergy(simulateMoons(moons, 1000)))

	x := []int{-6, 12, 9, -1, 0, 0, 0, 0}
	y := []int{2, -14, 5, -4, 0, 0, 0, 0}
	z := []int{-9, -4, -6, 9, 0, 0, 0, 0}

	xSteps := getSteps(x)
	ySteps := getSteps(y)
	zSteps := getSteps(z)
	fmt.Printf("x steps:%d\ny steps:%d\nz steps:%d\n", xSteps, ySteps, zSteps)
	fmt.Printf("Lowest Common Multiple: %d\n", lcm([]int{xSteps, ySteps, zSteps}))
}

func lcm(numbers []int) int {
	lcmPrimeFactors := make(map[int]int)
	for _, v := range numbers {
		factors := factorise(v)
		//take the higest power of the common factors
		for p, f := range factors {
			_, exists := lcmPrimeFactors[p]
			if !exists {
				lcmPrimeFactors[p] = f
			}
			lcmPrimeFactors[p] = int(math.Max(float64(lcmPrimeFactors[p]), float64(f)))
		}
	}
	lcm := 1
	//munge them all together
	for p, f := range lcmPrimeFactors {
		lcm = lcm * int(math.Pow(float64(p), float64(f)))
	}
	return lcm
}

func factorise(n int) map[int]int {
	factors := make([]int, 0)
	primeFactors := make(map[int]int)
	for n%2 == 0 {
		factors = append(factors, 2)
		n = n / 2
	}

	for i := 3; i < n*n; i = i + 2 {
		if n%i == 0 {
			factors = append(factors, i)
			n = n / i
		}
	}

	if n > 2 {
		factors = append(factors, n)
	}

	//put it in a nicely digestible format - a map of the component primes and their powers
	for _, f := range factors {
		_, exists := primeFactors[f]
		if !exists {
			primeFactors[f] = 0
		}
		primeFactors[f]++
	}

	return primeFactors
}

func getSteps(foo []int) int {
	states := make(map[int][][]int)

	steps := 0
	for {
		hash := hash(foo)
		copyForStorage := make([]int, len(foo))
		copy(copyForStorage, foo)
		foos, exists := states[hash]
		if exists {
			for _, v := range foos {
				if compareSlices(foo, v) {
					return steps
				}
			}

			states[hash] = append(foos, copyForStorage)

		} else {
			states[hash] = [][]int{copyForStorage}
		}
		for i := 0; i < 4; i++ {
			for j := 0; j < 4; j++ {
				if foo[i] < foo[j] {
					foo[i+4]++
				} else if foo[i] > foo[j] {
					foo[i+4]--
				}
			}
		}
		for i := 0; i < 4; i++ {
			foo[i] += foo[i+4]
		}
		steps++
	}
}

func compareSlices(left, right []int) bool {
	if len(left) != len(right) {
		return false
	}

	for i, v := range left {
		if v != right[i] {
			return false
		}
	}

	return true
}

func hash(foo []int) int {
	h := len(foo)
	for _, i := range foo {
		h = h*31 + i
	}
	return h
}

func areEqual(left, right []moon) bool {
	if len(left) != len(right) {
		return false
	}

	for i, v := range left {
		if v != right[i] {
			return false
		}
	}

	return true
}

type vector3 struct {
	x, y, z int
}

type moon struct {
	name string
	p, v vector3
}

func getTotalEnergy(moons []moon) int {
	total := 0
	for _, moon := range moons {
		total += getEnergy(moon)
	}
	return total
}

func getEnergy(moon moon) int {
	return sumAbs(moon.p) * sumAbs(moon.v)
}

func sumAbs(v vector3) int {
	return int(math.Abs(float64(v.x)) + math.Abs(float64(v.y)) + math.Abs(float64(v.z)))
}

func simulateMoons(moons []moon, steps int) []moon {
	for i := 0; i < steps; i++ {
		adjustVelocity(moons)
		applyVelocity(moons)
	}

	return moons
}

func adjustVelocity(moons []moon) {
	for i, moon := range moons {
		for j, otherMoon := range moons {
			if i == j {
				continue
			}

			if moon.p.x < otherMoon.p.x {
				moon.v.x++
			} else if moon.p.x > otherMoon.p.x {
				moon.v.x--
			}

			if moon.p.y < otherMoon.p.y {
				moon.v.y++
			} else if moon.p.y > otherMoon.p.y {
				moon.v.y--
			}

			if moon.p.z < otherMoon.p.z {
				moon.v.z++
			} else if moon.p.z > otherMoon.p.z {
				moon.v.z--
			}
		}

		moons[i] = moon
	}
}

func applyVelocity(moons []moon) {
	for i, moon := range moons {
		moon.p.x += moon.v.x
		moon.p.y += moon.v.y
		moon.p.z += moon.v.z
		moons[i] = moon
	}

}
