package main

import (
	"jvmgo/ch07/rtda"
	"fmt"
	"jvmgo/ch07/instructions/base"
	"jvmgo/ch07/instructions"
	"jvmgo/ch07/rtda/heap"
)

func interpret(method *heap.Method)  {
	thread := rtda.NewThread()
	frame := thread.NewFrame(method)
	thread.PushFrame(frame)
	defer catchErr(frame)

	loop(thread, method.Code())

}

func catchErr(frame *rtda.Frame) {
	if r := recover(); r != nil {
		//fmt.Printf("LocalVars: %v\n",frame.LocalVars())
		//fmt.Printf("OperandStack:%v\n", frame.OperandStack())
		//panic(r)
	}
}

func loop(thread *rtda.Thread, bytecode []byte){
	frame := thread.PopFrame()
	reader := &base.ByteCodeReader{}
	for {
		pc := frame.NextPC()
		thread.SetPc(pc)

		//decode
		reader.Reset(bytecode,pc)
		opcode := reader.ReadUint8()
		inst := instructions.NewInstruction(opcode)
		inst.FetchOperands(reader)
		frame.SetNextPC(reader.PC())

		//execute
		fmt.Printf("pc:%2d ,inst:%T %v\n",pc,inst, inst)
		inst.Execute(frame)
	}
}

