package validations

import (
	"errors"
	"net"
	"net/url"
	"regexp"
	"strings"
)

var (
	// ErrEmptyIPAddress is returned when IP address is empty.
	ErrEmptyIPAddress = errors.New("IP address cannot be empty")
	// ErrInvalidIPAddress is returned when IP address is invalid.
	ErrInvalidIPAddress = errors.New("invalid IP address")
	// ErrEmptyHostname is returned when hostname is empty.
	ErrEmptyHostname = errors.New("hostname cannot be empty")
	// ErrHostnameTooLong is returned when hostname is too long.
	ErrHostnameTooLong = errors.New("hostname is too long")
	// ErrInvalidHostname is returned when hostname is invalid.
	ErrInvalidHostname = errors.New("invalid hostname")
	// ErrInvalidIPv6Address is returned when IPv6 address is invalid.
	ErrInvalidIPv6Address = errors.New("invalid IPv6 address")
	// ErrExpectedIPv6Address is returned when IPv6 address is expected.
	ErrExpectedIPv6Address = errors.New("IPv6 address expected")
	// ErrInvalidIPv4Address is returned when IPv4 address is invalid.
	ErrInvalidIPv4Address = errors.New("invalid IPv4 address")
	// ErrExpectedIPv4Address is returned when IPv4 address is expected.
	ErrExpectedIPv4Address = errors.New("IPv4 address expected")
	// ErrEmptyMACAddress is returned when MAC address is empty.
	ErrEmptyMACAddress = errors.New("MAC address cannot be empty")
	// ErrInvalidMACAddress is returned when MAC address is invalid.
	ErrInvalidMACAddress = errors.New("invalid MAC address")
	// ErrEmptyURL is returned when URL is empty.
	ErrEmptyURL = errors.New("URL cannot be empty")
	// ErrInvalidURL is returned when URL is invalid.
	ErrInvalidURL = errors.New("invalid URL")
	// ErrInvalidCIDRv4 is returned when a CIDRv4 address is invalid.
	ErrInvalidCIDRv4 = errors.New("invalid CIDRv4 address")
	// ErrInvalidCIDRv6 is returned when a CIDRv6 address is invalid.
	ErrInvalidCIDRv6 = errors.New("invalid CIDRv6 address")
	// ErrInvalidDataURL is returned when a data URL is invalid.
	ErrInvalidDataURL = errors.New("invalid data URL")
	// ErrInvalidFQDN is returned when a fully qualified domain name is invalid.
	ErrInvalidFQDN = errors.New("invalid FQDN")
	// ErrInvalidRFC952Hostname is returned when an RFC 952 hostname is invalid.
	ErrInvalidRFC952Hostname = errors.New("invalid RFC 952 hostname")
	// ErrInvalidTCP4Addr is returned when ValidateTCP4Addr is given an invalid TCPv4 address.
	ErrInvalidTCP4Addr = errors.New("invalid TCPv4 address")
	// ErrInvalidTCP6Addr is returned when ValidateTCP6Addr is given an invalid TCPv6 address.
	ErrInvalidTCP6Addr = errors.New("invalid TCPv6 address")
	// ErrInvalidTCPAddr is returned when ValidateTCPAddr is given an invalid TCP address.
	ErrInvalidTCPAddr = errors.New("invalid TCP address")
	// ErrInvalidUDP4Addr is returned when ValidateUDP4Addr is given an invalid UDPv4 address.
	ErrInvalidUDP4Addr = errors.New("invalid UDPv4 address")
	// ErrInvalidUDP6Addr is returned when ValidateUDP6Addr is given an invalid UDPv6 address.
	ErrInvalidUDP6Addr = errors.New("invalid UDPv6 address")
	// ErrInvalidUDPAddr is returned when ValidateUDPAddr is given an invalid UDP address.
	ErrInvalidUDPAddr = errors.New("invalid UDP address")
)

// ValidateIPAddress validates an IPv4 or IPv6 address.
func ValidateIPAddress(ipAddress string) error {
	if ipAddress == "" {
		return ErrEmptyIPAddress
	}
	if net.ParseIP(ipAddress) == nil {
		return ErrInvalidIPAddress
	}
	return nil
}

// ValidateHostname validates a hostname.
func ValidateHostname(hostname string) error {
	if hostname == "" {
		return ErrEmptyHostname
	}
	if len(hostname) > 255 {
		return ErrHostnameTooLong
	}
	if hostname[len(hostname)-1] == '.' {
		hostname = hostname[:len(hostname)-1]
	}
	allowed := regexp.MustCompile(`^([a-zA-Z0-9]|[a-zA-Z0-9][a-zA-Z0-9\-]*[a-zA-Z0-9])$`)
	for _, label := range strings.Split(hostname, ".") {
		if len(label) > 63 || !allowed.MatchString(label) {
			return ErrInvalidHostname
		}
	}
	return nil
}

