package pb

import (
	"github.com/golang/protobuf/proto"
)

// Codec struct
type Codec struct {
}

// Marshal returns the PB encoding of v.
func (c *Codec) Marshal(buf []byte, v interface{}) ([]byte, error) {
	return proto.Marshal(v.(proto.Message))
}

// Unmarshal parses the PB-encoded data and stores the result in the value pointed to by v.
func (c *Codec) Unmarshal(data []byte, v interface{}) error {
	return proto.Unmarshal(data, v.(proto.Message))
}
