package bytes

import (
	"errors"
)

type Student struct {
	Name string
	Age int32
	Address string
}

func (s *Student)Encode()([]byte, error)  {
	if len(s.Name)>255||len(s.Address)>255||s.Age>255{
		return nil,errors.New("the length is  too long")
	}
	var buffer=make([]byte,2+len(s.Name)+len(s.Address))
	copy(buffer[0:], []byte{byte(len(s.Name))})
	copy(buffer[1:], []byte(s.Name))
	copy(buffer[1+len(s.Name):], []byte{byte(s.Age)})
	copy(buffer[2+len(s.Name):], []byte(s.Address))
	return buffer,nil
}
func (s *Student)Decode(data []byte) error  {
	if len(data)<3{
		return errors.New("the data is too short")
	}
	name_length:=int(data[0])
	s.Name=string(data[1:1+name_length])
	s.Age=int32(data[1+name_length])
	s.Address=string(data[2+name_length:])
	return nil
}
