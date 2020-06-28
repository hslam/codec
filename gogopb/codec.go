package gogopb

import (
	gogoproto "github.com/gogo/protobuf/proto"
	"github.com/golang/protobuf/proto"
)

// gogoprotobuf defines the interface for gogo's protobuf.
type gogoprotobuf interface {
	Size() (n int)
	Marshal() (data []byte, err error)
	MarshalTo(buf []byte) (int, error)
	Unmarshal(data []byte) error
}

// Codec struct
type Codec struct {
}

// Marshal returns the GOGOPB encoding of v.
func (c *Codec) Marshal(buf []byte, v interface{}) ([]byte, error) {
	if p, ok := v.(gogoprotobuf); ok {
		size := p.Size()
		if cap(buf) >= size {
			buf = buf[:size]
			n, err := p.MarshalTo(buf)
			return buf[:n], err
		}
		return p.Marshal()
	}
	return gogoproto.Marshal(v.(proto.Message))
}

// Unmarshal parses the GOGOPB-encoded data and stores the result in the value pointed to by v.
func (c *Codec) Unmarshal(data []byte, v interface{}) error {
	if p, ok := v.(gogoprotobuf); ok {
		return p.Unmarshal(data)
	}
	return gogoproto.Unmarshal(data, v.(proto.Message))
}
