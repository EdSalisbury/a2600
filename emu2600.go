package main

import (
	"fmt"
	"log"
	"os"

	"github.com/edsalisbury/a2600/mos6502"
)

// SR Flags (bit 7 to bit 0)
// N	Negative
// V	Overflow
// -	ignored
// B	Break
// D	Decimal (use BCD for arithmetics)
// I	Interrupt (IRQ disable)
// Z	Zero
// C	Carry

func getAbs(mem *map[uint16]uint8, pc *uint16) uint16 {
	addr := uint16((*mem)[*pc+1])<<8 + uint16((*mem)[*pc])
	fmt.Printf("$%04x", addr)
	*pc += 2
	return addr
}

func getImm(mem *map[uint16]uint8, pc *uint16) uint8 {
	val := (*mem)[*pc]
	fmt.Printf("#%02x", val)
	*pc += 1
	return val
}

func getZpgX(mem *map[uint16]uint8, x *uint8, pc *uint16) uint16 {
	addr := uint16(*x) + uint16((*mem)[*pc])
	fmt.Printf("$%02x,X", (*mem)[*pc])
	*pc += 1
	return addr
}

func getRel(mem *map[uint16]uint8, pc *uint16) uint8 {
	addr := (*mem)[*pc] ^ 0xff + 1
	fmt.Printf("$%02x", addr)
	*pc += 1
	return addr
}

// Is this the correct order?
func pushUint16(mem *map[uint16]uint8, addr uint16, sp *uint8) {
	top := 0x0100 + uint16(*sp)
	(*mem)[top] = uint8(addr)
	(*mem)[top-1] = uint8(addr >> 8)
	*sp -= 2
}

// Is this the correct order?
func pullUint16(mem *map[uint16]uint8, sp *uint8) uint16 {
	top := 0x0100 + uint16(*sp+1)
	addr := uint16((*mem)[top])<<8 + uint16((*mem)[top+1])
	*sp += 2
	return addr
}

func main() {
	program, err := os.ReadFile("combat.bin")
	if err != nil {
		log.Fatal(err)
	}
	mem := make(map[uint16]uint8)
	for i := 0; i < len(program); i++ {
		address := 0xf000 + uint16(i)
		mem[address] = program[i]
	}

	var clk uint64 = 0
	var pc uint16 = 0xf000
	var ac uint8 = 0
	var x uint8 = 0
	var y uint8 = 0
	var sr uint8 = 0
	var sp uint8 = 0

	done := false
	for !done {
		code := mem[pc]
		opcode := mos6502.Opcodes[code]
		fmt.Printf("CLK: %08d PC: %04x AC: %02x X: %02x Y: %02x SR: %08b SP: %02x | %02x %s ", clk, pc, ac, x, y, sr, sp, code, opcode.Inst)
		pc++
		switch code {
		case 0x20: // JSR
			addr := getAbs(&mem, &pc)
			pushUint16(&mem, pc, &sp)
			pc = addr
		case 0x60: // RTS
			pc = pullUint16(&mem, &sp)
		case 0x78: // SEI
			sr |= 0b00000100
		case 0x8d: // STA
			addr := getAbs(&mem, &pc)
			mem[addr] = ac
		case 0x95: // STA
			addr := getZpgX(&mem, &x, &pc)
			mem[addr] = ac
		case 0x9a: // TXS
			sp = x
		case 0xa2: // LDX
			x = getImm(&mem, &pc)
		case 0xa9: // LDA
			ac = getImm(&mem, &pc)
		case 0xd0: // BNE
			addr := getRel(&mem, &pc)
			if sr&0b00000010 == 0 {
				pc -= uint16(addr)
			}
		case 0xd8: // CLD
			sr &= 0b11110111
		case 0xe8: // INX
			x++
			if x == 0 {
				sr |= 0b00000010
			}
		default:
			done = true
		}
		if opcode.Cycles == 0 {
			fmt.Println("\nNo cycles for instruction!")
			done = true
		}
		clk += uint64(opcode.Cycles)
		if opcode.Mode == mos6502.Impl {
			fmt.Printf("\t")
		}

		fmt.Printf("\t; STK: ")
		for i := 0x0100; i <= 0x01ff; i++ {
			if mem[uint16(i)] != 0 {
				fmt.Printf("%02x", mem[uint16(i)])
			}
		}
		fmt.Println()
	}

}
