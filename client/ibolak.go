package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
)

type IbolakBody struct {
	Mobile string `json:"mobile"`
}

func newIbolak(mobileNumber int) *IbolakBody {
	return &IbolakBody{
		Mobile: fmt.Sprintf("0%d", mobileNumber),
	}
}

func IbolakRequest(mobileNumber int, m map[string]bool, wg *sync.WaitGroup) {
	defer wg.Done()

	targetBody := newIbolak(mobileNumber)
	jsonByte, err := json.Marshal(targetBody)
	if err != nil {
		fmt.Println(fmt.Errorf(err.Error()))
	}

	resp, err := http.Post("https://ibolak.com/api/v1/auth/send-login-code", "application/json", bytes.NewBuffer([]byte(jsonByte)))
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
		m["ibolak"] = true
		return
	} else if resp.StatusCode == 403 {
		resp, err := http.Post("https://ibolak.com/api/v1/auth/send-register-code", "application/json", bytes.NewBuffer([]byte(jsonByte)))
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
			m["ibolak"] = true
			return
		}
		m["ibolak"] = false
	}
	m["ibolak"] = false

	fmt.Println("ibolak")
}
