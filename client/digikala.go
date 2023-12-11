package client

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"sync"
)

type OtpBody struct {
	BackUrl  string `json:"backUrl"`
	OtpCall  bool   `json:"otp_call"`
	Username string `json:"username"`
}

func DigikalaRequest(mobileNumber int, m map[string]bool, wg *sync.WaitGroup) {
	data := fmt.Sprintf(`{"backUrl":"/","otp_call" : "false","username":"%d"}`, mobileNumber)
	var jsonStr = []byte(data)

	client := &http.Client{}
	req, err := http.NewRequest("POST", "https://api.digikala.com/v1/user/authenticate/", bytes.NewBuffer(jsonStr))
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	if resp.StatusCode == 200 {
		m["digikala"] = true
		return
	}
	m["digikala"] = false

	defer wg.Done()
}
