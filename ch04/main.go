package main

import "fmt"
import (
	"jvmgo/ch04/rtda"
)

func main() {
	cmd := parseCmd()

	if cmd.versionFlag {
		fmt.Println("version 0.0.1")
	} else if cmd.helpFlag || cmd.class == "" {
		printUsage()
	} else {
		startJVM(cmd)
	}
}

func startJVM(cmd *Cmd) {
	//cp := classpath.Parse(cmd.XjreOption, cmd.cpOption)
	//fmt.Printf("classpath:%v class:%v args:%v\n",
	//	cp, cmd.class, cmd.args)
	//
	//className := strings.Replace(cmd.class, ".", "/", -1)
	//classData, _, err := cp.ReadClass(className)
	//if err != nil {
	//	fmt.Printf("Could not find or load main class %s\n", cmd.class)
	//	return
	//}
	////print data
	//fmt.Printf("class data:%v\n", classData)
	//
	//cf := loadClass(className,cp)
	//fmt.Println(cmd.class)
	//printClassInfo(cf)
	//fmt.Println(cf.MajorVersion())
	//fmt.Println(cf.AccessFlags())
	//fmt.Println(cf.ThisClass())
	//fmt.Println(cf.SuperClass())
	//fmt.Println(cf.ConstantPool())

	frame := rtda.NewFrame(100,100)
	testLocalVars(frame.LocalVars())
	testOperandVars(frame.OperandStack())
}


func testLocalVars(vars rtda.LocalVars){
	vars.SetInt(0,100)
	vars.SetInt(1,-100)

	vars.SetLong(2,2997924580)
	vars.SetLong(4,-2997924580)

	vars.SetFloat(6,3.1415926)
	vars.SetDouble(7,2.71828182845)
	vars.SetRef(9,nil)

	println(vars.GetInt(0))
	println(vars.GetInt(1))
	println(vars.GetLong(2))
	println(vars.GetLong(4))
	println(vars.GetFloat(6))
	println(vars.GetDouble(7))
	println(vars.GetRef(9))

}



func testOperandVars(vars rtda.OperandStack){
	vars.PushInt(100)
	vars.PushInt(-100)

	vars.PushLong(2997924580)
	vars.PushLong(-2997924580)

	vars.PushFloat(3.1415926)
	vars.PushDouble(2.71828182845)
	vars.PushRef(nil)

	println(vars.PopRef())
	println(vars.PopDouble())
	println(vars.PopFloat())
	println(vars.PopLong())
	println(vars.PopLong())
	println(vars.PopInt())
	println(vars.PopInt())






}
