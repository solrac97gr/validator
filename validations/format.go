package validations

import (
	"encoding/base64"
	"errors"
)

var (
	ErrInvalidBase64       = errors.New("invalid base64 encoding")
	ErrInvalidBase64URL    = errors.New("invalid base64url encoding")
	ErrInvalidBase64RawURL = errors.New("invalid base64rawurl encoding")
	ErrInvalidBIC          = errors.New("invalid BIC")
)

// IsBase64 checks if the given string is a valid base64 encoding.
func IsBase64(str string) error {
	_, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return ErrInvalidBase64
	}
	return nil
}

// IsBase64URL checks if the given string is a valid base64url encoding.
func IsBase64URL(str string) error {
	_, err := base64.URLEncoding.DecodeString(str)
	if err != nil {
		return ErrInvalidBase64URL
	}
	return nil
}

// IsBase64RawURL checks if the given string is a valid base64rawurl encoding.
func IsBase64RawURL(str string) error {
	_, err := base64.RawURLEncoding.DecodeString(str)
	if err != nil {
		return ErrInvalidBase64RawURL
	}
	return nil
}

// IsBIC checks if the given string is a valid BIC (Bank Identifier Code).
func IsBIC(str string) error {
	if len(str) != 8 {
		return ErrInvalidBIC
	}
	for i := 0; i < len(str); i++ {
		c := str[i]
		if c < 'A' || c > 'Z' {
			return ErrInvalidBIC
		}
	}
	return nil
}
