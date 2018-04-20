package main

import "sync/atomic"

type count32 int32

var cntr count32

func (c *count32) increment() int32 {
	return atomic.AddInt32((*int32)(c), 1)
}

func (c *count32) getLast() int32 {
	return atomic.LoadInt32((*int32)(c))
}
