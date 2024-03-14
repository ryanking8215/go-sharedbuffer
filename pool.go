package sharedbuffer

import "sync"

type Pool struct {
	sync.Pool
}

func NewPool(bufSize int) *Pool {
	p := Pool{}
	p.Pool.New = func() interface{} {
		return New(bufSize, 0, p.put)
	}
	return &p
}

func (p *Pool) put(b *Buffer) {
	p.Pool.Put(b)
}

func (p *Pool) Get() *Buffer {
	return p.Pool.Get().(*Buffer)
}
