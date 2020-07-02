# codec
[![GoDoc](https://godoc.org/github.com/hslam/codec?status.svg)](https://godoc.org/github.com/hslam/codec)
[![Build Status](https://travis-ci.org/hslam/codec.svg?branch=master)](https://travis-ci.org/hslam/codec)
[![codecov](https://codecov.io/gh/hslam/codec/branch/master/graph/badge.svg)](https://codecov.io/gh/hslam/codec)
[![Go Report Card](https://goreportcard.com/badge/github.com/hslam/codec)](https://goreportcard.com/report/github.com/hslam/codec)
[![GitHub release](https://img.shields.io/github/release/hslam/codec.svg)](https://github.com/hslam/codec/releases/latest)
[![LICENSE](https://img.shields.io/github/license/hslam/codec.svg?style=flat-square)](https://github.com/hslam/codec/blob/master/LICENSE)

Package codec implements encoding and decoding of multiple codecs

## Feature
* bytes
* [code](https://github.com/hslam/code "code")
* [gencode](https://github.com/andyleap/gencode "gencode")
* [gogopb](https://github.com/gogo/protobuf "gogopb")
* [msgp](https://github.com/tinylib/msgp "msgp")
* [pb](github.com/golang/protobuf "pb")
* json
* xml
* gob


## Get started

### Install
```
go get github.com/hslam/codec
```
### Import
```
import "github.com/hslam/codec"
```
### Usage
#### Example
```
package main

import (
	"fmt"
	"github.com/hslam/codec"
	"github.com/hslam/codec/example/code"
	"github.com/hslam/codec/example/codepb"
	"github.com/hslam/codec/example/gencode"
	"github.com/hslam/codec/example/gogopb"
	"github.com/hslam/codec/example/model"
	"github.com/hslam/codec/example/msgp"
	"github.com/hslam/codec/example/pb"
	codecpb "github.com/hslam/codec/pb"
)

func main() {
	BYTES()
	CODE()
	GENCODE()
	CODEPB()
	MSGP()
	GOGOPB()
	PB()
	JSON()
	XML()
	GOB()
}

//BYTES Example
func BYTES() {
	var obj = []byte{128, 8, 128, 8, 195, 245, 72, 64, 74, 216, 18, 77, 251, 33, 9, 64, 10, 72, 101, 108, 108, 111, 87, 111, 114, 108, 100, 1, 1, 255, 2, 1, 128, 1, 255}
	c := &codec.BYTESCodec{}
	var buf = make([]byte, 512)
	data, _ := c.Marshal(buf, &obj)
	fmt.Printf("bytes Encode：length-%d,hex-%x\n", len(data), data)
	var objCopy []byte
	c.Unmarshal(data, &objCopy)
	fmt.Println("bytes Decode：", objCopy)
}

//CODE Example
func CODE() {
	var obj = code.Object{A: 1024, B: 1024, C: 3.14, D: 3.1415926, E: "HelloWorld", F: true, G: []byte{255}, H: [][]byte{{128}, {255}}}
	c := &codec.CODECodec{}
	var buf = make([]byte, 512)
	data, _ := c.Marshal(buf, &obj)
	fmt.Printf("code Encode：length-%d,hex-%x\n", len(data), data)
	var objCopy code.Object
	c.Unmarshal(data, &objCopy)
	fmt.Println("code Decode：", objCopy)
}

//GENCODE Example
func GENCODE() {
	var obj = gencode.Object{A: 1024, B: 1024, C: 3.14, D: 3.1415926, E: "HelloWorld", F: true, G: []byte{255}, H: [][]byte{{128}, {255}}}
	c := &codec.CODECodec{}
	var buf = make([]byte, 512)
	data, _ := c.Marshal(buf, &obj)
	fmt.Printf("gencode Encode：length-%d,hex-%x\n", len(data), data)
	var objCopy gencode.Object
	c.Unmarshal(data, &objCopy)
	fmt.Println("gencode Decode：", objCopy)
}

//CODEPB Example
func CODEPB() {
	var obj = codepb.Object{A: 1024, B: 1024, C: 3.14, D: 3.1415926, E: "HelloWorld", F: true, G: []byte{255}, H: [][]byte{{128}, {255}}}
	c := &codec.CODECodec{}
	var buf = make([]byte, 512)
	data, _ := c.Marshal(buf, &obj)
	fmt.Printf("codepb Encode：length-%d,hex-%x\n", len(data), data)
	var objCopy codepb.Object
	c.Unmarshal(data, &objCopy)
	fmt.Println("codepb Decode：", objCopy)
}

//MSGP Example
func MSGP() {
	var obj = msgp.Object{A: 1024, B: 1024, C: 3.14, D: 3.1415926, E: "HelloWorld", F: true, G: []byte{255}, H: [][]byte{{128}, {255}}}
	c := &codec.MSGPCodec{}
	var buf = make([]byte, 512)
	data, _ := c.Marshal(buf, &obj)
	fmt.Printf("msgp Encode：length-%d,hex-%x\n", len(data), data)
	var objCopy msgp.Object
	c.Unmarshal(data, &objCopy)
	fmt.Println("msgp Decode：", objCopy)
}

//GOGOPB Example
func GOGOPB() {
	var obj = gogopb.Object{A: 1024, B: 1024, C: 3.14, D: 3.1415926, E: "HelloWorld", F: true, G: []byte{255}, H: [][]byte{{128}, {255}}}
	c := &codec.GOGOPBCodec{}
	var buf = make([]byte, 512)
	data, _ := c.Marshal(buf, &obj)
	fmt.Printf("gogopb Encode：length-%d,hex-%x\n", len(data), data)
	var objCopy gogopb.Object
	c.Unmarshal(data, &objCopy)
	fmt.Println("gogopb Decode：", objCopy)
}

//PB Example
func PB() {
	var obj = pb.Object{A: 1024, B: 1024, C: 3.14, D: 3.1415926, E: "HelloWorld", F: true, G: []byte{255}, H: [][]byte{{128}, {255}}}
	c := &codecpb.Codec{}
	var buf = make([]byte, 512)
	data, _ := c.Marshal(buf, &obj)
	fmt.Printf("pb Encode：length-%d,hex-%x\n", len(data), data)
	var objCopy pb.Object
	c.Unmarshal(data, &objCopy)
	fmt.Println("pb Decode：", objCopy)
}

//JSON Example
func JSON() {
	var obj = model.Object{A: 1024, B: 1024, C: 3.14, D: 3.1415926, E: "HelloWorld", F: true, G: []byte{255}, H: [][]byte{{128}, {255}}}
	c := &codec.JSONCodec{}
	var buf = make([]byte, 512)
	data, _ := c.Marshal(buf, &obj)
	fmt.Printf("json Encode：length-%d,hex-%x\n", len(data), data)
	var objCopy model.Object
	c.Unmarshal(data, &objCopy)
	fmt.Println("json Decode：", objCopy)
}

//XML Example
func XML() {
	var obj = model.Object{A: 1024, B: 1024, C: 3.14, D: 3.1415926, E: "HelloWorld", F: true}
	c := &codec.XMLCodec{}
	var buf = make([]byte, 512)
	data, _ := c.Marshal(buf, &obj)
	fmt.Printf("xml Encode：length-%d,hex-%x\n", len(data), data)
	var objCopy model.Object
	c.Unmarshal(data, &objCopy)
	fmt.Println("xml Decode：", objCopy)
}

//GOB Example
func GOB() {
	var obj = model.Object{A: 1024, B: 1024, C: 3.14, D: 3.1415926, E: "HelloWorld", F: true, G: []byte{255}, H: [][]byte{{128}, {255}}}
	c := &codec.GOBCodec{}
	var buf = make([]byte, 512)
	data, _ := c.Marshal(buf, &obj)
	fmt.Printf("gob Encode：length-%d,hex-%x\n", len(data), data)
	var objCopy model.Object
	c.Unmarshal(data, &objCopy)
	fmt.Println("gob Decode：", objCopy)
}
```

### Output
```
bytes Encode：length-35,hex-80088008c3f548404ad8124dfb2109400a48656c6c6f576f726c640101ff02018001ff
bytes Decode： [128 8 128 8 195 245 72 64 74 216 18 77 251 33 9 64 10 72 101 108 108 111 87 111 114 108 100 1 1 255 2 1 128 1 255]
code Encode：length-35,hex-80088008c3f548404ad8124dfb2109400a48656c6c6f576f726c640101ff02018001ff
code Decode： {1024 1024 3.14 3.1415926 HelloWorld true [255] [[128] [255]]}
gencode Encode：length-43,hex-000400000004000000000000c3f548404ad8124dfb2109400a48656c6c6f576f726c640101ff02018001ff
gencode Decode： {1024 1024 3.14 3.1415926 HelloWorld true [255] [[128] [255]]}
codepb Encode：length-43,hex-0880081080081dc3f54840214ad8124dfb2109402a0a48656c6c6f576f726c6430013a01ff4201804201ff
codepb Decode： {1024 1024 3.14 3.1415926 HelloWorld true [255] [[128] [255]]}
msgp Encode：length-59,hex-88a141cd0400a142cd0400a143ca4048f5c3a144cb400921fb4d12d84aa145aa48656c6c6f576f726c64a146c3a147c401ffa14892c40180c401ff
msgp Decode： {1024 1024 3.14 3.1415926 HelloWorld true [255] [[128] [255]]}
gogopb Encode：length-43,hex-0880081080081dc3f54840214ad8124dfb2109402a0a48656c6c6f576f726c6430013a01ff4201804201ff
gogopb Decode： {1024 1024 3.14 3.1415926 HelloWorld true [255] [[128] [255]]}
pb Encode：length-43,hex-0880081080081dc3f54840214ad8124dfb2109402a0a48656c6c6f576f726c6430013a01ff4201804201ff
pb Decode： {1024 1024 3.14 3.1415926 HelloWorld true [255] [[128] [255]] {} [] 0}
json Encode：length-99,hex-7b2241223a313032342c2242223a313032342c2243223a332e31342c2244223a332e313431353932362c2245223a2248656c6c6f576f726c64222c2246223a747275652c2247223a222f773d3d222c2248223a5b2267413d3d222c222f773d3d225d7d
json Decode： {1024 1024 3.14 3.1415926 HelloWorld true [255] [[128] [255]]}
xml Encode：length-94,hex-3c4f626a6563743e3c413e313032343c2f413e3c423e313032343c2f423e3c433e332e31343c2f433e3c443e332e313431353932363c2f443e3c453e48656c6c6f576f726c643c2f453e3c463e747275653c2f463e3c2f4f626a6563743e
xml Decode： {1024 1024 3.14 3.1415926 HelloWorld true [] []}
gob Encode：length-146,hex-45ff81030101064f626a65637401ff82000108010141010600010142010600010143010800010144010800010145010c00010146010200010147010a0001014801ff8400000017ff83020101095b5d5b5d75696e743801ff8400010a000033ff8201fe040001fe040001fb60b81e094001f84ad8124dfb210940010a48656c6c6f576f726c6401010101ff0102018001ff00
gob Decode： {1024 1024 3.14 3.1415926 HelloWorld true [255] [[128] [255]]}
```

### Benchmark
go test -v -run="none" -bench=. -benchtime=1s
```
goos: darwin
goarch: amd64
pkg: github.com/hslam/codec
BenchmarkMarshalBYTES-4       	1000000000	         0.387 ns/op
BenchmarkMarshalCODE-4        	18549778	        63.6 ns/op
BenchmarkMarshalGENCODE-4     	21588637	        53.9 ns/op
BenchmarkMarshalCODEPB-4      	18010692	        64.6 ns/op
BenchmarkMarshalMSGP-4        	14500226	        81.9 ns/op
BenchmarkMarshalGOGOPB-4      	12407948	        95.2 ns/op
BenchmarkMarshalPB-4          	 3934147	       301 ns/op
BenchmarkMarshalJSON-4        	 1263380	       954 ns/op
BenchmarkMarshalXML-4         	  287898	      4031 ns/op
BenchmarkMarshalGOB-4         	  195620	      5773 ns/op
BenchmarkUnmarshalBYTES-4     	1000000000	         0.956 ns/op
BenchmarkUnmarshalCODE-4      	23460816	        50.3 ns/op
BenchmarkUnmarshalGENCODE-4   	16897834	        69.9 ns/op
BenchmarkUnmarshalCODEPB-4    	16739599	        69.3 ns/op
BenchmarkUnmarshalMSGP-4      	 3988837	       301 ns/op
BenchmarkUnmarshalGOGOPB-4    	 2672503	       408 ns/op
BenchmarkUnmarshalPB-4        	 2301373	       518 ns/op
BenchmarkUnmarshalJSON-4      	  419853	      2880 ns/op
BenchmarkUnmarshalXML-4       	  141288	      8157 ns/op
BenchmarkUnmarshalGOB-4       	   47277	     25616 ns/op
BenchmarkRoundtripBYTES-4     	1000000000	         1.10 ns/op
BenchmarkRoundtripCODE-4      	10118427	       117 ns/op
BenchmarkRoundtripGENCODE-4   	 9350256	       126 ns/op
BenchmarkRoundtripCODEPB-4    	 8032484	       141 ns/op
BenchmarkRoundtripMSGP-4      	 2962431	       402 ns/op
BenchmarkRoundtripGOGOPB-4    	 2567640	       464 ns/op
BenchmarkRoundtripPB-4        	 1364863	       869 ns/op
BenchmarkRoundtripJSON-4      	  288883	      4073 ns/op
BenchmarkRoundtripXML-4       	   92517	     12539 ns/op
BenchmarkRoundtripGOB-4       	   37496	     31739 ns/op
PASS
ok  	github.com/hslam/codec	40.974s
```

### License
This package is licensed under a MIT license (Copyright (c) 2019 Meng Huang)

### Authors
codec was written by Meng Huang.
