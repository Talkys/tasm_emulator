package assembler

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	FOUR_BYTES  = 8 * 4
	THREE_BYTES = 8 * 3
	TWO_BYTES   = 8 * 2
	ONE_BYTE    = 8 * 1
)

var (
	opcode_map = map[string]string{
		"nop":  "0x00",
		"li":   "0x01",
		"and":  "0x02",
		"or":   "0x03",
		"xor":  "0x04",
		"not":  "0x05",
		"nand": "0x06",
		"nor":  "0x07",
		"add":  "0x0A",
		"sub":  "0x0B",
		"addi": "0x0C",
		"subi": "0x0D",
		"sll":  "0x0E",
		"srl":  "0x0F",
		"beq":  "0x1A",
		"lw":   "0x1B",
		"sw":   "0x1C",
		"sys":  "0x1D",
		"jmp":  "0x1E",
		"hlt":  "0x1F",
	}

	reg_map = map[string]string{
		"$0": "0x00",
		"$1": "0x01",
		"$2": "0x02",
		"$3": "0x03",
		"$4": "0x04",
		"$5": "0x05",
		"$6": "0x06",
	}
)

func ui32fr4ui8(num [4]uint8) uint32 {
	output := uint32(num[0]) << THREE_BYTES
	output += uint32(num[1]) << TWO_BYTES
	output += uint32(num[2]) << ONE_BYTE
	output += uint32(num[3])

	return output
}

func parse_op(op_str string) uint8 {
	var op uint8
	if strings.Contains(op_str, "x") {
		op_str = strings.ReplaceAll(op_str, "x", "")
		res, _ := strconv.ParseInt(op_str, 16, 8)
		op = uint8(res)
	} else {
		res, _ := strconv.ParseInt(op_str, 10, 8)
		op = uint8(res)
	}
	return op
}

func Assemble(program []string) ([]uint32, []uint32) {

	code := []uint32{}

	program = pre_process_macros(program)
	program, jmp_table := label_macros(program)

	for l, line := range program {
		for opcode, value := range opcode_map {
			line = strings.Replace(line, opcode, value, -1)
		}

		for reg, ridx := range reg_map {
			line = strings.Replace(line, reg, ridx, -1)
		}
		program[l] = line
	}

	for _, line := range program {
		args := strings.Split(line, ",")
		if len(args) != 4 {
			panic("Linha inválida {" + line + "}\n")
		}

		opcode := parse_op(args[0])
		op1 := parse_op(args[1])
		op2 := parse_op(args[2])
		op3 := parse_op(args[3])

		inst := ui32fr4ui8([4]uint8{opcode, op1, op2, op3})
		code = append(code, inst)
	}
	return code, jmp_table
}

func pre_process_macros(program []string) []string {
	replace_map := map[string]string{
		"nop": "nop,0,0,0",
		"hlt": "hlt,0,0,0",
	}

	for l, line := range program {
		line = strings.ReplaceAll(line, " ", "")
		for key, value := range replace_map {
			line = strings.ReplaceAll(line, key, value)
		}
		program[l] = line
	}

	return program
}

func label_macros(program []string) ([]string, []uint32) {
	label_map := map[string]int{}
	jmp_table := []uint32{}

	//Inserir endereços das labels e trocar elas por nop
	for l, line := range program {
		if line[0] == ':' {
			jmp_table = append(jmp_table, uint32(4*l))
			label_map[line] = len(jmp_table) - 1
			program[l] = "nop,0,0,0"
		}
	}

	for l, line := range program {
		for key, value := range label_map {
			if strings.Contains(line, key) {
				op1 := uint8((value >> TWO_BYTES) & 0xFF)
				op2 := uint8((value >> ONE_BYTE) & 0xFF)
				op3 := uint8((value) & 0xFF)

				inst := fmt.Sprintf(",0x%02X,", op1)
				inst += fmt.Sprintf("0x%02X,", op2)
				inst += fmt.Sprintf("0x%02X", op3)

				// Em teoria a label deve virar 0xXX,0xYY,0xZZ
				program[l] = strings.ReplaceAll(program[l], key, inst)
			}
		}
	}

	return program, jmp_table
}
