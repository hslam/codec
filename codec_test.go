package codec

import (
	"testing"
	"github.com/hslam/codec/example/model"
	"github.com/hslam/codec/example/pb"
	"github.com/hslam/codec/example/gogopb"
	"github.com/hslam/codec/example/gencode"
	"github.com/hslam/codec/example/msgp"
	"github.com/hslam/codec/example/code"
)

func BenchmarkEncodeBytes(t *testing.B) {
	var obj=[]byte{0,4,0,0,0,4,0,0,0,0,0,0,195,245,72,64,74,216,18,77,251,33,9,64,10,72,101,108,108,111,87,111,114,108,100,1,1,0}
	var c=BytesCodec{}
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		c.Encode(&obj)
	}
}

//func BenchmarkEncodeCodeNoReflect(t *testing.B) {
//	buf:=make([]byte,512)
//	var obj=code.Object{A:1024,B:1024,C:3.14,D:3.1415926,E:"HelloWorld",F:true,G:[]byte{0}}
//	t.ResetTimer()
//	for i := 0; i < t.N; i++ {
//		obj.Marshal(buf)
//	}
//}

func BenchmarkEncodeCode(t *testing.B) {
	var obj=code.Object{A:1024,B:1024,C:3.14,D:3.1415926,E:"HelloWorld",F:true,G:[]byte{0}}
	var c=CodeCodec{Buffer:make([]byte,512)}
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		c.Encode(&obj)
	}
}


//func BenchmarkEncodeGencodeNoReflect(t *testing.B) {
//	buf:=make([]byte,512)
//	var obj=gencode.Object{A:1024,B:1024,C:3.14,D:3.1415926,E:"HelloWorld",F:true,G:[]byte{0}}
//	t.ResetTimer()
//	for i := 0; i < t.N; i++ {
//		obj.Marshal(buf)
//	}
//}

func BenchmarkEncodeGencode(t *testing.B) {
	var obj=gencode.Object{A:1024,B:1024,C:3.14,D:3.1415926,E:"HelloWorld",F:true,G:[]byte{0}}
	var c=CodeCodec{Buffer:make([]byte,512)}
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		c.Encode(&obj)
	}
}

//func BenchmarkEncodeGogoProtoNoReflect(t *testing.B) {
//	var obj=gogopb.Object{A:1024,B:1024,C:3.14,D:3.1415926,E:"HelloWorld",F:true,G:[]byte{0}}
//	t.ResetTimer()
//	for i := 0; i < t.N; i++ {
//		obj.Marshal()
//	}
//}

func BenchmarkEncodeGogoProto(t *testing.B) {
	var obj=gogopb.Object{A:1024,B:1024,C:3.14,D:3.1415926,E:"HelloWorld",F:true,G:[]byte{0}}
	var c=GogoProtoCodec{}
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		c.Encode(&obj)
	}
}

//func BenchmarkEncodeMsgpNoReflect(t *testing.B) {
//	buf:=make([]byte,512)
//	var obj=msgp.Object{A:1024,B:1024,C:3.14,D:3.1415926,E:"HelloWorld",F:true,G:[]byte{0}}
//	t.ResetTimer()
//	for i := 0; i < t.N; i++ {
//		obj.MarshalMsg(buf)
//	}
//}

func BenchmarkEncodeMsgp(t *testing.B) {
	var obj=msgp.Object{A:1024,B:1024,C:3.14,D:3.1415926,E:"HelloWorld",F:true,G:[]byte{0}}
	var c=MsgpCodec{Buffer:make([]byte,512)}
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		c.Encode(&obj)
	}
}

func BenchmarkEncodeProto(t *testing.B) {
	var obj=pb.Object{A:1024,B:1024,C:3.14,D:3.1415926,E:"HelloWorld",F:true,G:[]byte{0}}
	var c=ProtoCodec{}
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		c.Encode(&obj)
	}
}

