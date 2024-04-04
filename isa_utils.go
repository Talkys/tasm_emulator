package main

func nop(c *cpu, op1, op2, op3 uint32) {
	c.pc += PC_INCREMENT
}

func li(c *cpu, op1, op2, op3 uint32) {
	// mov literal -> $1
	literal := (op1 << ONE_BYTE) + op2
	c.reg[op3] = literal
	c.pc += PC_INCREMENT
}

func hlt(c *cpu, op1, op2, op3 uint32) {
	c.halt = true
	c.pc += PC_INCREMENT
}
