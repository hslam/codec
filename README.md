# codec
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
	data, _ := c.Encode(&obj)
	fmt.Printf("bytes Encode：length-%d,hex-%x\n", len(data), data)
	var objCopy []byte
	c.Decode(data, &objCopy)
	fmt.Println("bytes Decode：", objCopy)
}

//CODE Example
func CODE() {
	var obj = code.Object{A: 1024, B: 1024, C: 3.14, D: 3.1415926, E: "HelloWorld", F: true, G: []byte{255}, H: [][]byte{{128}, {255}}}
	c := &codec.CODECodec{Buffer: make([]byte, 512)}
	data, _ := c.Encode(&obj)
	fmt.Printf("code Encode：length-%d,hex-%x\n", len(data), data)
	var objCopy code.Object
	c.Decode(data, &objCopy)
	fmt.Println("code Decode：", objCopy)
}

//GENCODE Example
func GENCODE() {
	var obj = gencode.Object{A: 1024, B: 1024, C: 3.14, D: 3.1415926, E: "HelloWorld", F: true, G: []byte{255}, H: [][]byte{{128}, {255}}}
	c := &codec.CODECodec{Buffer: make([]byte, 512)}
	data, _ := c.Encode(&obj)
	fmt.Printf("gencode Encode：length-%d,hex-%x\n", len(data), data)
	var objCopy gencode.Object
	c.Decode(data, &objCopy)
	fmt.Println("gencode Decode：", objCopy)
}

//CODEPB Example
func CODEPB() {
	var obj = codepb.Object{A: 1024, B: 1024, C: 3.14, D: 3.1415926, E: "HelloWorld", F: true, G: []byte{255}, H: [][]byte{{128}, {255}}}
	c := &codec.CODECodec{Buffer: make([]byte, 512)}
	data, _ := c.Encode(&obj)
	fmt.Printf("codepb Encode：length-%d,hex-%x\n", len(data), data)
	var objCopy codepb.Object
	c.Decode(data, &objCopy)
	fmt.Println("codepb Decode：", objCopy)
}

//MSGP Example
func MSGP() {
	var obj = msgp.Object{A: 1024, B: 1024, C: 3.14, D: 3.1415926, E: "HelloWorld", F: true, G: []byte{255}, H: [][]byte{{128}, {255}}}
	c := &codec.MSGPCodec{Buffer: make([]byte, 512)}
	data, _ := c.Encode(&obj)
	fmt.Printf("msgp Encode：length-%d,hex-%x\n", len(data), data)
	var objCopy msgp.Object
	c.Decode(data, &objCopy)
	fmt.Println("msgp Decode：", objCopy)
}

//GOGOPB Example
func GOGOPB() {
	var obj = gogopb.Object{A: 1024, B: 1024, C: 3.14, D: 3.1415926, E: "HelloWorld", F: true, G: []byte{255}, H: [][]byte{{128}, {255}}}
	c := &codec.GOGOPBCodec{Buffer: make([]byte, 512)}
	data, _ := c.Encode(&obj)
	fmt.Printf("gogopb Encode：length-%d,hex-%x\n", len(data), data)
	var objCopy gogopb.Object
	c.Decode(data, &objCopy)
	fmt.Println("gogopb Decode：", objCopy)
}

//PB Example
func PB() {
	var obj = pb.Object{A: 1024, B: 1024, C: 3.14, D: 3.1415926, E: "HelloWorld", F: true, G: []byte{255}, H: [][]byte{{128}, {255}}}
	c := &codec.PBCodec{}
	data, _ := c.Encode(&obj)
	fmt.Printf("pb Encode：length-%d,hex-%x\n", len(data), data)
	var objCopy pb.Object
	c.Decode(data, &objCopy)
	fmt.Println("pb Decode：", objCopy)
}

