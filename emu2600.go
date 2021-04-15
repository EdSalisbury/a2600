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

	var top uint16 = 0x01ff
	done := false
	for !done {
		code := mem[pc]
		opcode := mos6502.Opcodes[code]
		fmt.Printf("STK: ")
		for i := 0x0100; i <= 0x01ff; i++ {
			if mem[uint16(i)] != 0 {
				fmt.Printf("%02x", mem[uint16(i)])
			}
		}
		fmt.Printf("\n")
		fmt.Printf("CLK: %08d TOP: %04x PC: %04x AC: %02x X: %02x Y: %02x SR: %08b SP: %02x | %02x %s ", clk, top, pc, ac, x, y, sr, sp, code, opcode.Inst)

		switch code {
		case 0x20: // JSR
			mem[top] = uint8(pc + 3)
			top--
			sp--
			mem[top] = uint8((pc + 3) >> 8)
			top--
			sp--
			pc++
			lo := mem[pc]
			pc++
			hi := mem[pc]
			fmt.Printf("%02x%02x", hi, lo)
			pc = (uint16(hi) << 8) + uint16(lo) - 1
		case 0x60: // RTS
			var addr uint16
			top++
			sp++
			addr = uint16(mem[top]) << 8
			top++
			sp++
			addr += uint16(mem[top]) - 1
			pc = addr
		case 0x78: // SEI
			sr |= 0b00000100
		case 0x95: // STA
			pc++
			fmt.Printf("$%02x,X", mem[pc])
			addr := uint16(x) + uint16(mem[pc])
			fmt.Printf(" ; mem[%04x] = %02x", addr, ac)
			mem[addr] = ac
		case 0x9a: // TXS
			sp = x
		case 0xa2: // LDX
			pc++
			fmt.Printf("#%02x", mem[pc])
			x = mem[pc]
		case 0xa9: // LDA
			pc++
			fmt.Printf("#%02x", mem[pc])
			ac = mem[pc]
		case 0xd0: // BNE
			pc++
			fmt.Printf("#%02x", mem[pc])
			addr := mem[pc] ^ 0xff + 1
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
		clk += uint64(opcode.Cycles)
		pc++
		fmt.Println()
	}

	//var operand uint16 = 0
	//if opcode, ok := mos6502.Opcodes[code]; ok {

	// switch mode := opcode.Mode; mode {
	// case mos6502.Imm:
	// 	pc++
	// 	operand = uint16(mem[pc])
	// case mos6502.Rel:
	// 	pc++
	// 	operand = uint16(mem[pc])
	// case mos6502.Abs:
	// 	pc++
	// 	lo := uint16(mem[pc])
	// 	pc++
	// 	hi := uint16(mem[pc])
	// 	operand = (hi * 256) + lo
	// case mos6502.AbsX:
	// 	pc++
	// 	lo := uint16(mem[pc])
	// 	pc++
	// 	hi := uint16(mem[pc])
	// 	operand = (hi * 256) + lo
	// case mos6502.AbsY:
	// 	pc++
	// 	lo := uint16(mem[pc])
	// 	pc++
	// 	hi := uint16(mem[pc])
	// 	operand = (hi * 256) + lo
	// case mos6502.Zpg:
	// 	pc++
	// 	operand = uint16(mem[pc])
	// case mos6502.ZpgX:
	// 	pc++
	// 	operand = uint16(mem[pc])
	// case mos6502.ZpgY:
	// 	pc++
	// 	operand = uint16(mem[pc])
	// case mos6502.IndY:
	// 	pc++
	// 	operand = uint16(mem[pc])
	// }

	//}
	// opcode := mos6502.Opcodes[code].Inst

}
