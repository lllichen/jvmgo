package heap

import "jvmgo/ch09/classfile"

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
		methods[i] = newMethod(class, cfMethod)
	}
	return methods
}

func newMethod(class *Class, info *classfile.MemberInfo) *Method {
	method := &Method{}
	method.class = class
	method.copyMemberInfo(info)
	method.copyAttributes(info)
	md := parseMethodDescriptor(method.descriptor)
	method.calcArgSlotCount(md.parameterTypes)
	if method.IsNative() {
		method.injectCodeAttribute(md.returnType)
	}
	return method
}

func (method *Method) injectCodeAttribute(returnType string) {
	method.maxStack = 4
	method.maxLocals = method.argSlotCount
	switch returnType[0] {
	case 'V':
		method.code = []byte{0xfe, 0xb1} //return
	case 'D':
		method.code = []byte{0xfe, 0xaf} //dreturn
	case 'F':
		method.code = []byte{0xfe, 0xae} //freturn
	case 'J':
		method.code = []byte{0xfe, 0xad} //lreturn
	case 'L':
		method.code = []byte{0xfe, 0xb0} //areturn
	default:
		method.code = []byte{0xfe, 0xac} //ireturn
	}
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

func (method *Method) IsNative() bool {
	return 0 != method.accessFlags&ACC_NATIVE
}

func (method *Method) ArgSlotCount() uint {
	return method.argSlotCount
}
func (method *Method) calcArgSlotCount(paramTypes []string) {
	for _, paramType := range paramTypes{
		method.argSlotCount++
		if paramType == "J" || paramType == "D" {
			method.argSlotCount++
		}
	}
	if !method.IsStatic() {
		method.argSlotCount++
	}
}


