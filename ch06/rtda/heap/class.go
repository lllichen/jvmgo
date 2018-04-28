package heap

import "jvmgo/ch06/classfile"

type Class struct {
	accessFlags uint16
	name string
	superClassName string
	interfaceNames []string
	constantsPool *classfile.ConstantPool
	fields []*Field
	methods []*Method
	//loader *ClassLoader
	superClass *Class
	interfaces []*Class
	instanceSlotCount uint
	staticSlotCount uint
	//staticVars *Slots
}


func newClass(file *classfile.ClassFile) *Class {
	class := &Class{}
	class.accessFlags = file.AccessFlags()
	class.name = file.ClassName()
	class.superClassName = file.SuperClassName()
	class.interfaces = file.InterfaceNames()
	class.constantsPool = newConstantPool(class,file.ConstantPool())
	class.fields = newFields(class,file.Fields())
	class.methods = newMethods(class,file.Methods())
	return class
}

