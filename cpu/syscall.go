package cpu

import "fmt"

// Syscall do sistema
// Separado pq obviamente vai ficar bem grande
// A única que vai usar importações também
func sys(c *cpu, op1, op2, op3 uint32) {
	/*
		op1 = 00 decide by reg op2 (same values)
		op1 = 01 print reg op3
		op1 = 02 print reg op3 as char
	*/

	var call int
	if op1 == 0 {
		call = int(c.reg[op2])
	} else {
		call = int(op1)
	}

	switch call {
	case 1:
		fmt.Print(c.reg[op3])
	case 2:
		fmt.Printf("%c", rune(int(c.reg[op3])))
	}
	c.pc += PC_INCREMENT
}
