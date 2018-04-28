package heap

import (
	"jvmgo/ch06/classpath"
	"fmt"
	"jvmgo/ch06/classfile"
)

type ClassLoader struct {
	cp       *classpath.Classpath
	classMap map[string]*Class //loaded classes
}

func NewClassLoader(pool *classpath.Classpath) *ClassLoader {
	return &ClassLoader{
		cp:       pool,
		classMap: make(map[string]*Class),
	}
}

func (classLoader *ClassLoader) LoadClass(name string) *Class{
	if class, ok := classLoader.classMap[name];ok {
		return class
	}
	return classLoader.loadNonArrayClass(name)
}

func (classLoader *ClassLoader) loadNonArrayClass(name string) *Class {
	data, entry := classLoader.readClass(name)
	class := classLoader.defineClass(data)
	link(class)
	fmt.Printf("[Loaded %s from %s]\n",name,entry )
	return class
}

func (classLoader *ClassLoader) readClass(name string) ([]byte,classpath.Entry) {
	data,entry,err := classLoader.cp.ReadClass(name)
	if err != nil {
		panic("java.lang.ClassNotFoundException: "+name)
	}
	return data,entry
}

func (classLoader *ClassLoader) defineClass(data []byte) *Class {
	class := parseClass(data)
	class.loader = classLoader
	resolveSuperClass(class)
	resolveInterfaces(class)
	classLoader.classMap[class.name] = class
	return class
}

func parseClass (data []byte) *Class {
	cf,err := classfile.Parse(data)
	if err != nil {
		panic("java.lang.ClassFormatError")
	}
	return newClass(cf)
}

