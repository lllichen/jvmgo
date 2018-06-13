package main

import (
	"jvmgo/ch07/rtda"
	"fmt"
	"jvmgo/ch07/instructions/base"
	"jvmgo/ch07/instructions"
	"jvmgo/ch07/rtda/heap"
)

func interpret(method *heap.Method, logInst bool)  {
	thread := rtda.NewThread()
	frame := thread.NewFrame(method)
	thread.PushFrame(frame)
	defer catchErr(thread)

	loop(thread,logInst)

}

func catchErr(thread *rtda.Thread) {
	if r := recover(); r != nil {
		logFrames(thread)
		panic(r)
	}
}

func logFrames(thread *rtda.Thread) {
	for !thread.IsStackEmpty() {
		frame := thread.PopFrame()
		method := frame.Method()
		className := method.Class().Name()
		fmt.Printf(">> pc:%4d %v.%v%v \n", frame.NextPC(), className, method.Name(), method.Descriptor())
	}
}

func logInstruction(frame *rtda.Frame, inst base.Instruction){
	method := frame.Method()
	className := method.Class().Name()
	methodName := method.Name()
	pc := frame.Thread().PC()
	fmt.Printf("%v.%v() #%2d %T %v\n",className, methodName,pc,inst,inst )
}

func loop(thread *rtda.Thread, logInst bool){
	reader := &base.ByteCodeReader{}
	for {
		frame := thread.CurrentFrame()
		pc := frame.NextPC()
		thread.SetPc(pc)

		//decode
		reader.Reset(frame.Method().Code(),pc)
		opcode := reader.ReadUint8()
		inst := instructions.NewInstruction(opcode)
		inst.FetchOperands(reader)
		frame.SetNextPC(reader.PC())

		if logInst {
			logInstruction(frame, inst)
		}
		//execute
		inst.Execute(frame)
		if thread.IsStackEmpty(){
			break
		}
	}
}
