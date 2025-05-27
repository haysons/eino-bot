package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// HTTPGet 发起http get请求
func HTTPGet[T any](url string) (*T, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http status code %d", resp.StatusCode)
	}
	res := new(T)
	if err = json.NewDecoder(resp.Body).Decode(res); err != nil {
		return nil, err
	}
	return res, nil
}
