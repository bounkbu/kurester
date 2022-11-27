package util

import (
	"io"
	"io/ioutil"
	"net/http"
)

func NewHTTPRequest(url string, method string, body io.Reader) ([]byte, error) {
	request, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	res, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}

	responseData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return responseData, nil
}
