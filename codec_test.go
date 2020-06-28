// Copyright (c) 2019 Meng Huang (mhboy@outlook.com)
// This package is licensed under a MIT license that can be found in the LICENSE file.

// Package codec implements encoding and decoding of multiple codecs
package codec

import (
	"github.com/hslam/codec/example/code"
	"github.com/hslam/codec/example/codepb"
	"github.com/hslam/codec/example/gencode"
	"github.com/hslam/codec/example/gogopb"
	"github.com/hslam/codec/example/model"
	"github.com/hslam/codec/example/msgp"
	"github.com/hslam/codec/example/pb"
	codecpb "github.com/hslam/codec/pb"
	"testing"
)

func TestBYTESCodec(t *testing.T) {
	var obj = []byte{128, 8, 128, 8, 195, 245, 72, 64, 74, 216, 18, 77, 251, 33, 9, 64, 10, 72, 101, 108, 108, 111, 87, 111, 114, 108, 100, 1, 1, 255, 2, 1, 128, 1, 255}
	var c = BYTESCodec{}
	var buf = make([]byte, 512)
	var objCopy []byte
	data, _ := c.Marshal(buf, &obj)
	c.Unmarshal(data, &objCopy)
}

func TestCODECodec(t *testing.T) {
	var obj = code.Object{A: 1024, B: 1024, C: 3.14, D: 3.1415926, E: "HelloWorld", F: true, G: []byte{255}, H: [][]byte{{128}, {255}}}
	var c = CODECodec{}
	var buf = make([]byte, 512)
	var objCopy code.Object
	data, _ := c.Marshal(buf, &obj)
	c.Unmarshal(data, &objCopy)
}

func TestGENCODECodec(t *testing.T) {
	var obj = gencode.Object{A: 1024, B: 1024, C: 3.14, D: 3.1415926, E: "HelloWorld", F: true, G: []byte{255}, H: [][]byte{{128}, {255}}}
	var c = CODECodec{}
	var buf = make([]byte, 512)
	var objCopy gencode.Object
	data, _ := c.Marshal(buf, &obj)
	c.Unmarshal(data, &objCopy)
}
func TestGOGOPBCodec(t *testing.T) {
	{
		var obj = gogopb.Object{A: 1024, B: 1024, C: 3.14, D: 3.1415926, E: "HelloWorld", F: true, G: []byte{255}, H: [][]byte{{128}, {255}}}
		var c = GOGOPBCodec{}
		var buf = make([]byte, 512)
		var objCopy gogopb.Object
		data, _ := c.Marshal(buf, &obj)
		c.Unmarshal(data, &objCopy)
	}
	{
		var obj = gogopb.Object{A: 1024, B: 1024, C: 3.14, D: 3.1415926, E: "HelloWorld", F: true, G: []byte{255}, H: [][]byte{{128}, {255}}}
		var c = GOGOPBCodec{}
		var buf = make([]byte, 0)
		var objCopy gogopb.Object
		data, _ := c.Marshal(buf, &obj)
		c.Unmarshal(data, &objCopy)
	}
	{
		var obj = model.Object{A: 1024, B: 1024, C: 3.14, D: 3.1415926, E: "HelloWorld", F: true, G: []byte{255}, H: [][]byte{{128}, {255}}}
		var c = GOGOPBCodec{}
		var buf = make([]byte, 512)
		var objCopy model.Object
		if _, err := c.Marshal(buf, &obj); err != ErrorGOGOPB {
			t.Error(ErrorGOGOPB)
		}
		if err := c.Unmarshal(nil, &objCopy); err != ErrorGOGOPB {
			t.Error(ErrorGOGOPB)
		}
	}
}

func TestMSGPCodec(t *testing.T) {
	var obj = msgp.Object{A: 1024, B: 1024, C: 3.14, D: 3.1415926, E: "HelloWorld", F: true, G: []byte{255}, H: [][]byte{{128}, {255}}}
	var c = MSGPCodec{}
	var buf = make([]byte, 512)
	var objCopy msgp.Object
	data, _ := c.Marshal(buf, &obj)
	c.Unmarshal(data, &objCopy)
}

