package codec

type JsonCodecPool struct {
	c chan *JsonCodec
}

func NewJsonCodecPool(total int) (*JsonCodecPool) {
	return &JsonCodecPool{
		c: make(chan *JsonCodec, total),
	}
}

func (cp *JsonCodecPool) Get() (c *JsonCodec) {
	select {
	case c = <-cp.c:
	default:
		c = &JsonCodec{}
	}
	return
}

func (cp *JsonCodecPool) Put(c *JsonCodec) {
	select {
	case cp.c <- c:
	default:
	}
}

type ProtoCodecPool struct {
	c chan *ProtoCodec
}

func NewProtoCodecPool(total int) (*ProtoCodecPool) {
	return &ProtoCodecPool{
		c: make(chan *ProtoCodec, total),
	}
}

func (cp *ProtoCodecPool) Get() (c *ProtoCodec) {
	select {
	case c = <-cp.c:
	default:
		c = &ProtoCodec{}
	}
	return
}

func (cp *ProtoCodecPool) Put(c *ProtoCodec) {
	select {
	case cp.c <- c:
	default:
	}
}


type XmlCodecPool struct {
	c chan *XmlCodec
}

func NewXmlCodecPool(total int) (*XmlCodecPool) {
	return &XmlCodecPool{
		c: make(chan *XmlCodec, total),
	}
}

func (cp *XmlCodecPool) Get() (c *XmlCodec) {
	select {
	case c = <-cp.c:
	default:
		c = &XmlCodec{}
	}
	return
}

func (cp *XmlCodecPool) Put(c *XmlCodec) {
	select {
	case cp.c <- c:
	default:
	}
}


type GobCodecPool struct {
	c chan *GobCodec
}

func NewGobCodecPool(total int) (*GobCodecPool) {
	return &GobCodecPool{
		c: make(chan *GobCodec, total),
	}
}

func (cp *GobCodecPool) Get() (c *GobCodec) {
	select {
	case c = <-cp.c:
	default:
		c = &GobCodec{}
	}
	return
}

func (cp *GobCodecPool) Put(c *GobCodec) {
	select {
	case cp.c <- c:
	default:
	}
}


type BytesCodecPool struct {
	c chan *BytesCodec
}

func NewBytesCodecPool(total int) (*BytesCodecPool) {
	return &BytesCodecPool{
		c: make(chan *BytesCodec, total),
	}
}

func (cp *BytesCodecPool) Get() (c *BytesCodec) {
	select {
	case c = <-cp.c:
	default:
		c = &BytesCodec{}
	}
	return
}

func (cp *BytesCodecPool) Put(c *BytesCodec) {
	select {
	case cp.c <- c:
	default:
	}
}


type GogoProtoCodecPool struct {
	c chan *GogoProtoCodec
}

func NewGogoProtoCodecPool(total int) (*GogoProtoCodecPool) {
	return &GogoProtoCodecPool{
		c: make(chan *GogoProtoCodec, total),
	}
}

func (cp *GogoProtoCodecPool) Get() (c *GogoProtoCodec) {
	select {
	case c = <-cp.c:
	default:
		c = &GogoProtoCodec{}
	}
	return
}

func (cp *GogoProtoCodecPool) Put(c *GogoProtoCodec) {
	select {
	case cp.c <- c:
	default:
	}
}

type GencodeCodecPool struct {
	c chan *GencodeCodec
	w int
}

func NewGencodeCodecPool(total int, width int) (*GencodeCodecPool) {
	return &GencodeCodecPool{
		c: make(chan *GencodeCodec, total),
		w: width,
	}
}

func (cp *GencodeCodecPool) Get() (c *GencodeCodec) {
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

func (cp *GencodeCodecPool) Put(c *GencodeCodec) {
	select {
	case cp.c <- c:
	default:
	}
}

type MsgpCodecPool struct {
	c chan *MsgpCodec
	w int
}

func NewMsgpCodecPool(total int, width int) (*MsgpCodecPool) {
	return &MsgpCodecPool{
		c: make(chan *MsgpCodec, total),
		w: width,
	}
}

func (cp *MsgpCodecPool) Get() (c *MsgpCodec) {
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

func (cp *MsgpCodecPool) Put(c *MsgpCodec) {
	select {
	case cp.c <- c:
	default:
	}
}