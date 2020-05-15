// Package ctr implements an application wide
// source for unique IDs.
package ctr

import (
	"fmt"
	"sync/atomic"
)

type count32 int32

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

// ------------------------

// New returns a new specific counter
func New() *count32 {
	cntr := count32(0)
	return &cntr
}

func (c *count32) Increment() int32 {
	return c.increment()
}

// GetLast returns the most recent (current) counter
func (c *count32) GetLast() int32 {
	return c.getLast()
}

// GetLastStr returns the most recent (current) counter
func (c *count32) GetLastStr() string {
	return fmt.Sprintf("%v", c.getLast())
}

// IncrementStr returns the next counter as string
func (c *count32) IncrementStr() string {
	return fmt.Sprintf("%v", c.increment())
}

// Reset turns the clock back to zero
func (c *count32) Reset() {
	c.reset()
}

// ------------------------
// Global counter across packages
var cntr count32

// Increment returns the next counter
func Increment() int32 {
	return cntr.Increment()
}

// GetLast returns the most recent (current) counter
func GetLast() int32 {
	return cntr.GetLast()
}

// GetLastStr returns the most recent (current) counter
func GetLastStr() string {
	return cntr.GetLastStr()
}

// IncrementStr returns the next counter as string
func IncrementStr() string {
	return cntr.IncrementStr()
}

// Reset turns the clock back to zero
func Reset() {
	cntr.Reset()
}
