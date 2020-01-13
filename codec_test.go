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
	"testing"
)

func BenchmarkEncodeBYTES(t *testing.B) {
	var obj = []byte{128, 8, 128, 8, 195, 245, 72, 64, 74, 216, 18, 77, 251, 33, 9, 64, 10, 72, 101, 108, 108, 111, 87, 111, 114, 108, 100, 1, 1, 255, 2, 1, 128, 1, 255}
	var c = BYTESCodec{}
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		c.Encode(&obj)
	}
}

func BenchmarkEncodeCODE(t *testing.B) {
	var obj = code.Object{A: 1024, B: 1024, C: 3.14, D: 3.1415926, E: "HelloWorld", F: true, G: []byte{255}, H: [][]byte{{128}, {255}}}
	var c = CODECodec{Buffer: make([]byte, 512)}
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		c.Encode(&obj)
	}
}

func BenchmarkEncodeGENCODE(t *testing.B) {
	var obj = gencode.Object{A: 1024, B: 1024, C: 3.14, D: 3.1415926, E: "HelloWorld", F: true, G: []byte{255}, H: [][]byte{{128}, {255}}}
	var c = CODECodec{Buffer: make([]byte, 512)}
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		c.Encode(&obj)
	}
}

func BenchmarkEncodeCODEPB(t *testing.B) {
	var obj = codepb.Object{A: 1024, B: 1024, C: 3.14, D: 3.1415926, E: "HelloWorld", F: true, G: []byte{255}, H: [][]byte{{128}, {255}}}
	var c = CODECodec{Buffer: make([]byte, 512)}
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		c.Encode(&obj)
	}
}

func BenchmarkEncodeMSGP(t *testing.B) {
	var obj = msgp.Object{A: 1024, B: 1024, C: 3.14, D: 3.1415926, E: "HelloWorld", F: true, G: []byte{255}, H: [][]byte{{128}, {255}}}
	var c = MSGPCodec{Buffer: make([]byte, 512)}
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		c.Encode(&obj)
	}
}

func BenchmarkEncodeGOGOPB(t *testing.B) {
	var obj = gogopb.Object{A: 1024, B: 1024, C: 3.14, D: 3.1415926, E: "HelloWorld", F: true, G: []byte{255}, H: [][]byte{{128}, {255}}}
	var c = GOGOPBCodec{Buffer: make([]byte, 512)}
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		c.Encode(&obj)
	}
}

func BenchmarkEncodePB(t *testing.B) {
	var obj = pb.Object{A: 1024, B: 1024, C: 3.14, D: 3.1415926, E: "HelloWorld", F: true, G: []byte{255}, H: [][]byte{{128}, {255}}}
	var c = PBCodec{}
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		c.Encode(&obj)
	}
}

func BenchmarkEncodeJSON(t *testing.B) {
	var obj = model.Object{A: 1024, B: 1024, C: 3.14, D: 3.1415926, E: "HelloWorld", F: true, G: []byte{255}, H: [][]byte{{128}, {255}}}
	var c = JSONCodec{}
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		c.Encode(&obj)
	}
}

func BenchmarkEncodeXML(t *testing.B) {
	var obj = model.Object{A: 1024, B: 1024, C: 3.14, D: 3.1415926, E: "HelloWorld", F: true}
	var c = XMLCodec{}
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		c.Encode(&obj)
	}
}

func BenchmarkEncodeGOB(t *testing.B) {
	var obj = model.Object{A: 1024, B: 1024, C: 3.14, D: 3.1415926, E: "HelloWorld", F: true, G: []byte{255}, H: [][]byte{{128}, {255}}}
	var c = GOBCodec{}
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		c.Encode(&obj)
	}
}

func BenchmarkDecodeBYTES(t *testing.B) {
	var obj = []byte{128, 8, 128, 8, 195, 245, 72, 64, 74, 216, 18, 77, 251, 33, 9, 64, 10, 72, 101, 108, 108, 111, 87, 111, 114, 108, 100, 1, 1, 255, 2, 1, 128, 1, 255}
	var c = BYTESCodec{}
	data, _ := c.Encode(&obj)
	var objCopy []byte
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		c.Decode(data, &objCopy)
	}
}

