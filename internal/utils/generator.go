package utils

import (
	"github.com/google/uuid"
	"math/rand"
	"sync/atomic"
	"time"
)

var counter int64 = time.Now().UnixNano()
var rng = rand.New(rand.NewSource(time.Now().UnixNano()))

func GenerateID() int64 {
	return atomic.AddInt64(&counter, 1)
}
func GenerateInt64() int64 {
	return rng.Int63()
}

func GenerateRandomFloat(min, max float64) float64 {
	return min + rng.Float64()*(max-min)
}

func GenerateUUID() string {
	return uuid.New().String()
}

func GenerateString(n int) string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	bytes := make([]byte, n)
	for i := range bytes {
		bytes[i] = letters[rng.Intn(len(letters))]
	}
	return string(bytes)
}
