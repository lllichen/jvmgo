package instructions

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/instructions/constants"
)

var (
	nop = &NOP{}

)

func NewInstruction(opcode byte) base.Instruction {
	switch opcode {
	case 0x00 : return &constants.NOP{}


	}
}
