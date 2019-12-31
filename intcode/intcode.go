package intcode

//Program - represents a program and its code
type Program struct {
	Code         []int
	Input        chan int
	Output       chan int
	Name         string
	Halted       chan int
	relativeBase int
}

func (p *Program) extendMemory(targetAddress int) {
	newCode := make([]int, targetAddress+1)
	copy(newCode, p.Code)
	p.Code = newCode
}

//Run - execute a program
func (p Program) Run() {
	pointer := 0
	for {
		instruction := p.Code[pointer]

		opCode := instruction % 100
		p1Mode := (instruction / 100) % 10
		p2Mode := (instruction / 1000) % 10
		p3Mode := (instruction / 10000) % 10

		switch opCode {
		case 1:
			slice := p.Code[pointer : pointer+4]
			p1 := p.getValue(p1Mode, slice[1])
			p2 := p.getValue(p2Mode, slice[2])

			p.setValue(p3Mode, p1+p2, slice[3])

			pointer = pointer + 4
		case 2:
			slice := p.Code[pointer : pointer+4]
			p1 := p.getValue(p1Mode, slice[1])
			p2 := p.getValue(p2Mode, slice[2])

			p.setValue(p3Mode, p1*p2, slice[3])

			pointer = pointer + 4
		case 3:
			slice := p.Code[pointer : pointer+2]

			value := <-p.Input
			p.setValue(p1Mode, value, slice[1])

			pointer = pointer + 2
		case 4:
			slice := p.Code[pointer : pointer+2]
			value := p.getValue(p1Mode, slice[1])

			p.Output <- value
			pointer = pointer + 2
		case 5:
			slice := p.Code[pointer : pointer+3]
			value := p.getValue(p1Mode, slice[1])
			newPointer := p.getValue(p2Mode, slice[2])

			if value != 0 {
				pointer = newPointer
			} else {
				pointer = pointer + 3
			}
		case 6:
			slice := p.Code[pointer : pointer+3]
			value := p.getValue(p1Mode, slice[1])
			newPointer := p.getValue(p2Mode, slice[2])
			if value == 0 {
				pointer = newPointer
			} else {
				pointer = pointer + 3
			}
		case 7:
			slice := p.Code[pointer : pointer+4]
			value := 0
			if p.getValue(p1Mode, slice[1]) < p.getValue(p2Mode, slice[2]) {
				value = 1
			}

			p.setValue(p3Mode, value, slice[3])

			pointer = pointer + 4
		case 8:
			slice := p.Code[pointer : pointer+4]
			value := 0
			if p.getValue(p1Mode, slice[1]) == p.getValue(p2Mode, slice[2]) {
				value = 1
			}

			p.setValue(p3Mode, value, slice[3])

			pointer = pointer + 4
		case 9:
			slice := p.Code[pointer : pointer+2]
			adjustment := p.getValue(p1Mode, slice[1])
			p.relativeBase += adjustment

			pointer = pointer + 2
		case 99:
			if p.Halted != nil {
				p.Halted <- 1
			}
			return
		}
	}
}

func (p *Program) getValue(mode int, value int) int {
	switch mode {
	case 2:
		targetAddress := p.relativeBase + value
		if targetAddress >= len(p.Code) {
			return 0
		}
		return p.Code[targetAddress]
	case 1:
		return value
	default:
		if value >= len(p.Code) {
			return 0
		}
		return p.Code[value]
	}
}

func (p *Program) setValue(mode int, value int, target int) {
	var targetAddress int
	switch mode {
	case 2:
		targetAddress = p.relativeBase + target
	default:
		targetAddress = target
	}

	if targetAddress >= len(p.Code) {
		p.extendMemory(targetAddress)
	}

	p.Code[targetAddress] = value
}
