// Copyright (c) 2019 Meng Huang (mhboy@outlook.com)
// This package is licensed under a MIT license that can be found in the LICENSE file.

package codec

// Pool defines the interface of codec pool.
type Pool interface {
	Get() Codec
	Put(c Codec)
}

// JSONCodecPool implements a pool of JSONCodec in the form of a bounded channel.
type JSONCodecPool struct {
	c chan Codec
}

// NewJSONCodecPool creates a new JSONCodecPool bounded to the given total.
func NewJSONCodecPool(total int) *JSONCodecPool {
	return &JSONCodecPool{
		c: make(chan Codec, total),
	}
}

// Get gets a JSONCodec from the JSONCodecPool, or creates a new one if none are
// available in the pool.
func (cp *JSONCodecPool) Get() (c Codec) {
	select {
	case c = <-cp.c:
		// reuse existing Codec
	default:
		// create new Codec
		c = &JSONCodec{}
	}
	return
}

// Put returns the given JSONCodec to the JSONCodecPool.
func (cp *JSONCodecPool) Put(c Codec) {
	select {
	case cp.c <- c:
		// Codec went back into pool
	default:
		// Codec didn't go back into pool, just discard
	}
}

// PBCodecPool implements a pool of PBCodec in the form of a bounded channel.
type PBCodecPool struct {
	c chan Codec
}

// NewPBCodecPool creates a new PBCodecPool bounded to the given total.
func NewPBCodecPool(total int) *PBCodecPool {
	return &PBCodecPool{
		c: make(chan Codec, total),
	}
}

// Get gets a PBCodec from the PBCodecPool, or creates a new one if none are
// available in the pool.
func (cp *PBCodecPool) Get() (c Codec) {
	select {
	case c = <-cp.c:
		// reuse existing Codec
	default:
		// create new Codec
		c = &PBCodec{}
	}
	return
}

// Put returns the given PBCodec to the PBCodecPool.
func (cp *PBCodecPool) Put(c Codec) {
	select {
	case cp.c <- c:
		// Codec went back into pool
	default:
		// Codec didn't go back into pool, just discard
	}
}

// XMLCodecPool implements a pool of XMLCodec in the form of a bounded channel.
type XMLCodecPool struct {
	c chan Codec
}

// NewXMLCodecPool creates a new XMLCodecPool bounded to the given total.
func NewXMLCodecPool(total int) *XMLCodecPool {
	return &XMLCodecPool{
		c: make(chan Codec, total),
	}
}

// Get gets a XMLCodec from the XMLCodecPool, or creates a new one if none are
// available in the pool.
func (cp *XMLCodecPool) Get() (c Codec) {
	select {
	case c = <-cp.c:
		// reuse existing Codec
	default:
		// create new Codec
		c = &XMLCodec{}
	}
	return
}

// Put returns the given XMLCodec to the XMLCodecPool.
func (cp *XMLCodecPool) Put(c Codec) {
	select {
	case cp.c <- c:
		// Codec went back into pool
	default:
		// Codec didn't go back into pool, just discard
	}
}

// GOBCodecPool implements a pool of GOBCodec in the form of a bounded channel.
type GOBCodecPool struct {
	c chan Codec
}

// NewGOBCodecPool creates a new GOBCodecPool bounded to the given total.
func NewGOBCodecPool(total int) *GOBCodecPool {
	return &GOBCodecPool{
		c: make(chan Codec, total),
	}
}

// Get gets a GOBCodec from the GOBCodecPool, or creates a new one if none are
// available in the pool.
func (cp *GOBCodecPool) Get() (c Codec) {
	select {
	case c = <-cp.c:
		// reuse existing Codec
	default:
		// create new Codec
		c = &GOBCodec{}
	}
	return
}

// Put returns the given GOBCodec to the GOBCodecPool.
func (cp *GOBCodecPool) Put(c Codec) {
	select {
	case cp.c <- c:
		// Codec went back into pool
	default:
		// Codec didn't go back into pool, just discard
	}
}

// BYTESCodecPool implements a pool of BYTESCodec in the form of a bounded channel.
type BYTESCodecPool struct {
	c chan Codec
}

