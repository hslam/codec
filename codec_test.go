package codec

import (
	"testing"
	"hslam.com/git/x/codec/example/model"
	"hslam.com/git/x/codec/example/pb"
	"hslam.com/git/x/codec/example/gogopb"
	"hslam.com/git/x/codec/example/bytes"
	"hslam.com/git/x/codec/example/gencode"
	"hslam.com/git/x/codec/example/msgp"
)

func BenchmarkEncodeOnlyBytes(t *testing.B) {
	var obj =[]byte{4,77,111,114,116,18,69,97,114,116,104}
	var c=BytesCodec{}
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		c.Encode(&obj)
	}
}

func BenchmarkEncodeBytesNoReflect(t *testing.B) {
	buf:=make([]byte,100)
	object:=bytes.Student{Name:"Mort",Age:18,Address:"Earth"}
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		object.Marshal(buf)
	}
}

func BenchmarkEncodeBytes(t *testing.B) {
	buf:=make([]byte,100)
	object:=bytes.Student{Name:"Mort",Age:18,Address:"Earth"}
	var obj []byte
	var c=BytesCodec{}
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		obj,_=object.Marshal(buf)
		c.Encode(&obj)
	}
}

func BenchmarkEncodeGencodeNoReflect(t *testing.B) {
	buf:=make([]byte,100)
	object:=gencode.Student{Name:"Mort",Age:18,Address:"Earth"}
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		object.Marshal(buf)
	}
}

func BenchmarkEncodeGencode(t *testing.B) {
	var obj= gencode.Student{Name:"Mort",Age:18,Address:"Earth"}
	var c=GencodeCodec{Buffer:make([]byte,100)}
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		c.Encode(&obj)
	}
}

func BenchmarkEncodeGogoProtoNoReflect(t *testing.B) {
	var obj= gogopb.Student{Name:"Mort",Age:18,Address:"Earth"}
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		obj.Marshal()
	}
}

func BenchmarkEncodeGogoProto(t *testing.B) {
	var obj= gogopb.Student{Name:"Mort",Age:18,Address:"Earth"}
	var c=GogoProtoCodec{}
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		c.Encode(&obj)
	}
}

func BenchmarkEncodeMsgpNoReflect(t *testing.B) {
	buf:=make([]byte,100)
	object:=msgp.Student{Name:"Mort",Age:18,Address:"Earth"}
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		object.MarshalMsg(buf)
	}
}

func BenchmarkEncodeMsgp(t *testing.B) {
	var obj= msgp.Student{Name:"Mort",Age:18,Address:"Earth"}
	var c=MsgpCodec{Buffer:make([]byte,100)}
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		c.Encode(&obj)
	}
}

func BenchmarkEncodeProto(t *testing.B) {
	var obj=pb.Student{Name:"Mort",Age:18,Address:"Earth"}
	var c=ProtoCodec{}
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		c.Encode(&obj)
	}
}

func BenchmarkEncodeJson(t *testing.B) {
	var obj=model.Student{Name:"Mort",Age:18,Address:"Earth"}
	var c=JsonCodec{}
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		c.Encode(&obj)
	}
}

func BenchmarkEncodeXml(t *testing.B) {
	var obj=model.Student{Name:"Mort",Age:18,Address:"Earth"}
	var c=XmlCodec{}
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		c.Encode(&obj)
	}
}

func BenchmarkEncodeGob(t *testing.B) {
	var obj=model.Student{Name:"Mort",Age:18,Address:"Earth"}
	var c=GobCodec{}
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		c.Encode(&obj)
	}
}

func BenchmarkDecodeOnlyBytes(t *testing.B) {
	var obj =[]byte{4,77,111,114,116,18,69,97,114,116,104}
	var c=BytesCodec{}
	data,_:=c.Encode(&obj)
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		var obj_copy []byte
		c.Decode(data,&obj_copy)
	}
}

func BenchmarkDecodeBytesNoReflect(t *testing.B) {
	buf:=make([]byte,100)
	object:=bytes.Student{Name:"Mort",Age:18,Address:"Earth"}
	var data []byte
	data,_=object.Marshal(buf)
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		var object_copy =&bytes.Student{}
		object_copy.Unmarshal(data)
	}
}