func BenchmarkEncodeJson(t *testing.B) {
	var obj=model.Object{A:1024,B:1024,C:3.14,D:3.1415926,E:"HelloWorld",F:true,G:[]byte{0}}
	var c=JsonCodec{}
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		c.Encode(&obj)
	}
}

func BenchmarkEncodeXml(t *testing.B) {
	var obj=model.Object{A:1024,B:1024,C:3.14,D:3.1415926,E:"HelloWorld",F:true,G:[]byte{0}}
	var c=XmlCodec{}
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		c.Encode(&obj)
	}
}

func BenchmarkEncodeGob(t *testing.B) {
	var obj=model.Object{A:1024,B:1024,C:3.14,D:3.1415926,E:"HelloWorld",F:true,G:[]byte{0}}
	var c=GobCodec{}
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		c.Encode(&obj)
	}
}

func BenchmarkDecodeBytes(t *testing.B) {
	var obj=[]byte{0,4,0,0,0,4,0,0,0,0,0,0,195,245,72,64,74,216,18,77,251,33,9,64,10,72,101,108,108,111,87,111,114,108,100,1,1,0}
	var c=BytesCodec{}
	data,_:=c.Encode(&obj)
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		var obj_copy []byte
		c.Decode(data,&obj_copy)
	}
}

//func BenchmarkDecodeCodeNoReflect(t *testing.B) {
//	var obj=code.Object{A:1024,B:1024,C:3.14,D:3.1415926,E:"HelloWorld",F:true,G:[]byte{0}}
//	data,_:=obj.Marshal(make([]byte,512))
//	t.ResetTimer()
//	for i := 0; i < t.N; i++ {
//		var obj_copy =&code.Object{}
//		obj_copy.Unmarshal(data)
//	}
//}

func BenchmarkDecodeCode(t *testing.B) {
	var obj=code.Object{A:1024,B:1024,C:3.14,D:3.1415926,E:"HelloWorld",F:true,G:[]byte{0}}
	var c=CodeCodec{Buffer:make([]byte,512)}
	data,_:=c.Encode(&obj)
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		var obj_copy code.Object
		c.Decode(data,&obj_copy)
	}
}

//func BenchmarkDecodeGencodeNoReflect(t *testing.B) {
//	var obj=gencode.Object{A:1024,B:1024,C:3.14,D:3.1415926,E:"HelloWorld",F:true,G:[]byte{0}}
//	data,_:=obj.Marshal(make([]byte,512))
//	t.ResetTimer()
//	for i := 0; i < t.N; i++ {
//		var obj_copy =&gencode.Object{}
//		obj_copy.Unmarshal(data)
//	}
//}

func BenchmarkDecodeGencode(t *testing.B) {
	var obj=gencode.Object{A:1024,B:1024,C:3.14,D:3.1415926,E:"HelloWorld",F:true,G:[]byte{0}}
	var c=CodeCodec{Buffer:make([]byte,512)}
	data,_:=c.Encode(&obj)
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		var obj_copy gencode.Object
		c.Decode(data,&obj_copy)
	}
}

//func BenchmarkDecodeGogoProtoNoReflect(t *testing.B) {
//	var obj=gogopb.Object{A:1024,B:1024,C:3.14,D:3.1415926,E:"HelloWorld",F:true,G:[]byte{0}}
//	data,_:=obj.Marshal()
//	t.ResetTimer()
//	for i := 0; i < t.N; i++ {
//		var obj_copy =&gogopb.Object{}
//		obj_copy.Unmarshal(data)
//	}
//}

func BenchmarkDecodeGogoProto(t *testing.B) {
	var obj=gogopb.Object{A:1024,B:1024,C:3.14,D:3.1415926,E:"HelloWorld",F:true,G:[]byte{0}}
	var c=GogoProtoCodec{}
	data,_:=c.Encode(&obj)
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		var obj_copy gogopb.Object
		c.Decode(data,&obj_copy)
	}
}

