package utils

import (
	"sync/atomic"
	"time"
)

var counter int64 = time.Now().UnixNano()

func GenerateTransactionID() int64 {
	return atomic.AddInt64(&counter, 1)
}
