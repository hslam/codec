# codec
A codec library written in golang.

## Feature
* bytes
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
go get hslam.com/git/x/codec
```
### Import
```
import "hslam.com/git/x/codec"
```
### Usage
#### Example
```
package main

import (
	"fmt"
	"hslam.com/git/x/codec"
	"hslam.com/git/x/codec/example/model"
	"hslam.com/git/x/codec/example/pb"
	"hslam.com/git/x/codec/example/bytes"
	"hslam.com/git/x/codec/example/gencode"
	"hslam.com/git/x/codec/example/msgp"
	"hslam.com/git/x/codec/example/gogopb"
)

func main(){
	var c codec.Codec
	var data []byte

	//bytes_noreflect
	object_noreflect:=bytes.Student{Name:"Meng",Age:18,Address:"Earth"}
	data,_=object_noreflect.Marshal(nil)
	fmt.Printf("bytes_noreflect Encode：%x\n",data)
	var object_noreflect_copy =&bytes.Student{}
	object_noreflect_copy.Unmarshal(data)
	fmt.Println("bytes_noreflect Decode：",object_noreflect_copy)

	//bytes
	object:=bytes.Student{Name:"Meng",Age:18,Address:"Earth"}
	obj_bytes,_:=object.Marshal(nil)
	c=&codec.BytesCodec{}
	data,_=c.Encode(&obj_bytes)
	fmt.Printf("bytes Encode：%x\n",data)
	var object_copy =&bytes.Student{}
	var obj_bytes_copy []byte
	c.Decode(data,&obj_bytes_copy)
	object_copy.Unmarshal(obj_bytes_copy)
	fmt.Println("bytes Decode：",object_copy)

	//gencode_noreflect
	obj_gencode_noreflect:= gencode.Student{Name:"Meng",Age:18,Address:"Earth"}
	data,_=obj_gencode_noreflect.Marshal(nil)
	fmt.Printf("gencode_noreflect Encode：%x\n",data)
	var obj_gencode_noreflect_cp=&gencode.Student{}
	obj_gencode_noreflect_cp.Unmarshal(data)
	fmt.Println("gencode_noreflect Decode：",obj_gencode_noreflect_cp)

	//gencode
	obj_gencode:= gencode.Student{Name:"Meng",Age:18,Address:"Earth"}
	c=&codec.GencodeCodec{}
	data,_=c.Encode(&obj_gencode)
	fmt.Printf("gencode Encode：%x\n",data)
	var obj_gencode_cp gencode.Student
	c.Decode(data,&obj_gencode_cp)
	fmt.Println("gencode Decode：",obj_gencode_cp)

	//gogopb_noreflect
	obj_gogopb_noreflect:= gogopb.Student{Name:"Meng",Age:18,Address:"Earth"}
	data,_=obj_gogopb_noreflect.Marshal()
	fmt.Printf("gogoproto_noreflect Encode：%x\n",data)
	var obj_gogopb_noreflect_cp= gogopb.Student{}
	obj_gogopb_noreflect_cp.Unmarshal(data)
	fmt.Println("gogoproto_noreflect Decode：",obj_gogopb_noreflect_cp)

	//gogoproto
	obj_gogopb:= gogopb.Student{Name:"Meng",Age:18,Address:"Earth"}
	c=&codec.GogoProtoCodec{}
	data,_=c.Encode(&obj_gogopb)
	fmt.Printf("gogoproto Encode：%x\n",data)
	var obj_gogopb_cp gogopb.Student
	c.Decode(data,&obj_gogopb_cp)
	fmt.Println("gogoproto Decode：",obj_gogopb_cp)

	//msgp_noreflect
	obj_msgp_noreflect:= msgp.Student{Name:"Meng",Age:18,Address:"Earth"}
	data,_=obj_msgp_noreflect.MarshalMsg(nil)
	fmt.Printf("msgp_noreflect Encode：%x\n",data)
	var obj_msgp_noreflect_cp=&msgp.Student{}
	obj_msgp_noreflect_cp.UnmarshalMsg(data)
	fmt.Println("msgp_noreflect Decode：",obj_msgp_noreflect_cp)

	//msgp
	obj_msgp:= msgp.Student{Name:"Meng",Age:18,Address:"Earth"}
	c=&codec.MsgpCodec{}
	data,_=c.Encode(&obj_msgp)
	fmt.Printf("msgp Encode：%x\n",data)
	var obj_msgp_cp msgp.Student
	c.Decode(data,&obj_msgp_cp)
	fmt.Println("msgp Decode：",obj_msgp_cp)

	//proto
	obj_pb:=pb.Student{Name:"Meng",Age:18,Address:"Earth"}
	c=&codec.ProtoCodec{}
	data,_=c.Encode(&obj_pb)
	fmt.Printf("proto Encode：%x\n",data)
	var obj_pb_cp pb.Student
	c.Decode(data,&obj_pb_cp)
	fmt.Println("proto Decode：",obj_pb_cp)

	var obj=&model.Student{Name:"Meng",Age:18,Address:"Earth"}

	//json
	c=&codec.JsonCodec{}
	data,_=c.Encode(&obj)
	fmt.Printf("json Encode：%x\n",data)
	var obj_json *model.Student
	c.Decode(data,&obj_json)
	fmt.Println("json Decode：",obj_json)

	//xml
	c=&codec.XmlCodec{}
	data,_=c.Encode(&obj)
	fmt.Printf("xml Encode：%x\n",data)
	var obj_xml *model.Student
	c.Decode(data,&obj_xml)
	fmt.Println("xml Decode：",obj_xml)

	//gob
	c=&codec.GobCodec{}
	data,_=c.Encode(&obj)
	fmt.Printf("gob Encode：%x\n",data)
	var obj_gob *model.Student
	c.Decode(data,&obj_gob)
	fmt.Println("gob Decode：",obj_gob)
}
```

### Output
```
bytes_noreflect Encode：044d6f7274124561727468
bytes_noreflect Decode： &{Meng 18 Earth}
bytes Encode：044d6f7274124561727468
bytes Decode： &{Meng 18 Earth}
gencode_noreflect Encode：044d6f727412000000054561727468
gencode_noreflect Decode： &{Meng 18 Earth}
gencode Encode：044d6f727412000000054561727468
gencode Decode： {Meng 18 Earth}
gogoproto_noreflect Encode：0a044d6f727410121a054561727468
gogoproto_noreflect Decode： {Meng 18 Earth}
gogoproto Encode：0a044d6f727410121a054561727468
gogoproto Decode： {Meng 18 Earth}
msgp_noreflect Encode：83a44e616d65a44d6f7274a341676512a741646472657373a54561727468
msgp_noreflect Decode： &{Meng 18 Earth}
msgp Encode：83a44e616d65a44d6f7274a341676512a741646472657373a54561727468
msgp Decode： {Meng 18 Earth}
proto Encode：0a044d6f727410121a054561727468
proto Decode： {Meng 18 Earth {} [] 0}
json Encode：7b224e616d65223a224d6f7274222c22416765223a31382c2241646472657373223a224561727468227d
json Decode： &{Meng 18 Earth}
xml Encode：3c53747564656e743e3c4e616d653e4d6f72743c2f4e616d653e3c4167653e31383c2f4167653e3c416464726573733e45617274683c2f416464726573733e3c2f53747564656e743e
xml Decode： &{Meng 18 Earth}
gob Encode：32ff810301010753747564656e7401ff8200010301044e616d65010c000103416765010400010741646472657373010c00000012ff8201044d6f727401240105456172746800
gob Decode： &{Meng 18 Earth}
```

### Benchmark
go test -v -run="none" -bench=. -benchtime=1s
```
goos: darwin
goarch: amd64
pkg: hslam.com/git/x/codec
BenchmarkEncodeOnlyBytes-4               	1000000000	         0.619 ns/op
BenchmarkEncodeBytesNoReflect-4          	57501272	        20.9 ns/op
BenchmarkEncodeBytes-4                   	55385464	        21.6 ns/op
BenchmarkEncodeGencodeNoReflect-4        	78135106	        15.2 ns/op
BenchmarkEncodeGencode-4                 	49197519	        24.4 ns/op
BenchmarkEncodeGogoProtoNoReflect-4      	21952580	        50.7 ns/op
BenchmarkEncodeGogoProto-4               	11942618	        99.5 ns/op
BenchmarkEncodeMsgpNoReflect-4           	13640454	        83.7 ns/op
BenchmarkEncodeMsgp-4                    	11700638	       101 ns/op
BenchmarkEncodeProto-4                   	 8212214	       142 ns/op
BenchmarkEncodeJson-4                    	 3611998	       333 ns/op
BenchmarkEncodeXml-4                     	  449294	      2483 ns/op
BenchmarkEncodeGob-4                     	  367005	      3645 ns/op
BenchmarkDecodeOnlyBytes-4               	720849464	         1.69 ns/op
BenchmarkDecodeBytesNoReflect-4          	32786457	        38.1 ns/op
BenchmarkDecodeBytes-4                   	30563037	        40.1 ns/op
BenchmarkDecodeGencodeNoReflect-4        	24874905	        49.4 ns/op
BenchmarkDecodeGencode-4                 	12512212	        93.1 ns/op
BenchmarkDecodeGogoProtoNoReflect-4      	19986607	        57.0 ns/op
BenchmarkDecodeGogoProto-4               	 8013654	       150 ns/op
BenchmarkDecodeMsgpNoReflect-4           	12076669	       105 ns/op
BenchmarkDecodeMsgp-4                    	 7704907	       174 ns/op
BenchmarkDecodeProto-4                   	 6432402	       181 ns/op
BenchmarkDecodeJson-4                    	  944318	      1330 ns/op
BenchmarkDecodeXml-4                     	  200362	      6512 ns/op
BenchmarkDecodeGob-4                     	   53953	     21825 ns/op
BenchmarkRoundtripOnlyBytes-4            	722997019	         1.66 ns/op
BenchmarkRoundtripBytesNoReflect-4       	18536655	        60.6 ns/op
BenchmarkRoundtripBytes-4                	19376626	        63.1 ns/op
BenchmarkRoundtripGencodeNoReflect-4     	18699894	        62.0 ns/op
BenchmarkRoundtripGencode-4              	 8545854	       132 ns/op
BenchmarkRoundtripGogoProtoNoReflect-4   	10241878	       112 ns/op
BenchmarkRoundtripGogoProto-4            	 4525212	       262 ns/op
BenchmarkRoundtripMsgpNoReflect-4        	 5603804	       207 ns/op
BenchmarkRoundtripMsgp-4                 	 4277578	       296 ns/op
BenchmarkRoundtripProto-4                	 3608727	       338 ns/op
BenchmarkRoundtripJson-4                 	  661342	      1746 ns/op
BenchmarkRoundtripXml-4                  	  141186	      8160 ns/op
BenchmarkRoundtripGob-4                  	   50650	     23135 ns/op
PASS
ok  	hslam.com/git/x/codec	51.592s
```

### Licence
This package is licenced under a MIT licence (Copyright (c) 2019 Meng Huang)

### Authors
codec was written by Meng Huang.
