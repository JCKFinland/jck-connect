package helper

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
)

func JSONRequest(
	method string,
	url string,
	body any,
) (*http.Request, error) {

	payload, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	req := httptest.NewRequest(
		method,
		url,
		bytes.NewReader(payload),
	)

	req.Header.Set(
		"Content-Type",
		"application/json",
	)

	return req, nil
}

func WithBearerToken(
	req *http.Request,
	token string,
) {

	req.Header.Set(
		"Authorization",
		"Bearer "+token,
	)
}
