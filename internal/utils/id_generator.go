package utils

import "sync/atomic"

var counter int64

func IdGenerator() int64 {
	return atomic.AddInt64(&counter, 1)
}
