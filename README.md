go test -v -run="none" -bench=. -benchtime=1s
```
goos: darwin
goarch: amd64
pkg: hslam.com/mgit/Mort/codec
BenchmarkEncodeJson-4    	 5000000	       381 ns/op
BenchmarkEncodeXml-4     	 1000000	      2420 ns/op
BenchmarkEncodeProto-4   	  500000	      3609 ns/op
BenchmarkEncodeGob-4     	  500000	      3246 ns/op
BenchmarkDecodeJson-4    	 1000000	      1316 ns/op
BenchmarkDecodeXml-4     	  300000	      5484 ns/op
BenchmarkDecodeProto-4   	  200000	     11410 ns/op
BenchmarkDecodeGob-4     	  100000	     20217 ns/op
BenchmarkCodecJson-4     	 1000000	      1787 ns/op
BenchmarkCodecXml-4      	  200000	      8093 ns/op
BenchmarkCodecProto-4    	  100000	     15428 ns/op
BenchmarkCodecGob-4      	   50000	     23888 ns/op
PASS
ok  	hslam.com/mgit/Mort/codec	22.580s
```

example
```
json 序列化后：7b224e616d65223a22e5bca0e4b889222c22416765223a31382c2241646472657373223a22e6b19fe88b8fe79c81227d
json 反序列化后： &{张三 18 江苏省}
xml 序列化后：3c53747564656e743e3c4e616d653ee5bca0e4b8893c2f4e616d653e3c4167653e31383c2f4167653e3c416464726573733ee6b19fe88b8fe79c813c2f416464726573733e3c2f53747564656e743e
xml 反序列化后： &{张三 18 江苏省}
gob 序列化后：32ff810301010753747564656e7401ff8200010301044e616d65010c000103416765010400010741646472657373010c00000018ff820106e5bca0e4b88901240109e6b19fe88b8fe79c8100
gob 反序列化后： &{张三 18 江苏省}
proto 序列化后：0a06e5bca0e4b88910121a09e6b19fe88b8fe79c81
proto 反序列化后： {张三 18 江苏省 {} [] 0}
```