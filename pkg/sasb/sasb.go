// Package sasb is the library implementation.
package sasb

import "fmt"

// StringAsBytes creates a stringified version of a byte representation of a
// string, including binary, octal, and hex. Binary, octal, and hexadecimal
// are 0-padded.
func StringAsBytes(s string) (binary, octal, decimal, hexadecimal []string) {
	asBytes := []byte(s)

	for _, v := range asBytes {
		binary = append(binary, fmt.Sprintf("%08b", v))
		octal = append(octal, fmt.Sprintf("%03o", v))
		decimal = append(decimal, fmt.Sprintf("%d", v))
		hexadecimal = append(hexadecimal, fmt.Sprintf("%02x", v))
	}

	return
}