func TestPBCodec(t *testing.T) {
	var obj = pb.Object{A: 1024, B: 1024, C: 3.14, D: 3.1415926, E: "HelloWorld", F: true, G: []byte{255}, H: [][]byte{{128}, {255}}}
	var c = codecpb.Codec{}
	var buf = make([]byte, 512)
	var objCopy pb.Object
	data, _ := c.Marshal(buf, &obj)
	c.Unmarshal(data, &objCopy)
}

func TestJSONCodec(t *testing.T) {
	var obj = model.Object{A: 1024, B: 1024, C: 3.14, D: 3.1415926, E: "HelloWorld", F: true, G: []byte{255}, H: [][]byte{{128}, {255}}}
	var c = JSONCodec{}
	var buf = make([]byte, 512)
	var objCopy *model.Object
	data, _ := c.Marshal(buf, &obj)
	c.Unmarshal(data, &objCopy)
}

func TestXMLCodec(t *testing.T) {
	var obj = model.Object{A: 1024, B: 1024, C: 3.14, D: 3.1415926, E: "HelloWorld", F: true}
	var c = XMLCodec{}
	var buf = make([]byte, 512)
	var objCopy *model.Object
	data, _ := c.Marshal(buf, &obj)
	c.Unmarshal(data, &objCopy)
}

func TestGOBCodec(t *testing.T) {
	var obj = model.Object{A: 1024, B: 1024, C: 3.14, D: 3.1415926, E: "HelloWorld", F: true, G: []byte{255}, H: [][]byte{{128}, {255}}}
	var c = GOBCodec{}
	var buf = make([]byte, 512)
	var objCopy *model.Object
	data, _ := c.Marshal(buf, &obj)
	c.Unmarshal(data, &objCopy)
}

func BenchmarkMarshalBYTES(t *testing.B) {
	var obj = []byte{128, 8, 128, 8, 195, 245, 72, 64, 74, 216, 18, 77, 251, 33, 9, 64, 10, 72, 101, 108, 108, 111, 87, 111, 114, 108, 100, 1, 1, 255, 2, 1, 128, 1, 255}
	var c = BYTESCodec{}
	var buf = make([]byte, 512)
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		c.Marshal(buf, &obj)
	}
}

func BenchmarkMarshalCODE(t *testing.B) {
	var obj = code.Object{A: 1024, B: 1024, C: 3.14, D: 3.1415926, E: "HelloWorld", F: true, G: []byte{255}, H: [][]byte{{128}, {255}}}
	var c = CODECodec{}
	var buf = make([]byte, 512)
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		c.Marshal(buf, &obj)
	}
}

func BenchmarkMarshalGENCODE(t *testing.B) {
	var obj = gencode.Object{A: 1024, B: 1024, C: 3.14, D: 3.1415926, E: "HelloWorld", F: true, G: []byte{255}, H: [][]byte{{128}, {255}}}
	var c = CODECodec{}
	var buf = make([]byte, 512)
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		c.Marshal(buf, &obj)
	}
}

func BenchmarkMarshalCODEPB(t *testing.B) {
	var obj = codepb.Object{A: 1024, B: 1024, C: 3.14, D: 3.1415926, E: "HelloWorld", F: true, G: []byte{255}, H: [][]byte{{128}, {255}}}
	var c = CODECodec{}
	var buf = make([]byte, 512)
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		c.Marshal(buf, &obj)
	}
}

func BenchmarkMarshalMSGP(t *testing.B) {
	var obj = msgp.Object{A: 1024, B: 1024, C: 3.14, D: 3.1415926, E: "HelloWorld", F: true, G: []byte{255}, H: [][]byte{{128}, {255}}}
	var c = MSGPCodec{}
	var buf = make([]byte, 512)
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		c.Marshal(buf, &obj)
	}
}

