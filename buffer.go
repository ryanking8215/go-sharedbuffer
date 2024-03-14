package sharedbuffer

import "sync/atomic"

const negativeRC = "sharedbuffer: negative reference counter"

type DoneFunc func(b *Buffer)

type Buffer struct {
	b    []byte
	rc   int32
	done DoneFunc
}

func New(size int, rc int32, done DoneFunc) *Buffer {
	if rc < 0 {
		panic(negativeRC)
	}

	return &Buffer{
		b:    make([]byte, size),
		rc:   rc,
		done: done,
	}
}

func (b *Buffer) Add(delta int32) {
	if v := atomic.AddInt32(&b.rc, delta); v < 0 {
		panic(negativeRC)
	}
}

func (b *Buffer) RC() int32 {
	return b.rc
}

func (b *Buffer) Bytes() []byte {
	return b.b
}

func (b *Buffer) Done() {
	v := atomic.AddInt32(&b.rc, -1)
	if v == 0 && b.done != nil {
		b.done(b)
	} else if v < 0 {
		panic(negativeRC)
	}
}
