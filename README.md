# codec
A codec library written in golang.

## Feature
* json
* xml
* proto
* gob
* bytes

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
	var obj=model.Student{Name:"张三",Age:18,Address:"江苏省"}
	//json
	c=codec.JsonCodec{}
	data,_=c.Encode(&obj)
	fmt.Printf("json 序列化后：%x\n",data)
	var obj_json *model.Student
	c.Decode(data,&obj_json)
	fmt.Println("json 反序列化后：",obj_json)

	//xml
	c=codec.XmlCodec{}
	data,_=c.Encode(&obj)
	fmt.Printf("xml 序列化后：%x\n",data)
	var obj_xml *model.Student
	c.Decode(data,&obj_xml)
	fmt.Println("xml 反序列化后：",obj_xml)

	//gob
	c=codec.GobCodec{}
	data,_=c.Encode(&obj)
	fmt.Printf("gob 序列化后：%x\n",data)
	var obj_gob *model.Student
	c.Decode(data,&obj_gob)
	fmt.Println("gob 反序列化后：",obj_gob)

	obj1:=pb.Student{Name:"张三",Age:18,Address:"江苏省"}
	//proto
	c=codec.ProtoCodec{}
	data,_=c.Encode(&obj1)
	fmt.Printf("proto 序列化后：%x\n",data)
	var obj_pb pb.Student
	c.Decode(data,&obj_pb)
	fmt.Println("proto 反序列化后：",obj_pb)

	obj2:=[]byte{123}
	//bytes
	c=codec.BytesCodec{}
	data,_=c.Encode(&obj2)
	fmt.Printf("proto 序列化后：%x\n",data)
	var obj_bytes []byte
	c.Decode(data,&obj_bytes)
	fmt.Println("proto 反序列化后：",obj_bytes)
}
```

### Output
```
json 序列化后：7b224e616d65223a22e5bca0e4b889222c22416765223a31382c2241646472657373223a22e6b19fe88b8fe79c81227d
json 反序列化后： &{张三 18 江苏省}
xml 序列化后：3c53747564656e743e3c4e616d653ee5bca0e4b8893c2f4e616d653e3c4167653e31383c2f4167653e3c416464726573733ee6b19fe88b8fe79c813c2f416464726573733e3c2f53747564656e743e
xml 反序列化后： &{张三 18 江苏省}
gob 序列化后：32ff810301010753747564656e7401ff8200010301044e616d65010c000103416765010400010741646472657373010c00000018ff820106e5bca0e4b88901240109e6b19fe88b8fe79c8100
gob 反序列化后： &{张三 18 江苏省}
proto 序列化后：0a06e5bca0e4b88910121a09e6b19fe88b8fe79c81
proto 反序列化后： {张三 18 江苏省 {} [] 0}
proto 序列化后：7b
proto 反序列化后： [123]
```

### Benchmark
go test -v -run="none" -bench=. -benchtime=1s
```
goos: darwin
goarch: amd64
pkg: hslam.com/git/x/codec
BenchmarkEncodeBytes-4   	500000000	         3.03 ns/op
BenchmarkEncodeProto-4   	10000000	       156 ns/op
BenchmarkEncodeJson-4    	 5000000	       381 ns/op
BenchmarkEncodeXml-4     	 1000000	      2421 ns/op
BenchmarkEncodeGob-4     	  500000	      3257 ns/op
BenchmarkDecodeBytes-4   	30000000	        53.3 ns/op
BenchmarkDecodeProto-4   	10000000	       184 ns/op
BenchmarkDecodeJson-4    	 1000000	      1305 ns/op
BenchmarkDecodeXml-4     	  300000	      5466 ns/op
BenchmarkDecodeGob-4     	  100000	     19676 ns/op
BenchmarkCodecBytes-4    	30000000	        55.9 ns/op
BenchmarkCodecProto-4    	 5000000	       358 ns/op
BenchmarkCodecJson-4     	 1000000	      1779 ns/op
BenchmarkCodecXml-4      	  200000	      8111 ns/op
BenchmarkCodecGob-4      	  100000	     23577 ns/op
PASS
ok  	hslam.com/git/x/codec	28.877s
```

### Licence
This package is licenced under a MIT licence (Copyright (c) 2019 Mort Huang)

### Authors
timer was written by Mort Huang.
