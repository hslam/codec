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
	Bytes()
	Code()
	Gencode()
	GogoProto()
	Msgp()
	Proto()
	Json()
	Xml()
	Gob()
}
func Bytes()  {
	var obj=[]byte{0,4,0,0,0,4,0,0,0,0,0,0,195,245,72,64,74,216,18,77,251,33,9,64,10,72,101,108,108,111,87,111,114,108,100,1,1,0}
	c:=&codec.BytesCodec{}
	data,_:=c.Encode(&obj)
	fmt.Printf("bytes Encode：length-%d,hex-%x\n",len(data),data)
	var obj_bytes_copy []byte
	c.Decode(data,&obj_bytes_copy)
	fmt.Println("bytes Decode：",obj_bytes_copy)
}

func Code()  {
	var obj=code.Object{A:1024,B:1024,C:3.14,D:3.1415926,E:"HelloWorld",F:true,G:[]byte{0}}
	c:=&codec.CodeCodec{make([]byte,512)}
	data,_:=c.Encode(&obj)
	fmt.Printf("code Encode：length-%d,hex-%x\n",len(data),data)
	var obj_cp code.Object
	c.Decode(data,&obj_cp)
	fmt.Println("code Decode：",obj_cp)
}

func Gencode()  {
	var obj=gencode.Object{A:1024,B:1024,C:3.14,D:3.1415926,E:"HelloWorld",F:true,G:[]byte{0}}
	c:=&codec.CodeCodec{make([]byte,512)}
	data,_:=c.Encode(&obj)
	fmt.Printf("gencode Encode：length-%d,hex-%x\n",len(data),data)
	var obj_cp gencode.Object
	c.Decode(data,&obj_cp)
	fmt.Println("gencode Decode：",obj_cp)
}

func GogoProto()  {
	var obj=gogopb.Object{A:1024,B:1024,C:3.14,D:3.1415926,E:"HelloWorld",F:true,G:[]byte{0}}
	c:=&codec.GogoProtoCodec{}
	data,_:=c.Encode(&obj)
	fmt.Printf("gogoproto Encode：length-%d,hex-%x\n",len(data),data)
	var obj_cp gogopb.Object
	c.Decode(data,&obj_cp)
	fmt.Println("gogoproto Decode：",obj_cp)
}

func Msgp()  {
	var obj=msgp.Object{A:1024,B:1024,C:3.14,D:3.1415926,E:"HelloWorld",F:true,G:[]byte{0}}
	c:=&codec.MsgpCodec{}
	data,_:=c.Encode(&obj)
	fmt.Printf("msgp Encode：length-%d,hex-%x\n",len(data),data)
	var obj_cp msgp.Object
	c.Decode(data,&obj_cp)
	fmt.Println("msgp Decode：",obj_cp)
}

func Proto()  {
	var obj=pb.Object{A:1024,B:1024,C:3.14,D:3.1415926,E:"HelloWorld",F:true,G:[]byte{0}}
	c:=&codec.ProtoCodec{}
	data,_:=c.Encode(&obj)
	fmt.Printf("proto Encode：length-%d,hex-%x\n",len(data),data)
	var obj_cp pb.Object
	c.Decode(data,&obj_cp)
	fmt.Println("proto Decode：",obj_cp)
}

func Json()  {
	var obj=model.Object{A:1024,B:1024,C:3.14,D:3.1415926,E:"HelloWorld",F:true,G:[]byte{0}}
	c:=&codec.JsonCodec{}
	data,_:=c.Encode(&obj)
	fmt.Printf("json Encode：length-%d,hex-%x\n",len(data),data)
	var obj_cp model.Object
	c.Decode(data,&obj_cp)
	fmt.Println("json Decode：",obj_cp)
}

func Xml()  {
	var obj=model.Object{A:1024,B:1024,C:3.14,D:3.1415926,E:"HelloWorld",F:true,G:[]byte{0}}
	c:=&codec.XmlCodec{}
	data,_:=c.Encode(&obj)
	fmt.Printf("xml Encode：length-%d,hex-%x\n",len(data),data)
	var obj_cp model.Object
	c.Decode(data,&obj_cp)
	fmt.Println("xml Decode：",obj_cp)
}

func Gob()  {
	var obj=model.Object{A:1024,B:1024,C:3.14,D:3.1415926,E:"HelloWorld",F:true,G:[]byte{0}}
	c:=&codec.GobCodec{}
	data,_:=c.Encode(&obj)
	fmt.Printf("gob Encode：length-%d,hex-%x\n",len(data),data)
	var obj_cp model.Object
	c.Decode(data,&obj_cp)
	fmt.Println("gob Decode：",obj_cp)
}