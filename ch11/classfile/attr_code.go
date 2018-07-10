package classfile

type ExceptionTableEntry struct {
	startPc uint16
	endPc uint16
	handlerPc uint16
	catchType uint16
}

func (entry *ExceptionTableEntry) StartPc() uint16 {
	return entry.startPc
}

func (entry *ExceptionTableEntry) EndPc() uint16 {
	return entry.endPc
}

func (entry *ExceptionTableEntry) HandlerPc() uint16 {
	return entry.handlerPc
}

func (entry *ExceptionTableEntry) CatchType() uint16 {
	return entry.catchType
}

type CodeAttribute struct {
	cp ConstantPool
	maxStack uint16
	maxLocals uint16
	code []byte
	exceptionTable []*ExceptionTableEntry
	attributes []AttributeInfo
}


func (ca *CodeAttribute) readInfo(reader *ClassReader){
	ca.maxStack = reader.readUint16()
	ca.maxLocals = reader.readUint16()
	codeLength := reader.readUint32()
	ca.code = reader.readBytes(codeLength)
	ca.exceptionTable = readExceptionTable(reader)
	ca.attributes = readAttributes(reader,ca.cp)
}
func (ca *CodeAttribute) MaxLocals() uint {
	return uint(ca.maxLocals)
}

func (ca *CodeAttribute) ExceptionTable() []*ExceptionTableEntry {
	return ca.exceptionTable
}
func (ca *CodeAttribute) MaxStack() uint {
	return uint(ca.maxStack)
}
func (ca *CodeAttribute) Code() []byte {
	return ca.code
}


func (ca *CodeAttribute) LineNumberTableAttribute() *LineNumberTableAttribute {
	for _, attrInfo := range ca.attributes {
		switch attrInfo.(type) {
		case *LineNumberTableAttribute:
			return attrInfo.(*LineNumberTableAttribute)
		}
	}
	return nil
}



func readExceptionTable(reader *ClassReader) []*ExceptionTableEntry {
	exceptionTableLength := reader.readUint16()
	exceptionTable := make([]*ExceptionTableEntry,exceptionTableLength)
	for i := range exceptionTable {
		exceptionTable[i] = &ExceptionTableEntry{
			startPc:reader.readUint16(),
			endPc:reader.readUint16(),
			handlerPc:reader.readUint16(),
			catchType:reader.readUint16(),
		}
	}
	return exceptionTable
}