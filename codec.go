// Copyright (c) 2019 Meng Huang (mhboy@outlook.com)
// This package is licensed under a MIT license that can be found in the LICENSE file.

// Package codec implements encoding and decoding of multiple codecs
package codec

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"encoding/xml"
	"errors"
)

// Codec defines the interface for encoding/decoding.
type Codec interface {
	Marshal(buf []byte, v interface{}) ([]byte, error)
	Unmarshal(data []byte, v interface{}) error
}

// JSONCodec struct
type JSONCodec struct {
}

// Marshal returns the JSON encoding of v.
func (c *JSONCodec) Marshal(buf []byte, v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

// Unmarshal parses the JSON-encoded data and stores the result in the value pointed to by v.
func (c *JSONCodec) Unmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

// XMLCodec struct
type XMLCodec struct {
}

// Marshal returns the XML encoding of v.
func (c *XMLCodec) Marshal(buf []byte, v interface{}) ([]byte, error) {
	return xml.Marshal(v)
}

// Unmarshal parses the XML-encoded data and stores the result in the value pointed to by v.
func (c *XMLCodec) Unmarshal(data []byte, v interface{}) error {
	return xml.Unmarshal(data, v)
}

// GOBCodec struct
type GOBCodec struct {
}

// Marshal returns the GOB encoding of v.
func (c *GOBCodec) Marshal(buf []byte, v interface{}) ([]byte, error) {
	var buffer = bytes.NewBuffer(buf)
	buffer.Reset()
	gob.NewEncoder(buffer).Encode(v)
	return buffer.Bytes(), nil
}

// Unmarshal parses the GOB-encoded data and stores the result in the value pointed to by v.
func (c *GOBCodec) Unmarshal(data []byte, v interface{}) error {
	return gob.NewDecoder(bytes.NewReader(data)).Decode(v)
}

// BYTESCodec struct
type BYTESCodec struct {
}

// Marshal returns the BYTES encoding of v.
func (c *BYTESCodec) Marshal(buf []byte, v interface{}) ([]byte, error) {
	return *v.(*[]byte), nil
}

// Unmarshal parses the BYTES-encoded data and stores the result in the value pointed to by v.
func (c *BYTESCodec) Unmarshal(data []byte, v interface{}) error {
	*v.(*[]byte) = data
	return nil
}

// GoGoProtobuf defines the interface for gogo's protobuf.
type GoGoProtobuf interface {
	Size() (n int)
	Marshal() (data []byte, err error)
	MarshalTo(buf []byte) (int, error)
	Unmarshal(data []byte) error
}

// GOGOPBCodec struct
type GOGOPBCodec struct {
}

//ErrorGOGOPB is the error that v is not GoGoProtobuf
var ErrorGOGOPB = errors.New("is not GoGoProtobuf")

// Marshal returns the GOGOPB encoding of v.
func (c *GOGOPBCodec) Marshal(buf []byte, v interface{}) ([]byte, error) {
	if p, ok := v.(GoGoProtobuf); ok {
		size := p.Size()
		if cap(buf) >= size {
			buf = buf[:size]
			n, err := p.MarshalTo(buf)
			return buf[:n], err
		}
		return p.Marshal()
	}
	return nil, ErrorGOGOPB
}

// Unmarshal parses the GOGOPB-encoded data and stores the result in the value pointed to by v.
func (c *GOGOPBCodec) Unmarshal(data []byte, v interface{}) error {
	if p, ok := v.(GoGoProtobuf); ok {
		return p.Unmarshal(data)
	}
	return ErrorGOGOPB
}

// Code defines the interface for code.
type Code interface {
	Marshal(buf []byte) ([]byte, error)
	Unmarshal(buf []byte) (uint64, error)
}

//ErrorCODE is the error that v is not Code
var ErrorCODE = errors.New("is not Code")

// CODECodec struct
type CODECodec struct {
}

// Marshal returns the CODE encoding of v.
func (c *CODECodec) Marshal(buf []byte, v interface{}) ([]byte, error) {
	if p, ok := v.(Code); ok {
		return p.Marshal(buf)
	}
	return nil, ErrorCODE
}

// Unmarshal parses the CODE-encoded data and stores the result in the value pointed to by v.
func (c *CODECodec) Unmarshal(data []byte, v interface{}) error {
	if p, ok := v.(Code); ok {
		_, err := p.Unmarshal(data)
		return err
	}
	return ErrorCODE
}

// MsgPack defines the interface for msgp.
type MsgPack interface {
	MarshalMsg(buf []byte) ([]byte, error)
	UnmarshalMsg(bts []byte) (o []byte, err error)
}

//ErrorMSGP is the error that v is not MSGP
var ErrorMSGP = errors.New("is not MSGP")

// MSGPCodec struct
type MSGPCodec struct {
}

// Marshal returns the MSGP encoding of v.
func (c *MSGPCodec) Marshal(buf []byte, v interface{}) ([]byte, error) {
	if p, ok := v.(MsgPack); ok {
		return p.MarshalMsg(buf)
	}
	return nil, ErrorMSGP
}

// Unmarshal parses the MSGP-encoded data and stores the result in the value pointed to by v.
func (c *MSGPCodec) Unmarshal(data []byte, v interface{}) error {
	if p, ok := v.(MsgPack); ok {
		_, err := p.UnmarshalMsg(data)
		return err
	}
	return ErrorMSGP
}
