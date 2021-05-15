package main

import (
	"fmt"
	"math/rand"
)

func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyz1234567890")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func GeneratePorts() (int, int) {
	max := 32767 // half of 65535 because other half is for apiPort
	min := 1024
	// port := fmt.Sprintf("http://localhost:%d", RandomIntegerwithinRange)
	RandomInt1 := rand.Intn(max-min) + min // range is min to max
	RandomInt2 := RandomInt1 + 32767
	return RandomInt1, RandomInt2 // demoPort, apiPort
}
