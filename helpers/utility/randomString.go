package utility

import (
	"math/rand"
	"time"
)

func RandStringBytes(n int) string {
	rand.Seed(time.Now().UTC().UnixNano())
	bytes := make([]byte, n)
	for i := 0; i < n; i++ {
		bytes[i] = byte(randInt(65, 90))
	}
	return string(bytes)
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}
