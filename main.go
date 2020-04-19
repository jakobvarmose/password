package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"strings"
)

func choose(s string) string {
	max := big.NewInt(int64(len(s)))
	r, err := rand.Int(rand.Reader, max)
	if err != nil {
		panic(err)
	}
	return s[r.Int64() : r.Int64()+1]
}

func main() {
	lowercase := "abcdefghijkmnopqrstuvwxyz"
	uppercase := "ABCDEFGHJKLMNPQRSTUVWXYZ"
	numbers := "23456789"
	symbols := "!\"#$%&'()*+,-./:;<=>?@[\\]^_{|}~"

	password1 := ""
	for {
		for i := 0; i < 8; i++ {
			password1 += choose(lowercase + numbers)
		}
		if !strings.ContainsAny(password1, lowercase) {
			password1 = ""
			continue
		}
		if !strings.ContainsAny(password1, numbers) {
			password1 = ""
			continue
		}
		break
	}
	password1 += choose(uppercase)
	password1 += choose(symbols)
	fmt.Println(password1)
}
