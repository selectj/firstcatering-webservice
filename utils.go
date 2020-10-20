package main

import (
	"math/rand"
	"time"
)

var characters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func randomString(length int) string {
	seq := make([]rune, length)
	for i := range seq {
		seq[i] = characters[rand.Intn(len(characters))]
	}

	return string(seq)
}

func getCurrentTimeMillis() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

func hasCurrentCardSession(cardID string) bool {
	for _, session := range sessions {
		if session.Card.ID == cardID {
			return true
		}
	}
	return false
}

func doesSessionExist(sid string) bool {
	for _, session := range sessions {
		if session.ID == sid {
			return true
		}
	}
	return false
}

func getCurrentCardSession(cardID string) Session {
	var currentSession Session
	for _, session := range sessions {
		if session.Card.ID == cardID {
			currentSession = session
		}
	}

	return currentSession
}

func removeFromSlice(slice []Session, i int) []Session {
	slice[len(slice)-1], slice[i] = slice[i], slice[len(slice)-1]
	return slice[:len(slice)-1]
}
