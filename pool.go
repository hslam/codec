package codec

type CodecPool interface {
	Get()(Codec)
	Put(c Codec)
}

type JsonCodecPool struct {
	c chan Codec
}

func NewJsonCodecPool(total int) (*JsonCodecPool) {
	return &JsonCodecPool{
		c: make(chan Codec, total),
	}
}

func (cp *JsonCodecPool) Get() (c Codec) {
	select {
	case c = <-cp.c:
	default:
		c = &JsonCodec{}
	}
	return
}

func (cp *JsonCodecPool) Put(c Codec) {
	select {
	case cp.c <- c:
	default:
	}
}

type ProtoCodecPool struct {
	c chan Codec
}

func NewProtoCodecPool(total int) (*ProtoCodecPool) {
	return &ProtoCodecPool{
		c: make(chan Codec, total),
	}
}

func (cp *ProtoCodecPool) Get() (c Codec) {
	select {
	case c = <-cp.c:
	default:
		c = &ProtoCodec{}
	}
	return
}

func (cp *ProtoCodecPool) Put(c Codec) {
	select {
	case cp.c <- c:
	default:
	}
}

type XmlCodecPool struct {
	c chan Codec
}

func NewXmlCodecPool(total int) (*XmlCodecPool) {
	return &XmlCodecPool{
		c: make(chan Codec, total),
	}
}

func (cp *XmlCodecPool) Get() (c Codec) {
	select {
	case c = <-cp.c:
	default:
		c = &XmlCodec{}
	}
	return
}

func (cp *XmlCodecPool) Put(c Codec) {
	select {
	case cp.c <- c:
	default:
	}
}


type GobCodecPool struct {
	c chan Codec
}

func NewGobCodecPool(total int) (*GobCodecPool) {
	return &GobCodecPool{
		c: make(chan Codec, total),
	}
}

func (cp *GobCodecPool) Get() (c Codec) {
	select {
	case c = <-cp.c:
	default:
		c = &GobCodec{}
	}
	return
}

func (cp *GobCodecPool) Put(c Codec) {
	select {
	case cp.c <- c:
	default:
	}
}


type BytesCodecPool struct {
	c chan Codec
}

func NewBytesCodecPool(total int) (*BytesCodecPool) {
	return &BytesCodecPool{
		c: make(chan Codec, total),
	}
}

func (cp *BytesCodecPool) Get() (c Codec) {
	select {
	case c = <-cp.c:
	default:
		c = &BytesCodec{}
	}
	return
}

func (cp *BytesCodecPool) Put(c Codec) {
	select {
	case cp.c <- c:
	default:
	}
}

type GogoProtoCodecPool struct {
	c chan Codec
}

func NewGogoProtoCodecPool(total int) (*GogoProtoCodecPool) {
	return &GogoProtoCodecPool{
		c: make(chan Codec, total),
	}
}

func (cp *GogoProtoCodecPool) Get() (c Codec) {
	select {
	case c = <-cp.c:
	default:
		c = &GogoProtoCodec{}
	}
	return
}

func (cp *GogoProtoCodecPool) Put(c Codec) {
	select {
	case cp.c <- c:
	default:
	}
}

type GencodeCodecPool struct {
	c chan Codec
	w int
}

func NewGencodeCodecPool(total int, width int) (*GencodeCodecPool) {
	return &GencodeCodecPool{
		c: make(chan Codec, total),
		w: width,
	}
}

func (cp *GencodeCodecPool) Get() (c Codec) {
	select {
	case c = <-cp.c:
	default:
		if cp.w>0{
			c = &GencodeCodec{make([]byte,cp.w)}
		}else {
			c = &GencodeCodec{}
		}
	}
	return
}

func (cp *GencodeCodecPool) Put(c Codec) {
	select {
	case cp.c <- c:
	default:
	}
}

type MsgpCodecPool struct {
	c chan Codec
	w int
}

func NewMsgpCodecPool(total int, width int) (*MsgpCodecPool) {
	return &MsgpCodecPool{
		c: make(chan Codec, total),
		w: width,
	}
}

func (cp *MsgpCodecPool) Get() (c Codec) {
	select {
	case c = <-cp.c:
	default:
		if cp.w>0{
			c = &MsgpCodec{make([]byte,cp.w)}
		}else {
			c = &MsgpCodec{}
		}
	}
	return
}

func (cp *MsgpCodecPool) Put(c Codec) {
	select {
	case cp.c <- c:
	default:
	}
}