func BenchmarkMarshalGOGOPB(t *testing.B) {
	var obj = gogopb.Object{A: 1024, B: 1024, C: 3.14, D: 3.1415926, E: "HelloWorld", F: true, G: []byte{255}, H: [][]byte{{128}, {255}}}
	var c = GOGOPBCodec{}
	var buf = make([]byte, 512)
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		c.Marshal(buf, &obj)
	}
}

func BenchmarkMarshalPB(t *testing.B) {
	var obj = pb.Object{A: 1024, B: 1024, C: 3.14, D: 3.1415926, E: "HelloWorld", F: true, G: []byte{255}, H: [][]byte{{128}, {255}}}
	var c = codecpb.Codec{}
	var buf = make([]byte, 512)
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		c.Marshal(buf, &obj)
	}
}

func BenchmarkMarshalJSON(t *testing.B) {
	var obj = model.Object{A: 1024, B: 1024, C: 3.14, D: 3.1415926, E: "HelloWorld", F: true, G: []byte{255}, H: [][]byte{{128}, {255}}}
	var c = JSONCodec{}
	var buf = make([]byte, 512)
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		c.Marshal(buf, &obj)
	}
}

func BenchmarkMarshalXML(t *testing.B) {
	var obj = model.Object{A: 1024, B: 1024, C: 3.14, D: 3.1415926, E: "HelloWorld", F: true}
	var c = XMLCodec{}
	var buf = make([]byte, 512)
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		c.Marshal(buf, &obj)
	}
}

func BenchmarkMarshalGOB(t *testing.B) {
	var obj = model.Object{A: 1024, B: 1024, C: 3.14, D: 3.1415926, E: "HelloWorld", F: true, G: []byte{255}, H: [][]byte{{128}, {255}}}
	var c = GOBCodec{}
	var buf = make([]byte, 512)
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		c.Marshal(buf, &obj)
	}
}

func BenchmarkUnmarshalBYTES(t *testing.B) {
	var obj = []byte{128, 8, 128, 8, 195, 245, 72, 64, 74, 216, 18, 77, 251, 33, 9, 64, 10, 72, 101, 108, 108, 111, 87, 111, 114, 108, 100, 1, 1, 255, 2, 1, 128, 1, 255}
	var c = BYTESCodec{}
	var buf = make([]byte, 512)
	data, _ := c.Marshal(buf, &obj)
	var objCopy []byte
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		c.Unmarshal(data, &objCopy)
	}
}

func BenchmarkUnmarshalCODE(t *testing.B) {
	var obj = code.Object{A: 1024, B: 1024, C: 3.14, D: 3.1415926, E: "HelloWorld", F: true, G: []byte{255}, H: [][]byte{{128}, {255}}}
	var c = CODECodec{}
	var buf = make([]byte, 512)
	data, _ := c.Marshal(buf, &obj)
	var objCopy code.Object
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		c.Unmarshal(data, &objCopy)
	}
}

func BenchmarkUnmarshalGENCODE(t *testing.B) {
	var obj = gencode.Object{A: 1024, B: 1024, C: 3.14, D: 3.1415926, E: "HelloWorld", F: true, G: []byte{255}, H: [][]byte{{128}, {255}}}
	var c = CODECodec{}
	var buf = make([]byte, 512)
	data, _ := c.Marshal(buf, &obj)
	var objCopy gencode.Object
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		c.Unmarshal(data, &objCopy)
	}
}

func BenchmarkUnmarshalCODEPB(t *testing.B) {
	var obj = codepb.Object{A: 1024, B: 1024, C: 3.14, D: 3.1415926, E: "HelloWorld", F: true, G: []byte{255}, H: [][]byte{{128}, {255}}}
	var c = CODECodec{}
	var buf = make([]byte, 512)
	data, _ := c.Marshal(buf, &obj)
	var objCopy codepb.Object
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		c.Unmarshal(data, &objCopy)
	}
}

