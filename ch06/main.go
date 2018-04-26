package main

import "fmt"
import (
	"jvmgo/ch05/classpath"
	"strings"
	"jvmgo/ch05/classfile"
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
	cp := classpath.Parse(cmd.XjreOption, cmd.cpOption)
	//fmt.Printf("classpath:%v class:%v args:%v\n",
	//	cp, cmd.class, cmd.args)
	//
	className := strings.Replace(cmd.class, ".", "/", -1)
	//classData, _, err := cp.ReadClass(className)
	//if err != nil {
	//	fmt.Printf("Could not find or load main class %s\n", cmd.class)
	//	return
	//}
	////print data
	//fmt.Printf("class data:%v\n", classData)
	//
	cf := loadClass(className,cp)
	mainMethod := getMainMethod(cf)
	if mainMethod != nil {
		//interpret(mainMethod)
	}else {
		fmt.Printf("Main method not found in class %s\n",cmd.class)
	}


}
func loadClass(className string, cp *classpath.Classpath) *classfile.ClassFile {
	classData,_, err := cp.ReadClass(className)
	if err != nil {
		panic(err)
	}
	cf, err := classfile.Parse(classData)
	if err != nil {
		panic(err)
	}
	return cf
}

func getMainMethod(cf *classfile.ClassFile) *classfile.MemberInfo{
	for _,m := range cf.Methods(){
		if m.Name() == "main" && m.Descriptor() == "[Ljava/lang/String;)V" {
			return m
		}
	}
	return nil
}

