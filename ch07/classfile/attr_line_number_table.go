package classfile

type LineNumberTableEntry struct {
	startPc uint16
	lineNumber uint16
}

type LineNumberTableAttribute struct {
	LineNumberTable []*LineNumberTableEntry
}

func (lta *LineNumberTableAttribute) readInfo (reader *ClassReader) {
	lineNumberLength := reader.readUint16()
	lta.LineNumberTable = make([]*LineNumberTableEntry,lineNumberLength)
	for i := range lta.LineNumberTable {
		lta.LineNumberTable[i] = &LineNumberTableEntry{
			startPc:reader.readUint16(),
			lineNumber:reader.readUint16(),
		}
	}

}