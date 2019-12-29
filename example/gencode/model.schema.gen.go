package gencode

import (
	"io"
	"time"
	"unsafe"
)

var (
	_ = unsafe.Sizeof(0)
	_ = io.ReadFull
	_ = time.Now()
)

type Object struct {
	A uint32
	B uint64
	C float32
	D float64
	E string
	F bool
	G []byte
}

func (d *Object) Size() (s uint64) {

	{
		l := uint64(len(d.E))

		{

			t := l
			for t >= 0x80 {
				t >>= 7
				s++
			}
			s++

		}
		s += l
	}
	{
		l := uint64(len(d.G))

		{

			t := l
			for t >= 0x80 {
				t >>= 7
				s++
			}
			s++

		}
		s += l
	}
	s += 25
	return
}
func (d *Object) Marshal(buf []byte) ([]byte, error) {
	size := d.Size()
	{
		if uint64(cap(buf)) >= size {
			buf = buf[:size]
		} else {
			buf = make([]byte, size)
		}
	}
	i := uint64(0)

	{

		buf[0+0] = byte(d.A >> 0)

		buf[1+0] = byte(d.A >> 8)

		buf[2+0] = byte(d.A >> 16)

		buf[3+0] = byte(d.A >> 24)

	}
	{

		buf[0+4] = byte(d.B >> 0)

		buf[1+4] = byte(d.B >> 8)

		buf[2+4] = byte(d.B >> 16)

		buf[3+4] = byte(d.B >> 24)

		buf[4+4] = byte(d.B >> 32)

		buf[5+4] = byte(d.B >> 40)

		buf[6+4] = byte(d.B >> 48)

		buf[7+4] = byte(d.B >> 56)

	}
	{

		v := *(*uint32)(unsafe.Pointer(&(d.C)))

		buf[0+12] = byte(v >> 0)

		buf[1+12] = byte(v >> 8)

		buf[2+12] = byte(v >> 16)

		buf[3+12] = byte(v >> 24)

	}
	{

		v := *(*uint64)(unsafe.Pointer(&(d.D)))

		buf[0+16] = byte(v >> 0)

		buf[1+16] = byte(v >> 8)

		buf[2+16] = byte(v >> 16)

		buf[3+16] = byte(v >> 24)

		buf[4+16] = byte(v >> 32)

		buf[5+16] = byte(v >> 40)

		buf[6+16] = byte(v >> 48)

		buf[7+16] = byte(v >> 56)

	}
	{
		l := uint64(len(d.E))

		{

			t := uint64(l)

			for t >= 0x80 {
				buf[i+24] = byte(t) | 0x80
				t >>= 7
				i++
			}
			buf[i+24] = byte(t)
			i++

		}
		copy(buf[i+24:], d.E)
		i += l
	}
	{
		if d.F {
			buf[i+24] = 1
		} else {
			buf[i+24] = 0
		}
	}
	{
		l := uint64(len(d.G))

		{

			t := uint64(l)

			for t >= 0x80 {
				buf[i+25] = byte(t) | 0x80
				t >>= 7
				i++
			}
			buf[i+25] = byte(t)
			i++

		}
		copy(buf[i+25:], d.G)
		i += l
	}
	return buf[:i+25], nil
}

func (d *Object) Unmarshal(buf []byte) (uint64, error) {
	i := uint64(0)

	{

		d.A = 0 | (uint32(buf[i+0+0]) << 0) | (uint32(buf[i+1+0]) << 8) | (uint32(buf[i+2+0]) << 16) | (uint32(buf[i+3+0]) << 24)

	}
	{

		d.B = 0 | (uint64(buf[i+0+4]) << 0) | (uint64(buf[i+1+4]) << 8) | (uint64(buf[i+2+4]) << 16) | (uint64(buf[i+3+4]) << 24) | (uint64(buf[i+4+4]) << 32) | (uint64(buf[i+5+4]) << 40) | (uint64(buf[i+6+4]) << 48) | (uint64(buf[i+7+4]) << 56)

	}
	{

		v := 0 | (uint32(buf[i+0+12]) << 0) | (uint32(buf[i+1+12]) << 8) | (uint32(buf[i+2+12]) << 16) | (uint32(buf[i+3+12]) << 24)
		d.C = *(*float32)(unsafe.Pointer(&v))

	}
	{

		v := 0 | (uint64(buf[i+0+16]) << 0) | (uint64(buf[i+1+16]) << 8) | (uint64(buf[i+2+16]) << 16) | (uint64(buf[i+3+16]) << 24) | (uint64(buf[i+4+16]) << 32) | (uint64(buf[i+5+16]) << 40) | (uint64(buf[i+6+16]) << 48) | (uint64(buf[i+7+16]) << 56)
		d.D = *(*float64)(unsafe.Pointer(&v))

	}
	{
		l := uint64(0)

		{

			bs := uint8(7)
			t := uint64(buf[i+24] & 0x7F)
			for buf[i+24]&0x80 == 0x80 {
				i++
				t |= uint64(buf[i+24]&0x7F) << bs
				bs += 7
			}
			i++

			l = t

		}
		d.E = string(buf[i+24 : i+24+l])
		i += l
	}
	{
		d.F = buf[i+24] == 1
	}
	{
		l := uint64(0)

		{

			bs := uint8(7)
			t := uint64(buf[i+25] & 0x7F)
			for buf[i+25]&0x80 == 0x80 {
				i++
				t |= uint64(buf[i+25]&0x7F) << bs
				bs += 7
			}
			i++

			l = t

		}
		if uint64(cap(d.G)) >= l {
			d.G = d.G[:l]
		} else {
			d.G = make([]byte, l)
		}
		copy(d.G, buf[i+25:])
		i += l
	}
	return i + 25, nil
}
