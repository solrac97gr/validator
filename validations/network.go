package validations

import (
	"fmt"
	"net"
	"net/url"
	"regexp"
	"strings"
)

var (
	errEmptyIPAddress      = fmt.Errorf("IP address cannot be empty")
	errInvalidIPAddress    = fmt.Errorf("invalid IP address")
	errEmptyHostname       = fmt.Errorf("hostname cannot be empty")
	errHostnameTooLong     = fmt.Errorf("hostname is too long")
	errInvalidHostname     = fmt.Errorf("invalid hostname")
	errInvalidIPv6Address  = fmt.Errorf("invalid IPv6 address")
	errExpectedIPv6Address = fmt.Errorf("IPv6 address expected")
	errInvalidIPv4Address  = fmt.Errorf("invalid IPv4 address")
	errExpectedIPv4Address = fmt.Errorf("IPv4 address expected")
	errEmptyMACAddress     = fmt.Errorf("MAC address cannot be empty")
	errInvalidMACAddress   = fmt.Errorf("invalid MAC address")
	errEmptyURL            = fmt.Errorf("URL cannot be empty")
	errInvalidURL          = fmt.Errorf("invalid URL")
)

// ValidateIPAddress validates an IPv4 or IPv6 address.
func ValidateIPAddress(ipAddress string) error {
	if ipAddress == "" {
		return errEmptyIPAddress
	}
	if net.ParseIP(ipAddress) == nil {
		return errInvalidIPAddress
	}
	return nil
}

// ValidateHostname validates a hostname.
func ValidateHostname(hostname string) error {
	if hostname == "" {
		return errEmptyHostname
	}
	if len(hostname) > 255 {
		return errHostnameTooLong
	}
	if hostname[len(hostname)-1] == '.' {
		hostname = hostname[:len(hostname)-1]
	}
	allowed := regexp.MustCompile(`^([a-zA-Z0-9]|[a-zA-Z0-9][a-zA-Z0-9\-]*[a-zA-Z0-9])$`)
	for _, label := range strings.Split(hostname, ".") {
		if len(label) > 63 || !allowed.MatchString(label) {
			return errInvalidHostname
		}
	}
	return nil
}

// ValidateIPv6Address validates an IPv6 address.
func ValidateIPv6Address(ipAddress string) error {
	if ipAddress == "" {
		return errEmptyIPAddress
	}
	ip := net.ParseIP(ipAddress)
	if ip == nil {
		return errInvalidIPv6Address
	}
	if ip.To4() != nil {
		return errExpectedIPv6Address
	}
	return nil
}

// ValidateIPv4Address validates an IPv4 address.
func ValidateIPv4Address(ipAddress string) error {
	if ipAddress == "" {
		return errEmptyIPAddress
	}
	ip := net.ParseIP(ipAddress)
	if ip == nil {
		return errInvalidIPv4Address
	}
	if ip.To4() == nil {
		return errExpectedIPv4Address
	}
	return nil
}

// ValidateMACAddress validates a MAC address.
func ValidateMACAddress(macAddress string) error {
	if macAddress == "" {
		return errEmptyMACAddress
	}
	macAddressRegex := regexp.MustCompile(`^([0-9A-Fa-f]{2}[:-]){5}([0-9A-Fa-f]{2})$`)
	if !macAddressRegex.MatchString(macAddress) {
		return errInvalidMACAddress
	}
	return nil
}

// ValidateURL validates a URL.
func ValidateURL(urlString string) error {
	if urlString == "" {
		return errEmptyURL
	}
	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return errInvalidURL
	}
	return nil
}
