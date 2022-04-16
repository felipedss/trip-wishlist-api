package rest

import (
	"net/url"
)

func BuildQueryParam(path string, query url.Values) (string, error) {
	parsedURL, err := url.Parse(path)
	if err != nil {
		return "", err
	}
	parsedURL.RawQuery = query.Encode()
	return parsedURL.String(), nil
}
