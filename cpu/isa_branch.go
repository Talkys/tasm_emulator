package cpu

func jmp(c *cpu, op1, op2, op3 uint32) {
	// jmp 16 bit addr_literal
	addr := uint32(0)
	addr += op1 << ONE_BYTE
	addr += op2

	c.pc = int(addr & TWO_BYTE_MASK)
}

func beq(c *cpu, op1, op2, op3 uint32) {
	// beq if $2 = $3 jumps a line
	if c.reg[op1] == c.reg[op2] {
		c.pc += PC_INCREMENT
	}

	c.pc += PC_INCREMENT
}
