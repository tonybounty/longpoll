package longpoll

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
)

// RunWithParm run a long polling GET request with parameters and return result
func RunWithParm(ctx context.Context, url string, param map[string]string) ([]byte, error) {
	url += "?"
	for index, value := range param {
		url += "&" + index + "=" + value
	}
	return Run(ctx, url)
}

// Run make a long polling request and return result
func Run(ctx context.Context, url string) ([]byte, error) {
	client := http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	res, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("%v\b", err)
	}
	defer res.Body.Close()

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("%v\b", err)
	}

	return b, nil
}