func BenchmarkUnmarshalMSGP(t *testing.B) {
	var obj = msgp.Object{A: 1024, B: 1024, C: 3.14, D: 3.1415926, E: "HelloWorld", F: true, G: []byte{255}, H: [][]byte{{128}, {255}}}
	var c = MSGPCodec{}
	var buf = make([]byte, 512)
	data, _ := c.Marshal(buf, &obj)
	var objCopy msgp.Object
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		c.Unmarshal(data, &objCopy)
	}
}

func BenchmarkUnmarshalGOGOPB(t *testing.B) {
	var obj = gogopb.Object{A: 1024, B: 1024, C: 3.14, D: 3.1415926, E: "HelloWorld", F: true, G: []byte{255}, H: [][]byte{{128}, {255}}}
	var c = GOGOPBCodec{}
	var buf = make([]byte, 512)
	data, _ := c.Marshal(buf, &obj)
	var objCopy gogopb.Object
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		c.Unmarshal(data, &objCopy)
	}
}

func BenchmarkUnmarshalPB(t *testing.B) {
	var obj = pb.Object{A: 1024, B: 1024, C: 3.14, D: 3.1415926, E: "HelloWorld", F: true, G: []byte{255}, H: [][]byte{{128}, {255}}}
	var c = codecpb.Codec{}
	var buf = make([]byte, 512)
	data, _ := c.Marshal(buf, &obj)
	var objCopy pb.Object
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		c.Unmarshal(data, &objCopy)
	}
}

func BenchmarkUnmarshalJSON(t *testing.B) {
	var obj = model.Object{A: 1024, B: 1024, C: 3.14, D: 3.1415926, E: "HelloWorld", F: true, G: []byte{255}, H: [][]byte{{128}, {255}}}
	var c = JSONCodec{}
	var buf = make([]byte, 512)
	data, _ := c.Marshal(buf, &obj)
	var objCopy *model.Object
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		c.Unmarshal(data, &objCopy)
	}
}

func BenchmarkUnmarshalXML(t *testing.B) {
	var obj = model.Object{A: 1024, B: 1024, C: 3.14, D: 3.1415926, E: "HelloWorld", F: true}
	var c = XMLCodec{}
	var buf = make([]byte, 512)
	data, _ := c.Marshal(buf, &obj)
	var objCopy *model.Object
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		c.Unmarshal(data, &objCopy)
	}
}

func BenchmarkUnmarshalGOB(t *testing.B) {
	var obj = model.Object{A: 1024, B: 1024, C: 3.14, D: 3.1415926, E: "HelloWorld", F: true, G: []byte{255}, H: [][]byte{{128}, {255}}}
	var c = GOBCodec{}
	var buf = make([]byte, 512)
	data, _ := c.Marshal(buf, &obj)
	var objCopy *model.Object
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		c.Unmarshal(data, &objCopy)
	}
}

func BenchmarkRoundtripBYTES(t *testing.B) {
	var obj = []byte{128, 8, 128, 8, 195, 245, 72, 64, 74, 216, 18, 77, 251, 33, 9, 64, 10, 72, 101, 108, 108, 111, 87, 111, 114, 108, 100, 1, 1, 255, 2, 1, 128, 1, 255}
	var c = BYTESCodec{}
	var buf = make([]byte, 512)
	var objCopy []byte
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		data, _ := c.Marshal(buf, &obj)
		c.Unmarshal(data, &objCopy)
	}
}

func BenchmarkRoundtripCODE(t *testing.B) {
	var obj = code.Object{A: 1024, B: 1024, C: 3.14, D: 3.1415926, E: "HelloWorld", F: true, G: []byte{255}, H: [][]byte{{128}, {255}}}
	var c = CODECodec{}
	var buf = make([]byte, 512)
	var objCopy code.Object
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		data, _ := c.Marshal(buf, &obj)
		c.Unmarshal(data, &objCopy)
	}
}

