# codec
A codec library written in golang.

## Feature
* bytes
* gogoproto
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
	"hslam.com/git/x/codec/example/gogopb"
	"hslam.com/git/x/codec/example/bytes"
)

func main(){
	var c codec.Codec
	var data []byte

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

	//gogopbnoreflect
	obj_gogopbnoreflect:=gogopb.Student{Name:"Mort",Age:18,Address:"Earth"}
	data,_=obj_gogopbnoreflect.Marshal()
	fmt.Printf("gogoprotonoreflect Encode：%x\n",data)
	var obj_gogopbnoreflect_cp=gogopb.Student{}
	obj_gogopbnoreflect_cp.Unmarshal(data)
	fmt.Println("gogoprotonoreflect Decode：",obj_gogopbnoreflect_cp)

	//gogoproto
	obj_gogopb:=gogopb.Student{Name:"Mort",Age:18,Address:"Earth"}
	c=codec.GoGoProtoCodec{}
	data,_=c.Encode(&obj_gogopb)
	fmt.Printf("gogoproto Encode：%x\n",data)
	var obj_gogopb_cp gogopb.Student
	c.Decode(data,&obj_gogopb_cp)
	fmt.Println("gogoproto Decode：",obj_gogopb_cp)

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
bytes Encode：044d6f7274124561727468
bytes Decode： &{Mort 18 Earth}
gogoprotonoreflect Encode：0a044d6f727410121a054561727468
gogoprotonoreflect Decode： {Mort 18 Earth}
gogoproto Encode：0a044d6f727410121a054561727468
gogoproto Decode： {Mort 18 Earth}
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
BenchmarkEncodeGoGoProtoNoReflect-4   	30000000	        49.7 ns/op
BenchmarkEncodeBytes-4                	30000000	        50.9 ns/op
BenchmarkEncodeGoGoProto-4            	20000000	        70.6 ns/op
BenchmarkEncodeProto-4                	10000000	       138 ns/op
BenchmarkEncodeJson-4                 	 5000000	       348 ns/op
BenchmarkEncodeXml-4                  	  500000	      2462 ns/op
BenchmarkEncodeGob-4                  	  500000	      3297 ns/op
BenchmarkDecodeGoGoProtoNoReflect-4   	30000000	        53.3 ns/op
BenchmarkDecodeBytes-4                	20000000	        86.4 ns/op
BenchmarkDecodeGoGoProto-4            	10000000	       128 ns/op
BenchmarkDecodeProto-4                	10000000	       174 ns/op
BenchmarkDecodeJson-4                 	 1000000	      1261 ns/op
BenchmarkDecodeXml-4                  	  300000	      5300 ns/op
BenchmarkDecodeGob-4                  	  100000	     20057 ns/op
BenchmarkCodecGoGoProtoNoReflect-4    	20000000	       107 ns/op
BenchmarkCodecBytes-4                 	10000000	       147 ns/op
BenchmarkCodecGoGoProto-4             	10000000	       205 ns/op
BenchmarkCodecProto-4                 	 5000000	       327 ns/op
BenchmarkCodecJson-4                  	 1000000	      1685 ns/op
BenchmarkCodecXml-4                   	  200000	      8033 ns/op
BenchmarkCodecGob-4                   	   50000	     23625 ns/op
PASS
ok  	hslam.com/git/x/codec	36.110s
```

### Licence
This package is licenced under a MIT licence (Copyright (c) 2019 Mort Huang)

### Authors
codec was written by Mort Huang.
