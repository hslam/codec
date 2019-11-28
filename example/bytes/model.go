package bytes

import "errors"

var buf []byte
func init() {
	buf=make([]byte,100)
}
type Student struct {
	Name string
	Age int32
	Address string
}

func (s *Student)Marshal()([]byte, error)  {
	name_length:=len(s.Name)
	address_length:=len(s.Address)
	if (name_length>255||address_length>255||s.Age>255){
		return nil,errors.New("too big")
	}
	length:=2+name_length+address_length
	var buffer=buf
	if cap(buffer) >= length {
		buffer = buffer[:length]
	} else {
		buffer = make([]byte, length)
	}
	buffer[0]=byte(name_length)
	copy(buffer[1:1+name_length], []byte(s.Name))
	buffer[1+name_length]=byte(s.Age)
	copy(buffer[2+name_length:length], []byte(s.Address))
	return buffer,nil
}
func (s *Student)Unmarshal(data []byte) error  {
	if (len(data)<3){
		return errors.New("too short")
	}
	name_length:=int(data[0])
	s.Name=string(data[1:1+name_length])
	s.Age=int32(data[1+name_length])
	s.Address=string(data[2+name_length:])
	return nil
}
