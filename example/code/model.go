package code

import (
	"github.com/hslam/code"
)

type Object struct {
	A uint32
	B uint64
	C float32
	D float64
	E string
	F bool
	G []byte
}

func (o *Object)Marshal(buf []byte)([]byte, error)  {
	var size uint64
	size+=4
	size+=8
	size+=4
	size+=8
	size+=code.SizeofString(o.E)
	size+=1
	size+=code.SizeofBytes(o.G)
	if uint64(cap(buf)) >= size {
		buf = buf[:size]
	} else {
		buf = make([]byte, size)
	}
	var offset uint64
	var n  uint64
	n = code.EncodeUint32(buf[offset:],o.A)
	offset+=n
	n = code.EncodeUint64(buf[offset:],o.B)
	offset+=n
	n = code.EncodeFloat32(buf[offset:],o.C)
	offset+=n
	n = code.EncodeFloat64(buf[offset:],o.D)
	offset+=n
	n = code.EncodeString(buf[offset:],o.E)
	offset+=n
	n = code.EncodeBool(buf[offset:],o.F)
	offset+=n
	n = code.EncodeBytes(buf[offset:],o.G)
	offset+=n
	return buf,nil
}
func (o *Object)Unmarshal(data []byte) (uint64, error)  {
	var offset uint64
	var n uint64
	n=code.DecodeUint32(data[offset:],&o.A)
	offset+=n
	n=code.DecodeUint64(data[offset:],&o.B)
	offset+=n
	n=code.DecodeFloat32(data[offset:],&o.C)
	offset+=n
	n=code.DecodeFloat64(data[offset:],&o.D)
	offset+=n
	n=code.DecodeString(data[offset:],&o.E)
	offset+=n
	n=code.DecodeBool(data[offset:],&o.F)
	offset+=n
	n=code.DecodeBytes(data[offset:],&o.G)
	offset+=n
	return offset,nil
}
