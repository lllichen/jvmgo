package main

import "jvmgo/ch05/classfile"

func interpret(info *classfile.MemberInfo)  {
	codeAttr := info.CodeAttribute()
}
