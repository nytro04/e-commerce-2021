package random_utils

import (
	"fmt"
	"math/rand"
	"strings"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

// RandomInt generates a random number between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1) //returns numbers btw min and max
}

// RandomString generates a random string of length n
func RandomString(n int) string {
	var sb strings.Builder

	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

// RandomOwner a random Username
func RandomUser() string {
	return RandomString(6)
}

// RandomMoney generates a random price
func RandomPrice() int64 {
	return RandomInt(299, 1000)
}

// RandomEmail generates a emails
func RandomEmail() string {
	return fmt.Sprintf("%s@email.com", RandomString(6))
}
