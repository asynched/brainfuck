package internal

import "fmt"

// Parser struct is used to parse the input string into a list of Runnable
// instructions
type Parser struct {
	input  []byte
	offset int
}

// NewParser creates a new Parser with the given input string
func NewParser(input []byte) *Parser {
	return &Parser{input: input, offset: 0}
}

// Parse parses the input string into a list of Runnable instructions.
// It returns a list of Runnable instructions or panics if an unexpected
// character is found.
func (p *Parser) Parse() []Runnable {
	instructions := make([]Runnable, 0)

	for p.offset < len(p.input) {
		switch p.input[p.offset] {
		case '+':
			instructions = append(instructions, InstructionIncr)
		case '-':
			instructions = append(instructions, InstructionDecr)
		case '>':
			instructions = append(instructions, InstructionNext)
		case '<':
			instructions = append(instructions, InstructionPrev)
		case '.':
			instructions = append(instructions, InstructionOut)
		case ',':
			instructions = append(instructions, InstructionIn)
		case '[':
			loop := &Loop{Instructions: make([]Runnable, 0)}

			p.offset++

			for p.input[p.offset] != ']' {
				instructions := p.Parse()
				loop.Instructions = append(loop.Instructions, instructions...)
			}

			instructions = append(instructions, loop)
		case ']':
			return instructions
		case '\n', '\t', ' ', '\r':
			// Ignore whitespace
		default:
			panic(fmt.Errorf("unexpected character '%c' at position %d", p.input[p.offset], p.offset))
		}

		p.offset++
	}

	return instructions
}
