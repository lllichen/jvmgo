package heap

import (
	"jvmgo/ch10/classfile"
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
	initStarted bool

	jClass *Object
}

func (class *Class) JClass() *Object{
	return class.jClass
}
func (class *Class) JavaName() string{
	return strings.Replace(class.name,"/",".",-1)
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
	return class.IsPublic() || class.GetPackageName() == other.GetPackageName()
}

func (class *Class) GetPackageName() string{
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
func  NewObject(class *Class) *Object {
	return newObject(class)
}

func (class *Class) StaticVars() Slots {
	return class.staticVars
}
func (class *Class) Name() string{
	return class.name
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


func (class *Class) getMethod(name, descriptor string, isStatic bool) *Method {
	for c := class; c != nil; c = c.superClass {
		for _, method := range c.methods {
			if method.IsStatic() == isStatic &&
				method.name == name &&
				method.descriptor == descriptor {

				return method
			}
		}
	}
	return nil
}



func (class *Class) SuperClass() *Class {
	return class.superClass
}

func (class *Class) GetClinitMethod() *Method {
	return class.getStaticMethod("<clinit>", "()V")
}

func (class *Class) InitStarted() bool {
	return class.initStarted
}


func (class *Class) StartInit() {
	class.initStarted = true
}
func (class *Class) Loader() *ClassLoader {
	return class.loader
}

func (class *Class) ArrayClass() *Class {
	arrayClassName := getArrayClassName(class.name)
	return class.loader.LoadClass(arrayClassName)
}

func (class *Class) isJlObject() bool {
	return class.name == "java/lang/Object"
}
func (class *Class) isJlCloneable() bool {
	return class.name == "java/lang/Cloneable"
}
func (class *Class) isJioSerializable() bool {
	return class.name == "java/io/Serializable"
}

func (class *Class) getField(name, descriptor string, isStatic bool) *Field {
	for c := class; c!= nil; c = c.superClass {
		for _,field := range c.fields {
			if field.IsStatic() == isStatic && field.name == name && field.descriptor == descriptor {
				return field
			}
		}
	}
	return  nil
}


func (class *Class) IsPrimitive() bool {
	_, ok := primitiveTypes[class.name]
	return ok
}


func (class *Class) GetInstanceMethod(name, descriptor string) *Method {
	return class.getMethod(name, descriptor, false)
}

func (class *Class) GetRefVar(fieldName, fieldDescriptor string) *Object {
	field := class.getField(fieldName, fieldDescriptor, true)
	return class.staticVars.GetRef(field.slotId)
}
func (class *Class) SetRefVar(fieldName, fieldDescriptor string, ref *Object) {
	field := class.getField(fieldName, fieldDescriptor, true)
	class.staticVars.SetRef(field.slotId, ref)
}
