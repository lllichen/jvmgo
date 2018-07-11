package heap

import (
	"jvmgo/ch11/classfile"
)

type Method struct {
	ClassMember
	maxStack                uint
	maxLocals               uint
	code                    []byte
	exceptionTable          ExceptionTable // todo: rename
	lineNumberTable         *classfile.LineNumberTableAttribute
	exceptions              *classfile.ExceptionsAttribute // todo: rename
	parameterAnnotationData []byte                         // RuntimeVisibleParameterAnnotations_attribute
	annotationDefaultData   []byte                         // AnnotationDefault_attribute
	parsedDescriptor        *MethodDescriptor
	argSlotCount            uint
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
		method.lineNumberTable = codeAttr.LineNumberTableAttribute()
		method.exceptionTable = newExceptionTable(codeAttr.ExceptionTable(), method.class.constantPool)
	}
}

func (method *Method) GetLineNumber(pc int) int {
	if method.IsNative() {
		return -2
	}
	if method.lineNumberTable == nil {
		return -1
	}
	return method.lineNumberTable.GetLineNumber(pc)
}

func (method *Method) FindExceptionHandler(exClass *Class, pc int) int{
	handler := method.exceptionTable.findExceptionHandler(exClass,pc)
	if handler != nil {
		return handler.handlerPc
	}
	return -1
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

func (method *Method) ParameterAnnotationData() []byte {
	return method.parameterAnnotationData
}
func (method *Method) AnnotationDefaultData() []byte {
	return method.annotationDefaultData
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



func (method *Method) isConstructor() bool {
	return !method.IsStatic() && method.name == "<init>"
}


// reflection
func (method *Method) ParameterTypes() []*Class {
	if method.argSlotCount == 0 {
		return nil
	}

	paramTypes := method.parsedDescriptor.parameterTypes
	paramClasses := make([]*Class, len(paramTypes))
	for i, paramType := range paramTypes {
		paramClassName := toClassName(paramType)
		paramClasses[i] = method.class.loader.LoadClass(paramClassName)
	}

	return paramClasses
}

func (method *Method) ReturnType() *Class {
	returnType := method.parsedDescriptor.returnType
	returnClassName := toClassName(returnType)
	return method.class.loader.LoadClass(returnClassName)
}


func (method *Method) ExceptionTypes() []*Class {
	if method.exceptions == nil {
		return nil
	}

	exIndexTable := method.exceptions.ExceptionIndexTable()
	exClasses := make([]*Class, len(exIndexTable))
	cp := method.class.constantPool

	for i, exIndex := range exIndexTable {
		classRef := cp.GetConstant(uint(exIndex)).(*ClassRef)
		exClasses[i] = classRef.ResolvedClass()
	}

	return exClasses
}

func (method *Method) isClinit() bool {
	return method.IsStatic() && method.name == "<clinit>"
}
