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
		"li, 0, 10, $1",
		"li, 0, 0x0B, $2",
		"beq, $1, $2, 0",
		//False
		"jmp :wasfalse",
		//True
		"jmp :wastrue",
		":wasfalse",
		"add, $1, $2, $3",
		"sys, 0x01, 0, $3",
		"sys, 0x02, 0, $1",
		"jmp :end",
		":wastrue",
		"add, $1, $1, $3",
		"sys, 0x01, 0, $3",
		"sys, 0x02, 0, $1",
		"jmp :end",
		":end",
		"hlt",
	}

	binary, jmp_table := assembler.Assemble(program)
	for _, line := range binary {
		fmt.Printf("0x%08X\n", line)
	}
	fmt.Println("--------------")

	c := cpu.New_cpu()
	c.Load_program(binary, jmp_table)
	c.Exec_program()
}
