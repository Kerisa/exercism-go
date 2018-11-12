package cipher

import (
	"strings"
	"regexp"
)

// Cipher interface must have encode/decode method
type Cipher interface {
	Encode(string) string
	Decode(string) string
}

// Caesar represent for Caesar Cipher
type Caesar struct {
	shift int
}

// Vigenere represent for Vigenère cipher
type Vigenere struct {
	key string
}

// NewCaesar generates a Caesar Cipher with a fixed shift distance of 3
func NewCaesar() *Caesar {
	return &Caesar{shift:3}
}

// NewShift generates a Caesar Cipher with a flexible shift distance
func NewShift(shift int) *Caesar {
	if shift >= 26 || shift <= -26 || shift == 0 {
		return nil
	}
	return &Caesar{shift:shift}
} 

// NewVigenere generates a Vigenère cipher
func NewVigenere(key string) *Vigenere {
	key = regexp.MustCompile("^a*$").ReplaceAllString(key, "")
	if regexp.MustCompile("[^a-z]+").Match([]byte(key)) {
		return nil
	}
	if len(key) == 0 {
		return nil
	}
	return &Vigenere{key:key}
}

// Encode encodes string using Caesar Cipher
func (C *Caesar)Encode(in string) string {
	var out string
	for _,c := range strings.ToLower(regexp.MustCompile("[^a-zA-Z]*").ReplaceAllString(in, "")) {
		offset := byte((int(c - 'a') + C.shift) % 26 + 26) % 26
		out += string('a' + offset)
	}
	return out
}

// Decode encodes string using Caesar Cipher
func (C *Caesar)Decode(in string) string {
	var out string
	for _,c := range strings.ToLower(in) {
		offset := byte((int(c - 'a') - C.shift) % 26 + 26) % 26
		out += string('a' + offset)
	}
	return out
}

// Encode encodes string using Vigenère cipher
func (V *Vigenere)Encode(in string) string {
	var out string
	for i,c := range strings.ToLower(regexp.MustCompile("[^a-zA-Z]*").ReplaceAllString(in, "")) {
		shift := int(V.key[i % len(V.key)] - 'a')
		offset := (byte((int(c - 'a') + shift) % 26) + 26 ) % 26
		out += string('a' + offset)
	}
	return out
}

// Decode encodes string using Vigenère cipher
func (V *Vigenere)Decode(in string) string {
	var out string
	for i,c := range strings.ToLower(in) {
		shift := int(V.key[i % len(V.key)] - 'a')
		offset := (byte((int(c - 'a') - shift) % 26) + 26 ) % 26
		out += string('a' + offset)
	}
	return out
}