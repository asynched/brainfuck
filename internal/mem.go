package internal

// Memory struct is used to store the memory buffer and the current offset of
// the memory pointer.
type Memory struct {
	Buffer [4096]byte
	Offset int
}

// NewMemory creates a new Memory struct with the buffer initialized to 0 and the
// offset set to 0.
func NewMemory() *Memory {
	return &Memory{
		Buffer: [4096]byte{},
		Offset: 0,
	}
}
