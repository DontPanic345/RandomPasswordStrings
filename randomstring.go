package main

import (
	"fmt"
	"strings"
	"math/big"
	"crypto/rand"
)

const lower = "abcdefghijklmnopqrstuvwxyz"
const upper = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const number = "0123456789"
const symbol = "!@#$%^&*"
const choiceset = lower + upper + number + symbol

func RandStringBytes(n int) string {
	biglength := big.NewInt(int64(len(choiceset)))
    bytes := make([]byte, n)
    for i := range bytes {
		randbig, err := rand.Int(rand.Reader, biglength)
		if (err != nil) {
			panic(err)
		}
		randNum := int(randbig.Int64())
        bytes[i] = choiceset[randNum]
    }
    return string(bytes)
}

func CheckIfContainsAllGroups(str string) bool {
	l := strings.ContainsAny(lower, str)
	u := strings.ContainsAny(upper, str)
	n := strings.ContainsAny(number, str)
	s := strings.ContainsAny(symbol, str)

	return l && u && n && s 
}

func main() {
	s := RandStringBytes(7)
	count := 1
	for ; CheckIfContainsAllGroups(s) == false; count++ {
		s = RandStringBytes(7)
	}
	fmt.Printf("%s\n", s)
	fmt.Printf("%v\n", count)
}
