package codec

import (
	"testing"
	"hslam.com/git/x/codec/example/model"
	"hslam.com/git/x/codec/example/pb"
	"hslam.com/git/x/codec/example/gogopb"
	"hslam.com/git/x/codec/example/bytes"
)

func BenchmarkEncodeBytesNoReflect(t *testing.B) {
	object:=bytes.Student{Name:"Mort",Age:18,Address:"Earth"}
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		object.Marshal()
	}
}

func BenchmarkEncodeBytes(t *testing.B) {
	object:=bytes.Student{Name:"Mort",Age:18,Address:"Earth"}
	var obj []byte
	var c=BytesCodec{}
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		obj,_=object.Marshal()
		c.Encode(&obj)
	}
}

func BenchmarkEncodeGoGoProtoNoReflect(t *testing.B) {
	var obj=gogopb.Student{Name:"Mort",Age:18,Address:"Earth"}
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		obj.Marshal()
	}
}

func BenchmarkEncodeGoGoProto(t *testing.B) {
	var obj=gogopb.Student{Name:"Mort",Age:18,Address:"Earth"}
	var c=GoGoProtoCodec{}
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
func BenchmarkDecodeBytesNoReflect(t *testing.B) {
	object:=bytes.Student{Name:"Mort",Age:18,Address:"Earth"}
	var data []byte
	data,_=object.Marshal()
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		var object_copy =&bytes.Student{}
		object_copy.Unmarshal(data)
	}
}
func BenchmarkDecodeBytes(t *testing.B) {
	object:=bytes.Student{Name:"Mort",Age:18,Address:"Earth"}
	var obj []byte
	var c=BytesCodec{}
	obj,_=object.Marshal()
	data,_:=c.Encode(&obj)
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		var obj_copy []byte
		c.Decode(data,&obj_copy)
		var object_copy =&bytes.Student{}
		object_copy.Unmarshal(data)
	}
}

func BenchmarkDecodeGoGoProtoNoReflect(t *testing.B) {
	var obj=gogopb.Student{Name:"Mort",Age:18,Address:"Earth"}
	data,_:=obj.Marshal()
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		var obj_copy =&gogopb.Student{}
		obj_copy.Unmarshal(data)
	}
}

func BenchmarkDecodeGoGoProto(t *testing.B) {
	var obj=gogopb.Student{Name:"Mort",Age:18,Address:"Earth"}
	var c=GoGoProtoCodec{}
	data,_:=c.Encode(&obj)
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		var obj_copy gogopb.Student
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

func BenchmarkCodecBytesNoReflect(t *testing.B) {
	object:=bytes.Student{Name:"Mort",Age:18,Address:"Earth"}
	var data []byte
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		data,_=object.Marshal()
		var object_copy =&bytes.Student{}
		object_copy.Unmarshal(data)
	}
}

func BenchmarkCodecBytes(t *testing.B) {
	object:=bytes.Student{Name:"Mort",Age:18,Address:"Earth"}
	var obj []byte
	var c=BytesCodec{}
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		obj,_=object.Marshal()
		data,_:=c.Encode(&obj)
		var obj_copy []byte
		c.Decode(data,&obj_copy)
		var object_copy =&bytes.Student{}
		object_copy.Unmarshal(data)
	}
}

func BenchmarkCodecGoGoProtoNoReflect(t *testing.B) {
	var obj=gogopb.Student{Name:"Mort",Age:18,Address:"Earth"}
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		data,_:=obj.Marshal()
		var obj_copy =&gogopb.Student{}
		obj_copy.Unmarshal(data)
	}
}

func BenchmarkCodecGoGoProto(t *testing.B) {
	var obj=gogopb.Student{Name:"Mort",Age:18,Address:"Earth"}
	var c=GoGoProtoCodec{}
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		data,_:=c.Encode(&obj)
		var obj_copy gogopb.Student
		c.Decode(data,&obj_copy)
	}
}

func BenchmarkCodecProto(t *testing.B) {
	var obj=pb.Student{Name:"Mort",Age:18,Address:"Earth"}
	var c=ProtoCodec{}
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		data,_:=c.Encode(&obj)
		var obj_copy pb.Student
		c.Decode(data,&obj_copy)
	}
}

func BenchmarkCodecJson(t *testing.B) {
	var obj=model.Student{Name:"Mort",Age:18,Address:"Earth"}
	var c=JsonCodec{}
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		data,_:=c.Encode(&obj)
		var obj_copy *model.Student
		c.Decode(data,&obj_copy)
	}
}

func BenchmarkCodecXml(t *testing.B) {
	var obj=model.Student{Name:"Mort",Age:18,Address:"Earth"}
	var c=XmlCodec{}
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		data,_:=c.Encode(&obj)
		var obj_copy *model.Student
		c.Decode(data,&obj_copy)
	}
}

func BenchmarkCodecGob(t *testing.B) {
	var obj=model.Student{Name:"Mort",Age:18,Address:"Earth"}
	var c=GobCodec{}
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		data,_:=c.Encode(&obj)
		var obj_copy *model.Student
		c.Decode(data,&obj_copy)
	}
}