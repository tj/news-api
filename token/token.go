package token

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"strings"
)

// helpers.
var (
	encode = base64.URLEncoding.EncodeToString
	decode = base64.URLEncoding.DecodeString
)

// Sign a value.
func Sign(secret, value string) string {
	m := hmac.New(sha256.New, []byte(secret))
	m.Write([]byte(value))
	return encode(m.Sum(nil)) + "." + encode([]byte(value))
}

// Unsign a value.
func Unsign(secret, msg string) (string, bool) {
	p := strings.Split(msg, ".")
	if len(p) != 2 {
		return "", false
	}

	signature, err := decode(p[0])
	if err != nil {
		return "", false
	}

	payload, err := decode(p[1])
	if err != nil {
		return "", false
	}

	m := hmac.New(sha256.New, []byte(secret))
	m.Write(payload)
	expected := m.Sum(nil)

	if !hmac.Equal(signature, expected) {
		return "", false
	}

	return string(payload), true
}
