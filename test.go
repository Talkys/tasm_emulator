package main

import (
	"fmt"
	"tasm/assembler"
)

func main() {

	//c := cpu.New_cpu()
	//c.Load_program(program2)
	//c.Exec_program()
	//fmt.Println(int32(c.reg[1]))
	//fmt.Println("")

	program := []string{
		"li,0x05,$1,0",
		"add,$1,$1,$1",
	}

	for _, line := range assembler.Assemble(program) {
		fmt.Println(line)
	}
}
