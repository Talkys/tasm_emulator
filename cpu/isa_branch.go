package cpu

func jmp(c *cpu, op1, op2, op3 uint32) {
	// jmp 16 bit addr_literal
	addr := uint32(0)
	addr += op1 << TWO_BYTES
	addr += op2 << ONE_BYTE
	addr += op3

	jmp_addr := c.jmp_table[addr]

	c.pc = int(jmp_addr)
}

func beq(c *cpu, op1, op2, op3 uint32) {
	// beq if $1 = $2 jumps a line
	if c.reg[op1] == c.reg[op2] {
		c.pc += PC_INCREMENT
	}

	c.pc += PC_INCREMENT
}