//func BenchmarkDecodeMsgpNoReflect(t *testing.B) {
//	var obj=msgp.Object{A:1024,B:1024,C:3.14,D:3.1415926,E:"HelloWorld",F:true,G:[]byte{0}}
//	data,_:=obj.MarshalMsg(make([]byte,512))
//	t.ResetTimer()
//	for i := 0; i < t.N; i++ {
//		var obj_copy =&msgp.Object{}
//		obj_copy.UnmarshalMsg(data)
//	}
//}

func BenchmarkDecodeMsgp(t *testing.B) {
	var obj=msgp.Object{A:1024,B:1024,C:3.14,D:3.1415926,E:"HelloWorld",F:true,G:[]byte{0}}
	var c=MsgpCodec{make([]byte,512)}
	data,_:=c.Encode(&obj)
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		var obj_copy msgp.Object
		c.Decode(data,&obj_copy)
	}
}

func BenchmarkDecodeProto(t *testing.B) {
	var obj=pb.Object{A:1024,B:1024,C:3.14,D:3.1415926,E:"HelloWorld",F:true,G:[]byte{0}}
	var c=ProtoCodec{}
	data,_:=c.Encode(&obj)
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		var obj_copy pb.Object
		c.Decode(data,&obj_copy)
	}
}

func BenchmarkDecodeJson(t *testing.B) {
	var obj=model.Object{A:1024,B:1024,C:3.14,D:3.1415926,E:"HelloWorld",F:true,G:[]byte{0}}
	var c=JsonCodec{}
	data,_:=c.Encode(&obj)
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		var obj_copy *model.Object
		c.Decode(data,&obj_copy)
	}
}

func BenchmarkDecodeXml(t *testing.B) {
	var obj=model.Object{A:1024,B:1024,C:3.14,D:3.1415926,E:"HelloWorld",F:true,G:[]byte{0}}
	var c=XmlCodec{}
	data,_:=c.Encode(&obj)
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		var obj_copy *model.Object
		c.Decode(data,&obj_copy)
	}
}

func BenchmarkDecodeGob(t *testing.B) {
	var obj=model.Object{A:1024,B:1024,C:3.14,D:3.1415926,E:"HelloWorld",F:true,G:[]byte{0}}
	var c=GobCodec{}
	data,_:=c.Encode(&obj)
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		var obj_copy *model.Object
		c.Decode(data,&obj_copy)
	}
}

func BenchmarkRoundtripBytes(t *testing.B) {
	var obj=[]byte{0,4,0,0,0,4,0,0,0,0,0,0,195,245,72,64,74,216,18,77,251,33,9,64,10,72,101,108,108,111,87,111,114,108,100,1,1,0}
	var c=BytesCodec{}
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		data,_:=c.Encode(&obj)
		var obj_copy []byte
		c.Decode(data,&obj_copy)
	}
}

//func BenchmarkRoundtripCodeNoReflect(t *testing.B) {
//	var obj=code.Object{A:1024,B:1024,C:3.14,D:3.1415926,E:"HelloWorld",F:true,G:[]byte{0}}
//	buf:=make([]byte,512)
//	t.ResetTimer()
//	for i := 0; i < t.N; i++ {
//		data,_:=obj.Marshal(buf)
//		var obj_copy =&code.Object{}
//		obj_copy.Unmarshal(data)
//	}
//}

func BenchmarkRoundtripCode(t *testing.B) {
	var obj=code.Object{A:1024,B:1024,C:3.14,D:3.1415926,E:"HelloWorld",F:true,G:[]byte{0}}
	var c=CodeCodec{Buffer:make([]byte,512)}
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		data,_:=c.Encode(&obj)
		var obj_copy code.Object
		c.Decode(data,&obj_copy)
	}
}

//func BenchmarkRoundtripGencodeNoReflect(t *testing.B) {
//	var obj=gencode.Object{A:1024,B:1024,C:3.14,D:3.1415926,E:"HelloWorld",F:true,G:[]byte{0}}
//	buf:=make([]byte,100)
//	t.ResetTimer()
//	for i := 0; i < t.N; i++ {
//		data,_:=obj.Marshal(buf)
//		var obj_copy =&gencode.Object{}
//		obj_copy.Unmarshal(data)
//	}
//}

