# codec
A codec library written in golang.

## Feature
* bytes
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
)

func main(){
	var c codec.Codec
	var data []byte

	//bytes
	obj_bytes:=[]byte{123}
	c=codec.BytesCodec{}
	data,_=c.Encode(&obj_bytes)
	fmt.Printf("bytes Encode：%x\n",data)
	var obj_bytes_copy []byte
	c.Decode(data,&obj_bytes_copy)
	fmt.Println("bytes Decode：",obj_bytes_copy)


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
bytes Encode：4d6f72742c31382c4561727468
bytes Decode： [77 111 114 116 44 49 56 44 69 97 114 116 104]
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
BenchmarkEncodeBytes-4       	500000000	         3.24 ns/op
BenchmarkEncodeGoGoProto-4   	20000000	        79.2 ns/op
BenchmarkEncodeProto-4       	10000000	       140 ns/op
BenchmarkEncodeJson-4        	 5000000	       353 ns/op
BenchmarkEncodeXml-4         	  500000	      2485 ns/op
BenchmarkEncodeGob-4         	  500000	      3307 ns/op
BenchmarkDecodeBytes-4       	30000000	        50.9 ns/op
BenchmarkDecodeGoGoProto-4   	10000000	       129 ns/op
BenchmarkDecodeProto-4       	10000000	       172 ns/op
BenchmarkDecodeJson-4        	 1000000	      1286 ns/op
BenchmarkDecodeXml-4         	  300000	      5550 ns/op
BenchmarkDecodeGob-4         	  100000	     19884 ns/op
BenchmarkCodecBytes-4        	30000000	        56.4 ns/op
BenchmarkCodecGoGoProto-4    	10000000	       208 ns/op
BenchmarkCodecProto-4        	 5000000	       332 ns/op
BenchmarkCodecJson-4         	 1000000	      1682 ns/op
BenchmarkCodecXml-4          	  200000	      8144 ns/op
BenchmarkCodecGob-4          	  100000	     23850 ns/op
PASS
ok  	hslam.com/git/x/codec	32.508s
```

### Licence
This package is licenced under a MIT licence (Copyright (c) 2019 Mort Huang)

### Authors
codec was written by Mort Huang.