// ValidateIPv6Address validates an IPv6 address.
func ValidateIPv6Address(ipAddress string) error {
	if ipAddress == "" {
		return ErrEmptyIPAddress
	}
	ip := net.ParseIP(ipAddress)
	if ip == nil {
		return ErrInvalidIPv6Address
	}
	if ip.To4() != nil {
		return ErrExpectedIPv6Address
	}
	return nil
}

// ValidateIPv4Address validates an IPv4 address.
func ValidateIPv4Address(ipAddress string) error {
	if ipAddress == "" {
		return ErrEmptyIPAddress
	}
	ip := net.ParseIP(ipAddress)
	if ip == nil {
		return ErrInvalidIPv4Address
	}
	if ip.To4() == nil {
		return ErrExpectedIPv4Address
	}
	return nil
}

// ValidateMACAddress validates a MAC address.
func ValidateMACAddress(macAddress string) error {
	if macAddress == "" {
		return ErrEmptyMACAddress
	}
	macAddressRegex := regexp.MustCompile(`^([0-9A-Fa-f]{2}[:-]){5}([0-9A-Fa-f]{2})$`)
	if !macAddressRegex.MatchString(macAddress) {
		return ErrInvalidMACAddress
	}
	return nil
}

// ValidateURL validates a URL.
func ValidateURL(urlString string) error {
	if urlString == "" {
		return ErrEmptyURL
	}
	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return ErrInvalidURL
	}
	return nil
}

var ()

// ValidateCIDRv4 validates a CIDRv4 address.
func ValidateCIDRv4(cidr string) error {
	_, _, err := net.ParseCIDR(cidr)
	if err != nil {
		return ErrInvalidCIDRv4
	}
	return nil
}

// ValidateCIDRv6 validates a CIDRv6 address.
func ValidateCIDRv6(cidr string) error {
	_, _, err := net.ParseCIDR(cidr)
	if err != nil {
		return ErrInvalidCIDRv6
	}
	return nil
}

// ValidateDataURL validates a data URL.
func ValidateDataURL(dataURL string) error {
	pattern := "^data:[a-z]+/[a-z]+(;[a-z-]+=[a-z-]+)*;base64,[a-zA-Z0-9/+=]+$"
	matched, err := regexp.MatchString(pattern, dataURL)
	if err != nil || !matched {
		return ErrInvalidDataURL
	}
	return nil
}

// ValidateFQDN validates a fully qualified domain name.
func ValidateFQDN(fqdn string) error {
	_, err := net.LookupHost(fqdn)
	if err != nil {
		return ErrInvalidFQDN
	}
	return nil
}

// ValidateRFC952 validates an RFC 952 hostname.
func ValidateRFC952(hostname string) error {
	pattern := "^[a-zA-Z][a-zA-Z0-9-]{0,22}[a-zA-Z0-9]$"
	matched, err := regexp.MatchString(pattern, hostname)
	if err != nil || !matched {
		return ErrInvalidRFC952Hostname
	}
	return nil
}

// ValidateTCP4Addr checks if the given address is a valid TCPv4 address.
func ValidateTCP4Addr(addr string) error {
	if _, err := net.ResolveTCPAddr("tcp4", addr); err != nil {
		return ErrInvalidTCP4Addr
	}
	return nil
}

// ValidateTCP6Addr checks if the given address is a valid TCPv6 address.
func ValidateTCP6Addr(addr string) error {
	if _, err := net.ResolveTCPAddr("tcp6", addr); err != nil {
		return ErrInvalidTCP6Addr
	}
	return nil
}

// ValidateTCPAddr checks if the given address is a valid TCP address.
func ValidateTCPAddr(addr string) error {
	if _, err := net.ResolveTCPAddr("tcp", addr); err != nil {
		return ErrInvalidTCPAddr
	}
	return nil
}

// ValidateUDP4Addr checks if the given address is a valid UDPv4 address.
func ValidateUDP4Addr(addr string) error {
	if _, err := net.ResolveUDPAddr("udp4", addr); err != nil {
		return ErrInvalidUDP4Addr
	}
	return nil
}

// ValidateUDP6Addr checks if the given address is a valid UDPv6 address.
func ValidateUDP6Addr(addr string) error {
	if _, err := net.ResolveUDPAddr("udp6", addr); err != nil {
		return ErrInvalidUDP6Addr
	}
	return nil
}

// ValidateUDPAddr checks if the given address is a valid UDP address.
func ValidateUDPAddr(addr string) error {
	if _, err := net.ResolveUDPAddr("udp", addr); err != nil {
		return ErrInvalidUDPAddr
	}
	return nil
}
