package day24

import "fmt"

//DoTheThing that day 24 is supposed to do
func DoTheThing() {
	state := "#.#...###....#.###..#...."
	states := make(map[int]struct{})

	states[encodeState(state)] = struct{}{}
	for {
		state = simulate(state)
		encoded := encodeState(state)
		_, exists := states[encoded]
		if exists {
			fmt.Println(encoded)
			break
		}
		states[encoded] = struct{}{}
	}

	//now do the horrible recursive shit
	state = "#.#...###...?#.###..#...."
	plutonianStates := make(map[int][]rune)
	minLevel := 0
	maxLevel := 0
	plutonianStates[0] = []rune(state)
	for i := 0; i < 200; i++ {
		plutonianStates, minLevel, maxLevel = simulatePlutonian(plutonianStates, minLevel, maxLevel)
	}

	totalBugs := 0
	for _, v := range plutonianStates {
		for _, c := range v {
			if c == '#' {
				totalBugs++
			}
		}
	}

	fmt.Println(totalBugs)
}

func encodeState(initial string) int {
	state := 0
	for i, c := range []rune(initial) {
		if c == '#' {
			state = state | (1 << i)
		}
	}
	return state
}

func simulatePlutonian(levels map[int][]rune, min, max int) (map[int][]rune, int, int) {
	if shouldGoDown(levels[min]) {
		min--
		levels[min] = []rune("............?............")
	}
	if shouldGoUp(levels[max]) {
		max++
		levels[max] = []rune("............?............")
	}

	newLevels := make(map[int][]rune)

	for i, level := range levels {
		newLevel := make([]rune, len(level))
		up, upExists := levels[i+1]
		down, downExists := levels[i-1]

		for j, c := range level {
			if c == '?' {
				continue
			}
			//count adjacent bugs
			adjacent := 0

			//j%5==0 - look at level+1, index 11
			if j%5 == 0 {
				if upExists && up[11] == '#' {
					adjacent++
				}
			} else if level[j-1] == '#' {
				adjacent++
			}

			//j%5==4 - look at level+1, index 13
			if j%5 == 4 {
				if upExists && up[13] == '#' {
					adjacent++
				}
			} else if level[j+1] == '#' {
				adjacent++
			}

			//j<5    - look at level+1, index 7
			if j < 5 {
				if upExists && up[7] == '#' {
					adjacent++
				}
			} else if level[j-5] == '#' {
				adjacent++
			}

			//j>19   - look at level+1, index 17
			if j > 19 {
				if upExists && up[17] == '#' {
					adjacent++
				}
			} else if level[j+5] == '#' {
				adjacent++
			}

			if downExists {
				//j==7   - look at level-1, index<5
				if j == 7 {
					for k := 0; k < 5; k++ {
						if down[k] == '#' {
							adjacent++
						}
					}
				}

				//j==17  - look at level-1, index>19
				if j == 17 {
					for k := 20; k < 25; k++ {
						if down[k] == '#' {
							adjacent++
						}
					}
				}

				//j==11  - look at level-1, index%5==0
				if j == 11 {
					for k := 0; k < 25; k += 5 {
						if down[k] == '#' {
							adjacent++
						}
					}
				}

				//j==13  - look at level-1, index%5==4
				if j == 13 {
					for k := 4; k < 25; k += 5 {
						if down[k] == '#' {
							adjacent++
						}
					}
				}
			}

			switch {
			case c == '#' && adjacent == 1:
				newLevel[j] = '#'
			case c == '#':
				newLevel[j] = '.'
			case c == '.' && (adjacent == 1 || adjacent == 2):
				newLevel[j] = '#'
			case c == '.':
				newLevel[j] = '.'
			}
		}
		newLevels[i] = newLevel
	}

	return newLevels, min, max
}

func shouldGoUp(level []rune) bool {
	perimeter := []int{0, 1, 2, 3, 4, 5, 9, 10, 14, 15, 19, 20, 21, 22, 23, 24}
	for _, i := range perimeter {
		if level[i] == '#' {
			return true
		}
	}
	return false
}

func shouldGoDown(level []rune) bool {
	insides := []int{7, 11, 13, 17}
	for _, i := range insides {
		if level[i] == '#' {
			return true
		}
	}
	return false
}

func simulate(state string) string {
	chars := []rune(state)
	newChars := make([]rune, len(state))

	for i, c := range chars {
		//count adjacent bugs
		adjacent := 0
		//adjacent cells are i-1,i+1,i-5,i+5
		//i-1 - ignore if i%5==0
		if i%5 != 0 && chars[i-1] == '#' {
			adjacent++
		}

		//i+1 - ignore if i%5==4
		if i%5 != 4 && chars[i+1] == '#' {
			adjacent++
		}

		//i-5 - ignore if i<5
		if i >= 5 && chars[i-5] == '#' {
			adjacent++
		}

		//i+5 - ignore if i>19
		if i < 20 && chars[i+5] == '#' {
			adjacent++
		}

		switch {
		case c == '#' && adjacent == 1:
			newChars[i] = '#'
		case c == '#':
			newChars[i] = '.'
		case c == '.' && (adjacent == 1 || adjacent == 2):
			newChars[i] = '#'
		case c == '.':
			newChars[i] = '.'
		}
	}

	return string(newChars)
}
