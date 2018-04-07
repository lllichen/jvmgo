package classfile

type AttributeInfo interface {
	readInfo (reader *ClassReader)
}

//func readAttribute(reader *ClassReader,cp ConstantPool) AttributeInfo {
//	attrNameIndex := reader.readUint16()
//	attrName := cp.getUtf8(attrNameIndex)
//	attrLen := reader.readUint32()
//	attrInfo := newAttributeInfo(attrName,attrLen,cp)
//	attrInfo.readInfo(reader)
//	return attrinfo
//}