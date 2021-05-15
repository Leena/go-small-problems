//https://leetcode.com/problems/validate-ip-address/

package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func validIPAddress(IP string) string {
	if strings.Contains(IP, string(".")) {
		if validateIPv4(IP) {
			return "IPv4"
		}
	} else if strings.Contains(IP, string(":")) {
		if validateIPv6(IP) {
			return "IPv6"
		}
	}
	return "Neither"
}

func validateIPv4(IP string) bool {
	isNotAlpha := regexp.MustCompile(`^[A-Za-z]`).MatchString
	for _, letter := range IP {
		if isNotAlpha(string(letter)) {
			return false
		}
	}

	parts := strings.Split(IP, ".")
	if len(parts) != 4 {
		return false
	}

	partsValidator := func(IPSegment string) bool {
		digits, err := strconv.Atoi(IPSegment)
		if err != nil || string(IPSegment[0]) == "0" && len(IPSegment) > 1 {
			return false
		}

		if digits >= 0 && digits <= 255 {
			return true
		}

		if strconv.Itoa(digits) != IPSegment {
			return false
		}

		return false
	}

	for _, part := range parts {
		if !partsValidator(part) {
			return false
		}
	}

	return true
}

func validateIPv6(IP string) bool {

	parts := strings.Split(IP, ":")
	if len(parts) != 8 {
		return false
	}

	isHexa := ":1234567890abcdefABCDEF"
	for _, letter := range IP {
		if !strings.ContainsRune(isHexa, letter) {
			return false
		}
	}

	for _, part := range parts {
		size := len(part)
		if !(size >= 1 && size <= 4) {
			return false
		}
	}

	return true
}

func main() {
	IP1 := "192.168.1.1"
	IP2 := "192.168.1.0"
	IP3 := "192.168.01.1"
	IP4 := "192.168.1.00"
	IP5 := "192.168@1.1"
	IP6 := "172.16.254.1"
	IP7 := "2001:0db8:85a3:0000:0000:8a2e:0370:7334"
	IP8 := "2001:db8:85a3:0:0:8A2E:0370:7334"
	IP9 := "2001:0db8:85a3::8A2E:037j:7334"
	IP10 := "02001:0db8:85a3:0000:0000:8a2e:0370:7334"
	IP11 := "256.256.256.256"
	IP12 := "2001:0db8:85a3:0:0:8A2E:0370:7334:"
	IP13 := "1e1.4.5.6"

	fmt.Println(validIPAddress(IP1) == "IPv4")
	fmt.Println(validIPAddress(IP2) == "IPv4")
	fmt.Println(validIPAddress(IP3) == "Neither")
	fmt.Println(validIPAddress(IP4) == "Neither")
	fmt.Println(validIPAddress(IP5) == "Neither")
	fmt.Println(validIPAddress(IP6) == "IPv4")
	fmt.Println(validIPAddress(IP7) == "IPv6")
	fmt.Println(validIPAddress(IP8) == "IPv6")
	fmt.Println(validIPAddress(IP9) == "Neither")
	fmt.Println(validIPAddress(IP10) == "Neither")
	fmt.Println(validIPAddress(IP11) == "Neither")
	fmt.Println(validIPAddress(IP12) == "Neither")
	fmt.Println(validIPAddress(IP13) == "Neither")
}
