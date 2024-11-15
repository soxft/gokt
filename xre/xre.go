package xre

import "regexp"

// IsEmail checks if the provided email string is a valid email address format.
// It returns true if the email matches the pattern, otherwise false.
func IsEmail(email string) bool {
	b, _ := regexp.MatchString("^([a-z0-9_.-]+)@([da-z.-]+).([a-z.]{2,6})$", email)
	return b
}

// IsUserName checks if the provided id string is a valid username format.
// A valid username contains only alphanumeric characters and is between 5 and 30 characters long.
// It returns true if the id matches the pattern, otherwise false.
func IsUserName(id string) bool {
	b, _ := regexp.MatchString("^[0-9a-zA-Z]{5,30}$", id)
	return b
}

// IsPassword checks if the provided password string is a valid password format.
// A valid password is between 8 and 64 characters long.
// It returns true if the password length is within the valid range, otherwise false.
func IsPassword(password string) bool {
	if len(password) < 8 || len(password) > 64 {
		return false
	}
	return true
}

// IsDomain checks if the provided domain string is a valid domain format.
// It returns true if the domain matches the pattern, otherwise false.
func IsDomain(domain string) bool {
	b, _ := regexp.MatchString("^[a-zA-Z0-9][a-zA-Z0-9-]{0,61}[a-zA-Z0-9]\\.[a-zA-Z.]{2,6}$", domain)
	return b
}
