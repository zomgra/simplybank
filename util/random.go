package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijclmnopqrstuvwxyz"

func inti() {
	rand.Seed(int64(time.Now().UnixNano()))
}

func RandonInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

func RandomOwner() string {
	return RandomString(6)
}
func RandomMoney() int64 {
	return RandonInt(0, 9999)
}
func RandomCurrency() string {
	currency := []string{"EUR", "CAD", "USD"}
	n := len(currency)
	return currency[rand.Intn(n)]
}