func BenchmarkRoundtripGENCODE(t *testing.B) {
	var obj = gencode.Object{A: 1024, B: 1024, C: 3.14, D: 3.1415926, E: "HelloWorld", F: true, G: []byte{255}, H: [][]byte{{128}, {255}}}
	var c = CODECodec{}
	var buf = make([]byte, 512)
	var objCopy gencode.Object
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		data, _ := c.Marshal(buf, &obj)
		c.Unmarshal(data, &objCopy)
	}
}

func BenchmarkRoundtripCODEPB(t *testing.B) {
	var obj = codepb.Object{A: 1024, B: 1024, C: 3.14, D: 3.1415926, E: "HelloWorld", F: true, G: []byte{255}, H: [][]byte{{128}, {255}}}
	var c = CODECodec{}
	var buf = make([]byte, 512)
	var objCopy codepb.Object
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		data, _ := c.Marshal(buf, &obj)
		c.Unmarshal(data, &objCopy)
	}
}

func BenchmarkRoundtripMSGP(t *testing.B) {
	var obj = msgp.Object{A: 1024, B: 1024, C: 3.14, D: 3.1415926, E: "HelloWorld", F: true, G: []byte{255}, H: [][]byte{{128}, {255}}}
	var c = MSGPCodec{}
	var buf = make([]byte, 512)
	var objCopy msgp.Object
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		data, _ := c.Marshal(buf, &obj)
		c.Unmarshal(data, &objCopy)
	}
}

func BenchmarkRoundtripGOGOPB(t *testing.B) {
	var obj = gogopb.Object{A: 1024, B: 1024, C: 3.14, D: 3.1415926, E: "HelloWorld", F: true, G: []byte{255}, H: [][]byte{{128}, {255}}}
	var c = GOGOPBCodec{}
	var buf = make([]byte, 512)
	var objCopy gogopb.Object
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		data, _ := c.Marshal(buf, &obj)
		c.Unmarshal(data, &objCopy)
	}
}

func BenchmarkRoundtripPB(t *testing.B) {
	var obj = pb.Object{A: 1024, B: 1024, C: 3.14, D: 3.1415926, E: "HelloWorld", F: true, G: []byte{255}, H: [][]byte{{128}, {255}}}
	var c = codecpb.Codec{}
	var buf = make([]byte, 512)
	var objCopy pb.Object
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		data, _ := c.Marshal(buf, &obj)
		c.Unmarshal(data, &objCopy)
	}
}

func BenchmarkRoundtripJSON(t *testing.B) {
	var obj = model.Object{A: 1024, B: 1024, C: 3.14, D: 3.1415926, E: "HelloWorld", F: true, G: []byte{255}, H: [][]byte{{128}, {255}}}
	var c = JSONCodec{}
	var buf = make([]byte, 512)
	var objCopy *model.Object
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		data, _ := c.Marshal(buf, &obj)
		c.Unmarshal(data, &objCopy)
	}
}

func BenchmarkRoundtripXML(t *testing.B) {
	var obj = model.Object{A: 1024, B: 1024, C: 3.14, D: 3.1415926, E: "HelloWorld", F: true}
	var c = XMLCodec{}
	var buf = make([]byte, 512)
	var objCopy *model.Object
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		data, _ := c.Marshal(buf, &obj)
		c.Unmarshal(data, &objCopy)
	}
}

func BenchmarkRoundtripGOB(t *testing.B) {
	var obj = model.Object{A: 1024, B: 1024, C: 3.14, D: 3.1415926, E: "HelloWorld", F: true, G: []byte{255}, H: [][]byte{{128}, {255}}}
	var c = GOBCodec{}
	var buf = make([]byte, 512)
	var objCopy *model.Object
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		data, _ := c.Marshal(buf, &obj)
		c.Unmarshal(data, &objCopy)
	}
}
