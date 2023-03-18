package validations

import (
	"errors"
	"regexp"
	"strings"
)

var (
	ErrNotAlpha               = errors.New("string is not alpha")
	ErrNotAlphanumeric        = errors.New("string is not alphanumeric")
	ErrNotAlphanumericUnicode = errors.New("string is not alphanumeric unicode")
	ErrNotAlphaUnicode        = errors.New("string is not alpha unicode")
	ErrNotASCIICode           = errors.New("string is not ASCII code")
	ErrNotBoolean             = errors.New("string is not a boolean")
	ErrDoesNotContain         = errors.New("string does not contain the substring")
	ErrDoesNotContainAny      = errors.New("string does not contain any of the substrings")
	ErrDoesNotContainRune     = errors.New("string does not contain the rune")
	ErrEndsWith               = errors.New("string ends with the substring")
	ErrEndsNotWith            = errors.New("string does not end with the substring")
	ErrExcludes               = errors.New("string includes the substring")
	ErrExcludesAll            = errors.New("string does not include all of the substrings")
	ErrIncludesAll            = errors.New("string does not include all of the substrings")
	ErrExcludesRune           = errors.New("string includes the rune")
	ErrNotLowerCase           = errors.New("string is not lowercase")
	ErrNotUpercase            = errors.New("string is not uppercase")
	ErrNotMultibyte           = errors.New("string does not contain one or more multibyte characters")
	ErrNotNumeric             = errors.New("string is not numeric")
	ErrNotPrintableASCII      = errors.New("string contains non-printable ASCII characters")
	ErrStartsWith             = errors.New("string starts with the substring")
	ErrNotStartsWith          = errors.New("string does not start with the substring")
)

// StringIsAlpha checks if a string contains only letters.
func StringIsAlpha(s string) error {
	if regexp.MustCompile(`^[a-zA-Z]+$`).MatchString(s) {
		return nil
	}
	return ErrNotAlpha
}

// StringIsAlphanumeric checks if a string contains only numbers.
func StringIsAlphanumeric(s string) error {
	if !regexp.MustCompile(`^[a-zA-Z0-9]+$`).MatchString(s) {
		return ErrNotAlphanumeric
	}
	return nil
}

// StringIsAlphaUnicode checks if a string contains only unicode letters.
func StringIsAlphanumericUnicode(s string) error {
	if !regexp.MustCompile(`^[\p{L}\p{N}]+$`).MatchString(s) {
		return ErrNotAlphanumericUnicode
	}
	return nil
}

// StringIsAlphaUnicode checks if a string contains only unicode letters.
func StringIsAlphaUnicode(s string) error {
	if !regexp.MustCompile(`^[\p{L}]+$`).MatchString(s) {
		return ErrNotAlphaUnicode
	}
	return nil
}

// StringIsASCIICode checks if a string contains only ASCII characters.
func StringIsASCIICode(s string) error {
	if !regexp.MustCompile(`^[\x00-\x7F]+$`).MatchString(s) {
		return ErrNotASCIICode
	}
	return nil
}

// StringIsBoolean checks if a string is a boolean.
func StringIsBoolean(s string) error {
	if !regexp.MustCompile(`^(?i)(true|false)$`).MatchString(s) {
		return ErrNotBoolean
	}
	return nil
}

// StringContains checks if a string contains a substring.
func StringContains(s, substr string) error {
	if !strings.Contains(s, substr) {
		return ErrDoesNotContain
	}
	return nil
}

// StringContainsAny checks if a string contains any of the substrings.
func StringContainsAny(s string, substrs ...string) error {
	for _, substr := range substrs {
		if strings.Contains(s, substr) {
			return nil
		}
	}
	return ErrDoesNotContainAny
}

// StringContainsRune checks if a string contains a rune.
func StringContainsRune(s string, r rune) error {
	if !strings.ContainsRune(s, r) {
		return ErrDoesNotContainRune
	}
	return nil
}

// StringEndsNotWith checks if a string does not end with a substring.
func StringEndsNotWith(s, substr string) error {
	if strings.HasSuffix(s, substr) {
		return ErrEndsWith
	}
	return nil
}

// StringEndsWith checks if a string ends with a substring.
func StringEndsWith(s, substr string) error {
	if !strings.HasSuffix(s, substr) {
		return ErrEndsNotWith
	}
	return nil
}

// StringExcludes checks if a string excludes a substring.
func StringExcludes(s, substr string) error {
	if strings.Contains(s, substr) {
		return ErrExcludes
	}
	return nil
}

// StringExcludesAll checks if a string excludes all of the substrings.
func StringExcludesAll(s string, substrs ...string) error {
	for _, substr := range substrs {
		if strings.Contains(s, substr) {
			return ErrExcludesAll
		}
	}
	return nil
}

// StringExcludesRune checks if a string excludes a rune.
func StringExcludesRune(s string, r rune) error {
	if strings.ContainsRune(s, r) {
		return ErrExcludesRune
	}
	return nil
}

// StringIsLowerCase checks if a string is lowercase.
func StringIsLowerCase(s string) error {
	if s != strings.ToLower(s) {
		return ErrNotLowerCase
	}
	return nil
}

// StringIsUpperCase checks if a string is uppercase.
func StringIsUpperCase(s string) error {
	if s != strings.ToUpper(s) {
		return ErrNotUpercase
	}
	return nil
}

// StringIsMultibyte checks if a string contains one or more multibyte characters.
func StringIsMultibyte(s string) error {
	if len(s) == len([]rune(s)) {
		return ErrNotMultibyte
	}
	return nil
}

// StringIsNumeric checks if a string contains only numbers.
func StringIsNumeric(s string) error {
	if !regexp.MustCompile(`^[0-9]+$`).MatchString(s) {
		return ErrNotNumeric
	}
	return nil
}

// StringIsPrintableASCII checks if a string contains only printable ASCII characters.
func StringIsPrintableASCII(s string) error {
	if !regexp.MustCompile(`^[\x20-\x7E]+$`).MatchString(s) {
		return ErrNotPrintableASCII
	}
	return nil
}

// StringStartsNotWith checks if a string does not start with a substring.
func StringStartsNotWith(s, substr string) error {
	if strings.HasPrefix(s, substr) {
		return ErrStartsWith
	}
	return nil
}

// StringStartsWith checks if a string starts with a substring.
func StringStartsWith(s, substr string) error {
	if !strings.HasPrefix(s, substr) {
		return ErrNotStartsWith
	}
	return nil
}
