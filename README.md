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
	var c codec.Codec
	var data []byte

	//bytes
	obj_bytes:=[]byte{4,77,111,114,116,18,5,69,97,114,116,104}
	c=&codec.BytesCodec{}
	data,_=c.Encode(&obj_bytes)
	fmt.Printf("bytes Encode：length-%d,hex-%x\n",len(data),data)
	var obj_bytes_copy []byte
	c.Decode(data,&obj_bytes_copy)
	fmt.Println("bytes Decode：",obj_bytes_copy)

	//code
	obj_code:= code.Student{Name:"Mort",Age:18,Address:"Earth"}
	c=&codec.CodeCodec{make([]byte,100)}
	data,_=c.Encode(&obj_code)
	fmt.Printf("code Encode：length-%d,hex-%x\n",len(data),data)
	var obj_code_cp code.Student
	c.Decode(data,&obj_code_cp)
	fmt.Println("code Decode：",obj_code_cp)

	//gencode
	obj_gencode:= gencode.Student{Name:"Mort",Age:18,Address:"Earth"}
	c=&codec.CodeCodec{make([]byte,100)}
	data,_=c.Encode(&obj_gencode)
	fmt.Printf("gencode Encode：length-%d,hex-%x\n",len(data),data)
	var obj_gencode_cp gencode.Student
	c.Decode(data,&obj_gencode_cp)
	fmt.Println("gencode Decode：",obj_gencode_cp)

	//gogoproto
	obj_gogopb:= gogopb.Student{Name:"Mort",Age:18,Address:"Earth"}
	c=&codec.GogoProtoCodec{}
	data,_=c.Encode(&obj_gogopb)
	fmt.Printf("gogoproto Encode：length-%d,hex-%x\n",len(data),data)
	var obj_gogopb_cp gogopb.Student
	c.Decode(data,&obj_gogopb_cp)
	fmt.Println("gogoproto Decode：",obj_gogopb_cp)

	//msgp
	obj_msgp:= msgp.Student{Name:"Mort",Age:18,Address:"Earth"}
	c=&codec.MsgpCodec{}
	data,_=c.Encode(&obj_msgp)
	fmt.Printf("msgp Encode：length-%d,hex-%x\n",len(data),data)
	var obj_msgp_cp msgp.Student
	c.Decode(data,&obj_msgp_cp)
	fmt.Println("msgp Decode：",obj_msgp_cp)

	//proto
	obj_pb:=pb.Student{Name:"Mort",Age:18,Address:"Earth"}
	c=&codec.ProtoCodec{}
	data,_=c.Encode(&obj_pb)
	fmt.Printf("proto Encode：length-%d,hex-%x\n",len(data),data)
	var obj_pb_cp pb.Student
	c.Decode(data,&obj_pb_cp)
	fmt.Println("proto Decode：",obj_pb_cp)

	var obj=&model.Student{Name:"Mort",Age:18,Address:"Earth"}

	//json
	c=&codec.JsonCodec{}
	data,_=c.Encode(&obj)
	fmt.Printf("json Encode：length-%d,hex-%x\n",len(data),data)
	var obj_json *model.Student
	c.Decode(data,&obj_json)
	fmt.Println("json Decode：",obj_json)

	//xml
	c=&codec.XmlCodec{}
	data,_=c.Encode(&obj)
	fmt.Printf("xml Encode：length-%d,hex-%x\n",len(data),data)
	var obj_xml *model.Student
	c.Decode(data,&obj_xml)
	fmt.Println("xml Decode：",obj_xml)

	//gob
	c=&codec.GobCodec{}
	data,_=c.Encode(&obj)
	fmt.Printf("gob Encode：length-%d,hex-%x\n",len(data),data)
	var obj_gob *model.Student
	c.Decode(data,&obj_gob)
	fmt.Println("gob Decode：",obj_gob)
}
```

### Output
```
bytes Encode：length-12,hex-044d6f727412054561727468
bytes Decode： [4 77 111 114 116 18 5 69 97 114 116 104]
code Encode：length-15,hex-044d6f727412000000054561727468
code Decode： {Mort 18 Earth}
gencode Encode：length-15,hex-044d6f727412000000054561727468
gencode Decode： {Mort 18 Earth}
gogoproto Encode：length-15,hex-0a044d6f727410121a054561727468
gogoproto Decode： {Mort 18 Earth}
msgp Encode：length-30,hex-83a44e616d65a44d6f7274a341676512a741646472657373a54561727468
msgp Decode： {Mort 18 Earth}
proto Encode：length-15,hex-0a044d6f727410121a054561727468
proto Decode： {Mort 18 Earth {} [] 0}
json Encode：length-42,hex-7b224e616d65223a224d6f7274222c22416765223a31382c2241646472657373223a224561727468227d
json Decode： &{Mort 18 Earth}
xml Encode：length-73,hex-3c53747564656e743e3c4e616d653e4d6f72743c2f4e616d653e3c4167653e31383c2f4167653e3c416464726573733e45617274683c2f416464726573733e3c2f53747564656e743e
xml Decode： &{Mort 18 Earth}
gob Encode：length-70,hex-32ff810301010753747564656e7401ff8200010301044e616d65010c000103416765010400010741646472657373010c00000012ff8201044d6f727401240105456172746800
gob Decode： &{Mort 18 Earth}
```

### Benchmark
go test -v -run="none" -bench=. -benchtime=1s
```
goos: darwin
goarch: amd64
pkg: github.com/hslam/codec
BenchmarkEncodeBytes-4                 	1000000000	         0.405 ns/op
BenchmarkEncodeCodeNoReflect-4         	70534636	        17.2 ns/op
BenchmarkEncodeCode-4                  	47127092	        25.7 ns/op
BenchmarkEncodeGencodeNoReflect-4      	76368139	        15.4 ns/op
BenchmarkEncodeGencode-4               	48860787	        25.2 ns/op
BenchmarkEncodeGogoProto-4             	 9893486	       101 ns/op
BenchmarkEncodeMsgp-4                  	12350995	        98.2 ns/op
BenchmarkEncodeProto-4                 	 7483719	       141 ns/op
BenchmarkEncodeJson-4                  	 3576153	       334 ns/op
BenchmarkEncodeXml-4                   	  507687	      2424 ns/op
BenchmarkEncodeGob-4                   	  374424	      3272 ns/op
BenchmarkDecodeBytes-4                 	760777468	         1.56 ns/op
BenchmarkDecodeCodeNoReflect-4         	80665164	        14.6 ns/op
BenchmarkDecodeCode-4                  	17899275	        60.1 ns/op
BenchmarkDecodeGencodeNoReflect-4      	26136415	        44.2 ns/op
BenchmarkDecodeGencode-4               	12538338	        93.2 ns/op
BenchmarkDecodeGogoProto-4             	 8020778	       147 ns/op
BenchmarkDecodeMsgp-4                  	 7080954	       165 ns/op
BenchmarkDecodeProto-4                 	 6728606	       174 ns/op
BenchmarkDecodeJson-4                  	  987219	      1232 ns/op
BenchmarkDecodeXml-4                   	  217090	      5324 ns/op
BenchmarkDecodeGob-4                   	   60376	     19282 ns/op
BenchmarkRoundtripBytes-4              	746814086	         1.57 ns/op
BenchmarkRoundtripCodeNoReflect-4      	38134770	        31.7 ns/op
BenchmarkRoundtripCode-4               	13226766	        85.6 ns/op
BenchmarkRoundtripGencodeNoReflect-4   	19248774	        58.4 ns/op
BenchmarkRoundtripGencode-4            	 9241761	       126 ns/op
BenchmarkRoundtripGogoProto-4          	 4803943	       246 ns/op
BenchmarkRoundtripMsgp-4               	 4301529	       276 ns/op
BenchmarkRoundtripProto-4              	 3554709	       330 ns/op
BenchmarkRoundtripJson-4               	  752407	      1635 ns/op
BenchmarkRoundtripXml-4                	  139837	      8150 ns/op
BenchmarkRoundtripGob-4                	   50524	     22882 ns/op
PASS
ok  	github.com/hslam/codec	43.591s
```

### Licence
This package is licenced under a MIT licence (Copyright (c) 2019 Meng Huang)

### Authors
codec was written by Meng Huang.
