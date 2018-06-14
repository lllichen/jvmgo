package heap

import "jvmgo/ch08/classfile"

type Method struct {
	ClassMember
	maxStack uint
	maxLocals uint
	code []byte
	argSlotCount uint
}


func newMethods(class *Class, cfMethods []*classfile.MemberInfo) []*Method {
	methods := make([]*Method ,len(cfMethods))
	for i, cfMethod := range cfMethods {
		methods[i] = &Method{}
		methods[i].class = class
		methods[i].copyMemberInfo(cfMethod)
		methods[i].copyAttributes(cfMethod)
		methods[i].calcArgSlotCount()
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
func (method *Method) Code() []byte {
	return method.code
}

func (self *Method) IsNative() bool {
	return 0 != self.accessFlags&ACC_NATIVE
}

func (method *Method) ArgSlotCount() uint {
	return method.argSlotCount
}
func (method *Method) calcArgSlotCount() {
	parsedDescriptor := parseMethodDescriptor(method.descriptor)
	for _, paramType := range parsedDescriptor.parameterTypes {
		method.argSlotCount++
		if paramType == "J" || paramType == "D" {
			method.argSlotCount++
		}
	}
	if !method.IsStatic() {
		method.argSlotCount++
	}
}


