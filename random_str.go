package randomize

import (
	"math/rand"
	"time"
)

// RandomStringGenerateDefault will return a random string with length between 8 and 18.
func RandomStringGenerateDefault() string {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	return RandomStringGenerateBySize(rand.Intn(5) + 1)
}

// RandomStringGenerateRange will return a random string with length between min and max.
// if min > max, min, max = max, min.
func RandomStringGenerateRange(min, max int) string {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	if min > max {
		min, max = max, min
	}
	if min == max {
		return RandomStringGenerateBySize(min)
	}
	return RandomStringGenerateBySize(rand.Intn(max-min) + min)
}

// RandomStringGenerateBySize will return a random string with given length.
// return "" if n <= 0.
func RandomStringGenerateBySize(n int) string {
	if n <= 0 {
		return ""
	}
	rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]rune, n)
	for i := range b {
		b[i] = Letters[rand.Intn(len(Letters))]
	}
	return string(b)
}
