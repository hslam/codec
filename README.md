# codec
A codec library is written in golang.

## Feature
* bytes
* [code](https://github.com/hslam/code "code")
* [gencode](https://github.com/andyleap/gencode "gencode")
* [gogoproto](https://github.com/gogo/protobuf "gogoproto")
* [msgp](https://github.com/tinylib/msgp "msgp")
* [proto](github.com/golang/protobuf "proto")
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
	"github.com/hslam/codec/example/model"
	"github.com/hslam/codec/example/pb"
	"github.com/hslam/codec/example/gencode"
	"github.com/hslam/codec/example/msgp"
	"github.com/hslam/codec/example/gogopb"
	"github.com/hslam/codec/example/code"
)

func main(){
	Bytes()
	Code()
	Gencode()
	GogoProto()
	Msgp()
	Proto()
	Json()
	Xml()
	Gob()
}
func Bytes()  {
	var obj=[]byte{0,4,0,0,0,4,0,0,0,0,0,0,195,245,72,64,74,216,18,77,251,33,9,64,10,72,101,108,108,111,87,111,114,108,100,1,1,0}
	c:=&codec.BytesCodec{}
	data,_:=c.Encode(&obj)
	fmt.Printf("bytes Encode：length-%d,hex-%x\n",len(data),data)
	var obj_bytes_copy []byte
	c.Decode(data,&obj_bytes_copy)
	fmt.Println("bytes Decode：",obj_bytes_copy)
}

func Code()  {
	var obj=code.Object{A:1024,B:1024,C:3.14,D:3.1415926,E:"HelloWorld",F:true,G:[]byte{0}}
	c:=&codec.CodeCodec{make([]byte,512)}
	data,_:=c.Encode(&obj)
	fmt.Printf("code Encode：length-%d,hex-%x\n",len(data),data)
	var obj_cp code.Object
	c.Decode(data,&obj_cp)
	fmt.Println("code Decode：",obj_cp)
}

func Gencode()  {
	var obj=gencode.Object{A:1024,B:1024,C:3.14,D:3.1415926,E:"HelloWorld",F:true,G:[]byte{0}}
	c:=&codec.CodeCodec{make([]byte,512)}
	data,_:=c.Encode(&obj)
	fmt.Printf("gencode Encode：length-%d,hex-%x\n",len(data),data)
	var obj_cp gencode.Object
	c.Decode(data,&obj_cp)
	fmt.Println("gencode Decode：",obj_cp)
}

func GogoProto()  {
	var obj=gogopb.Object{A:1024,B:1024,C:3.14,D:3.1415926,E:"HelloWorld",F:true,G:[]byte{0}}
	c:=&codec.GogoProtoCodec{}
	data,_:=c.Encode(&obj)
	fmt.Printf("gogoproto Encode：length-%d,hex-%x\n",len(data),data)
	var obj_cp gogopb.Object
	c.Decode(data,&obj_cp)
	fmt.Println("gogoproto Decode：",obj_cp)
}

func Msgp()  {
	var obj=msgp.Object{A:1024,B:1024,C:3.14,D:3.1415926,E:"HelloWorld",F:true,G:[]byte{0}}
	c:=&codec.MsgpCodec{}
	data,_:=c.Encode(&obj)
	fmt.Printf("msgp Encode：length-%d,hex-%x\n",len(data),data)
	var obj_cp msgp.Object
	c.Decode(data,&obj_cp)
	fmt.Println("msgp Decode：",obj_cp)
}

func Proto()  {
	var obj=pb.Object{A:1024,B:1024,C:3.14,D:3.1415926,E:"HelloWorld",F:true,G:[]byte{0}}
	c:=&codec.ProtoCodec{}
	data,_:=c.Encode(&obj)
	fmt.Printf("proto Encode：length-%d,hex-%x\n",len(data),data)
	var obj_cp pb.Object
	c.Decode(data,&obj_cp)
	fmt.Println("proto Decode：",obj_cp)
}

func Json()  {
	var obj=model.Object{A:1024,B:1024,C:3.14,D:3.1415926,E:"HelloWorld",F:true,G:[]byte{0}}
	c:=&codec.JsonCodec{}
	data,_:=c.Encode(&obj)
	fmt.Printf("json Encode：length-%d,hex-%x\n",len(data),data)
	var obj_cp model.Object
	c.Decode(data,&obj_cp)
	fmt.Println("json Decode：",obj_cp)
}

func Xml()  {
	var obj=model.Object{A:1024,B:1024,C:3.14,D:3.1415926,E:"HelloWorld",F:true,G:[]byte{0}}
	c:=&codec.XmlCodec{}
	data,_:=c.Encode(&obj)
	fmt.Printf("xml Encode：length-%d,hex-%x\n",len(data),data)
	var obj_cp model.Object
	c.Decode(data,&obj_cp)
	fmt.Println("xml Decode：",obj_cp)
}

func Gob()  {
	var obj=model.Object{A:1024,B:1024,C:3.14,D:3.1415926,E:"HelloWorld",F:true,G:[]byte{0}}
	c:=&codec.GobCodec{}
	data,_:=c.Encode(&obj)
	fmt.Printf("gob Encode：length-%d,hex-%x\n",len(data),data)
	var obj_cp model.Object
	c.Decode(data,&obj_cp)
	fmt.Println("gob Decode：",obj_cp)
}
```

### Output
```
bytes Encode：length-38,hex-000400000004000000000000c3f548404ad8124dfb2109400a48656c6c6f576f726c64010100
bytes Decode： [0 4 0 0 0 4 0 0 0 0 0 0 195 245 72 64 74 216 18 77 251 33 9 64 10 72 101 108 108 111 87 111 114 108 100 1 1 0]
code Encode：length-38,hex-000400000004000000000000c3f548404ad8124dfb2109400a48656c6c6f576f726c64010100
code Decode： {1024 1024 3.14 3.1415926 HelloWorld true [0]}
gencode Encode：length-38,hex-000400000004000000000000c3f548404ad8124dfb2109400a48656c6c6f576f726c64010100
gencode Decode： {1024 1024 3.14 3.1415926 HelloWorld true [0]}
gogoproto Encode：length-37,hex-0880081080081dc3f54840214ad8124dfb2109402a0a48656c6c6f576f726c6430013a0100
gogoproto Decode： {1024 1024 3.14 3.1415926 HelloWorld true [0]}
msgp Encode：length-50,hex-87a141cd0400a142cd0400a143ca4048f5c3a144cb400921fb4d12d84aa145aa48656c6c6f576f726c64a146c3a147c40100
msgp Decode： {1024 1024 3.14 3.1415926 HelloWorld true [0]}
proto Encode：length-37,hex-0880081080081dc3f54840214ad8124dfb2109402a0a48656c6c6f576f726c6430013a0100
proto Decode： {1024 1024 3.14 3.1415926 HelloWorld true [0] {} [] 0}
json Encode：length-79,hex-7b2241223a313032342c2242223a313032342c2243223a332e31342c2244223a332e313431353932362c2245223a2248656c6c6f576f726c64222c2246223a747275652c2247223a2241413d3d227d
json Decode： {1024 1024 3.14 3.1415926 HelloWorld true [0]}
xml Encode：length-104,hex-3c4f626a6563743e3c413e313032343c2f413e3c423e313032343c2f423e3c433e332e31343c2f433e3c443e332e313431353932363c2f443e3c453e48656c6c6f576f726c643c2f453e3c463e747275653c2f463e3c473eefbfbd3c2f473e3c2f4f626a6563743e
xml Decode： {1024 1024 3.14 3.1415926 HelloWorld true [239 191 189]}
gob Encode：length-109,hex-3eff81030101064f626a65637401ff82000107010141010600010142010600010143010800010144010800010145010c00010146010200010147010a0000002dff8201fe040001fe040001fb60b81e094001f84ad8124dfb210940010a48656c6c6f576f726c64010101010000
gob Decode： {1024 1024 3.14 3.1415926 HelloWorld true [0]}
```

### Benchmark
go test -v -run="none" -bench=. -benchtime=1s
```
goos: darwin
goarch: amd64
pkg: github.com/hslam/codec
BenchmarkEncodeBytes-4          	1000000000	         0.404 ns/op
BenchmarkEncodeCode-4           	33418732	        36.3 ns/op
BenchmarkEncodeGencode-4        	33693255	        33.7 ns/op
BenchmarkEncodeGogoProto-4      	10051132	       116 ns/op
BenchmarkEncodeMsgp-4           	 4890691	       239 ns/op
BenchmarkEncodeProto-4          	 5634568	       208 ns/op
BenchmarkEncodeJson-4           	 1508983	       806 ns/op
BenchmarkEncodeXml-4            	  277222	      4356 ns/op
BenchmarkEncodeGob-4            	  246861	      4569 ns/op
BenchmarkDecodeBytes-4          	746690206	         1.60 ns/op
BenchmarkDecodeCode-4           	15226951	        76.0 ns/op
BenchmarkDecodeGencode-4        	10616595	       117 ns/op
BenchmarkDecodeGogoProto-4      	 5676855	       223 ns/op
BenchmarkDecodeMsgp-4           	 6719329	       180 ns/op
BenchmarkDecodeProto-4          	 4791656	       245 ns/op
BenchmarkDecodeJson-4           	  490012	      2525 ns/op
BenchmarkDecodeXml-4            	  117184	     10828 ns/op
BenchmarkDecodeGob-4            	   47283	     25213 ns/op
BenchmarkRoundtripBytes-4       	704668684	         1.69 ns/op
BenchmarkRoundtripCode-4        	 9274729	       117 ns/op
BenchmarkRoundtripGencode-4     	 7427210	       161 ns/op
BenchmarkRoundtripGogoProto-4   	 3131545	       409 ns/op
BenchmarkRoundtripMsgp-4        	 3603441	       322 ns/op
BenchmarkRoundtripProto-4       	 2369629	       609 ns/op
BenchmarkRoundtripJson-4        	  332361	      3895 ns/op
BenchmarkRoundtripXml-4         	   81940	     14719 ns/op
BenchmarkRoundtripGob-4         	   39615	     29841 ns/op
PASS
ok  	github.com/hslam/codec	38.072s
```

### Licence
This package is licenced under a MIT licence (Copyright (c) 2019 Meng Huang)

### Authors
codec was written by Meng Huang.
