// Copyright (c) 2019 Meng Huang (mhboy@outlook.com)
// This package is licensed under a MIT license that can be found in the LICENSE file.

package codec

import (
	"github.com/hslam/codec/example/gencode"
	"github.com/hslam/codec/example/gogopb"
	"github.com/hslam/codec/example/model"
	"github.com/hslam/codec/example/msgp"
	"github.com/hslam/codec/example/pb"
	"testing"
)

func TestBYTESCodecPool(t *testing.T) {
	var p Pool
	var obj = []byte{128, 8, 128, 8, 195, 245, 72, 64, 74, 216, 18, 77, 251, 33, 9, 64, 10, 72, 101, 108, 108, 111, 87, 111, 114, 108, 100, 1, 1, 255, 2, 1, 128, 1, 255}
	p = NewBYTESCodecPool(1)
	c := p.Get()
	ct := p.Get()
	data, _ := c.Encode(&obj)
	var objCopy []byte
	c.Decode(data, &objCopy)
	p.Put(c)
	p.Put(ct)
}

func TestCODECodecPool(t *testing.T) {
	var p Pool
	var obj = gencode.Object{
		A: 1024,
		B: 1024,
		C: 3.14,
		D: 3.1415926,
		E: "HelloWorld",
		F: true,
		G: []byte{255},
		H: [][]byte{{128}, {255}},
	}
	{
		p = NewCODECodecPool(1, 0)
		c := p.Get()
		ct := p.Get()
		data, _ := c.Encode(&obj)
		var objCopy = gencode.Object{}
		c.Decode(data, &objCopy)
		p.Put(c)
		p.Put(ct)
	}
	{
		p = NewCODECodecPool(1024, 65536)
		c := p.Get()
		data, _ := c.Encode(&obj)
		var objCopy = gencode.Object{}
		c.Decode(data, &objCopy)
		p.Put(c)
	}
}

func TestGOGOPBCodecPool(t *testing.T) {
	{
		var p Pool
		var obj = gogopb.Object{
			A: 1024,
			B: 1024,
			C: 3.14,
			D: 3.1415926,
			E: "HelloWorld",
			F: true,
			G: []byte{255},
			H: [][]byte{{128}, {255}},
		}
		{
			p = NewGOGOPBCodecPool(1, 0)
			c := p.Get()
			ct := p.Get()
			data, _ := c.Encode(&obj)
			var objCopy = gogopb.Object{}
			c.Decode(data, &objCopy)
			p.Put(c)
			p.Put(ct)
		}
		{
			p = NewGOGOPBCodecPool(1024, 65536)
			c := p.Get()
			data, _ := c.Encode(&obj)
			var objCopy = gogopb.Object{}
			c.Decode(data, &objCopy)
			p.Put(c)
		}
	}
	{
		var p Pool
		var obj = pb.Object{
			A: 1024,
			B: 1024,
			C: 3.14,
			D: 3.1415926,
			E: "HelloWorld",
			F: true,
			G: []byte{255},
			H: [][]byte{{128}, {255}},
		}
		p = NewGOGOPBCodecPool(1024, 0)
		c := p.Get()
		data, _ := c.Encode(&obj)
		var objCopy = pb.Object{}
		c.Decode(data, &objCopy)
		p.Put(c)
	}
}

func TestMSGPCodecPool(t *testing.T) {
	var p Pool
	var obj = msgp.Object{
		A: 1024,
		B: 1024,
		C: 3.14,
		D: 3.1415926,
		E: "HelloWorld",
		F: true,
		G: []byte{255},
		H: [][]byte{{128}, {255}},
	}
	{
		p = NewMSGPCodecPool(1, 0)
		c := p.Get()
		ct := p.Get()
		data, _ := c.Encode(&obj)
		var objCopy = msgp.Object{}
		c.Decode(data, &objCopy)
		p.Put(c)
		p.Put(ct)
	}
	{
		p = NewMSGPCodecPool(1024, 65536)
		c := p.Get()
		data, _ := c.Encode(&obj)
		var objCopy = msgp.Object{}
		c.Decode(data, &objCopy)
		p.Put(c)
	}
}

func TestPBCodecPool(t *testing.T) {
	var p Pool
	var obj = pb.Object{
		A: 1024,
		B: 1024,
		C: 3.14,
		D: 3.1415926,
		E: "HelloWorld",
		F: true,
		G: []byte{255},
		H: [][]byte{{128}, {255}},
	}
	p = NewPBCodecPool(1)
	c := p.Get()
	ct := p.Get()
	data, _ := c.Encode(&obj)
	var objCopy = pb.Object{}
	c.Decode(data, &objCopy)
	p.Put(c)
	p.Put(ct)
}

func TestJSONCodecPool(t *testing.T) {
	var p Pool
	var obj = model.Object{
		A: 1024,
		B: 1024,
		C: 3.14,
		D: 3.1415926,
		E: "HelloWorld",
		F: true,
		G: []byte{255},
		H: [][]byte{{128}, {255}},
	}
	p = NewJSONCodecPool(1)
	c := p.Get()
	ct := p.Get()
	data, _ := c.Encode(&obj)
	var objCopy = model.Object{}
	c.Decode(data, &objCopy)
	p.Put(c)
	p.Put(ct)
}

func TestXMLCodecPool(t *testing.T) {
	var p Pool
	var obj = model.Object{
		A: 1024,
		B: 1024,
		C: 3.14,
		D: 3.1415926,
		E: "HelloWorld",
		F: true,
		G: []byte{255},
		H: [][]byte{{128}, {255}},
	}
	p = NewXMLCodecPool(1)
	c := p.Get()
	ct := p.Get()
	data, _ := c.Encode(&obj)
	var objCopy = model.Object{}
	c.Decode(data, &objCopy)
	p.Put(c)
	p.Put(ct)
}

func TestGOBCodecPool(t *testing.T) {
	var p Pool
	var obj = model.Object{
		A: 1024,
		B: 1024,
		C: 3.14,
		D: 3.1415926,
		E: "HelloWorld",
		F: true,
		G: []byte{255},
		H: [][]byte{{128}, {255}},
	}
	p = NewGOBCodecPool(1)
	c := p.Get()
	ct := p.Get()
	data, _ := c.Encode(&obj)
	var objCopy = model.Object{}
	c.Decode(data, &objCopy)
	p.Put(c)
	p.Put(ct)
}
