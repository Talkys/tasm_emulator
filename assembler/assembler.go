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

func Assemble(program []string) []uint32 {

	code := []uint32{}

	for l, line := range program {
		for opcode, value := range opcode_map {
			line = strings.Replace(line, opcode, value, -1)
		}

		for reg, ridx := range reg_map {
			line = strings.Replace(line, reg, ridx, -1)
		}
		program[l] = line
		fmt.Println(program[l])
	}

	for _, line := range program {
		args := strings.Split(line, ",")
		if len(args) != 4 {
			panic("Linha inválida " + line + "\n")
		}

		op_str := args[0]
		var opcode uint8
		if strings.Contains(op_str, "x") {
			op_str = strings.ReplaceAll(op_str, "x", "")
			res, _ := strconv.ParseInt(op_str, 16, 8)
			opcode = uint8(res)
		} else {
			res, _ := strconv.ParseInt(op_str, 10, 8)
			opcode = uint8(res)
		}

		op1_str := args[1]
		var op1 uint8
		if strings.Contains(op1_str, "x") {
			op1_str = strings.ReplaceAll(op1_str, "x", "")
			res, _ := strconv.ParseInt(op1_str, 16, 8)
			op1 = uint8(res)
		} else {
			res, _ := strconv.ParseInt(op1_str, 10, 8)
			op1 = uint8(res)
		}

		op2_str := args[2]
		var op2 uint8
		if strings.Contains(op2_str, "x") {
			op2_str = strings.ReplaceAll(op2_str, "x", "")
			res, _ := strconv.ParseInt(op2_str, 16, 8)
			op2 = uint8(res)
		} else {
			res, _ := strconv.ParseInt(op2_str, 10, 8)
			op2 = uint8(res)
		}

		op3_str := args[3]
		var op3 uint8
		if strings.Contains(op3_str, "x") {
			op3_str = strings.ReplaceAll(op3_str, "x", "")
			res, _ := strconv.ParseInt(op3_str, 16, 8)
			op3 = uint8(res)
		} else {
			res, _ := strconv.ParseInt(op3_str, 10, 8)
			op3 = uint8(res)
		}

		inst := ui32fr4ui8([4]uint8{opcode, op1, op2, op3})
		code = append(code, inst)
	}
	return code
}
