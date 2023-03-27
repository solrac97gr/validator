package validations

import (
	"encoding/base64"
	"errors"
	"regexp"
)

var (
	ErrInvalidBase64           = errors.New("invalid base64 encoding")
	ErrInvalidBase64URL        = errors.New("invalid base64url encoding")
	ErrInvalidBase64RawURL     = errors.New("invalid base64rawurl encoding")
	ErrInvalidBIC              = errors.New("invalid BIC")
	ErrInvalidBCP47LanguageTag = errors.New("invalid BCP47 language tag")
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

// IsBCP47LanguageTag checks if the given string is a valid BCP47 language tag.
func IsBCP47LanguageTag(str string) error {
	// This function checks for the syntax of the language tag, but doesn't validate the subtags against the IANA Language Subtag Registry.

	r := regexp.MustCompile(`^[a-zA-Z]{1,8}(-[a-zA-Z0-9]{1,8})*$`)
	if !r.MatchString(str) {
		return ErrInvalidBCP47LanguageTag
	}
	return nil
}

// IsBTCAddress checks if the given string is a valid Bitcoin address.
func IsBTCAddress(str string) error {
	if len(str) < 26 || len(str) > 35 {
		return errors.New("invalid Bitcoin address")
	}

	// Check the characters are valid
	for _, c := range str {
		if !((c >= '0' && c <= '9') || (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z')) {
			return errors.New("invalid Bitcoin address")
		}
	}

	// Decode the base58 string
	decoded, err := base58Decode(str)
	if err != nil {
		return errors.New("invalid Bitcoin address")
	}

	// Check the length of the decoded string
	if len(decoded) != 25 {
		return errors.New("invalid Bitcoin address")
	}

	// Check the version byte (0x00 for mainnet, 0x6f for testnet)
	if decoded[0] != 0x00 && decoded[0] != 0x6f {
		return errors.New("invalid Bitcoin address")
	}

	// Check the checksum
	checksum := sha256(sha256(decoded[:21]))[:4]
	if string(checksum) != string(decoded[21:]) {
		return errors.New("invalid Bitcoin address")
	}

	return nil
}
