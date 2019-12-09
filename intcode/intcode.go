package intcode

//Program - represents a program and its code
type Program struct {
	Code   []int
	Input  chan int
	Output chan int
	Name   string
	Halted chan int
}

//Run - execute a program
func (p Program) Run() {
	RunIntCode(p)
}

//RunIntCode takes program input and does the needful
func RunIntCode(p Program) {
	//var intcode []int
	intcode := p.Code
	pointer := 0
	for {
		instruction := intcode[pointer]

		opCode := instruction % 100
		instruction = instruction / 100

		p1Mode := instruction % 10
		instruction = instruction / 10

		p2Mode := instruction % 10
		instruction = instruction / 10

		//p3Mode := instruction % 10

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
			// reader := bufio.NewReader(os.Stdin)
			// fmt.Println("Enter text:")
			// input, _ := reader.ReadString('\n')
			// asInt, err := strconv.Atoi(strings.TrimSpace(input))
			// if err != nil {
			// 	fmt.Println(err)
			// 	os.Exit(2)
			// }
			// fmt.Printf("Worker %v awaiting input\n", p)
			intcode[dest] = <-p.Input
			// fmt.Printf("Worker %v received %d\n", p, intcode[dest])
			pointer = pointer + 2
		case 4:
			slice := intcode[pointer : pointer+2]
			value := getValue(intcode, p1Mode, slice[1])
			// fmt.Printf("Worker %v Output: %d\n", p, value)
			p.Output <- value
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
			//fmt.Printf("Worker %v halting\n", p.Name)
			if p.Halted != nil {
				p.Halted <- 1
			}
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
