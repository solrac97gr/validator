package validations

import (
	"net"
	"net/url"
	"regexp"
)

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

// StringIsValidHostname checks if a string is a valid hostname.
// The string must not be longer than 255 characters.
// The string must not end with a dot.
// The string must only consist of valid characters.
// The string must resolve to at least one IP address.
func StringIsValidHostname(s string) bool {
	// Compile the regular expression.
	// The regular expression must match the entire string.
	re, err := regexp.Compile(`^(([a-zA-Z0-9]|[a-zA-Z0-9][a-zA-Z0-9\-]*[a-zA-Z0-9])\.)*([A-Za-z0-9]|[A-Za-z0-9][A-Za-z0-9\-]*[A-Za-z0-9])$`)
	if err != nil {
		return false
	}
	if len(s) > 255 {
		return false
	}
	if s[len(s)-1] == '.' {
		return false
	}
	if !re.MatchString(s) {
		return false
	}
	// Lookup the IP addresses for the hostname.
	ips, err := net.LookupIP(s)
	if err != nil {
		return false
	}
	if len(ips) == 0 {
		return false
	}
	return true
}

// StringIsIPv4 checks if a string is a valid IPv4 address.
func StringIsIPv4(s string) bool {
	ip := net.ParseIP(s)
	return ip != nil && ip.To4() != nil
}

// StringIsIPv6 checks if a string is a valid IPv6 address.
func StringIsIPv6(s string) bool {
	ip := net.ParseIP(s)
	return ip != nil && ip.To4() == nil
}

// StringIsURL checks if a string is a valid URL.
func StringIsURL(s string) bool {
	u, err := url.Parse(s)
	return err == nil && u.Scheme != "" && u.Host != ""
}
