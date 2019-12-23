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
	var age=uint32(s.Age)
	var l uint64
	var n uint64
	n=code.EncodeString(buf,s.Name)
	l+=n
	n=code.EncodeUint32(buf[l:],age)
	l+=n
	n=code.EncodeString(buf[l:],s.Address)
	l+=n
	return buf[:l],nil
}
func (s *Student)Unmarshal(data []byte) (uint64, error)  {
	var offset uint64
	size:=code.DecodeString(data,&s.Name)
	offset+=uint64(size)
	var age uint32
	size=code.DecodeUint32(data[offset:],&age)
	s.Age=int32(age)
	offset+=uint64(size)
	size=code.DecodeString(data[offset:],&s.Address)
	offset+=uint64(size)
	return offset,nil
}
