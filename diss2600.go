package main

import (
	"fmt"
	"log"
	"os"

	"github.com/edsalisbury/a2600/mos6502"
)

func main() {
	content, err := os.ReadFile("combat.bin")
	if err != nil {
		log.Fatal(err)
	}

	data := false

	for i := 0; i < len(content); i++ {
		hex := content[i]

		if data {
			fmt.Printf("%04x: %02x\tUNK\t\t; %08b\n", i, hex, hex)
			continue
		}
		if opcode, ok := mos6502.Opcodes[hex]; ok {
			switch mode := opcode.Mode; mode {
			case mos6502.Impl:
				fmt.Printf("%04x: %02x\t%s\t\t; %s\n", i, hex, opcode.Inst, opcode.Desc)
			case mos6502.Imm:
				i++
				addr := content[i]
				fmt.Printf("%04x: %02x\t%s #$%02x\t; %s\n", i, hex, opcode.Inst, addr, opcode.Desc)
			case mos6502.Rel:
				i++
				addr := content[i]
				fmt.Printf("%04x: %02x\t%s $%02x\t\t; %s\n", i, hex, opcode.Inst, addr, opcode.Desc)
			case mos6502.Abs:
				i++
				lo := content[i]
				i++
				hi := content[i]
				fmt.Printf("%04x: %02x\t%s $%02x%02x\t; %s\n", i, hex, opcode.Inst, hi, lo, opcode.Desc)
			case mos6502.AbsX:
				i++
				lo := content[i]
				i++
				hi := content[i]
				fmt.Printf("%04x: %02x\t%s $%02x%02x,X\t; %s\n", i, hex, opcode.Inst, hi, lo, opcode.Desc)
			case mos6502.AbsY:
				i++
				lo := content[i]
				i++
				hi := content[i]
				fmt.Printf("%04x: %02x\t%s $%02x%02x,Y\t; %s\n", i, hex, opcode.Inst, hi, lo, opcode.Desc)
			case mos6502.Zpg:
				i++
				addr := content[i]
				fmt.Printf("%04x: %02x\t%s $00%02x\t; %s\n", i, hex, opcode.Inst, addr, opcode.Desc)
			case mos6502.ZpgX:
				i++
				addr := content[i]
				fmt.Printf("%04x: %02x\t%s $00%02x,X\t; %s\n", i, hex, opcode.Inst, addr, opcode.Desc)
			case mos6502.ZpgY:
				i++
				addr := content[i]
				fmt.Printf("%04x: %02x\t%s $00%02x,Y\t; %s\n", i, hex, opcode.Inst, addr, opcode.Desc)
			case mos6502.IndY:
				i++
				addr := content[i]
				fmt.Printf("%04x: %02x\t%s ($%02x),Y\t; %s\n", i, hex, opcode.Inst, addr, opcode.Desc)

			}
		} else {
			fmt.Printf("%04x: %02x\tUNK\t\t; %08b\n", i, hex, hex)
			data = true
			//os.Exit(1)
		}
	}

}
