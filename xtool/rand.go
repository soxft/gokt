package xtool

import (
	"math/rand"
	"strings"
	"time"
)

const strList string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandStr(length int) string {
	var result strings.Builder
	result.Grow(length)

	r := rand.New(rand.NewSource(time.Now().UnixNano() + int64(rand.Intn(100))))
	for i := 0; i < length; i++ {
		result.WriteByte(strList[r.Int63()%int64(len(strList))])
	}
	return result.String()
}

func RandInt(length int) int {
	var code int
	for i := 0; i < 5; i++ {
		code += rand.Intn(10)
	}
	return code
}
