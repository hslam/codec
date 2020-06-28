package main

import (
	"fmt"
	"github.com/hslam/codec"
	"github.com/hslam/codec/example/code"
	"github.com/hslam/codec/example/codepb"
	"github.com/hslam/codec/example/gencode"
	"github.com/hslam/codec/example/gogopb"
	"github.com/hslam/codec/example/model"
	"github.com/hslam/codec/example/msgp"
	"github.com/hslam/codec/example/pb"
	codecpb "github.com/hslam/codec/pb"
)

func main() {
	BYTES()
	CODE()
	GENCODE()
	CODEPB()
	MSGP()
	GOGOPB()
	PB()
	JSON()
	XML()
	GOB()
}

//BYTES Example
func BYTES() {
	var obj = []byte{128, 8, 128, 8, 195, 245, 72, 64, 74, 216, 18, 77, 251, 33, 9, 64, 10, 72, 101, 108, 108, 111, 87, 111, 114, 108, 100, 1, 1, 255, 2, 1, 128, 1, 255}
	c := &codec.BYTESCodec{}
	var buf = make([]byte, 512)
	data, _ := c.Marshal(buf, &obj)
	fmt.Printf("bytes Encode：length-%d,hex-%x\n", len(data), data)
	var objCopy []byte
	c.Unmarshal(data, &objCopy)
	fmt.Println("bytes Decode：", objCopy)
}

//CODE Example
func CODE() {
	var obj = code.Object{A: 1024, B: 1024, C: 3.14, D: 3.1415926, E: "HelloWorld", F: true, G: []byte{255}, H: [][]byte{{128}, {255}}}
	c := &codec.CODECodec{}
	var buf = make([]byte, 512)
	data, _ := c.Marshal(buf, &obj)
	fmt.Printf("code Encode：length-%d,hex-%x\n", len(data), data)
	var objCopy code.Object
	c.Unmarshal(data, &objCopy)
	fmt.Println("code Decode：", objCopy)
}

//GENCODE Example
func GENCODE() {
	var obj = gencode.Object{A: 1024, B: 1024, C: 3.14, D: 3.1415926, E: "HelloWorld", F: true, G: []byte{255}, H: [][]byte{{128}, {255}}}
	c := &codec.CODECodec{}
	var buf = make([]byte, 512)
	data, _ := c.Marshal(buf, &obj)
	fmt.Printf("gencode Encode：length-%d,hex-%x\n", len(data), data)
	var objCopy gencode.Object
	c.Unmarshal(data, &objCopy)
	fmt.Println("gencode Decode：", objCopy)
}

//CODEPB Example
func CODEPB() {
	var obj = codepb.Object{A: 1024, B: 1024, C: 3.14, D: 3.1415926, E: "HelloWorld", F: true, G: []byte{255}, H: [][]byte{{128}, {255}}}
	c := &codec.CODECodec{}
	var buf = make([]byte, 512)
	data, _ := c.Marshal(buf, &obj)
	fmt.Printf("codepb Encode：length-%d,hex-%x\n", len(data), data)
	var objCopy codepb.Object
	c.Unmarshal(data, &objCopy)
	fmt.Println("codepb Decode：", objCopy)
}

//MSGP Example
func MSGP() {
	var obj = msgp.Object{A: 1024, B: 1024, C: 3.14, D: 3.1415926, E: "HelloWorld", F: true, G: []byte{255}, H: [][]byte{{128}, {255}}}
	c := &codec.MSGPCodec{}
	var buf = make([]byte, 512)
	data, _ := c.Marshal(buf, &obj)
	fmt.Printf("msgp Encode：length-%d,hex-%x\n", len(data), data)
	var objCopy msgp.Object
	c.Unmarshal(data, &objCopy)
	fmt.Println("msgp Decode：", objCopy)
}

//GOGOPB Example
func GOGOPB() {
	var obj = gogopb.Object{A: 1024, B: 1024, C: 3.14, D: 3.1415926, E: "HelloWorld", F: true, G: []byte{255}, H: [][]byte{{128}, {255}}}
	c := &codec.GOGOPBCodec{}
	var buf = make([]byte, 512)
	data, _ := c.Marshal(buf, &obj)
	fmt.Printf("gogopb Encode：length-%d,hex-%x\n", len(data), data)
	var objCopy gogopb.Object
	c.Unmarshal(data, &objCopy)
	fmt.Println("gogopb Decode：", objCopy)
}

//PB Example
func PB() {
	var obj = pb.Object{A: 1024, B: 1024, C: 3.14, D: 3.1415926, E: "HelloWorld", F: true, G: []byte{255}, H: [][]byte{{128}, {255}}}
	c := &codecpb.Codec{}
	var buf = make([]byte, 512)
	data, _ := c.Marshal(buf, &obj)
	fmt.Printf("pb Encode：length-%d,hex-%x\n", len(data), data)
	var objCopy pb.Object
	c.Unmarshal(data, &objCopy)
	fmt.Println("pb Decode：", objCopy)
}

//JSON Example
func JSON() {
	var obj = model.Object{A: 1024, B: 1024, C: 3.14, D: 3.1415926, E: "HelloWorld", F: true, G: []byte{255}, H: [][]byte{{128}, {255}}}
	c := &codec.JSONCodec{}
	var buf = make([]byte, 512)
	data, _ := c.Marshal(buf, &obj)
	fmt.Printf("json Encode：length-%d,hex-%x\n", len(data), data)
	var objCopy model.Object
	c.Unmarshal(data, &objCopy)
	fmt.Println("json Decode：", objCopy)
}

//XML Example
func XML() {
	var obj = model.Object{A: 1024, B: 1024, C: 3.14, D: 3.1415926, E: "HelloWorld", F: true}
	c := &codec.XMLCodec{}
	var buf = make([]byte, 512)
	data, _ := c.Marshal(buf, &obj)
	fmt.Printf("xml Encode：length-%d,hex-%x\n", len(data), data)
	var objCopy model.Object
	c.Unmarshal(data, &objCopy)
	fmt.Println("xml Decode：", objCopy)
}

//GOB Example
func GOB() {
	var obj = model.Object{A: 1024, B: 1024, C: 3.14, D: 3.1415926, E: "HelloWorld", F: true, G: []byte{255}, H: [][]byte{{128}, {255}}}
	c := &codec.GOBCodec{}
	var buf = make([]byte, 512)
	data, _ := c.Marshal(buf, &obj)
	fmt.Printf("gob Encode：length-%d,hex-%x\n", len(data), data)
	var objCopy model.Object
	c.Unmarshal(data, &objCopy)
	fmt.Println("gob Decode：", objCopy)
}
