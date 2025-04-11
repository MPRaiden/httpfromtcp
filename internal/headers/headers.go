package headers

import (
	"bytes"
	"errors"
	"fmt"
	"strings"
)

type Headers map[string]string

const crlf = "\r\n"

func (h Headers) Parse(data []byte) (n int, done bool, err error) {
	idx := bytes.Index(data, []byte(crlf))
	if idx == -1 {
		return 0, false, nil
	}

	if idx == 0 {
		return len(crlf), true, nil // Consume 2 bytes, signify parsing is done, no error
	}

	headers := string(data[:idx])
	parts := strings.SplitN(headers, ":", 2)

	if len(parts) != 2 {
		return 0, false, errors.New("func Parse() - invalid header format: missing ':' separator")
	}

	key := parts[0]
	key = strings.TrimSpace(key) // Finally clean valid keys

	err = validateHeaderKey(key)
	if err != nil {
		return 0, false, err
	}

	value := strings.TrimSpace(parts[1])

	lowerCaseKey := strings.ToLower(key)

	h[lowerCaseKey] = value

	return (idx + len(crlf)), false, nil
}

func validateHeaderKey(key string) error {
	if key == "" {
		return errors.New("func validateHeaderKey() - invalid header format: empty key")
	}

	// Check for valid characters
	for _, char := range key {
		// Check if character is in the allowed set:
		// A-Z, a-z, 0-9, or one of the special characters: !#$%&'*+-.^_`|~
		if !((char >= 'A' && char <= 'Z') ||
			(char >= 'a' && char <= 'z') ||
			(char >= '0' && char <= '9') ||
			strings.ContainsRune("!#$%&'*+-.^_`|~", char)) {
			return fmt.Errorf("invalid character in header key: %c", char)
		}
	}

	return nil
}

func NewHeaders() Headers {
	return make(Headers)
}
