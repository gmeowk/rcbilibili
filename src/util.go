package main

import (
	"errors"
	"log"
	"net/http"
	"time"
)

// Get Http请求
func Get(url string, params map[string]string, headers map[string]string) (*http.Response, error) {
	//new request
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Println(err)
		return nil, errors.New("new request is fail ")
	}
	//add params
	q := req.URL.Query()
	if params != nil {
		for key, val := range params {
			q.Add(key, val)
		}
		req.URL.RawQuery = q.Encode()
	}
	//add headers
	req.Header.Add("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.105 Safari/537.36")
	req.Header.Add("origin", "https://live.bilibili.com")
	req.Header.Add("referer", "https://live.bilibili.com")
	req.Header.Add("accept", "*/*")
	if headers != nil {
		for key, val := range headers {
			req.Header.Add(key, val)
		}
	}
	//http client
	client := &http.Client{
		Timeout: 30 * time.Second,
	}
	log.Printf("Go %s URL : %s \n", http.MethodGet, req.URL.String())
	return client.Do(req)
}
