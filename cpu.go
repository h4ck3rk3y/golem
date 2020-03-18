package cpu

// CPU ...
type CPU struct {
	memory [8192]byte //8K SRAM
}

// NewCPU ...
func NewCPU() CPU {
	c := CPU{}
	return c
}
