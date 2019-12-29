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
	var obj=[]byte{0,4,0,0,0,4,0,0,0,0,0,0,195,245,72,64,74,216,18,77,251,33,9,64,10,72,101,108,108,111,87,111,114,108,100,1,1,0}
	p=NewBytesCodecPool(1024)
	c:=p.Get()
	data,_:=c.Encode(&obj)
	var obj_copy []byte
	c.Decode(data,&obj_copy)
	p.Put(c)
}

func TestCodeCodecPool(t *testing.T) {
	var p CodecPool
	var obj=gencode.Object{
		A:1024,
		B:1024,
		C:3.14,
		D:3.1415926,
		E:"HelloWorld",
		F:true,
		G:[]byte{0},
	}
	p=NewCodeCodecPool(1024,65536)
	c:=p.Get()
	data,_:=c.Encode(&obj)
	var obj_copy =gencode.Object{}
	c.Decode(data,&obj_copy)
	p.Put(c)
}

func TestGogoProtoCodecPool(t *testing.T) {
	var p CodecPool
	var obj=gogopb.Object{
		A:1024,
		B:1024,
		C:3.14,
		D:3.1415926,
		E:"HelloWorld",
		F:true,
		G:[]byte{0},
	}
	p=NewGogoProtoCodecPool(1024)
	c:=p.Get()
	data,_:=c.Encode(&obj)
	var obj_copy =gogopb.Object{}
	c.Decode(data,&obj_copy)
	p.Put(c)
}

func TestMsgpCodecPool(t *testing.T) {
	var p CodecPool
	var obj=msgp.Object{
		A:1024,
		B:1024,
		C:3.14,
		D:3.1415926,
		E:"HelloWorld",
		F:true,
		G:[]byte{0},
	}
	p=NewMsgpCodecPool(1024,65536)
	c:=p.Get()
	data,_:=c.Encode(&obj)
	var obj_copy =msgp.Object{}
	c.Decode(data,&obj_copy)
	p.Put(c)
}

func TestProtoCodecPool(t *testing.T) {
	var p CodecPool
	var obj=pb.Object{
		A:1024,
		B:1024,
		C:3.14,
		D:3.1415926,
		E:"HelloWorld",
		F:true,
		G:[]byte{0},
	}
	p=NewProtoCodecPool(1024)
	c:=p.Get()
	data,_:=c.Encode(&obj)
	var obj_copy =pb.Object{}
	c.Decode(data,&obj_copy)
	p.Put(c)
}

func TestJsonCodecPool(t *testing.T) {
	var p CodecPool
	var obj=model.Object{
		A:1024,
		B:1024,
		C:3.14,
		D:3.1415926,
		E:"HelloWorld",
		F:true,
		G:[]byte{0},
	}
	p=NewJsonCodecPool(1024)
	c:=p.Get()
	data,_:=c.Encode(&obj)
	var obj_copy =model.Object{}
	c.Decode(data,&obj_copy)
	p.Put(c)
}

func TestXmlCodecPool(t *testing.T) {
	var p CodecPool
	var obj=model.Object{
		A:1024,
		B:1024,
		C:3.14,
		D:3.1415926,
		E:"HelloWorld",
		F:true,
		G:[]byte{0},
	}
	p=NewXmlCodecPool(1024)
	c:=p.Get()
	data,_:=c.Encode(&obj)
	var obj_copy =model.Object{}
	c.Decode(data,&obj_copy)
	p.Put(c)
}

func TestGobCodecPool(t *testing.T) {
	var p CodecPool
	var obj=model.Object{
		A:1024,
		B:1024,
		C:3.14,
		D:3.1415926,
		E:"HelloWorld",
		F:true,
		G:[]byte{0},
	}
	p=NewGobCodecPool(1024)
	c:=p.Get()
	data,_:=c.Encode(&obj)
	var obj_copy =model.Object{}
	c.Decode(data,&obj_copy)
	p.Put(c)
}