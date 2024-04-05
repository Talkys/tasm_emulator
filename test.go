package main

import (
	"fmt"
	"tasm/assembler"
	"tasm/cpu"
)

func main() {

	//c := cpu.New_cpu()
	//c.Load_program(program2)
	//c.Exec_program()
	//fmt.Println(int32(c.reg[1]))
	//fmt.Println("")

	program := []string{
		"li, 0x0A, $1, 0",
		"li, 0x06, $2, 0",
		"add, $1, $2, $3",
		"sys, 0x01, 0, $3",
		"hlt",
	}

	binary, jmp_table := assembler.Assemble(program)
	for _, line := range binary {
		fmt.Printf("0x%08X\n", line)
	}

	c := cpu.New_cpu()
	c.Load_program(binary, jmp_table)
}
