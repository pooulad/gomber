package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
)

type DigikalaBody struct {
	BackUrl  string `json:"backUrl"`
	OtpCall  bool   `json:"otp_call"`
	Username string `json:"username"`
}

func newDigikala(mobileNumber int) *DigikalaBody {
	return &DigikalaBody{
		BackUrl:  "/",
		OtpCall:  false,
		Username: fmt.Sprintf("%d", mobileNumber),
	}
}

func DigikalaRequest(mobileNumber int, m map[string]bool, wg *sync.WaitGroup) {
	defer wg.Done()
	
	targetBody := newDigikala(mobileNumber)
	jsonByte, err := json.Marshal(targetBody)
	if err != nil {
		log.Fatal(err)
	}
	
	client := &http.Client{}
	req, err := http.NewRequest("POST", "https://api.digikala.com/v1/user/authenticate/", bytes.NewBuffer(jsonByte))
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

	fmt.Println("digikala")
}
