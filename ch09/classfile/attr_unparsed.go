package classfile

type UnParsedAttribute struct {
	name string
	length uint32
	info []byte
}

func (ua *UnParsedAttribute) readInfo(reader *ClassReader)  {
	ua.info = reader.readBytes(ua.length)
}



