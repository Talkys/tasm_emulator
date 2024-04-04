package main

func sw(c *cpu, op1, op2, op3 uint32) {
	// sw reg op1 -> data[reg op2]
	for offset := range ui32to4ui8(c.reg[op1]) {
		c.data[c.reg[op2]+uint32(offset)] = ui32to4ui8(c.reg[op1])[offset]
	}
	c.pc += PC_INCREMENT
}

func lw(c *cpu, op1, op2, op3 uint32) {
	// lw data[reg op1] -> reg op2
	c.reg[op2] = ui32fr4ui8([4]uint8(c.data[c.reg[op1] : c.reg[op1]+4]))
	c.pc += PC_INCREMENT
}
