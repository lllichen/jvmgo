package control

import (
	"jvmgo/ch10/instructions/base"
	"jvmgo/ch10/rtda"
)

type LOOKUP_SWITCH struct {
	defaultOffset int32
	npairs  int32
	matchOffsets []int32
}

func (lookSwitch *LOOKUP_SWITCH) FetchOperands(reader *base.ByteCodeReader) {
	reader.SkipPadding()
	lookSwitch.defaultOffset = reader.ReadInt32()
	lookSwitch.npairs = reader.ReadInt32()
	lookSwitch.matchOffsets = reader.ReadInt32s(lookSwitch.npairs*2)
}

func (lookSwitch *LOOKUP_SWITCH) Execute(frame *rtda.Frame) {
	key := frame.OperandStack().PopInt()
	for i := int32(0); i< lookSwitch.npairs*2; i+=2 {
		if lookSwitch.matchOffsets[i] == key {
			offset := lookSwitch.matchOffsets[i+1]
			base.Branch(frame, int(offset))
			return
		}
	}
	base.Branch(frame,int(lookSwitch.defaultOffset))
}



