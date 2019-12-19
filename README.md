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

	//code_noreflect
	obj_code_noreflect:=code.Student{Name:"Mort",Age:18,Address:"Earth"}
	data,_=obj_code_noreflect.Marshal(nil)
	fmt.Printf("code_noreflect Encode：length-%d,hex-%x\n",len(data),data)
	var obj_code_noreflect_cp=&code.Student{}
	obj_code_noreflect_cp.Unmarshal(data)
	fmt.Println("code_noreflect Decode：",obj_code_noreflect_cp)

	//code
	obj_code:= code.Student{Name:"Mort",Age:18,Address:"Earth"}
	c=&codec.CodeCodec{}
	data,_=c.Encode(&obj_code)
	fmt.Printf("code Encode：length-%d,hex-%x\n",len(data),data)
	var obj_code_cp code.Student
	c.Decode(data,&obj_code_cp)
	fmt.Println("code Decode：",obj_code_cp)

	//gencode_noreflect
	obj_gencode_noreflect:= gencode.Student{Name:"Mort",Age:18,Address:"Earth"}
	data,_=obj_gencode_noreflect.Marshal(nil)
	fmt.Printf("gencode_noreflect Encode：length-%d,hex-%x\n",len(data),data)
	var obj_gencode_noreflect_cp=&gencode.Student{}
	obj_gencode_noreflect_cp.Unmarshal(data)
	fmt.Println("gencode_noreflect Decode：",obj_gencode_noreflect_cp)

	//gencode
	obj_gencode:= gencode.Student{Name:"Mort",Age:18,Address:"Earth"}
	c=&codec.CodeCodec{}
	data,_=c.Encode(&obj_gencode)
	fmt.Printf("gencode Encode：length-%d,hex-%x\n",len(data),data)
	var obj_gencode_cp gencode.Student
	c.Decode(data,&obj_gencode_cp)
	fmt.Println("gencode Decode：",obj_gencode_cp)

	//gogopb_noreflect
	obj_gogopb_noreflect:= gogopb.Student{Name:"Mort",Age:18,Address:"Earth"}
	data,_=obj_gogopb_noreflect.Marshal()
	fmt.Printf("gogoproto_noreflect Encode：length-%d,hex-%x\n",len(data),data)
	var obj_gogopb_noreflect_cp= gogopb.Student{}
	obj_gogopb_noreflect_cp.Unmarshal(data)
	fmt.Println("gogoproto_noreflect Decode：",obj_gogopb_noreflect_cp)

	//gogoproto
	obj_gogopb:= gogopb.Student{Name:"Mort",Age:18,Address:"Earth"}
	c=&codec.GogoProtoCodec{}
	data,_=c.Encode(&obj_gogopb)
	fmt.Printf("gogoproto Encode：length-%d,hex-%x\n",len(data),data)
	var obj_gogopb_cp gogopb.Student
	c.Decode(data,&obj_gogopb_cp)
	fmt.Println("gogoproto Decode：",obj_gogopb_cp)

	//msgp_noreflect
	obj_msgp_noreflect:= msgp.Student{Name:"Mort",Age:18,Address:"Earth"}
	data,_=obj_msgp_noreflect.MarshalMsg(nil)
	fmt.Printf("msgp_noreflect Encode：length-%d,hex-%x\n",len(data),data)
	var obj_msgp_noreflect_cp=&msgp.Student{}
	obj_msgp_noreflect_cp.UnmarshalMsg(data)
	fmt.Println("msgp_noreflect Decode：",obj_msgp_noreflect_cp)

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
code_noreflect Encode：length-12,hex-044d6f727412054561727468
code_noreflect Decode： &{Mort 18 Earth}
code Encode：length-12,hex-044d6f727412054561727468
code Decode： {Mort 18 Earth}
gencode_noreflect Encode：length-15,hex-044d6f727412000000054561727468
gencode_noreflect Decode： &{Mort 18 Earth}
gencode Encode：length-15,hex-044d6f727412000000054561727468
gencode Decode： {Mort 18 Earth}
gogoproto_noreflect Encode：length-15,hex-0a044d6f727410121a054561727468
gogoproto_noreflect Decode： {Mort 18 Earth}
gogoproto Encode：length-15,hex-0a044d6f727410121a054561727468
gogoproto Decode： {Mort 18 Earth}
msgp_noreflect Encode：length-30,hex-83a44e616d65a44d6f7274a341676512a741646472657373a54561727468
msgp_noreflect Decode： &{Mort 18 Earth}
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
BenchmarkEncodeBytes-4                   	1000000000	         0.396 ns/op
BenchmarkEncodeCodeNoReflect-4           	44250181	        26.1 ns/op
BenchmarkEncodeCode-4                    	35154648	        33.6 ns/op
BenchmarkEncodeGencodeNoReflect-4        	79097215	        15.2 ns/op
BenchmarkEncodeGencode-4                 	49227848	        25.3 ns/op
BenchmarkEncodeGogoProtoNoReflect-4      	22160108	        50.5 ns/op
BenchmarkEncodeGogoProto-4               	12349443	        95.5 ns/op
BenchmarkEncodeMsgpNoReflect-4           	13861084	        82.4 ns/op
BenchmarkEncodeMsgp-4                    	11951776	        97.4 ns/op
BenchmarkEncodeProto-4                   	 8286038	       141 ns/op
BenchmarkEncodeJson-4                    	 3527810	       332 ns/op
BenchmarkEncodeXml-4                     	  510241	      2350 ns/op
BenchmarkEncodeGob-4                     	  384271	      3113 ns/op
BenchmarkDecodeBytes-4                   	754935650	         1.58 ns/op
BenchmarkDecodeCodeNoReflect-4           	71890587	        16.5 ns/op
BenchmarkDecodeCode-4                    	18793789	        61.9 ns/op
BenchmarkDecodeGencodeNoReflect-4        	26502820	        44.3 ns/op
BenchmarkDecodeGencode-4                 	12645571	        92.4 ns/op
BenchmarkDecodeGogoProtoNoReflect-4      	21354776	        54.5 ns/op
BenchmarkDecodeGogoProto-4               	 8102512	       146 ns/op
BenchmarkDecodeMsgpNoReflect-4           	12120217	        99.8 ns/op
BenchmarkDecodeMsgp-4                    	 7671261	       160 ns/op
BenchmarkDecodeProto-4                   	 6770258	       175 ns/op
BenchmarkDecodeJson-4                    	  992134	      1232 ns/op
BenchmarkDecodeXml-4                     	  219211	      5297 ns/op
BenchmarkDecodeGob-4                     	   60319	     19831 ns/op
BenchmarkRoundtripBytes-4                	744601983	         1.57 ns/op
BenchmarkRoundtripCodeNoReflect-4        	28789464	        41.1 ns/op
BenchmarkRoundtripCode-4                 	11822552	       101 ns/op
BenchmarkRoundtripGencodeNoReflect-4     	19219207	        59.6 ns/op
BenchmarkRoundtripGencode-4              	 9354936	       126 ns/op
BenchmarkRoundtripGogoProtoNoReflect-4   	11069980	       109 ns/op
BenchmarkRoundtripGogoProto-4            	 4659026	       252 ns/op
BenchmarkRoundtripMsgpNoReflect-4        	 5807440	       203 ns/op
BenchmarkRoundtripMsgp-4                 	 4334667	       274 ns/op
BenchmarkRoundtripProto-4                	 3553176	       331 ns/op
BenchmarkRoundtripJson-4                 	  746168	      1651 ns/op
BenchmarkRoundtripXml-4                  	  143686	      8183 ns/op
BenchmarkRoundtripGob-4                  	   49285	     24268 ns/op
PASS
ok  	github.com/hslam/codec	49.804s
```

### Licence
This package is licenced under a MIT licence (Copyright (c) 2019 Meng Huang)

### Authors
codec was written by Meng Huang.
