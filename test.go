package main

import "cpu"

func main() {

	c := cpu.New_cpu()

	/*program1 := []uint32{
		0x01ABBA05, // li ABBA $5   ; Load IP A.B.B.A to $5
		0x01000006, // li 0 $6      ; Load 0 to $6
		0x1C050600, // sw $5 $6     ; Store $5 on mem at 0
		0x1B060100, // lw $6 $1     ; Load mem at 0 to $1
		0x01FFF802, // li FFF8 $2   ; Load mask F.F.F.8 to $2
		0x02020103, // and $2 $1 $3 ; $3 = $2 & $1 = Calculate net
		0x1D010003, // sys 01 $3    ; Print net on $3
		0x01000A04, // li 000A $4   ; Load '\n' to $4
		0x1D020004, // sys 02 $4    ; Print '\n'
		0x05020300, // not $2 $3    ; Invert $2 and put on $3
		0x02010301, // and $1 $3 $1 ; $1 = $1 & $3 = Calculate host
		0x1D010001, // sys 01 $1    ; Print host on $1
		0x1D020004, // sys 02 $4    ; Print '\n'
		0x1F000000, // hlt          ; Stop program
	}*/

	program2 := []uint32{
		0x01000401, // 00 Load 4 into R1
		0x01000402, // 01 Load 5 into R2
		0x01005403, // 02 Load 'T' into R3
		0x01004604, // 03 Load 'F' into R4
		0x1A010200, // 04 Branch if R1 == R2
		// IF 04 == False
		0x1E000000 + 0x0700*4, // 05 Jump to 07
		// IF 04 == True
		0x1E000000 + 0x0900*4, // 06 Jump to 09
		//Print F routine
		0x1D020004,            // 07 Print R4 char
		0x1E000000 + 0x0B00*4, // 08 Jump to 0B
		//Print T routine
		0x1D020003,            // 09 Print R3 char
		0x1E000000 + 0x0B00*4, // 0A Jump to 0B
		0x1F000000,            // 0B Halt
	}

	c.Load_program(program2)
	c.Exec_program()
	//fmt.Println(int32(c.reg[1]))
	//fmt.Println("")
}
