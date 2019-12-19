package codec

import (
	"testing"
	"github.com/hslam/codec/example/model"
	"github.com/hslam/codec/example/pb"
	"github.com/hslam/codec/example/gogopb"
	"github.com/hslam/codec/example/gencode"
	"github.com/hslam/codec/example/msgp"
)

func TestBytesCodecPool(t *testing.T) {
	var p CodecPool
	var obj =[]byte{4,77,111,114,116,18,69,97,114,116,104}
	p=NewBytesCodecPool(1024)
	c:=p.Get()
	data,_:=c.Encode(&obj)
	var obj_copy []byte
	c.Decode(data,&obj_copy)
	p.Put(c)
}

func TestCodeCodecPool(t *testing.T) {
	var p CodecPool
	var obj= gencode.Student{Name:"Mort",Age:18,Address:"Earth"}
	p=NewCodeCodecPool(1024,65536)
	c:=p.Get()
	data,_:=c.Encode(&obj)
	var obj_copy =gencode.Student{}
	c.Decode(data,&obj_copy)
	p.Put(c)
}

func TestGogoProtoCodecPool(t *testing.T) {
	var p CodecPool
	var obj= gogopb.Student{Name:"Mort",Age:18,Address:"Earth"}
	p=NewGogoProtoCodecPool(1024)
	c:=p.Get()
	data,_:=c.Encode(&obj)
	var obj_copy =gogopb.Student{}
	c.Decode(data,&obj_copy)
	p.Put(c)
}

func TestMsgpCodecPool(t *testing.T) {
	var p CodecPool
	var obj= msgp.Student{Name:"Mort",Age:18,Address:"Earth"}
	p=NewMsgpCodecPool(1024,65536)
	c:=p.Get()
	data,_:=c.Encode(&obj)
	var obj_copy =msgp.Student{}
	c.Decode(data,&obj_copy)
	p.Put(c)
}

func TestProtoCodecPool(t *testing.T) {
	var p CodecPool
	var obj=pb.Student{Name:"Mort",Age:18,Address:"Earth"}
	p=NewProtoCodecPool(1024)
	c:=p.Get()
	data,_:=c.Encode(&obj)
	var obj_copy =pb.Student{}
	c.Decode(data,&obj_copy)
	p.Put(c)
}

func TestJsonCodecPool(t *testing.T) {
	var p CodecPool
	var obj=model.Student{Name:"Mort",Age:18,Address:"Earth"}
	p=NewJsonCodecPool(1024)
	c:=p.Get()
	data,_:=c.Encode(&obj)
	var obj_copy =model.Student{}
	c.Decode(data,&obj_copy)
	p.Put(c)
}

func TestXmlCodecPool(t *testing.T) {
	var p CodecPool
	var obj=model.Student{Name:"Mort",Age:18,Address:"Earth"}
	p=NewXmlCodecPool(1024)
	c:=p.Get()
	data,_:=c.Encode(&obj)
	var obj_copy =model.Student{}
	c.Decode(data,&obj_copy)
	p.Put(c)
}

func TestGobCodecPool(t *testing.T) {
	var p CodecPool
	var obj=model.Student{Name:"Mort",Age:18,Address:"Earth"}
	p=NewGobCodecPool(1024)
	c:=p.Get()
	data,_:=c.Encode(&obj)
	var obj_copy =model.Student{}
	c.Decode(data,&obj_copy)
	p.Put(c)
}