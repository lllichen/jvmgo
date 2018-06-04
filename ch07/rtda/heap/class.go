package heap

import (
	"jvmgo/ch07/classfile"
	"strings"
)

type Class struct {
	accessFlags uint16
	name string
	superClassName string
	interfaceNames []string
	constantsPool *ConstantPool
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
	class.interfaceNames = file.InterfaceNames()
	class.constantsPool = newConstantPool(class,file.ConstantPool())
	class.fields = newFields(class,file.Fields())
	class.methods = newMethods(class,file.Methods())
	return class
}

func (class *Class) IsPublic() bool{
	return 0 != class.accessFlags&ACC_PUBLIC
}

func (class *Class) IsFinal() bool{
	return 0 != class.accessFlags&ACC_FINAL
}


func (class *Class) IsSuper() bool{
	return 0 != class.accessFlags&ACC_SUPER
}


func (class *Class) IsInterface() bool{
	return 0 != class.accessFlags&ACC_INTERFACE
}

func (class *Class) IsAbstract() bool{
	return 0 != class.accessFlags&ACC_ABSTRACT
}


func (class *Class) IsSynthetic() bool{
	return 0 != class.accessFlags&ACC_SYNTHETIC
}


func (class *Class) IsAnnotation() bool{
	return 0 != class.accessFlags&ACC_ANNOTATION
}


func (class *Class) IsEnum() bool{
	return 0 != class.accessFlags&ACC_ENUM
}


func (class *Class) isAccessibleTo(other *Class) bool {
	return class.IsPublic() || class.getPackageName() == other.getPackageName()
}

func (class *Class) getPackageName() string{
	if i := strings.LastIndex(class.name,"/"); i>= 0 {
		return class.name[:1]
	}
	return ""
}
func (class *Class) ConstantPool() *ConstantPool {
	return class.constantsPool
}
func (class *Class) NewObject() *Object {
	return newObject(class)
}
func  newObject(class *Class) *Object {
	return &Object{
		class : class,
		fields: newSlots(class.instanceSlotCount),
	}
}

func (class *Class) StaticVars() Slots {
	return class.staticVars
}
func (class *Class) GetMainMethod() *Method {
	return class.getStaticMethod("main", "([Ljava/lang/String;)V")
}
func (class *Class) getStaticMethod(name string, descriptor string) *Method {
	for _, method := range class.methods {
		if method.IsStatic() && method.name == name && method.descriptor == descriptor {
			return method
		}
	}
	return nil
}


