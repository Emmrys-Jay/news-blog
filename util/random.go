package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().Unix())
}

func RandomInt(n int64) int64 {
	num := rand.Int63n(n)
	return num
}

func RandomString(n int) string {
	k := len(alphabet)
	var s strings.Builder
	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		s.WriteByte(c)
	}

	return s.String()
}

func RandomOwner() string {
	return RandomString(6)
}

func RandomBalance() int64 {
	return RandomInt(10000)
}

func RandomCurrency() string {
	c := []string{EUR, NGN, USD, JPY, CHF, CAD}
	k := len(c)
	n := rand.Intn(k)
	return c[n]
}

func RandomEntryAmount() int64 {
	num := RandomInt(10000)
	s := []int{1, -1}

	n := s[rand.Intn(len(s))]
	num *= int64(n)
	return num
}
