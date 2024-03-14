# Usage
```
go get github.com/ryanking8215/sharedbuffer
```

# Shared Buffer
`Shared Buffer` is designed for the scenario that sharing buffer to several consumers.

It is a type of buffer with reference counter. When the reference counter is counting down to zero, the `Done` callback will be invoked.

If reference counter is negative, it panics.

```
import "github.com/ryanking8215/sharedbuffer"

sb := shared_buffer.New(1000, 2, func(sb *sharedbuffer.Buffer){
	// done callback
})
// after consumer1 consumes the buffer.
sb.Done()
// after consumer2 consumes the buffer.
sb.Done() // here done callback is invoked
```

# Pool
`Pool` is a easy and efficient way to use `Shared Buffer`. `Shared Buffer` is given back to pool automaticlly when reference counter is counting down to zero.

```
import "github.com/ryanking8215/sharedbuffer"

pool := shared_buffer.NewPool(1000)
sb := pool.Get()
sb.Add(2)

// after consumer1 consumes the buffer.
sb.Done()
// after consumer2 consumes the buffer.
sb.Done() // the sb is automaticlly given back to the pool.
```