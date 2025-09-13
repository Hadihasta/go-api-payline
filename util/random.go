package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"


func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandomInt Generator (generates a random integer between min and max)
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max - min + 1)
}	

// RandomString 
func RandomString(n int) string { 
		var sb strings.Builder
		k := len (alphabet)

		for i := 0; i < n; i++ {
			c:= alphabet[rand.Intn(k)]
			sb.WriteByte(c)
		}
		return sb.String()
}

// RandomName
func  RandomName() string { 
	return RandomString(6)
}

// RandomMoney
func RandomMoneyNominal() int64 { 
	return RandomInt(0,1000)
}

//RandomRole
func RandomRole() string{ 
	roles := []string{"customer", "super_admin", "owner"}
	n := len(roles)
	return roles[rand.Intn(n)]
}