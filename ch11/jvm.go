package main

import (
	"jvmgo/ch11/rtda/heap"
	"jvmgo/ch11/rtda"
	"jvmgo/ch11/classpath"
	"jvmgo/ch11/instructions/base"
)

type JVM struct {
	cmd *Cmd
	classLoader *heap.ClassLoader
	mainThread *rtda.Thread
}

func newJVM(cmd *Cmd) *JVM {
	cp := classpath.Parse(cmd.XjreOption,cmd.cpOption)
	classLoader := heap.NewClassLoader(cp, cmd.verboseClassFlag)

	return &JVM{
		cmd: cmd,
		classLoader:classLoader,
		mainThread:rtda.NewThread(),
	}
}

func (jvm *JVM) start()  {
	jvm.initVM()
	jvm.execMain()
}

func (jvm *JVM) initVM()  {
	vmClass := jvm.classLoader.LoadClass("sun/misc/VM")

	base.InitClass(jvm.mainThread,vmClass)
	interpret(jvm.mainThread, jvm.cmd.verboseClassFlag)
}