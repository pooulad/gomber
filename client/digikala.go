package client

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
)

type OtpBody struct {
	BackUrl  string `json:"backUrl"`
	OtpCall  bool   `json:"otp_call"`
	Username string `json:"username"`
}

func DigikalaRequest(mobileNumber int, m map[string]bool) {
	data := fmt.Sprintf(`{"backUrl":"/","otp_call" : "false","username":"%d"}`, mobileNumber)
	var jsonStr = []byte(data)

	client := &http.Client{}
	req, err := http.NewRequest("POST", "https://api.digikala.com/v1/user/authenticate/", bytes.NewBuffer(jsonStr))
	if err != nil {
		log.Fatal(err)
	}
	// req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:109.0) Gecko/20100101 Firefox/118.0")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	if resp.StatusCode == 200 {
		m["digikala"] = true
	}
	m["digikala"] = false
}
