package util

import (
	"math/rand"
	"strings"
)

const alphabet = "abcdefghijklmnoprstuvxyz"

// generates a random int between min to max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// generates a random string of length n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// Generates a random owner name
func RandomOwner() string {
	return RandomString(6)
}

// Generates a random money amount
func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

// Generates a random currency
func RandomCurrency() string {
	currencies := []string{EUR, USD, CAD}

	n := len(currencies)

	return currencies[rand.Intn(n)]
}
