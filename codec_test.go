package codec

import (
	"testing"
	"hslam.com/git/x/codec/example/model"
	"hslam.com/git/x/codec/example/model_copy"
	"hslam.com/git/x/codec/example/pb"
	"hslam.com/git/x/codec/example/pb_copy"
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

func BenchmarkDecodeJson(t *testing.B) {
	var obj=model.Student{Name:"张三",Age:18,Address:"江苏省"}
	var c=JsonCodec{}
	data,_:=c.Encode(&obj)
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		var obj_copy *model_copy.Student
		c.Decode(data,&obj_copy)
	}
}

func BenchmarkDecodeXml(t *testing.B) {
	var obj=model.Student{Name:"张三",Age:18,Address:"江苏省"}
	var c=XmlCodec{}
	data,_:=c.Encode(&obj)
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		var obj_copy *model_copy.Student
		c.Decode(data,&obj_copy)
	}
}

func BenchmarkDecodeProto(t *testing.B) {
	var obj=pb.Student{Name:"张三",Age:18,Address:"江苏省"}
	var c=XmlCodec{}
	data,_:=c.Encode(&obj)
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		var obj_copy *pb_copy.Student
		c.Decode(data,&obj_copy)
	}
}

func BenchmarkDecodeGob(t *testing.B) {
	var obj=model.Student{Name:"张三",Age:18,Address:"江苏省"}
	var c=GobCodec{}
	data,_:=c.Encode(&obj)
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		var obj_copy *model_copy.Student
		c.Decode(data,&obj_copy)
	}
}


func BenchmarkCodecJson(t *testing.B) {
	var obj=model.Student{Name:"张三",Age:18,Address:"江苏省"}
	var c=JsonCodec{}
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		data,_:=c.Encode(&obj)
		var obj_copy *model_copy.Student
		c.Decode(data,&obj_copy)
	}
}

func BenchmarkCodecXml(t *testing.B) {
	var obj=model.Student{Name:"张三",Age:18,Address:"江苏省"}
	var c=XmlCodec{}
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		data,_:=c.Encode(&obj)
		var obj_copy *model_copy.Student
		c.Decode(data,&obj_copy)
	}
}

func BenchmarkCodecProto(t *testing.B) {
	var obj=pb.Student{Name:"张三",Age:18,Address:"江苏省"}
	var c=XmlCodec{}
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		data,_:=c.Encode(&obj)
		var obj_copy *pb_copy.Student
		c.Decode(data,&obj_copy)
	}
}

func BenchmarkCodecGob(t *testing.B) {
	var obj=model.Student{Name:"张三",Age:18,Address:"江苏省"}
	var c=GobCodec{}
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		data,_:=c.Encode(&obj)
		var obj_copy *model_copy.Student
		c.Decode(data,&obj_copy)
	}
}