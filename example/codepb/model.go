package codepb

import (
	"fmt"
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
	size += code.MaxVarintBytes(uint64(o.A)) + 1
	size += code.MaxVarintBytes(o.B) + 1
	size += code.MaxFloat32Bytes(o.C) + 1
	size += code.MaxFloat64Bytes(o.D) + 1
	size += code.MaxStringBytes(o.E) + 1
	size += code.MaxBoolBytes(o.F) + 1
	size += code.MaxBytesBytes(o.G) + 1
	size += code.MaxBytesSliceBytes(o.H) + 1
	if uint64(cap(buf)) >= size {
		buf = buf[:size]
	} else {
		buf = make([]byte, size)
	}
	var offset uint64
	var n uint64
	if o.A != 0 {
		buf[offset] = 1<<3 | 0
		offset++
		n = code.EncodeVarint(buf[offset:], uint64(o.A))
		offset += n
	}
	if o.B != 0 {
		buf[offset] = 2<<3 | 0
		offset++
		n = code.EncodeVarint(buf[offset:], o.B)
		offset += n
	}
	if o.C != 0 {
		buf[offset] = 3<<3 | 5
		offset++
		n = code.EncodeFloat32(buf[offset:], o.C)
		offset += n
	}
	if o.D != 0 {
		buf[offset] = 4<<3 | 1
		offset++
		n = code.EncodeFloat64(buf[offset:], o.D)
		offset += n
	}
	if len(o.E) > 0 {
		buf[offset] = 5<<3 | 2
		offset++
		n = code.EncodeString(buf[offset:], o.E)
		offset += n
	}
	if o.F == true {
		buf[offset] = 6<<3 | 0
		offset++
		n = code.EncodeBool(buf[offset:], o.F)
		offset += n
	}
	if len(o.G) > 0 {
		buf[offset] = 7<<3 | 2
		offset++
		n = code.EncodeBytes(buf[offset:], o.G)
		offset += n
	}
	if len(o.H) > 0 {
		tag := byte(8<<3 | 2)
		for _, v := range o.H {
			buf[offset] = tag
			offset++
			n = code.EncodeBytes(buf[offset:], v)
			offset += n
		}
	}
	return buf[:offset], nil
}

//Unmarshal unmarshals the Object from buf and returns the number of bytes read (> 0).
func (o *Object) Unmarshal(data []byte) (uint64, error) {
	var length uint64 = uint64(len(data))
	var offset uint64
	var n uint64
	var tag uint64
	var fieldNumber int
	var wireType uint8
	o.H = o.H[:0]
	for {
		if offset < length {
			//n = code.DecodeVarint(data[offset:], &tag)
			//offset += n
			tag = uint64(data[offset])
			offset++
		} else {
			break
		}
		fieldNumber = int(tag >> 3)
		wireType = uint8(tag & 0x7)
		switch fieldNumber {
		case 1:
			if wireType != 0 {
				return 0, fmt.Errorf("proto: wrong wireType = %d for field A", wireType)
			}
			var t uint64
			n = code.DecodeVarint(data[offset:], &t)
			o.A = uint32(t)
			offset += n
		case 2:
			if wireType != 0 {
				return 0, fmt.Errorf("proto: wrong wireType = %d for field B", wireType)
			}
			n = code.DecodeVarint(data[offset:], &o.B)
			offset += n
		case 3:
			if wireType != 5 {
				return 0, fmt.Errorf("proto: wrong wireType = %d for field C", wireType)
			}
			n = code.DecodeFloat32(data[offset:], &o.C)
			offset += n
		case 4:
			if wireType != 1 {
				return 0, fmt.Errorf("proto: wrong wireType = %d for field D", wireType)
			}
			n = code.DecodeFloat64(data[offset:], &o.D)
			offset += n
		case 5:
			if wireType != 2 {
				return 0, fmt.Errorf("proto: wrong wireType = %d for field E", wireType)
			}
			n = code.DecodeString(data[offset:], &o.E)
			offset += n
		case 6:
			if wireType != 0 {
				return 0, fmt.Errorf("proto: wrong wireType = %d for field F", wireType)
			}
			n = code.DecodeBool(data[offset:], &o.F)
			offset += n
		case 7:
			if wireType != 2 {
				return 0, fmt.Errorf("proto: wrong wireType = %d for field G", wireType)
			}
			n = code.DecodeBytes(data[offset:], &o.G)
			offset += n
		case 8:
			if wireType != 2 {
				return 0, fmt.Errorf("proto: wrong wireType = %d for field H", wireType)
			}
			for {
				var b []byte
				n = code.DecodeBytes(data[offset:], &b)
				o.H = append(o.H, b)
				offset += n
				if offset < length {
					//n = code.DecodeVarint(data[offset:], &t)
					tmpTag := uint64(data[offset])
					if tmpTag == tag {
						//offset += n
						offset++
						continue
					}
				}
				break
			}
		}
	}
	return offset, nil
}