func BenchmarkDecodeCODE(t *testing.B) {
	var obj = code.Object{A: 1024, B: 1024, C: 3.14, D: 3.1415926, E: "HelloWorld", F: true, G: []byte{255}, H: [][]byte{{128}, {255}}}
	var c = CODECodec{Buffer: make([]byte, 512)}
	data, _ := c.Encode(&obj)
	var objCopy code.Object
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		c.Decode(data, &objCopy)
	}
}

func BenchmarkDecodeGENCODE(t *testing.B) {
	var obj = gencode.Object{A: 1024, B: 1024, C: 3.14, D: 3.1415926, E: "HelloWorld", F: true, G: []byte{255}, H: [][]byte{{128}, {255}}}
	var c = CODECodec{Buffer: make([]byte, 512)}
	data, _ := c.Encode(&obj)
	var objCopy gencode.Object
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		c.Decode(data, &objCopy)
	}
}

func BenchmarkDecodeCODEPB(t *testing.B) {
	var obj = codepb.Object{A: 1024, B: 1024, C: 3.14, D: 3.1415926, E: "HelloWorld", F: true, G: []byte{255}, H: [][]byte{{128}, {255}}}
	var c = CODECodec{Buffer: make([]byte, 512)}
	data, _ := c.Encode(&obj)
	var objCopy codepb.Object
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		c.Decode(data, &objCopy)
	}
}

func BenchmarkDecodeMSGP(t *testing.B) {
	var obj = msgp.Object{A: 1024, B: 1024, C: 3.14, D: 3.1415926, E: "HelloWorld", F: true, G: []byte{255}, H: [][]byte{{128}, {255}}}
	var c = MSGPCodec{Buffer: make([]byte, 512)}
	data, _ := c.Encode(&obj)
	var objCopy msgp.Object
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		c.Decode(data, &objCopy)
	}
}

func BenchmarkDecodeGOGOPB(t *testing.B) {
	var obj = gogopb.Object{A: 1024, B: 1024, C: 3.14, D: 3.1415926, E: "HelloWorld", F: true, G: []byte{255}, H: [][]byte{{128}, {255}}}
	var c = GOGOPBCodec{Buffer: make([]byte, 512)}
	data, _ := c.Encode(&obj)
	var objCopy gogopb.Object
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		c.Decode(data, &objCopy)
	}
}

func BenchmarkDecodePB(t *testing.B) {
	var obj = pb.Object{A: 1024, B: 1024, C: 3.14, D: 3.1415926, E: "HelloWorld", F: true, G: []byte{255}, H: [][]byte{{128}, {255}}}
	var c = PBCodec{}
	data, _ := c.Encode(&obj)
	var objCopy pb.Object
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		c.Decode(data, &objCopy)
	}
}

func BenchmarkDecodeJSON(t *testing.B) {
	var obj = model.Object{A: 1024, B: 1024, C: 3.14, D: 3.1415926, E: "HelloWorld", F: true, G: []byte{255}, H: [][]byte{{128}, {255}}}
	var c = JSONCodec{}
	data, _ := c.Encode(&obj)
	var objCopy *model.Object
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		c.Decode(data, &objCopy)
	}
}

func BenchmarkDecodeXML(t *testing.B) {
	var obj = model.Object{A: 1024, B: 1024, C: 3.14, D: 3.1415926, E: "HelloWorld", F: true}
	var c = XMLCodec{}
	data, _ := c.Encode(&obj)
	var objCopy *model.Object
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		c.Decode(data, &objCopy)
	}
}

func BenchmarkDecodeGOB(t *testing.B) {
	var obj = model.Object{A: 1024, B: 1024, C: 3.14, D: 3.1415926, E: "HelloWorld", F: true, G: []byte{255}, H: [][]byte{{128}, {255}}}
	var c = GOBCodec{}
	data, _ := c.Encode(&obj)
	var objCopy *model.Object
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		c.Decode(data, &objCopy)
	}
}

func BenchmarkRoundtripBYTES(t *testing.B) {
	var obj = []byte{128, 8, 128, 8, 195, 245, 72, 64, 74, 216, 18, 77, 251, 33, 9, 64, 10, 72, 101, 108, 108, 111, 87, 111, 114, 108, 100, 1, 1, 255, 2, 1, 128, 1, 255}
	var c = BYTESCodec{}
	var objCopy []byte
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		data, _ := c.Encode(&obj)
		c.Decode(data, &objCopy)
	}
}

