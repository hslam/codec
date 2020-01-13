// Copyright (c) 2019 Meng Huang (mhboy@outlook.com)
// This package is licensed under a MIT license that can be found in the LICENSE file.

// Package codec implements encoding and decoding of multiple codecs
package codec

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"encoding/xml"
	gogoproto "github.com/gogo/protobuf/proto"
	"github.com/golang/protobuf/proto"
)

// Codec defines the interface for encoding/decoding.
type Codec interface {
	Encode(v interface{}) ([]byte, error)
	Decode(data []byte, v interface{}) error
}

// JSONCodec struct
type JSONCodec struct {
}

// Encode returns the JSON encoding of v.
func (c *JSONCodec) Encode(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

// Decode parses the JSON-encoded data and stores the result in the value pointed to by v.
func (c *JSONCodec) Decode(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

// PBCodec struct
type PBCodec struct {
}

// Encode returns the PB encoding of v.
func (c *PBCodec) Encode(v interface{}) ([]byte, error) {
	return proto.Marshal(v.(proto.Message))
}

// Decode parses the PB-encoded data and stores the result in the value pointed to by v.
func (c *PBCodec) Decode(data []byte, v interface{}) error {
	return proto.Unmarshal(data, v.(proto.Message))
}

// XMLCodec struct
type XMLCodec struct {
}

// Encode returns the XML encoding of v.
func (c *XMLCodec) Encode(v interface{}) ([]byte, error) {
	return xml.Marshal(v)
}

// Decode parses the XML-encoded data and stores the result in the value pointed to by v.
func (c *XMLCodec) Decode(data []byte, v interface{}) error {
	return xml.Unmarshal(data, v)
}

// GOBCodec struct
type GOBCodec struct {
	buffer bytes.Buffer
}

// Encode returns the GOB encoding of v.
func (c *GOBCodec) Encode(v interface{}) ([]byte, error) {
	c.buffer.Reset()
	gob.NewEncoder(&c.buffer).Encode(v)
	return c.buffer.Bytes(), nil
}

// Decode parses the GOB-encoded data and stores the result in the value pointed to by v.
func (c *GOBCodec) Decode(data []byte, v interface{}) error {
	return gob.NewDecoder(bytes.NewReader(data)).Decode(v)
}

// BYTESCodec struct
type BYTESCodec struct {
}

// Encode returns the BYTES encoding of v.
func (c *BYTESCodec) Encode(v interface{}) ([]byte, error) {
	return *v.(*[]byte), nil
}

// Decode parses the BYTES-encoded data and stores the result in the value pointed to by v.
func (c *BYTESCodec) Decode(data []byte, v interface{}) error {
	*v.(*[]byte) = data
	return nil
}

// gogoprotobuf defines the interface for gogo's protobuf.
type gogoprotobuf interface {
	Size() (n int)
	Marshal() (data []byte, err error)
	MarshalTo(buf []byte) (int, error)
	Unmarshal(data []byte) error
}

// GOGOPBCodec struct
type GOGOPBCodec struct {
	Buffer []byte
}

// Encode returns the GOGOPB encoding of v.
func (c *GOGOPBCodec) Encode(v interface{}) ([]byte, error) {
	if p, ok := v.(gogoprotobuf); ok {
		size := p.Size()
		if cap(c.Buffer) >= size {
			c.Buffer = c.Buffer[:size]
			n, err := p.MarshalTo(c.Buffer)
			return c.Buffer[:n], err
		}
		return p.Marshal()
	}
	return gogoproto.Marshal(v.(proto.Message))
}

// Decode parses the GOGOPB-encoded data and stores the result in the value pointed to by v.
func (c *GOGOPBCodec) Decode(data []byte, v interface{}) error {
	if p, ok := v.(gogoprotobuf); ok {
		return p.Unmarshal(data)
	}
	return gogoproto.Unmarshal(data, v.(proto.Message))
}

// Code defines the interface for code.
type Code interface {
	Marshal(buf []byte) ([]byte, error)
	Unmarshal(buf []byte) (uint64, error)
}

// CODECodec struct
type CODECodec struct {
	Buffer []byte
}

// Encode returns the CODE encoding of v.
func (c *CODECodec) Encode(v interface{}) ([]byte, error) {
	return v.(Code).Marshal(c.Buffer)
}

// Decode parses the CODE-encoded data and stores the result in the value pointed to by v.
func (c *CODECodec) Decode(data []byte, v interface{}) error {
	_, err := v.(Code).Unmarshal(data)
	return err
}

// msgpack defines the interface for msgp.
type msgpack interface {
	MarshalMsg(buf []byte) ([]byte, error)
	UnmarshalMsg(bts []byte) (o []byte, err error)
}

// MSGPCodec struct
type MSGPCodec struct {
	Buffer []byte
}

// Encode returns the MSGP encoding of v.
func (c *MSGPCodec) Encode(v interface{}) ([]byte, error) {
	return v.(msgpack).MarshalMsg(c.Buffer[:0])
}

// Decode parses the MSGP-encoded data and stores the result in the value pointed to by v.
func (c *MSGPCodec) Decode(data []byte, v interface{}) error {
	_, err := v.(msgpack).UnmarshalMsg(data)
	return err
}
