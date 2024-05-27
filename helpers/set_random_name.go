package helpers

import "math/rand"

const letters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func GenerateStringName() string {
	b := make([]byte, 5)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