func BenchmarkRoundtripCODE(t *testing.B) {
	var obj = code.Object{A: 1024, B: 1024, C: 3.14, D: 3.1415926, E: "HelloWorld", F: true, G: []byte{255}, H: [][]byte{{128}, {255}}}
	var c = CODECodec{Buffer: make([]byte, 512)}
	var objCopy code.Object
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		data, _ := c.Encode(&obj)
		c.Decode(data, &objCopy)
	}
}

func BenchmarkRoundtripGENCODE(t *testing.B) {
	var obj = gencode.Object{A: 1024, B: 1024, C: 3.14, D: 3.1415926, E: "HelloWorld", F: true, G: []byte{255}, H: [][]byte{{128}, {255}}}
	var c = CODECodec{Buffer: make([]byte, 512)}
	var objCopy gencode.Object
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		data, _ := c.Encode(&obj)
		c.Decode(data, &objCopy)
	}
}

func BenchmarkRoundtripCODEPB(t *testing.B) {
	var obj = codepb.Object{A: 1024, B: 1024, C: 3.14, D: 3.1415926, E: "HelloWorld", F: true, G: []byte{255}, H: [][]byte{{128}, {255}}}
	var c = CODECodec{Buffer: make([]byte, 512)}
	var objCopy codepb.Object
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		data, _ := c.Encode(&obj)
		c.Decode(data, &objCopy)
	}
}

func BenchmarkRoundtripMSGP(t *testing.B) {
	var obj = msgp.Object{A: 1024, B: 1024, C: 3.14, D: 3.1415926, E: "HelloWorld", F: true, G: []byte{255}, H: [][]byte{{128}, {255}}}
	var c = MSGPCodec{Buffer: make([]byte, 512)}
	var objCopy msgp.Object
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		data, _ := c.Encode(&obj)
		c.Decode(data, &objCopy)
	}
}

func BenchmarkRoundtripGOGOPB(t *testing.B) {
	var obj = gogopb.Object{A: 1024, B: 1024, C: 3.14, D: 3.1415926, E: "HelloWorld", F: true, G: []byte{255}, H: [][]byte{{128}, {255}}}
	var c = GOGOPBCodec{Buffer: make([]byte, 512)}
	var objCopy gogopb.Object
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		data, _ := c.Encode(&obj)
		c.Decode(data, &objCopy)
	}
}

func BenchmarkRoundtripPB(t *testing.B) {
	var obj = pb.Object{A: 1024, B: 1024, C: 3.14, D: 3.1415926, E: "HelloWorld", F: true, G: []byte{255}, H: [][]byte{{128}, {255}}}
	var c = PBCodec{}
	var objCopy pb.Object
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		data, _ := c.Encode(&obj)
		c.Decode(data, &objCopy)
	}
}

func BenchmarkRoundtripJSON(t *testing.B) {
	var obj = model.Object{A: 1024, B: 1024, C: 3.14, D: 3.1415926, E: "HelloWorld", F: true, G: []byte{255}, H: [][]byte{{128}, {255}}}
	var c = JSONCodec{}
	var objCopy *model.Object
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		data, _ := c.Encode(&obj)
		c.Decode(data, &objCopy)
	}
}

func BenchmarkRoundtripXML(t *testing.B) {
	var obj = model.Object{A: 1024, B: 1024, C: 3.14, D: 3.1415926, E: "HelloWorld", F: true}
	var c = XMLCodec{}
	var objCopy *model.Object
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		data, _ := c.Encode(&obj)
		c.Decode(data, &objCopy)
	}
}

func BenchmarkRoundtripGOB(t *testing.B) {
	var obj = model.Object{A: 1024, B: 1024, C: 3.14, D: 3.1415926, E: "HelloWorld", F: true, G: []byte{255}, H: [][]byte{{128}, {255}}}
	var c = GOBCodec{}
	var objCopy *model.Object
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		data, _ := c.Encode(&obj)
		c.Decode(data, &objCopy)
	}
}
