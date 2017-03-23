package slowtest

import (
	"math/rand"
	"time"
)

// Slow sleeps for 0-3 seconds and returns true
func Slow(n string) bool {
	// seed pseudo-random number generator
	rand.Seed(time.Now().UTC().UnixNano())
	t := rand.Intn(3)
	time.Sleep(time.Duration(t) * time.Second)
	return true
}
