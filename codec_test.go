package codec

import (
	"testing"
	"hslam.com/git/x/codec/example/model"
	"hslam.com/git/x/codec/example/pb"
	"hslam.com/git/x/codec/example/gogopb"
)

func BenchmarkEncodeBytes(t *testing.B) {
	var obj=[]byte{77, 111, 114, 116, 44, 49, 56, 44, 69, 97, 114, 116, 104}
	var c=BytesCodec{}
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		c.Encode(&obj)
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

func BenchmarkDecodeBytes(t *testing.B) {
	var obj=[]byte{77, 111, 114, 116, 44, 49, 56, 44, 69, 97, 114, 116, 104}
	var c=BytesCodec{}
	data,_:=c.Encode(&obj)
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		var obj_copy []byte
		c.Decode(data,&obj_copy)
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

func BenchmarkCodecBytes(t *testing.B) {
	var obj=[]byte{77, 111, 114, 116, 44, 49, 56, 44, 69, 97, 114, 116, 104}
	var c=BytesCodec{}
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		data,_:=c.Encode(&obj)
		var obj_copy []byte
		c.Decode(data,&obj_copy)
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