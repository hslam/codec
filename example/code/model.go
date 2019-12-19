package code

import (
	"github.com/hslam/code"
)

type Student struct {
	Name string
	Age int32
	Address string
}

func (s *Student)Marshal(buf []byte)([]byte, error)  {
	var age=uint64(s.Age)
	name_length:=code.SizeofString(&s.Name)
	age_length:=code.SizeofVarint(&age)
	address_length:=code.SizeofString(&s.Address)
	var max int
	if name_length>max{
		max=name_length
	}
	if address_length>max{
		max=address_length
	}
	if age_length>max{
		max=age_length
	}
	length:=name_length+address_length+age_length
	var buffer []byte
	if cap(buf) >= length {
		buffer = buf[:length]
	} else {
		buffer = make([]byte, length)
	}
	code.EncodeString(buffer[:name_length],&s.Name)
	code.EncodeVarint(buffer[name_length:name_length+age_length],&age)
	code.EncodeString(buffer[name_length+age_length:],&s.Address)
	return buffer,nil
}
func (s *Student)Unmarshal(data []byte) (uint64, error)  {
	var offset uint64
	size:=code.DecodeString(data,&s.Name)
	offset+=uint64(size)
	var age uint64
	size=code.DecodeVarint(data[offset:],&age)
	s.Age=int32(age)
	offset+=uint64(size)
	size=code.DecodeString(data[offset:],&s.Address)
	offset+=uint64(size)
	return offset,nil
}
