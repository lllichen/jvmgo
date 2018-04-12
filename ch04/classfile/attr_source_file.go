package classfile

type SourceFileAttribute struct {
	cp ConstantPool
	sourceFileIndex uint16
}

func (sa *SourceFileAttribute) readInfo(reader *ClassReader){
	sa.sourceFileIndex = reader.readUint16()
}

func (sa *SourceFileAttribute) FileName( ) string {
	return sa.cp.getUtf8(sa.sourceFileIndex)
}


