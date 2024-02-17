package internal

import "fmt"

// Runnable is an interface that represents an instruction that can be executed
// on the memory buffer.
type Runnable interface {
	// Run runs the instruction on the given memory buffer
	Run(memory *Memory)
}

// Instruction is an enum of all possible operations that can be performed on
// the memory buffer.
type Instruction int

const (
	// Character '+' increments the value at the current memory cell
	InstructionIncr Instruction = iota
	// Character '-' decrements the value at the current memory cell
	InstructionDecr
	// Character '>' moves the memory pointer to the next cell
	InstructionNext
	// Character '<' moves the memory pointer to the previous cell
	InstructionPrev
	// Character '.' outputs the value at the current memory cell as a character
	InstructionOut
	// Character ',' reads a character from the input and stores it at the current
	// memory cell
	InstructionIn
)

// Run runs the instruction on the given memory buffer
func (i Instruction) Run(memory *Memory) {
	switch i {
	case InstructionIncr:
		memory.Buffer[memory.Offset]++
	case InstructionDecr:
		memory.Buffer[memory.Offset]--
	case InstructionNext:
		memory.Offset++
	case InstructionPrev:
		memory.Offset--
	case InstructionOut:
		fmt.Printf("%c", memory.Buffer[memory.Offset])
	case InstructionIn:
		var input byte

		_, err := fmt.Scanf("%c", &input)

		if err != nil {
			panic(err)
		}

		memory.Buffer[memory.Offset] = input
	}
}

// String returns the string representation of the instruction
func (i Instruction) String() string {
	switch i {
	case InstructionIncr:
		return "InstructionIncr"
	case InstructionDecr:
		return "InstructionDecr"
	case InstructionNext:
		return "InstructionNext"
	case InstructionPrev:
		return "InstructionPrev"
	case InstructionOut:
		return "InstructionOut"
	case InstructionIn:
		return "InstructionIn"
	default:
		panic(fmt.Errorf("unknown instruction: %d", i))
	}
}

// Loop is a list of instructions that are executed until the current memory
// cell is 0.
type Loop struct {
	Instructions []Runnable
}

// Run runs the loop on the given memory buffer.
func (l *Loop) Run(memory *Memory) {
	for memory.Buffer[memory.Offset] != 0 {
		for _, instruction := range l.Instructions {
			instruction.Run(memory)
		}
	}
}
