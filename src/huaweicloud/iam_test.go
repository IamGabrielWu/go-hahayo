package main

import (
	"io/ioutil"
	"net/http"
	"testing"
)

func TestGetAuthToken(t *testing.T) {

	project_id := "06d31f95838026432f6fc0123466feeb"
	endpoint := "ecs.cn-east-3.myhuaweicloud.com"

	client := &http.Client{}
	req, _ := http.NewRequest("GET", "https://iam.af-south-1.myhuaweicloud.com/v3/users", nil)
	req.Header.Set("X-Auth-Token", "value")
	req.Header.Set("Content-Type", "application/json;charset=utf8")
	resp, err := client.Do(req)

	if err != nil {
		t.Error(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		t.Error(err)
	}

	t.Log(string(body))
}