func BenchmarkRoundtripGencode(t *testing.B) {
	var obj=gencode.Object{A:1024,B:1024,C:3.14,D:3.1415926,E:"HelloWorld",F:true,G:[]byte{0}}
	var c=CodeCodec{Buffer:make([]byte,100)}
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		data,_:=c.Encode(&obj)
		var obj_copy gencode.Object
		c.Decode(data,&obj_copy)
	}
}

//func BenchmarkRoundtripGogoProtoNoReflect(t *testing.B) {
//	var obj=gogopb.Object{A:1024,B:1024,C:3.14,D:3.1415926,E:"HelloWorld",F:true,G:[]byte{0}}
//	t.ResetTimer()
//	for i := 0; i < t.N; i++ {
//		data,_:=obj.Marshal()
//		var obj_copy =&gogopb.Object{}
//		obj_copy.Unmarshal(data)
//	}
//}

func BenchmarkRoundtripGogoProto(t *testing.B) {
	var obj=gogopb.Object{A:1024,B:1024,C:3.14,D:3.1415926,E:"HelloWorld",F:true,G:[]byte{0}}
	var c=GogoProtoCodec{}
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		data,_:=c.Encode(&obj)
		var obj_copy gogopb.Object
		c.Decode(data,&obj_copy)
	}
}

//func BenchmarkRoundtripMsgpNoReflect(t *testing.B) {
//	var obj=msgp.Object{A:1024,B:1024,C:3.14,D:3.1415926,E:"HelloWorld",F:true,G:[]byte{0}}
//	buf:=make([]byte,100)
//	t.ResetTimer()
//	for i := 0; i < t.N; i++ {
//		data,_:=obj.MarshalMsg(buf)
//		var obj_copy =&msgp.Object{}
//		obj_copy.UnmarshalMsg(data)
//	}
//}

func BenchmarkRoundtripMsgp(t *testing.B) {
	var obj=msgp.Object{A:1024,B:1024,C:3.14,D:3.1415926,E:"HelloWorld",F:true,G:[]byte{0}}
	var c=MsgpCodec{Buffer:make([]byte,100)}
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		data,_:=c.Encode(&obj)
		var obj_copy msgp.Object
		c.Decode(data,&obj_copy)
	}
}

func BenchmarkRoundtripProto(t *testing.B) {
	var obj=pb.Object{A:1024,B:1024,C:3.14,D:3.1415926,E:"HelloWorld",F:true,G:[]byte{0}}
	var c=ProtoCodec{}
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		data,_:=c.Encode(&obj)
		var obj_copy pb.Object
		c.Decode(data,&obj_copy)
	}
}

func BenchmarkRoundtripJson(t *testing.B) {
	var obj=model.Object{A:1024,B:1024,C:3.14,D:3.1415926,E:"HelloWorld",F:true,G:[]byte{0}}
	var c=JsonCodec{}
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		data,_:=c.Encode(&obj)
		var obj_copy *model.Object
		c.Decode(data,&obj_copy)
	}
}

func BenchmarkRoundtripXml(t *testing.B) {
	var obj=model.Object{A:1024,B:1024,C:3.14,D:3.1415926,E:"HelloWorld",F:true,G:[]byte{0}}
	var c=XmlCodec{}
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		data,_:=c.Encode(&obj)
		var obj_copy *model.Object
		c.Decode(data,&obj_copy)
	}
}

func BenchmarkRoundtripGob(t *testing.B) {
	var obj=model.Object{A:1024,B:1024,C:3.14,D:3.1415926,E:"HelloWorld",F:true,G:[]byte{0}}
	var c=GobCodec{}
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		data,_:=c.Encode(&obj)
		var obj_copy *model.Object
		c.Decode(data,&obj_copy)
	}
}