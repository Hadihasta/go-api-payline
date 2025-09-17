package util

import (
	"fmt"
	"math/rand"
	"strings"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

// hanya perlu di versi golang di bawah 1.20 agar rand tidak generate 1 terus terusan
// func init() {
// 	rand.Seed(time.Now().UnixNano())
// }

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

// RandomEmail
func RandomEmail() string{
	return fmt.Sprint("%s@email.com", RandomString(6))
}


// RandomPhoneNumer
func RandomPhoneNumber() string {

	// Prefix Indonesia biasanya "08"
	phone := "08"

	// Panjang total nomor HP biasanya 10–13 digit
	length := rand.Intn(4) + 10 // hasil: 10, 11, 12, atau 13 digit

	// Generate sisa digit
	for i := len(phone); i < length; i++ {
		phone += fmt.Sprintf("%d", rand.Intn(10))
	}
	return phone
}

// random store_access
func RandomStoreAccess() string{ 
	access := []string{"menu_qr ", "no_acces", "pos"}
	n := len(access)
	return access[rand.Intn(n)]
}

// random store_access
func CreateRandomStore() string{ 
	store := []string{"cafe_kopi", "restaurant_naga", "restaurant_kopi"}
	n := len(store)
	return store[rand.Intn(n)]
}
