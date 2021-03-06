package msgp

// Code generated by github.com/tinylib/msgp DO NOT EDIT.

import (
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *Object) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "A":
			z.A, err = dc.ReadUint32()
			if err != nil {
				err = msgp.WrapError(err, "A")
				return
			}
		case "B":
			z.B, err = dc.ReadUint64()
			if err != nil {
				err = msgp.WrapError(err, "B")
				return
			}
		case "C":
			z.C, err = dc.ReadFloat32()
			if err != nil {
				err = msgp.WrapError(err, "C")
				return
			}
		case "D":
			z.D, err = dc.ReadFloat64()
			if err != nil {
				err = msgp.WrapError(err, "D")
				return
			}
		case "E":
			z.E, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "E")
				return
			}
		case "F":
			z.F, err = dc.ReadBool()
			if err != nil {
				err = msgp.WrapError(err, "F")
				return
			}
		case "G":
			z.G, err = dc.ReadBytes(z.G)
			if err != nil {
				err = msgp.WrapError(err, "G")
				return
			}
		case "H":
			var zb0002 uint32
			zb0002, err = dc.ReadArrayHeader()
			if err != nil {
				err = msgp.WrapError(err, "H")
				return
			}
			if cap(z.H) >= int(zb0002) {
				z.H = (z.H)[:zb0002]
			} else {
				z.H = make([][]byte, zb0002)
			}
			for za0001 := range z.H {
				z.H[za0001], err = dc.ReadBytes(z.H[za0001])
				if err != nil {
					err = msgp.WrapError(err, "H", za0001)
					return
				}
			}
		default:
			err = dc.Skip()
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *Object) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 8
	// write "A"
	err = en.Append(0x88, 0xa1, 0x41)
	if err != nil {
		return
	}
	err = en.WriteUint32(z.A)
	if err != nil {
		err = msgp.WrapError(err, "A")
		return
	}
	// write "B"
	err = en.Append(0xa1, 0x42)
	if err != nil {
		return
	}
	err = en.WriteUint64(z.B)
	if err != nil {
		err = msgp.WrapError(err, "B")
		return
	}
	// write "C"
	err = en.Append(0xa1, 0x43)
	if err != nil {
		return
	}
	err = en.WriteFloat32(z.C)
	if err != nil {
		err = msgp.WrapError(err, "C")
		return
	}
	// write "D"
	err = en.Append(0xa1, 0x44)
	if err != nil {
		return
	}
	err = en.WriteFloat64(z.D)
	if err != nil {
		err = msgp.WrapError(err, "D")
		return
	}
	// write "E"
	err = en.Append(0xa1, 0x45)
	if err != nil {
		return
	}
	err = en.WriteString(z.E)
	if err != nil {
		err = msgp.WrapError(err, "E")
		return
	}
	// write "F"
	err = en.Append(0xa1, 0x46)
	if err != nil {
		return
	}
	err = en.WriteBool(z.F)
	if err != nil {
		err = msgp.WrapError(err, "F")
		return
	}
	// write "G"
	err = en.Append(0xa1, 0x47)
	if err != nil {
		return
	}
	err = en.WriteBytes(z.G)
	if err != nil {
		err = msgp.WrapError(err, "G")
		return
	}
	// write "H"
	err = en.Append(0xa1, 0x48)
	if err != nil {
		return
	}
	err = en.WriteArrayHeader(uint32(len(z.H)))
	if err != nil {
		err = msgp.WrapError(err, "H")
		return
	}
	for za0001 := range z.H {
		err = en.WriteBytes(z.H[za0001])
		if err != nil {
			err = msgp.WrapError(err, "H", za0001)
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *Object) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 8
	// string "A"
	o = append(o, 0x88, 0xa1, 0x41)
	o = msgp.AppendUint32(o, z.A)
	// string "B"
	o = append(o, 0xa1, 0x42)
	o = msgp.AppendUint64(o, z.B)
	// string "C"
	o = append(o, 0xa1, 0x43)
	o = msgp.AppendFloat32(o, z.C)
	// string "D"
	o = append(o, 0xa1, 0x44)
	o = msgp.AppendFloat64(o, z.D)
	// string "E"
	o = append(o, 0xa1, 0x45)
	o = msgp.AppendString(o, z.E)
	// string "F"
	o = append(o, 0xa1, 0x46)
	o = msgp.AppendBool(o, z.F)
	// string "G"
	o = append(o, 0xa1, 0x47)
	o = msgp.AppendBytes(o, z.G)
	// string "H"
	o = append(o, 0xa1, 0x48)
	o = msgp.AppendArrayHeader(o, uint32(len(z.H)))
	for za0001 := range z.H {
		o = msgp.AppendBytes(o, z.H[za0001])
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Object) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "A":
			z.A, bts, err = msgp.ReadUint32Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "A")
				return
			}
		case "B":
			z.B, bts, err = msgp.ReadUint64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "B")
				return
			}
		case "C":
			z.C, bts, err = msgp.ReadFloat32Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "C")
				return
			}
		case "D":
			z.D, bts, err = msgp.ReadFloat64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "D")
				return
			}
		case "E":
			z.E, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "E")
				return
			}
		case "F":
			z.F, bts, err = msgp.ReadBoolBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "F")
				return
			}
		case "G":
			z.G, bts, err = msgp.ReadBytesBytes(bts, z.G)
			if err != nil {
				err = msgp.WrapError(err, "G")
				return
			}
		case "H":
			var zb0002 uint32
			zb0002, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "H")
				return
			}
			if cap(z.H) >= int(zb0002) {
				z.H = (z.H)[:zb0002]
			} else {
				z.H = make([][]byte, zb0002)
			}
			for za0001 := range z.H {
				z.H[za0001], bts, err = msgp.ReadBytesBytes(bts, z.H[za0001])
				if err != nil {
					err = msgp.WrapError(err, "H", za0001)
					return
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Object) Msgsize() (s int) {
	s = 1 + 2 + msgp.Uint32Size + 2 + msgp.Uint64Size + 2 + msgp.Float32Size + 2 + msgp.Float64Size + 2 + msgp.StringPrefixSize + len(z.E) + 2 + msgp.BoolSize + 2 + msgp.BytesPrefixSize + len(z.G) + 2 + msgp.ArrayHeaderSize
	for za0001 := range z.H {
		s += msgp.BytesPrefixSize + len(z.H[za0001])
	}
	return
}
