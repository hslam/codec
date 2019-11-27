# codec
A codec library written in golang.

## Feature
* bytes
* fastproto
* proto
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
	fmt.Printf("obj_gencode_noreflect Encode：%x\n",data)
	var obj_gencode_noreflect_cp=&gencode.Student{}
	obj_gencode_noreflect_cp.Unmarshal(data)
	fmt.Println("obj_gencode_noreflect Decode：",obj_gencode_noreflect_cp)

	//gencode
	obj_gencode:= gencode.Student{Name:"Mort",Age:18,Address:"Earth"}
	c=codec.GencodeCodec{}
	data,_=c.Encode(&obj_gencode)
	fmt.Printf("obj_gencode Encode：%x\n",data)
	var obj_gencode_cp gencode.Student
	c.Decode(data,&obj_gencode_cp)
	fmt.Println("obj_gencode Decode：",obj_gencode_cp)

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
obj_gencode_noreflect Encode：044d6f727412000000054561727468
obj_gencode_noreflect Decode： &{Mort 18 Earth}
obj_gencode Encode：044d6f727412000000054561727468
obj_gencode Decode： {Mort 18 Earth}
fastproto_noreflect Encode：0a044d6f727410121a054561727468
fastproto_noreflect Decode： {Mort 18 Earth}
fastproto Encode：0a044d6f727410121a054561727468
fastproto Decode： {Mort 18 Earth}
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
BenchmarkEncodeBytesNoReflect-4       	100000000	        20.9 ns/op
BenchmarkEncodeBytes-4                	100000000	        21.3 ns/op
BenchmarkEncodeGencodeNoReflect-4     	100000000	        15.3 ns/op
BenchmarkEncodeGencode-4              	50000000	        24.5 ns/op
BenchmarkEncodeFastProtoNoReflect-4   	30000000	        48.8 ns/op
BenchmarkEncodeFastProto-4            	20000000	        82.6 ns/op
BenchmarkEncodeProto-4                	10000000	       137 ns/op
BenchmarkEncodeJson-4                 	 5000000	       349 ns/op
BenchmarkEncodeXml-4                  	  500000	      2440 ns/op
BenchmarkEncodeGob-4                  	  500000	      3257 ns/op
BenchmarkDecodeBytesNoReflect-4       	50000000	        34.7 ns/op
BenchmarkDecodeBytes-4                	20000000	        53.0 ns/op
BenchmarkDecodeGencodeNoReflect-4     	30000000	        46.0 ns/op
BenchmarkDecodeGencode-4              	20000000	        87.8 ns/op
BenchmarkDecodeFastProtoNoReflect-4   	30000000	        52.6 ns/op
BenchmarkDecodeFastProto-4            	10000000	       137 ns/op
BenchmarkDecodeProto-4                	10000000	       171 ns/op
BenchmarkDecodeJson-4                 	 1000000	      1257 ns/op
BenchmarkDecodeXml-4                  	  300000	      5246 ns/op
BenchmarkDecodeGob-4                  	  100000	     19656 ns/op
BenchmarkCodecBytesNoReflect-4        	30000000	        58.4 ns/op
BenchmarkCodecBytes-4                 	20000000	        73.1 ns/op
BenchmarkCodecGencodeNoReflect-4      	30000000	        60.7 ns/op
BenchmarkCodecGencode-4               	10000000	       129 ns/op
BenchmarkCodecFastProtoNoReflect-4    	20000000	       106 ns/op
BenchmarkCodecFastProto-4             	10000000	       227 ns/op
BenchmarkCodecProto-4                 	 5000000	       327 ns/op
BenchmarkCodecJson-4                  	 1000000	      1686 ns/op
BenchmarkCodecXml-4                   	  200000	      7942 ns/op
BenchmarkCodecGob-4                   	  100000	     23231 ns/op
PASS
ok  	hslam.com/git/x/codec	52.491s
```

### Licence
This package is licenced under a MIT licence (Copyright (c) 2019 Mort Huang)

### Authors
codec was written by Mort Huang.
