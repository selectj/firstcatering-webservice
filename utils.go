package main

import "math/rand"

var characters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func randomString(length int) string {
	seq := make([]rune, length)
	for i := range seq {
		seq[i] = characters[rand.Intn(len(characters))]
	}

	return string(seq)
}
