package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	url := "http://api.tushare.pro"
	fmt.Println("URL:>", url)
	token := os.Getenv("token")
	data := "{" +
		"	\"api_name\": \"trade_cal\", " +
		"	\"token\": \"" + token + "\"," +
		"	\"params\": { " +
		"		\"exchange\":\"\"," +
		"		\"start_date\": \"20200321\", " +
		"		\"end_date\": \"20200323\", " +
		"		\"is_open\": \"0\" " +
		"		}," +
		"		\"fields\": \"exchange,cal_date,is_open,pretrade_date\" " +
		"}"

	fmt.Println(data)
	var jsonStr = []byte(data)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
}
