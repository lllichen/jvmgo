package main

import (
	"jvmgo/ch05/classfile"
	"jvmgo/ch05/rtda"
	"fmt"
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/instructions"
)

func interpret(info *classfile.MemberInfo)  {
	codeAttr := info.CodeAttribute()
	maxLocals := codeAttr.MaxLocals()
	maxStack := codeAttr.MaxStack()
	bytecode := codeAttr.Code()

	thread := rtda.NewThread()
	frame := thread.NewFrame(maxLocals,maxStack)
	thread.PushFrame(frame)
	defer catchErr(frame)

	loop(thread, bytecode)

}

func catchErr(frame *rtda.Frame) {
	if r := recover(); r != nil {
		fmt.Printf("LocalVars: %v\n",frame.LocalVars())
		fmt.Printf("OperandStack:%v\n", frame.OperandStack())
		panic(r)
	}
}

func loop(thread *rtda.Thread, bytecode []byte){
	frame := thread.PopFrame()
	reader := &base.ByteCodeReader{}
	for {
		pc := frame.NextPc()
		thread.SetPc(pc)

		//decode
		reader.Reset(bytecode,pc)
		opcode := reader.ReadUint8()
		inst := instructions.NewInstruction(opcode)

		//execute
		fmt.Printf("pc:%2d ,inst:%T %v\n",pc,inst, inst)
		inst.Execute(frame)
	}
}

