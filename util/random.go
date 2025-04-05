package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.NewSource(time.Now().UnixNano())
}

// RandomString generates a random string between min and max length.
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// RandomString generates a random string of specified length.
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

// RandomOwner generates the random owner name of length 6
func RandomOwner() string {
	return RandomString(6)
}

// RandomMoney generates the money between specified range
func RandomMoney() int64 {
	return RandomInt(0, 10000)
}

// RandomCurrency generates the random currency code
func RandomCurrency() string {
	currencies := []string{"USD", "EUR", "GBP", "JPY", "YUAN"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}
