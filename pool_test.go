package codec

import (
	"testing"
	"hslam.com/git/x/codec/example/model"
	"hslam.com/git/x/codec/example/pb"
	"hslam.com/git/x/codec/example/gogopb"
	"hslam.com/git/x/codec/example/gencode"
	"hslam.com/git/x/codec/example/msgp"
)

func TestBytesCodecPool(t *testing.T) {
	var obj =[]byte{4,77,111,114,116,18,69,97,114,116,104}
	p:=NewBytesCodecPool(10)
	c:=p.Get()
	data,_:=c.Encode(&obj)
	var obj_copy []byte
	c.Decode(data,&obj_copy)
	p.Put(c)
}

func TestGencodeCodecPool(t *testing.T) {
	var obj= gencode.Student{Name:"Mort",Age:18,Address:"Earth"}
	p:=NewGencodeCodecPool(10,0)
	c:=p.Get()
	data,_:=c.Encode(&obj)
	var obj_copy =gencode.Student{}
	c.Decode(data,&obj_copy)
	p.Put(c)
}

func TestGogoProtoCodecPool(t *testing.T) {
	var obj= gogopb.Student{Name:"Mort",Age:18,Address:"Earth"}
	p:=NewGogoProtoCodecPool(10)
	c:=p.Get()
	data,_:=c.Encode(&obj)
	var obj_copy =gogopb.Student{}
	c.Decode(data,&obj_copy)
	p.Put(c)
}

func TestMsgpCodecPool(t *testing.T) {
	var obj= msgp.Student{Name:"Mort",Age:18,Address:"Earth"}
	p:=NewMsgpCodecPool(10,0)
	c:=p.Get()
	data,_:=c.Encode(&obj)
	var obj_copy =msgp.Student{}
	c.Decode(data,&obj_copy)
	p.Put(c)
}

func TestProtoCodecPool(t *testing.T) {
	var obj=pb.Student{Name:"Mort",Age:18,Address:"Earth"}
	p:=NewProtoCodecPool(10)
	c:=p.Get()
	data,_:=c.Encode(&obj)
	var obj_copy =pb.Student{}
	c.Decode(data,&obj_copy)
	p.Put(c)
}

func TestJsonCodecPool(t *testing.T) {
	var obj=model.Student{Name:"Mort",Age:18,Address:"Earth"}
	p:=NewJsonCodecPool(10)
	c:=p.Get()
	data,_:=c.Encode(&obj)
	var obj_copy =model.Student{}
	c.Decode(data,&obj_copy)
	p.Put(c)
}

func TestXmlCodecPool(t *testing.T) {
	var obj=model.Student{Name:"Mort",Age:18,Address:"Earth"}
	p:=NewXmlCodecPool(10)
	c:=p.Get()
	data,_:=c.Encode(&obj)
	var obj_copy =model.Student{}
	c.Decode(data,&obj_copy)
	p.Put(c)
}

func TestGobCodecPool(t *testing.T) {
	var obj=model.Student{Name:"Mort",Age:18,Address:"Earth"}
	p:=NewGobCodecPool(10)
	c:=p.Get()
	data,_:=c.Encode(&obj)
	var obj_copy =model.Student{}
	c.Decode(data,&obj_copy)
	p.Put(c)
}