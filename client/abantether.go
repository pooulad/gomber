package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
)

type AbanththerBody struct {
	PhoneNumber string `json:"phoneNumber"`
}

func newAbanthther(mobileNumber int) *AbanththerBody {
	return &AbanththerBody{
		PhoneNumber: fmt.Sprintf("0%d", mobileNumber),
	}
}

func AbanththerRequest(mobileNumber int, m map[string]bool, wg *sync.WaitGroup) {
	defer wg.Done()

	targetBody := newAbanthther(mobileNumber)
	jsonByte, err := json.Marshal(targetBody)
	if err != nil {
		fmt.Println(fmt.Errorf(err.Error()))
	}

	resp, err := http.Post("https://abantether.com/users/reset-password/phone/send/", "application/json", bytes.NewBuffer([]byte(jsonByte)))
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
		m["arzplus"] = true
		return
	}
	m["arzplus"] = false

	fmt.Println("arzplus")
}
