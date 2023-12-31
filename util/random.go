package util

import (
	"math/rand"
	"strings"
)

// RandomOwner generates a random owner name
func RandomOwner() string {
	return randomString(6)
}

// RandomMoney generates a random amount of money
func RandomMoney() int64 {
	return randomInt(0, 1000)
}

// RandomCurrency generates a random currency code
func RandomCurrency() string {
	currencies := []string{"BRL", "CAD", "EUR", "USD"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}

// randomInt generates a random integer between min and max
func randomInt(min, max int64) int64 {
	return min + rand.Int63n(max - min + 1)
}

// randomString generates a random string of length n
func randomString(n int) string {
	const alphabet = "abcdefghijklmnopqrstuvwxyz"
	var sb strings.Builder
	k := len(alphabet)
	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}