func BenchmarkDecodeBytes(t *testing.B) {
	buf:=make([]byte,100)
	object:=bytes.Student{Name:"Mort",Age:18,Address:"Earth"}
	var obj []byte
	var c=BytesCodec{}
	obj,_=object.Marshal(buf)
	data,_:=c.Encode(&obj)
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		var obj_copy []byte
		c.Decode(data,&obj_copy)
		var object_copy =&bytes.Student{}
		object_copy.Unmarshal(data)
	}
}

func BenchmarkDecodeGencodeNoReflect(t *testing.B) {
	var obj= gencode.Student{Name:"Mort",Age:18,Address:"Earth"}
	data,_:=obj.Marshal(nil)
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		var obj_copy =&gencode.Student{}
		obj_copy.Unmarshal(data)
	}
}

func BenchmarkDecodeGencode(t *testing.B) {
	var obj= gencode.Student{Name:"Mort",Age:18,Address:"Earth"}
	var c=GencodeCodec{}
	data,_:=c.Encode(&obj)
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		var obj_copy gencode.Student
		c.Decode(data,&obj_copy)
	}
}

func BenchmarkDecodeGogoProtoNoReflect(t *testing.B) {
	var obj= gogopb.Student{Name:"Mort",Age:18,Address:"Earth"}
	data,_:=obj.Marshal()
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		var obj_copy =&gogopb.Student{}
		obj_copy.Unmarshal(data)
	}
}

func BenchmarkDecodeGogoProto(t *testing.B) {
	var obj= gogopb.Student{Name:"Mort",Age:18,Address:"Earth"}
	var c=GogoProtoCodec{}
	data,_:=c.Encode(&obj)
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		var obj_copy gogopb.Student
		c.Decode(data,&obj_copy)
	}
}

func BenchmarkDecodeMsgpNoReflect(t *testing.B) {
	var obj= msgp.Student{Name:"Mort",Age:18,Address:"Earth"}
	data,_:=obj.MarshalMsg(nil)
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		var obj_copy =&msgp.Student{}
		obj_copy.UnmarshalMsg(data)
	}
}

func BenchmarkDecodeMsgp(t *testing.B) {
	var obj= msgp.Student{Name:"Mort",Age:18,Address:"Earth"}
	var c=MsgpCodec{}
	data,_:=c.Encode(&obj)
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		var obj_copy msgp.Student
		c.Decode(data,&obj_copy)
	}
}

func BenchmarkDecodeProto(t *testing.B) {
	var obj=pb.Student{Name:"Mort",Age:18,Address:"Earth"}
	var c=ProtoCodec{}
	data,_:=c.Encode(&obj)
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		var obj_copy pb.Student
		c.Decode(data,&obj_copy)
	}
}

func BenchmarkDecodeJson(t *testing.B) {
	var obj=model.Student{Name:"Mort",Age:18,Address:"Earth"}
	var c=JsonCodec{}
	data,_:=c.Encode(&obj)
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		var obj_copy *model.Student
		c.Decode(data,&obj_copy)
	}
}

func BenchmarkDecodeXml(t *testing.B) {
	var obj=model.Student{Name:"Mort",Age:18,Address:"Earth"}
	var c=XmlCodec{}
	data,_:=c.Encode(&obj)
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		var obj_copy *model.Student
		c.Decode(data,&obj_copy)
	}
}

func BenchmarkDecodeGob(t *testing.B) {
	var obj=model.Student{Name:"Mort",Age:18,Address:"Earth"}
	var c=GobCodec{}
	data,_:=c.Encode(&obj)
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		var obj_copy *model.Student
		c.Decode(data,&obj_copy)
	}
}

func BenchmarkRoundtripOnlyBytes(t *testing.B) {
	var obj =[]byte{4,77,111,114,116,18,69,97,114,116,104}
	var c=BytesCodec{}
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		data,_:=c.Encode(&obj)
		var obj_copy []byte
		c.Decode(data,&obj_copy)
	}
}

