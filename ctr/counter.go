// Package ctr implements an application wide
// source for unique IDs.
package ctr

import (
	"fmt"
	"sync/atomic"
)

type count32 int32

var cntr count32 // Application wide source for unique IDs

// atomic.AddInt32 exposes hardware/CPU provided threadsafe counters.
// No lock() - unlock() required.
func (c *count32) increment() int32 {
	return atomic.AddInt32((*int32)(c), 1)
}

func (c *count32) getLast() int32 {
	return atomic.LoadInt32((*int32)(c))
}

func (c *count32) reset() {
	i := int32(0)
	atomic.StoreInt32((*int32)(c), i)
}

// Increment returns the next counter
func Increment() int32 {
	return cntr.increment()
}

// GetLast returns the most recent (current) counter
func GetLast() int32 {
	return cntr.increment()
}

// IncrementStr returns the next counter as string
func IncrementStr() string {
	return fmt.Sprintf("%v", cntr.increment())
}

// Reset turns the clock back to zero
func Reset() {
	cntr.reset()
}
