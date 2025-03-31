package request

import (
	"fmt"
	"io"
	"log"
	"strings"
)

type Request struct {
	RequestLine RequestLine
}

type RequestLine struct {
	HttpVersion   string
	RequestTarget string
	Method        string
}

func RequestFromReader(reader io.Reader) (*Request, error) {
	bytes, err := io.ReadAll(reader)
	if err != nil {
		log.Fatal(err)
	}

	requestStr := string(bytes)
	requestLine := strings.Split(requestStr, "\r\n")[0]
	reqLineParts := strings.Split(requestLine, " ")

	// Request line needs to have 3 parts, the method, target and HTTP version
	if len(reqLineParts) != 3 {
		return nil, fmt.Errorf("Invalid request line, expected 3 parts but got %d", len(reqLineParts))
	}

	method := reqLineParts[0]
	target := reqLineParts[1]
	httpVersion := reqLineParts[2]

	// Method can only have capital letters
	if strings.ToUpper(method) != method {
		return nil, fmt.Errorf("Invalid Http method, expected all upper case letters, got %s", method)
	}

	// HTTP version should only be "HTTP/1.1"
	if httpVersion != "HTTP/1.1" {
		return nil, fmt.Errorf("Invalid request line, expected HTTP/1.1 for Http version but got %s", httpVersion)
	}

	reqLine := RequestLine{
		HttpVersion:   strings.TrimPrefix(httpVersion, "HTTP/"),
		RequestTarget: target,
		Method:        method,
	}

	req := Request{
		RequestLine: reqLine,
	}

	return &req, nil
}
