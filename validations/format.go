package validations

import (
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"math/big"
	"regexp"
	"strconv"
)

var (
	ErrInvalidBase64           = errors.New("invalid base64 encoding")
	ErrInvalidBase64URL        = errors.New("invalid base64url encoding")
	ErrInvalidBase64RawURL     = errors.New("invalid base64rawurl encoding")
	ErrInvalidBIC              = errors.New("invalid BIC")
	ErrInvalidBCP47LanguageTag = errors.New("invalid BCP47 language tag")
	// ErrInvalidCreditCard is returned when the provided credit card number is invalid
	ErrInvalidCreditCard = errors.New("invalid credit card number")
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
	checksum := sha256C(sha256C(decoded[:21]))[:4]
	if string(checksum) != string(decoded[21:]) {
		return errors.New("invalid Bitcoin address")
	}

	return nil
}

func base58Decode(str string) ([]byte, error) {
	alphabet := "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"

	// Build the decoding table
	decTable := make(map[byte]int)
	for i := 0; i < len(alphabet); i++ {
		decTable[alphabet[i]] = i
	}

	// Convert the string to a byte slice
	bytes := []byte(str)

	// Decode the base58 string
	result := big.NewInt(0)
	base := big.NewInt(58)
	for i := 0; i < len(bytes); i++ {
		char := bytes[i]
		val, ok := decTable[char]
		if !ok {
			return nil, errors.New("invalid base58 character")
		}
		result.Mul(result, base)
		result.Add(result, big.NewInt(int64(val)))
	}

	// Convert the result to a byte slice
	decoded := result.Bytes()

	// Pad the byte slice with zeros
	padSize := 0
	for i := 0; i < len(bytes) && bytes[i] == alphabet[0]; i++ {
		padSize++
	}
	resultBytes := make([]byte, padSize+len(decoded))
	copy(resultBytes[padSize:], decoded)

	return resultBytes, nil
}

func sha256C(data []byte) []byte {
	hasher := sha256.New()
	hasher.Write(data)
	return hasher.Sum(nil)
}

func IsValidCreditCard(str string) error {
	// Remove any non-digit characters
	digits := ""
	for _, c := range str {
		if c >= '0' && c <= '9' {
			digits += string(c)
		}
	}

	// Check the length is valid
	length := len(digits)
	if length < 13 || length > 19 {
		return ErrInvalidCreditCard
	}

	// Convert the string to a list of digits
	digitsList := make([]int, length)
	for i := 0; i < length; i++ {
		digit, err := strconv.Atoi(string(digits[i]))
		if err != nil {
			return ErrInvalidCreditCard
		}
		digitsList[i] = digit
	}

	// Calculate the sum of the digits using Luhn's algorithm
	sum := 0
	for i := length - 1; i >= 0; i-- {
		digit := digitsList[i]

		if (length-i)%2 == 0 {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}

		sum += digit
	}

	// Check if the sum is divisible by 10
	if sum%10 != 0 {
		return ErrInvalidCreditCard
	}

	return nil
}
