package text

import "math/rand"

// RandomString generates a random string of n length. Based on https://stackoverflow.com/a/22892986/1260548
func RandomString(charset []rune, n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

func LettersUppercase() []rune {
	// remove vowels to make it less likely to generate something offensive
	return []rune("BCDFGHJKLMNPQRSTVWXZ")
}

func LettersLowercase() []rune {
	return []rune("bcdfghjklmnpqrstvwxz")
}

func Digits() []rune {
	return []rune("0123456789")
}
