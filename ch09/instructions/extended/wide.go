package extended

import (
	"jvmgo/ch09/instructions/base"
	"jvmgo/ch09/rtda"
	"jvmgo/ch09/instructions/loads"
	"jvmgo/ch09/instructions/math"
	"jvmgo/ch09/instructions/stores"
)

type WIDE struct {
	modifiedInstruction base.Instruction
}

func (wide *WIDE) FetchOperands(reader *base.ByteCodeReader) {
	opcode := reader.ReadUint8()
	switch opcode {
	case 0x15 :
		inst := &loads.ILOAD{}
		inst.Index = uint(reader.ReadUint16())
		wide.modifiedInstruction = inst
	case 0x16 :
		inst := &loads.LLOAD{}
		inst.Index = uint(reader.ReadUint16())
		wide.modifiedInstruction = inst
	case 0x17 :
		inst := &loads.FLOAD{}
		inst.Index = uint(reader.ReadUint16())
		wide.modifiedInstruction = inst
	case 0x18 :
		inst := &loads.DLOAD{}
		inst.Index = uint(reader.ReadUint16())
		wide.modifiedInstruction = inst
	case 0x19 :
		inst := &loads.ALOAD{}
		inst.Index = uint(reader.ReadUint16())
		wide.modifiedInstruction = inst
	case 0x36 :
		inst := &stores.ISTORE{}
		inst.Index = uint(reader.ReadUint16())
		wide.modifiedInstruction = inst
	case 0x37 :
		inst := &stores.LSTORE{}
		inst.Index = uint(reader.ReadUint16())
		wide.modifiedInstruction = inst
	case 0x38 :
		inst := &stores.FSTORE{}
		inst.Index = uint(reader.ReadUint16())
		wide.modifiedInstruction = inst
	case 0x39 :
		inst := &stores.DSTORE{}
		inst.Index = uint(reader.ReadUint16())
		wide.modifiedInstruction = inst
	case 0x3a :
		inst := &stores.ASTORE{}
		inst.Index = uint(reader.ReadUint16())
		wide.modifiedInstruction = inst
	case 0x84 :
		inst := &math.IINC{}
		inst.Index = uint(reader.ReadInt16())
		inst.Const = int32(reader.ReadInt16())
		wide.modifiedInstruction = inst
	case 0xa8 :
		panic("Unsupported opcode: oxa9!")
	}
}


func (wide *WIDE) Execute(frame *rtda.Frame) {
	wide.modifiedInstruction.Execute(frame)
}




