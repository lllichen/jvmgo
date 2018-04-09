package classfile


type ConstantNameAndTypeInfo struct {
	nameIndex uint16
	descriptorIndex uint16
}

func (cti *ConstantNameAndTypeInfo) readInfo(reader *ClassReader) {
	cti.nameIndex = reader.readUint16()
	cti.descriptorIndex = reader.readUint16()
}
