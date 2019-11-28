# codec
A codec library written in golang.

## Feature
* bytes
* [gencode](https://github.com/andyleap/gencode "gencode")
* [fastproto](https://github.com/gogo/protobuf "fastproto")
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
	"hslam.com/git/x/codec/example/fastpb"
	"hslam.com/git/x/codec/example/gencode"
	"hslam.com/git/x/codec/example/msgp"
)

func main(){
	var c codec.Codec
	var data []byte

	//bytes_noreflect
	object_noreflect:=bytes.Student{Name:"Mort",Age:18,Address:"Earth"}
	data,_=object_noreflect.Marshal()
	fmt.Printf("bytes_noreflect Encode：%x\n",data)
	var object_noreflect_copy =&bytes.Student{}
	object_noreflect_copy.Unmarshal(data)
	fmt.Println("bytes_noreflect Decode：",object_noreflect_copy)

	//bytes
	object:=bytes.Student{Name:"Mort",Age:18,Address:"Earth"}
	obj_bytes,_:=object.Marshal()
	c=codec.BytesCodec{}
	data,_=c.Encode(&obj_bytes)
	fmt.Printf("bytes Encode：%x\n",data)
	var object_copy =&bytes.Student{}
	var obj_bytes_copy []byte
	c.Decode(data,&obj_bytes_copy)
	object_copy.Unmarshal(obj_bytes_copy)
	fmt.Println("bytes Decode：",object_copy)

	//gencode_noreflect
	obj_gencode_noreflect:= gencode.Student{Name:"Mort",Age:18,Address:"Earth"}
	data,_=obj_gencode_noreflect.Marshal(nil)
	fmt.Printf("gencode_noreflect Encode：%x\n",data)
	var obj_gencode_noreflect_cp=&gencode.Student{}
	obj_gencode_noreflect_cp.Unmarshal(data)
	fmt.Println("gencode_noreflect Decode：",obj_gencode_noreflect_cp)

	//gencode
	obj_gencode:= gencode.Student{Name:"Mort",Age:18,Address:"Earth"}
	c=codec.GencodeCodec{}
	data,_=c.Encode(&obj_gencode)
	fmt.Printf("gencode Encode：%x\n",data)
	var obj_gencode_cp gencode.Student
	c.Decode(data,&obj_gencode_cp)
	fmt.Println("gencode Decode：",obj_gencode_cp)

	//fastpb_noreflect
	obj_fastpb_noreflect:= fastpb.Student{Name:"Mort",Age:18,Address:"Earth"}
	data,_=obj_fastpb_noreflect.Marshal()
	fmt.Printf("fastproto_noreflect Encode：%x\n",data)
	var obj_fastpb_noreflect_cp= fastpb.Student{}
	obj_fastpb_noreflect_cp.Unmarshal(data)
	fmt.Println("fastproto_noreflect Decode：",obj_fastpb_noreflect_cp)

	//fastproto
	obj_fastpb:= fastpb.Student{Name:"Mort",Age:18,Address:"Earth"}
	c=codec.FastProtoCodec{}
	data,_=c.Encode(&obj_fastpb)
	fmt.Printf("fastproto Encode：%x\n",data)
	var obj_fastpb_cp fastpb.Student
	c.Decode(data,&obj_fastpb_cp)
	fmt.Println("fastproto Decode：",obj_fastpb_cp)

	//msgp_noreflect
	obj_msgp_noreflect:= msgp.Student{Name:"Mort",Age:18,Address:"Earth"}
	data,_=obj_msgp_noreflect.MarshalMsg(nil)
	fmt.Printf("msgp_noreflect Encode：%x\n",data)
	var obj_msgp_noreflect_cp=&msgp.Student{}
	obj_msgp_noreflect_cp.UnmarshalMsg(data)
	fmt.Println("msgp_noreflect Decode：",obj_msgp_noreflect_cp)

	//msgp
	obj_msgp:= msgp.Student{Name:"Mort",Age:18,Address:"Earth"}
	c=codec.MsgpCodec{}
	data,_=c.Encode(&obj_msgp)
	fmt.Printf("msgp Encode：%x\n",data)
	var obj_msgp_cp msgp.Student
	c.Decode(data,&obj_msgp_cp)
	fmt.Println("msgp Decode：",obj_msgp_cp)

	//proto
	obj_pb:=pb.Student{Name:"Mort",Age:18,Address:"Earth"}
	c=codec.ProtoCodec{}
	data,_=c.Encode(&obj_pb)
	fmt.Printf("proto Encode：%x\n",data)
	var obj_pb_cp pb.Student
	c.Decode(data,&obj_pb_cp)
	fmt.Println("proto Decode：",obj_pb_cp)

	var obj=&model.Student{Name:"Mort",Age:18,Address:"Earth"}

	//json
	c=codec.JsonCodec{}
	data,_=c.Encode(&obj)
	fmt.Printf("json Encode：%x\n",data)
	var obj_json *model.Student
	c.Decode(data,&obj_json)
	fmt.Println("json Decode：",obj_json)

	//xml
	c=codec.XmlCodec{}
	data,_=c.Encode(&obj)
	fmt.Printf("xml Encode：%x\n",data)
	var obj_xml *model.Student
	c.Decode(data,&obj_xml)
	fmt.Println("xml Decode：",obj_xml)

	//gob
	c=codec.GobCodec{}
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
bytes_noreflect Decode： &{Mort 18 Earth}
bytes Encode：044d6f7274124561727468
bytes Decode： &{Mort 18 Earth}
gencode_noreflect Encode：044d6f727412000000054561727468
gencode_noreflect Decode： &{Mort 18 Earth}
gencode Encode：044d6f727412000000054561727468
gencode Decode： {Mort 18 Earth}
fastproto_noreflect Encode：0a044d6f727410121a054561727468
fastproto_noreflect Decode： {Mort 18 Earth}
fastproto Encode：0a044d6f727410121a054561727468
fastproto Decode： {Mort 18 Earth}
msgp_noreflect Encode：83a44e616d65a44d6f7274a341676512a741646472657373a54561727468
msgp_noreflect Decode： &{Mort 18 Earth}
msgp Encode：83a44e616d65a44d6f7274a341676512a741646472657373a54561727468
msgp Decode： {Mort 18 Earth}
proto Encode：0a044d6f727410121a054561727468
proto Decode： {Mort 18 Earth {} [] 0}
json Encode：7b224e616d65223a224d6f7274222c22416765223a31382c2241646472657373223a224561727468227d
json Decode： &{Mort 18 Earth}
xml Encode：3c53747564656e743e3c4e616d653e4d6f72743c2f4e616d653e3c4167653e31383c2f4167653e3c416464726573733e45617274683c2f416464726573733e3c2f53747564656e743e
xml Decode： &{Mort 18 Earth}
gob Encode：32ff810301010753747564656e7401ff8200010301044e616d65010c000103416765010400010741646472657373010c00000012ff8201044d6f727401240105456172746800
gob Decode： &{Mort 18 Earth}
```

### Benchmark
go test -v -run="none" -bench=. -benchtime=1s
```
goos: darwin
goarch: amd64
pkg: hslam.com/git/x/codec
BenchmarkEncodeBytesNoReflect-4       	100000000	        20.8 ns/op
BenchmarkEncodeBytes-4                	100000000	        22.0 ns/op
BenchmarkEncodeGencodeNoReflect-4     	100000000	        15.1 ns/op
BenchmarkEncodeGencode-4              	50000000	        25.9 ns/op
BenchmarkEncodeFastProtoNoReflect-4   	30000000	        48.7 ns/op
BenchmarkEncodeFastProto-4            	20000000	        82.4 ns/op
BenchmarkEncodeMsgpNoReflect-4        	20000000	        82.6 ns/op
BenchmarkEncodeMsgp-4                 	20000000	        99.0 ns/op
BenchmarkEncodeProto-4                	10000000	       140 ns/op
BenchmarkEncodeJson-4                 	 5000000	       346 ns/op
BenchmarkEncodeXml-4                  	  500000	      2459 ns/op
BenchmarkEncodeGob-4                  	  500000	      3273 ns/op
BenchmarkDecodeBytesNoReflect-4       	50000000	        34.7 ns/op
BenchmarkDecodeBytes-4                	30000000	        53.3 ns/op
BenchmarkDecodeGencodeNoReflect-4     	30000000	        43.0 ns/op
BenchmarkDecodeGencode-4              	20000000	        88.1 ns/op
BenchmarkDecodeFastProtoNoReflect-4   	30000000	        53.4 ns/op
BenchmarkDecodeFastProto-4            	10000000	       134 ns/op
BenchmarkDecodeMsgpNoReflect-4        	20000000	        96.4 ns/op
BenchmarkDecodeMsgp-4                 	10000000	       152 ns/op
BenchmarkDecodeProto-4                	10000000	       170 ns/op
BenchmarkDecodeJson-4                 	 1000000	      1251 ns/op
BenchmarkDecodeXml-4                  	  300000	      5251 ns/op
BenchmarkDecodeGob-4                  	  100000	     19665 ns/op
BenchmarkCodecBytesNoReflect-4        	30000000	        55.4 ns/op
BenchmarkCodecBytes-4                 	20000000	        73.5 ns/op
BenchmarkCodecGencodeNoReflect-4      	20000000	        58.4 ns/op
BenchmarkCodecGencode-4               	10000000	       129 ns/op
BenchmarkCodecFastProtoNoReflect-4    	20000000	       105 ns/op
BenchmarkCodecFastProto-4             	10000000	       223 ns/op
BenchmarkCodecMsgpNoReflect-4         	10000000	       197 ns/op
BenchmarkCodecMsgp-4                  	 5000000	       269 ns/op
BenchmarkCodecProto-4                 	 5000000	       331 ns/op
BenchmarkCodecJson-4                  	 1000000	      1700 ns/op
BenchmarkCodecXml-4                   	  200000	      8235 ns/op
BenchmarkCodecGob-4                   	   50000	     23655 ns/op
PASS
ok  	hslam.com/git/x/codec	62.621s
```

### Licence
This package is licenced under a MIT licence (Copyright (c) 2019 Mort Huang)

### Authors
codec was written by Mort Huang.
