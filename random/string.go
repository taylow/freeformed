package random

import "math/rand"

const (
	Alphabet     = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	Numbers      = "0123456789"
	AlphaNumeric = Alphabet + Numbers
)

type RandomStringGenerator func(int) string

// GenerateRandomAlphaString generates a random string of the given length using the alphabet
func GenerateRandomAlphaString(length int) string {
	str := make([]byte, length)
	for i := range str {
		str[i] = Alphabet[rand.Intn(len(Alphabet))]
	}
	return string(str)
}

// GenerateRandomAlphaNumericString generates a random string of the given length using the alphabet and numbers
func GenerateRandomAlphaNumericString(length int) string {
	str := make([]byte, length)
	for i := range str {
		str[i] = AlphaNumeric[rand.Intn(len(AlphaNumeric))]
	}
	return string(str)
}
