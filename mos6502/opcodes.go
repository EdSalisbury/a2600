package mos6502

const (
	A = iota
	Abs
	AbsX
	AbsY
	Imm
	Impl
	Ind
	XInd
	IndY
	Rel
	Zpg
	ZpgX
	ZpgY
)

type Opcode struct {
	Inst   string
	Mode   uint8
	Desc   string
	Cycles uint8
}

var Opcodes = map[byte]Opcode{
	0x00: {Inst: "BRK", Mode: Impl, Desc: "Break / Interrupt"},
	0x05: {Inst: "ORA", Mode: Zpg, Desc: "OR with Accumulator"},
	0x06: {Inst: "ASL", Mode: Zpg, Desc: "Arithmetic Shift Left"},
	0x08: {Inst: "PHP", Mode: Impl, Desc: "PusH Processor status onto stack"},
	0x09: {Inst: "ORA", Mode: Imm, Desc: "OR with Accumulator"},
	0x0a: {Inst: "ASL", Mode: Impl, Desc: "Arithmetic Shift Left"},
	0x0e: {Inst: "ASL", Mode: Abs, Desc: "Arithmetic Shift Left"},
	0x10: {Inst: "BPL", Mode: Rel, Desc: "Branch on PLus"},
	0x18: {Inst: "CLC", Mode: Impl, Desc: "CLear Carry"},
	0x20: {Inst: "JSR", Mode: Abs, Desc: "Jump SubRoutine", Cycles: 6},
	0x24: {Inst: "BIT", Mode: Zpg, Desc: "BIt Test"},
	0x25: {Inst: "AND", Mode: Zpg, Desc: "And"},
	0x29: {Inst: "AND", Mode: Imm, Desc: "And"},
	0x2a: {Inst: "ROL", Mode: Impl, Desc: "ROtate Left"},
	0x30: {Inst: "BMI", Mode: Rel, Desc: "Branch on MInus"},
	0x38: {Inst: "SEC", Mode: Impl, Desc: "SEt Carry"},
	0x45: {Inst: "EOR", Mode: Zpg, Desc: "Exclusive OR"},
	0x49: {Inst: "EOR", Mode: Imm, Desc: "Exclusive OR"},
	0x4a: {Inst: "LSR", Mode: Impl, Desc: "Logical Shift Right"},
	0x4c: {Inst: "JMP", Mode: Abs, Desc: "Jump"},
	0x50: {Inst: "BVC", Mode: Rel, Desc: "Branch on oVerflow Clear"},
	0x60: {Inst: "RTS", Mode: Impl, Desc: "Return from SubRoutine", Cycles: 6},
	0x65: {Inst: "ADC", Mode: Zpg, Desc: "ADd with Carry"},
	0x69: {Inst: "ADC", Mode: Imm, Desc: "ADd with Carry"},
	0x70: {Inst: "BVS", Mode: Rel, Desc: "Branch on oVerflow Set"},
	0x75: {Inst: "ADC", Mode: ZpgX, Desc: "ADd with Carry"},
	0x78: {Inst: "SEI", Mode: Impl, Desc: "Set Interrupt Disable", Cycles: 2},
	0x79: {Inst: "ADC", Mode: AbsY, Desc: "ADd with Carry"},
	0x84: {Inst: "STY", Mode: Zpg, Desc: "Store Y"},
	0x85: {Inst: "STA", Mode: Zpg, Desc: "STore Accumulator"},
	0x86: {Inst: "STX", Mode: Zpg, Desc: "STore X"},
	0x88: {Inst: "DEY", Mode: Impl, Desc: "DEcrement Y"},
	0x8a: {Inst: "TXA", Mode: Impl, Desc: "Transfer X to Accumulator"},
	0x8d: {Inst: "STA", Mode: Abs, Desc: "STore Accumulator"},
	0x90: {Inst: "BCC", Mode: Rel, Desc: "Branch on Carry Clear"},
	0x94: {Inst: "STY", Mode: ZpgX, Desc: "STore Y"},
	0x98: {Inst: "TYA", Mode: Impl, Desc: "Transfer Y to Accumulator"},
	0x99: {Inst: "STA", Mode: AbsY, Desc: "STore Accumulator"},
	0x9a: {Inst: "TXS", Mode: Impl, Desc: "Transfer X to Stack pointer", Cycles: 2},
	0x95: {Inst: "STA", Mode: ZpgX, Desc: "STore Accumulator", Cycles: 4},
	0xa0: {Inst: "LDY", Mode: Imm, Desc: "LoaD Y"},
	0xa2: {Inst: "LDX", Mode: Imm, Desc: "LoaD X", Cycles: 2},
	0xa4: {Inst: "LDY", Mode: Zpg, Desc: "LoaD Y"},
	0xa5: {Inst: "LDA", Mode: Zpg, Desc: "LoaD Accumulator"},
	0xa6: {Inst: "LDX", Mode: Zpg, Desc: "LoaD X"},
	0xa8: {Inst: "TAY", Mode: Impl, Desc: "Transfer Accumulator to Y"},
	0xa9: {Inst: "LDA", Mode: Imm, Desc: "LoaD Accumulator", Cycles: 2},
	0xaa: {Inst: "TAX", Mode: Impl, Desc: "Transfer Accumulator to X"},
	0xad: {Inst: "LDA", Mode: Abs, Desc: "LoaD Accumulator"},
	0xb0: {Inst: "BCS", Mode: Rel, Desc: "Branch on Carry Set"},
	0xb1: {Inst: "LDA", Mode: IndY, Desc: "LoaD Accumulator"},
	0xb5: {Inst: "LDA", Mode: ZpgX, Desc: "LoaD Accumulator"},
	0xb9: {Inst: "LDA", Mode: AbsY, Desc: "LoaD Accumulator"},
	0xba: {Inst: "TSX", Mode: Impl, Desc: "Transfer Stack pointer to X"},
	0xbd: {Inst: "LDA", Mode: AbsX, Desc: "LoaD Accumulator"},
	0xc8: {Inst: "INY", Mode: Impl, Desc: "INcrement Y"},
	0xc9: {Inst: "CMP", Mode: Imm, Desc: "CoMPare with accumulator"},
	0xca: {Inst: "DEX", Mode: Impl, Desc: "DEcrement X"},
	0xd0: {Inst: "BNE", Mode: Rel, Desc: "Branch on Not Equal", Cycles: 2},
	0xd5: {Inst: "CMP", Mode: ZpgX, Desc: "CoMPare with accumulator"},
	0xd6: {Inst: "DEC", Mode: ZpgX, Desc: "DECrement"},
	0xd8: {Inst: "CLD", Mode: Impl, Desc: "Clear Decimal", Cycles: 2},
	0xe0: {Inst: "CPX", Mode: Imm, Desc: "ComPare with X"},
	0xe5: {Inst: "SBC", Mode: Zpg, Desc: "SuBtract with Carry"},
	0xe6: {Inst: "INC", Mode: Zpg, Desc: "INCrement"},
	0xe8: {Inst: "INX", Mode: Impl, Desc: "INcrement X", Cycles: 2},
	0xe9: {Inst: "SBC", Mode: Imm, Desc: "SuBtract with Carry"},
	0xea: {Inst: "NOP", Mode: Impl, Desc: "No OPeration"},
	0xf0: {Inst: "BEQ", Mode: Rel, Desc: "Branch on EQual"},
	0xf6: {Inst: "INC", Mode: ZpgX, Desc: "INCrement"},
	0xf8: {Inst: "SED", Mode: Impl, Desc: "SEt Decimal"},
}
