package heap

import "jvmgo/ch06/classfile"

type Method struct {
	ClassMember
	maxStack uint
	maxLocals uint
	code []byte
}


func newMethods(class *Class, cfMethods []*classfile.MemberInfo) []*Method {
	methods := make([]*Method ,len(cfMethods))
	for i, cfMethod := range cfMethods {
		methods[i] = &Method{}
		methods[i].class = class
		methods[i].copyMemberInfo(cfMethod)
		methods[i].copyAttributes(cfMethod)
	}
	return methods
}

func (method *Method) copyAttributes(cfMethod *classfile.MemberInfo) {
	if codeAttr := cfMethod.CodeAttribute(); codeAttr != nil {
		method.maxStack = codeAttr.MaxStack()
		method.maxLocals = codeAttr.MaxLocals()
		method.code = codeAttr.Code()
	}
}
func (method *Method) MaxLocals() uint {
	return method.maxLocals
}
func (method *Method) MaxStack() uint {
	return method.maxStack
}
