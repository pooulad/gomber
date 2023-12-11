package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
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
		Username: fmt.Sprintf("0%d", mobileNumber),
	}
}

func DigikalaRequest(mobileNumber int, m map[string]bool, wg *sync.WaitGroup) {
	defer wg.Done()

	targetBody := newDigikala(mobileNumber)
	jsonByte, err := json.Marshal(targetBody)
	if err != nil {
		fmt.Println(fmt.Errorf(err.Error()))
	}

	resp, err := http.Post("https://api.digikala.com/v1/user/authenticate/", "application/json", bytes.NewBuffer([]byte(jsonByte)))
	if err != nil {
		fmt.Println(fmt.Errorf(err.Error()))
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(fmt.Errorf(err.Error()))
	}
	fmt.Println(string(b))

	if resp.StatusCode == 200 {
		m["digikala"] = true
		return
	}
	m["digikala"] = false

	fmt.Println("digikala")
}
