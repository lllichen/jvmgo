package classfile

import "fmt"

type ClassFile struct {
	//magic      uint32
	minorVersion uint16
	majorVersion uint16
	constantPool ConstantPool
	accessFlags uint16
	thisClass uint16
	superClass uint16
	interfaces []uint16
	fields  []*MemberInfo
	methods []*MemberInfo
	attributes []AttributeInfo

}


func Parse(classData []byte)(cf *ClassFile,err error) {
	//interface assertion
	defer func(){
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)
			// check err ,not err ,format to err
			if !ok {
				err = fmt.Errorf("%v",r)
			}
		}
	}()
	cr := &ClassReader{classData}
	cf = &ClassFile{}
	cf.read(cr)
	return
}

func (cf *ClassFile) read(reader *ClassReader) {
	//read magic
	cf.readAndCheckMagic(reader)
	//read version
	cf.readAndCheckVersion(reader)
	//read pool
	cf.constantPool = readConstantPool(reader)
	cf.accessFlags = reader.readUint16()
	cf.thisClass = reader.readUint16()
	cf.superClass = reader.readUint16()
	cf.interfaces = reader.readUint16s()
	cf.fields = readMembers(reader, cf.constantPool)
	cf.methods = readMembers(reader,cf.constantPool)
	cf.attributes = readAttributes(reader,cf.constantPool)

}

func (cf *ClassFile) MajorVersion() uint16{
	return cf.majorVersion

}

func (cf *ClassFile) MinorVersion() uint16{
	return cf.minorVersion

}


func (cf *ClassFile) ConstantPool() ConstantPool{
	return cf.constantPool
}

func (cf *ClassFile) AccessFlags() uint16{
	return cf.accessFlags
}

func (cf *ClassFile) ClassName() string{
	return cf.constantPool.getClassName(cf.thisClass)
}

func (cf *ClassFile) SuperClassName() string{
	if cf.superClass > 0 {
		return cf.constantPool.getClassName(cf.superClass)
	}
	return ""
}


func (cf *ClassFile) InterfaceNames() []string{
	interfaceName := make([]string,len(cf.interfaces))
	for i,cpIndex := range cf.interfaces {
		interfaceName[i] = cf.constantPool.getClassName(cpIndex)
	}
	return interfaceName
}


func (cf *ClassFile) Fields() []*MemberInfo{
	return cf.fields
}


func (cf *ClassFile) Methods() []*MemberInfo{
	return cf.methods
}

func (cf *ClassFile) readAndCheckMagic(reader *ClassReader) {
	magic := reader.readUint32()
	if magic != 0XCAFEBABE {
		panic("java.lang.ClassFormatError: magic!")
	}
}
func (cf *ClassFile) readAndCheckVersion(reader *ClassReader) {
	cf.minorVersion = reader.readUint16()
	cf.majorVersion = reader.readUint16()
	//fmt.Printf("minorVersion is %d\n",cf.minorVersion)
	//fmt.Printf("majorVersion is %d\n",cf.majorVersion)
	switch cf.majorVersion {
	case 45:
		return
	case 46,47,48,49,50,51,52:
		if cf.minorVersion == 0 {
			return
		}
	}
	panic("java.lang.UnsupportedClassVersionError!")
}







