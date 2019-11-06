package codec

import (
	"testing"
	"hslam.com/git/x/codec/example/model"
	"hslam.com/git/x/codec/example/pb"
)

func BenchmarkEncodeJson(t *testing.B) {
	var obj=model.Student{Name:"张三",Age:18,Address:"江苏省"}
	var c=JsonCodec{}
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		c.Encode(&obj)
	}
}

func BenchmarkEncodeXml(t *testing.B) {
	var obj=model.Student{Name:"张三",Age:18,Address:"江苏省"}
	var c=XmlCodec{}
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		c.Encode(&obj)
	}
}

func BenchmarkEncodeProto(t *testing.B) {
	var obj=pb.Student{Name:"张三",Age:18,Address:"江苏省"}
	var c=XmlCodec{}
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		c.Encode(&obj)
	}
}

func BenchmarkEncodeGob(t *testing.B) {
	var obj=model.Student{Name:"张三",Age:18,Address:"江苏省"}
	var c=GobCodec{}
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		c.Encode(&obj)
	}
}

func BenchmarkEncodeBytes(t *testing.B) {
	var obj=[]byte{123}
	var c=BytesCodec{}
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		c.Encode(&obj)
	}
}

func BenchmarkDecodeJson(t *testing.B) {
	var obj=model.Student{Name:"张三",Age:18,Address:"江苏省"}
	var c=JsonCodec{}
	data,_:=c.Encode(&obj)
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		var obj_copy *model.Student
		c.Decode(data,&obj_copy)
	}
}

func BenchmarkDecodeXml(t *testing.B) {
	var obj=model.Student{Name:"张三",Age:18,Address:"江苏省"}
	var c=XmlCodec{}
	data,_:=c.Encode(&obj)
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		var obj_copy *model.Student
		c.Decode(data,&obj_copy)
	}
}

func BenchmarkDecodeProto(t *testing.B) {
	var obj=pb.Student{Name:"张三",Age:18,Address:"江苏省"}
	var c=XmlCodec{}
	data,_:=c.Encode(&obj)
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		var obj_copy *pb.Student
		c.Decode(data,&obj_copy)
	}
}

func BenchmarkDecodeGob(t *testing.B) {
	var obj=model.Student{Name:"张三",Age:18,Address:"江苏省"}
	var c=GobCodec{}
	data,_:=c.Encode(&obj)
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		var obj_copy *model.Student
		c.Decode(data,&obj_copy)
	}
}

func BenchmarkDecodeBytes(t *testing.B) {
	var obj=[]byte{123}
	var c=BytesCodec{}
	data,_:=c.Encode(&obj)
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		var obj_copy []byte
		c.Decode(data,&obj_copy)
	}
}

func BenchmarkCodecJson(t *testing.B) {
	var obj=model.Student{Name:"张三",Age:18,Address:"江苏省"}
	var c=JsonCodec{}
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		data,_:=c.Encode(&obj)
		var obj_copy *model.Student
		c.Decode(data,&obj_copy)
	}
}

func BenchmarkCodecXml(t *testing.B) {
	var obj=model.Student{Name:"张三",Age:18,Address:"江苏省"}
	var c=XmlCodec{}
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		data,_:=c.Encode(&obj)
		var obj_copy *model.Student
		c.Decode(data,&obj_copy)
	}
}

func BenchmarkCodecProto(t *testing.B) {
	var obj=pb.Student{Name:"张三",Age:18,Address:"江苏省"}
	var c=XmlCodec{}
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		data,_:=c.Encode(&obj)
		var obj_copy *pb.Student
		c.Decode(data,&obj_copy)
	}
}

func BenchmarkCodecGob(t *testing.B) {
	var obj=model.Student{Name:"张三",Age:18,Address:"江苏省"}
	var c=GobCodec{}
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		data,_:=c.Encode(&obj)
		var obj_copy *model.Student
		c.Decode(data,&obj_copy)
	}
}

func BenchmarkCodecBytes(t *testing.B) {
	var obj=[]byte{}
	var c=BytesCodec{}
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		data,_:=c.Encode(&obj)
		var obj_copy []byte
		c.Decode(data,&obj_copy)
	}
}