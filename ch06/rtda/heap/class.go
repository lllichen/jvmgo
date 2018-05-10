package heap

import (
	"jvmgo/ch06/classfile"
	"strings"
)

type Class struct {
	accessFlags uint16
	name string
	superClassName string
	interfaceNames []string
	constantsPool *classfile.ConstantPool
	fields []*Field
	methods []*Method
	loader *ClassLoader
	superClass *Class
	interfaces []*Class
	instanceSlotCount uint
	staticSlotCount uint
	staticVars Slots
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

func (class *Class) isAccessibleTo(other *Class) bool {
	return class.isPublic() || class.getPackageName() == other.getPackageName()
}

func (class *Class) getPackageName() string{
	if i := strings.LastIndex(class.name,"/"); i>= 0 {
		return class.name[:1]
	}
	return ""
}
func (class *Class) ConstantPool() *classfile.ConstantPool {
	return class.constantsPool
}
func (class *Class) NewObject() *Object {
	return &Object{
		class : class,
		fields: newSlots(class.instanceSlotCount),
	}
}
func (class *Class) StaticVars() Slots {
	return class.staticVars
}


