package classfile

type ConstantMemberRefInfo struct {
	cp ConstantPool
	classIndex uint16
	nameAndTypeIndex uint16
}

func (cmi ConstantMemberRefInfo) readInfo (reader *ClassReader) {
	cmi.classIndex = reader.readUint16()
	cmi.nameAndTypeIndex = reader.readUint16()
}

func (cmi *ConstantMemberRefInfo) ClassName() string {
	return cmi.cp.getClassName(cmi.classIndex)
}


func (cmi *ConstantMemberRefInfo) NameAndDescritor() (string ,string) {
	return cmi.cp.getNameAndType(cmi.nameAndTypeIndex)
}

