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