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