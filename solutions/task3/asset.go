package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	var letterRunes = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, 1000)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	fmt.Println(string(b))
}
