package main

func and(c *cpu, op1, op2, op3 uint32) {
	// and $1 & $2 -> $3
	c.reg[op3] = c.reg[op1] & c.reg[op2]
	c.pc += PC_INCREMENT
}

func or(c *cpu, op1, op2, op3 uint32) {
	// or $1 | $2 -> $3
	c.reg[op3] = c.reg[op1] | c.reg[op2]
	c.pc += PC_INCREMENT
}

func xor(c *cpu, op1, op2, op3 uint32) {
	// xor $1 ^ $2 -> $3
	c.reg[op3] = c.reg[op1] ^ c.reg[op2]
	c.pc += PC_INCREMENT
}

func not(c *cpu, op1, op2, op3 uint32) {
	// not ^$1 -> $2
	c.reg[op2] = ^c.reg[op1]
	c.pc += PC_INCREMENT
}

func nand(c *cpu, op1, op2, op3 uint32) {
	// nand $1 ^& $2 -> $3
	c.reg[op3] = ^(c.reg[op1] & c.reg[op2])
	c.pc += PC_INCREMENT
}

func nor(c *cpu, op1, op2, op3 uint32) {
	// nor $1 ^| $2 -> $3
	c.reg[op3] = ^(c.reg[op1] | c.reg[op2])
	c.pc += PC_INCREMENT
}

func add(c *cpu, op1, op2, op3 uint32) {
	// add $1 + $2 -> $3
	c.reg[op3] = c.reg[op1] + c.reg[op2]
	c.pc += PC_INCREMENT
}

func sub(c *cpu, op1, op2, op3 uint32) {
	// sub $1 - $2 -> $3
	c.reg[op3] = c.reg[op1] - c.reg[op2]
	c.pc += PC_INCREMENT
}

func addi(c *cpu, op1, op2, op3 uint32) {
	// addi $1 + literal -> $2
	c.reg[op3] = c.reg[op1] + op2
	c.pc += PC_INCREMENT
}

func subi(c *cpu, op1, op2, op3 uint32) {
	//subi $1 - literal -> $2
	c.reg[op3] = c.reg[op1] - op2
	c.pc += PC_INCREMENT
}

func sll(c *cpu, op1, op2, op3 uint32) {
	// sll $1 << $2 -> $3
	c.reg[op3] = c.reg[op1] << c.reg[op2]
	c.pc += PC_INCREMENT
}

func srl(c *cpu, op1, op2, op3 uint32) {
	// sll $1 << $2 -> $3
	c.reg[op3] = c.reg[op1] >> c.reg[op2]
	c.pc += PC_INCREMENT
}
