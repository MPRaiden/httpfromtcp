package headers

import (
	"bytes"
	"errors"
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
	if strings.TrimSpace(key) != key {
		return 0, false, errors.New("func Parse() - invalid header format: spaces around key")
	}
	key = strings.TrimSpace(key) // Finally clean valid keys
	if key == "" {
		return 0, false, errors.New("func Parse() - invalid header format: empty key")
	}

	value := strings.TrimSpace(parts[1])

	h[key] = value

	return (idx + len(crlf)), false, nil
}

func NewHeaders() Headers {
	return make(Headers)
}
