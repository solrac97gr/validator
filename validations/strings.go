package validations

import "regexp"

// emailRegex is a regular expression for validating email addresses.
// It is based on the regex from the HTML5 specification.
// Used as var instead of const to avoid initialization loop.
// and to avoid the need to use regexp.MustCompile
var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9.!#$%&'*+/=?^_{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$`)

// StringIsPresent checks if a string is present.
func StringIsPresent(value string) bool {
	return value != ""
}

// StringIsBlank checks if a string is blank.
func StringIsBlank(value string) bool {
	return value == ""
}

// StringIsValidEmail checks if a string is a valid email address.
func StringIsValidEmail(email string) bool {
	return emailRegex.MatchString(email)
}
