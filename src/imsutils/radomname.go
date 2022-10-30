package imsutils

import (
	"math/rand"
	"strings"
	"time"
)

func GetRandstring(length int) string {
	if length < 1 {
		return ""
	}
	char := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	charArr := strings.Split(char, "")
	ran := rand.New(rand.NewSource(time.Now().Unix()))

	l := len(charArr)
	for i := l - 1; i > 0; i-- {
		r := ran.Intn(i)
		charArr[r], charArr[i] = charArr[i], charArr[r]
	}
	rchar := charArr[:length]
	return strings.Join(rchar, "")
}