func BenchmarkRoundtripBytesNoReflect(t *testing.B) {
	buf:=make([]byte,100)
	object:=bytes.Student{Name:"Mort",Age:18,Address:"Earth"}
	var data []byte
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		data,_=object.Marshal(buf)
		var object_copy =&bytes.Student{}
		object_copy.Unmarshal(data)
	}
}

func BenchmarkRoundtripBytes(t *testing.B) {
	buf:=make([]byte,100)
	object:=bytes.Student{Name:"Mort",Age:18,Address:"Earth"}
	var obj []byte
	var c=BytesCodec{}
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		obj,_=object.Marshal(buf)
		data,_:=c.Encode(&obj)
		var obj_copy []byte
		c.Decode(data,&obj_copy)
		var object_copy =&bytes.Student{}
		object_copy.Unmarshal(data)
	}
}

func BenchmarkRoundtripGencodeNoReflect(t *testing.B) {
	var obj= gencode.Student{Name:"Mort",Age:18,Address:"Earth"}
	buf:=make([]byte,100)
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		data,_:=obj.Marshal(buf)
		var obj_copy =&gencode.Student{}
		obj_copy.Unmarshal(data)
	}
}

func BenchmarkRoundtripGencode(t *testing.B) {
	var obj= gencode.Student{Name:"Mort",Age:18,Address:"Earth"}
	var c=GencodeCodec{Buffer:make([]byte,100)}
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		data,_:=c.Encode(&obj)
		var obj_copy gencode.Student
		c.Decode(data,&obj_copy)
	}
}

func BenchmarkRoundtripGogoProtoNoReflect(t *testing.B) {
	var obj= gogopb.Student{Name:"Mort",Age:18,Address:"Earth"}
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		data,_:=obj.Marshal()
		var obj_copy =&gogopb.Student{}
		obj_copy.Unmarshal(data)
	}
}

func BenchmarkRoundtripGogoProto(t *testing.B) {
	var obj= gogopb.Student{Name:"Mort",Age:18,Address:"Earth"}
	var c=GogoProtoCodec{}
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		data,_:=c.Encode(&obj)
		var obj_copy gogopb.Student
		c.Decode(data,&obj_copy)
	}
}

func BenchmarkRoundtripMsgpNoReflect(t *testing.B) {
	var obj= msgp.Student{Name:"Mort",Age:18,Address:"Earth"}
	buf:=make([]byte,100)
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		data,_:=obj.MarshalMsg(buf)
		var obj_copy =&msgp.Student{}
		obj_copy.UnmarshalMsg(data)
	}
}

func BenchmarkRoundtripMsgp(t *testing.B) {
	var obj= msgp.Student{Name:"Mort",Age:18,Address:"Earth"}
	var c=MsgpCodec{Buffer:make([]byte,100)}
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		data,_:=c.Encode(&obj)
		var obj_copy msgp.Student
		c.Decode(data,&obj_copy)
	}
}

func BenchmarkRoundtripProto(t *testing.B) {
	var obj=pb.Student{Name:"Mort",Age:18,Address:"Earth"}
	var c=ProtoCodec{}
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		data,_:=c.Encode(&obj)
		var obj_copy pb.Student
		c.Decode(data,&obj_copy)
	}
}

func BenchmarkRoundtripJson(t *testing.B) {
	var obj=model.Student{Name:"Mort",Age:18,Address:"Earth"}
	var c=JsonCodec{}
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		data,_:=c.Encode(&obj)
		var obj_copy *model.Student
		c.Decode(data,&obj_copy)
	}
}

func BenchmarkRoundtripXml(t *testing.B) {
	var obj=model.Student{Name:"Mort",Age:18,Address:"Earth"}
	var c=XmlCodec{}
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		data,_:=c.Encode(&obj)
		var obj_copy *model.Student
		c.Decode(data,&obj_copy)
	}
}

func BenchmarkRoundtripGob(t *testing.B) {
	var obj=model.Student{Name:"Mort",Age:18,Address:"Earth"}
	var c=GobCodec{}
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		data,_:=c.Encode(&obj)
		var obj_copy *model.Student
		c.Decode(data,&obj_copy)
	}
}