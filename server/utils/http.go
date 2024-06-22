package utils

import (
	"net/http"
)

func DoRequest(request *http.Request) (*http.Response, error) {
	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return resp, err
	}
	return resp, err
}