// NewBYTESCodecPool creates a new BYTESCodecPool bounded to the given total.
func NewBYTESCodecPool(total int) *BYTESCodecPool {
	return &BYTESCodecPool{
		c: make(chan Codec, total),
	}
}

// Get gets a BYTESCodec from the BYTESCodecPool, or creates a new one if none are
// available in the pool.
func (cp *BYTESCodecPool) Get() (c Codec) {
	select {
	case c = <-cp.c:
		// reuse existing Codec
	default:
		// create new Codec
		c = &BYTESCodec{}
	}
	return
}

// Put returns the given BYTESCodec to the BYTESCodecPool.
func (cp *BYTESCodecPool) Put(c Codec) {
	select {
	case cp.c <- c:
		// Codec went back into pool
	default:
		// Codec didn't go back into pool, just discard
	}
}

// GOGOPBCodecPool implements a pool of GOGOPBCodec in the form of a bounded channel.
type GOGOPBCodecPool struct {
	c chan Codec
	w int
}

// NewGOGOPBCodecPool creates a new GOGOPBCodecPool bounded to the given total.
func NewGOGOPBCodecPool(total int, width int) *GOGOPBCodecPool {
	return &GOGOPBCodecPool{
		c: make(chan Codec, total),
		w: width,
	}
}

// Get gets a GOGOPBCodec from the GOGOPBCodecPool, or creates a new one if none are
// available in the pool.
func (cp *GOGOPBCodecPool) Get() (c Codec) {
	select {
	case c = <-cp.c:
		// reuse existing Codec
	default:
		// create new Codec
		if cp.w > 0 {
			c = &GOGOPBCodec{make([]byte, cp.w)}
		} else {
			c = &GOGOPBCodec{}
		}
	}
	return
}

// Put returns the given GOGOPBCodec to the GOGOPBCodecPool.
func (cp *GOGOPBCodecPool) Put(c Codec) {
	select {
	case cp.c <- c:
		// Codec went back into pool
	default:
		// Codec didn't go back into pool, just discard
	}
}

// CODECodecPool implements a pool of CODECodec in the form of a bounded channel.
type CODECodecPool struct {
	c chan Codec
	w int
}

// NewCODECodecPool creates a new CODECodecPool bounded to the given total.
func NewCODECodecPool(total int, width int) *CODECodecPool {
	return &CODECodecPool{
		c: make(chan Codec, total),
		w: width,
	}
}

// Get gets a CODECodec from the CODECodecPool, or creates a new one if none are
// available in the pool.
func (cp *CODECodecPool) Get() (c Codec) {
	select {
	case c = <-cp.c:
		// reuse existing Codec
	default:
		// create new Codec
		if cp.w > 0 {
			c = &CODECodec{make([]byte, cp.w)}
		} else {
			c = &CODECodec{}
		}
	}
	return
}

// Put returns the given CODECodec to the CODECodecPool.
func (cp *CODECodecPool) Put(c Codec) {
	select {
	case cp.c <- c:
		// Codec went back into pool
	default:
		// Codec didn't go back into pool, just discard
	}
}

// MSGPCodecPool implements a pool of MSGPCodec in the form of a bounded channel.
type MSGPCodecPool struct {
	c chan Codec
	w int
}

// NewMSGPCodecPool creates a new MSGPCodecPool bounded to the given total.
func NewMSGPCodecPool(total int, width int) *MSGPCodecPool {
	return &MSGPCodecPool{
		c: make(chan Codec, total),
		w: width,
	}
}

// Get gets a MSGPCodec from the MSGPCodecPool, or creates a new one if none are
// available in the pool.
func (cp *MSGPCodecPool) Get() (c Codec) {
	select {
	case c = <-cp.c:
		// reuse existing Codec
	default:
		// create new Codec
		if cp.w > 0 {
			c = &MSGPCodec{make([]byte, cp.w)}
		} else {
			c = &MSGPCodec{}
		}
	}
	return
}

// Put returns the given MSGPCodec to the MSGPCodecPool.
func (cp *MSGPCodecPool) Put(c Codec) {
	select {
	case cp.c <- c:
		// Codec went back into pool
	default:
		// Codec didn't go back into pool, just discard
	}
}
