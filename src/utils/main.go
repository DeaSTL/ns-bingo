package utils

import "math/rand"


func GenID(length int) string {
	const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	id := make([]byte, length)

	for i := range id {
		id[i] = alphabet[rand.Intn(len(alphabet))]
	}
	return string(id)
}

