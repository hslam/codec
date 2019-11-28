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
	data,_=object_noreflect.Marshal(nil)
	fmt.Printf("bytes_noreflect Encode：%x\n",data)
	var object_noreflect_copy =&bytes.Student{}
	object_noreflect_copy.Unmarshal(data)
	fmt.Println("bytes_noreflect Decode：",object_noreflect_copy)

	//bytes
	object:=bytes.Student{Name:"Mort",Age:18,Address:"Earth"}
	obj_bytes,_:=object.Marshal(nil)
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