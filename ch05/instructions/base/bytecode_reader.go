package base

type ByteCodeReader struct {
	code []byte
	pc int
}

func (reader *ByteCodeReader) ReadInt16() int16 {
	return  int16(reader.readUint16())
}

func (reader *ByteCodeReader) readUint16() uint16 {
	byte1 := uint16(reader.readUint8())
	byte2 := uint16(reader.readUint8())
	return (byte1 <<8) |byte2
}

func (reader *ByteCodeReader) reset(code []byte, pc int){
	reader.code = code
	reader.pc = pc
}

func (reader *ByteCodeReader) readUint8() uint8 {
	i := reader.code[reader.pc]
	reader.pc++
	return i
}


func (reader *ByteCodeReader) readInt32() int32 {
	byte1 := int32(reader.readUint8())
	byte2 := int32(reader.readUint8())
	byte3 := int32(reader.readUint8())
	byte4 := int32(reader.readUint8())
	return (byte1<<24) | (byte2<<16) | (byte3<<8) | byte4
}
