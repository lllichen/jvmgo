package classfile

import (
	"encoding/binary"
	"fmt"
)

type ClassReader struct {
	data []byte
}

func (cr *ClassReader) readUint8() uint8{
	val := cr.data[0]
	cr.data = cr.data[1:]
	return val
}

func (cr *ClassReader) readUint16() uint16{
	fmt.Printf("%v\n",cr.data)

	val := binary.BigEndian.Uint16(cr.data)
	cr.data = cr.data[2:]
	return val
}

func (cr *ClassReader) readUint32() uint32{
	val := binary.BigEndian.Uint32(cr.data)
	cr.data = cr.data[4:]
	fmt.Printf("read magic %v\n",cr.data)
	return val
}

func (cr *ClassReader) readUint64() uint64{
	val := binary.BigEndian.Uint64(cr.data)
	cr.data = cr.data[8:]
	return val
}

func (cr *ClassReader) readUint16s() []uint16{
	n := cr.readUint16()
	s := make([]uint16,n)
	for i := range s {
		s[i] = cr.readUint16()
	}
	return s
}


func (cr *ClassReader) readBytes(n uint32) []byte{
	bytes := cr.data[:n]
	cr.data = cr.data[:n]
	return bytes
}
