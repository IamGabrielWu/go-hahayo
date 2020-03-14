package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestHttp_test(t *testing.T) {
	resp, err := http.Get("https://httpbin.org/get")
	if err != nil {
		t.Error(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		t.Error(err)
	}

	t.Log(string(body))

	requestBody, err := json.Marshal(map[string]string{
		"name":  "dz.wu",
		"email": "dz.wu@aftership.com",
	})

	if err != nil {
		t.Error(err)
	}

	resp, err = http.Post("https://httpbin.org/post", "application/json", bytes.NewBuffer(requestBody))

	if err != nil {
		t.Error(err)
	}

	defer resp.Body.Close()

	body, err = ioutil.ReadAll(resp.Body)

	if err != nil {
		t.Error(err)
	}

	t.Log(string(body))
}
