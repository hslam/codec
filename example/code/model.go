package code

import (
	"github.com/hslam/code"
)

//Object is a test struct
type Object struct {
	A uint32
	B uint64
	C float32
	D float64
	E string
	F bool
	G []byte
	H [][]byte
}

//Marshal marshals the Object into buf and returns the bytes.
func (o *Object) Marshal(buf []byte) ([]byte, error) {
	var size uint64
	size += code.MaxVarintBytes(uint64(o.A))
	size += code.MaxVarintBytes(o.B)
	size += code.MaxFloat32Bytes(o.C)
	size += code.MaxFloat64Bytes(o.D)
	size += code.MaxStringBytes(o.E)
	size += code.MaxBoolBytes(o.F)
	size += code.MaxBytesBytes(o.G)
	size += code.MaxBytesSliceBytes(o.H)
	if uint64(cap(buf)) >= size {
		buf = buf[:size]
	} else {
		buf = make([]byte, size)
	}
	var offset uint64
	var n uint64
	n = code.EncodeVarint(buf[offset:], uint64(o.A))
	offset += n
	n = code.EncodeVarint(buf[offset:], o.B)
	offset += n
	n = code.EncodeFloat32(buf[offset:], o.C)
	offset += n
	n = code.EncodeFloat64(buf[offset:], o.D)
	offset += n
	n = code.EncodeString(buf[offset:], o.E)
	offset += n
	n = code.EncodeBool(buf[offset:], o.F)
	offset += n
	n = code.EncodeBytes(buf[offset:], o.G)
	offset += n
	n = code.EncodeBytesSlice(buf[offset:], o.H)
	offset += n
	return buf[:offset], nil
}

//Unmarshal unmarshals the Object from buf and returns the number of bytes read (> 0).
func (o *Object) Unmarshal(data []byte) (uint64, error) {
	var offset uint64
	var n uint64
	var t uint64
	n = code.DecodeVarint(data[offset:], &t)
	o.A = uint32(t)
	offset += n
	n = code.DecodeVarint(data[offset:], &o.B)
	offset += n
	n = code.DecodeFloat32(data[offset:], &o.C)
	offset += n
	n = code.DecodeFloat64(data[offset:], &o.D)
	offset += n
	n = code.DecodeString(data[offset:], &o.E)
	offset += n
	n = code.DecodeBool(data[offset:], &o.F)
	offset += n
	n = code.DecodeBytes(data[offset:], &o.G)
	offset += n
	n = code.DecodeBytesSlice(data[offset:], &o.H)
	offset += n
	return offset, nil
}
