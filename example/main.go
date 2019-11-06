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
	var obj=&model.Student{Name:"张三",Age:18,Address:"江苏省"}
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