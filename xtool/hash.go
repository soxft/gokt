package xtool

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
)

// Md5 generates an MD5 hash for the given string.
// It returns the hexadecimal representation of the hash.
//
// Parameters:
// - str: The input string to hash.
//
// Returns:
// - A string containing the hexadecimal representation of the MD5 hash.
func Md5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

// Sha1 generates a SHA-1 hash for the given string.
// It returns the hexadecimal representation of the hash.
//
// Parameters:
// - str: The input string to hash.
//
// Returns:
// - A string containing the hexadecimal representation of the SHA-1 hash.
func Sha1(str string) string {
	h := sha1.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}
