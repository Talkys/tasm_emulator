package cpu

const (
	MEM_SIZE  = 4 * 1024 // 4kb = 1024 instructions of 32 bits
	DATA_SIZE = 65536    // 16 bit addr mem
	REG_NUM   = 256      // Number of registers
	OP_SIZE   = 32       // Number of operations

	FOUR_BYTES  = 8 * 4
	THREE_BYTES = 8 * 3
	TWO_BYTES   = 8 * 2
	ONE_BYTE    = 8 * 1

	ONE_BYTE_MASK = 0x000000FF
	TWO_BYTE_MASK = 0x0000FFFF

	PC_INCREMENT = 4 // Incrementar em 4 para compensar o tamanho de palavra (8 * 4 = 32 bits)
)

var (
	ISA = []func(c *cpu, op1, op2, op3 uint32){
		nop,  // OPCODE 00 0x00
		li,   // OPCODE 01 0x01
		and,  // OPCODE 02 0x02
		or,   // OPCODE 03 0x03
		xor,  // OPCODE 04 0x04
		not,  // OPCODE 05 0x05
		nand, // OPCODE 06 0x06
		nor,  // OPCODE 07 0x07
		nop,
		nop,
		add,  // OPCODE 10 0x0A
		sub,  // OPCODE 11 0x0B
		addi, // OPCODE 12 0x0C
		subi, // OPCODE 13 0x0D
		sll,  // OPCODE 14 0x0E
		srl,  // OPCODE 15 0x0F
		nop,
		nop,
		nop,
		nop,
		nop,
		nop,
		nop,
		nop,
		nop,
		nop,
		beq, // OPCODE 26 0x1A
		lw,  // OPCODE 27 0x1B
		sw,  // OPCODE 28 0x1C
		sys, // OPCODE 29 0x1D
		jmp, // OPCODE 30 0x1E
		hlt, // OPCODE 31 0x1F
	}
)

type cpu struct {
	mem  [MEM_SIZE]uint8
	pc   int
	reg  [REG_NUM]uint32
	halt bool
	data [DATA_SIZE]uint8
}

func New_cpu() cpu {
	var c cpu
	c.pc = 0
	c.halt = false
	for i := range MEM_SIZE {
		c.mem[i] = 0
	}
	for i := range REG_NUM {
		c.reg[i] = 0
	}
	return c
}

func (c *cpu) Exec_inst(inst uint32) {

	word := ui32to4ui8(inst)
	/*if word[0] != 0 {
		fmt.Print(c.pc)
		fmt.Print(" - ")
		fmt.Println(word)
	}*/
	opcode := int(word[0])
	op1 := uint32(word[1])
	op2 := uint32(word[2])
	op3 := uint32(word[3])

	if opcode >= 0 && opcode < OP_SIZE {
		ISA[opcode](c, op1, op2, op3)
	} else {
		c.pc += PC_INCREMENT
	}

	// Reg 0 deve ser sempre 0
	c.reg[0] = 0
}

func (c *cpu) Exec_program() {
	for !c.halt {
		word := c.mem[c.pc : c.pc+PC_INCREMENT]
		inst := ui32fr4ui8([4]uint8(word))
		c.Exec_inst(inst)
	}
}

func (c *cpu) Load_program(program []uint32) {
	program_size := len(program)
	if program_size*4 > MEM_SIZE {
		return
	}

	for i := range program_size {
		inst := program[i]
		word := ui32to4ui8(inst)

		base_addr := i * PC_INCREMENT
		for j := range word {
			c.mem[base_addr+j] = word[j]
		}
	}
}