//JSON Example
func JSON() {
	var obj = model.Object{A: 1024, B: 1024, C: 3.14, D: 3.1415926, E: "HelloWorld", F: true, G: []byte{255}, H: [][]byte{{128}, {255}}}
	c := &codec.JSONCodec{}
	data, _ := c.Encode(&obj)
	fmt.Printf("json Encode：length-%d,hex-%x\n", len(data), data)
	var objCopy model.Object
	c.Decode(data, &objCopy)
	fmt.Println("json Decode：", objCopy)
}

//XML Example
func XML() {
	var obj = model.Object{A: 1024, B: 1024, C: 3.14, D: 3.1415926, E: "HelloWorld", F: true}
	c := &codec.XMLCodec{}
	data, _ := c.Encode(&obj)
	fmt.Printf("xml Encode：length-%d,hex-%x\n", len(data), data)
	var objCopy model.Object
	c.Decode(data, &objCopy)
	fmt.Println("xml Decode：", objCopy)
}

//GOB Example
func GOB() {
	var obj = model.Object{A: 1024, B: 1024, C: 3.14, D: 3.1415926, E: "HelloWorld", F: true, G: []byte{255}, H: [][]byte{{128}, {255}}}
	c := &codec.GOBCodec{}
	data, _ := c.Encode(&obj)
	fmt.Printf("gob Encode：length-%d,hex-%x\n", len(data), data)
	var objCopy model.Object
	c.Decode(data, &objCopy)
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
BenchmarkEncodeBYTES-4        	1000000000	         0.388 ns/op
BenchmarkEncodeCODE-4         	18417238	        63.6 ns/op
BenchmarkEncodeGENCODE-4      	22390467	        51.2 ns/op
BenchmarkEncodeCODEPB-4       	17877290	        65.8 ns/op
BenchmarkEncodeMSGP-4         	13571006	        84.9 ns/op
BenchmarkEncodeGOGOPB-4       	15944592	        73.4 ns/op
BenchmarkEncodePB-4           	 4892133	       242 ns/op
BenchmarkEncodeJSON-4         	 1251392	       943 ns/op
BenchmarkEncodeXML-4          	  291318	      4071 ns/op
BenchmarkEncodeGOB-4          	  203014	      5645 ns/op
BenchmarkDecodeBYTES-4        	1000000000	         0.960 ns/op
BenchmarkDecodeCODE-4         	23268570	        50.4 ns/op
BenchmarkDecodeGENCODE-4      	17301549	        67.7 ns/op
BenchmarkDecodeCODEPB-4       	16613464	        70.9 ns/op
BenchmarkDecodeMSGP-4         	 3865035	       307 ns/op
BenchmarkDecodeGOGOPB-4       	 2505402	       426 ns/op
BenchmarkDecodePB-4           	 2665432	       450 ns/op
BenchmarkDecodeJSON-4         	  421455	      2856 ns/op
BenchmarkDecodeXML-4          	  142396	      8102 ns/op
BenchmarkDecodeGOB-4          	   47439	     24950 ns/op
BenchmarkRoundtripBYTES-4     	1000000000	         1.12 ns/op
BenchmarkRoundtripCODE-4      	10073326	       118 ns/op
BenchmarkRoundtripGENCODE-4   	 9897543	       120 ns/op
BenchmarkRoundtripCODEPB-4    	 8407045	       141 ns/op
BenchmarkRoundtripMSGP-4      	 2932954	       409 ns/op
BenchmarkRoundtripGOGOPB-4    	 2878573	       451 ns/op
BenchmarkRoundtripPB-4        	 1634523	       730 ns/op
BenchmarkRoundtripJSON-4      	  297882	      3951 ns/op
BenchmarkRoundtripXML-4       	   93462	     12792 ns/op
BenchmarkRoundtripGOB-4       	   37586	     31512 ns/op
PASS
ok  	github.com/hslam/codec	41.569s
```

### Licence
This package is licensed under a MIT license (Copyright (c) 2019 Meng Huang)

### Authors
codec was written by Meng Huang.
