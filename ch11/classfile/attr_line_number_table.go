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

func (lta *LineNumberTableAttribute) GetLineNumber(pc int) int {
	for i := len(lta.LineNumberTable) -1 ; i >= 0; i--{
		entry := lta.LineNumberTable[i]

		if pc >= int(entry.startPc) {
			return int(entry.lineNumber)
		}
	}
	return -1
}