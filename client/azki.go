package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
)

type AzkiBody struct {
	PhoneNumber string `json:"phoneNumber"`
}

func newAzki(mobileNumber int) *AzkiBody {
	return &AzkiBody{
		PhoneNumber: fmt.Sprintf("0%d", mobileNumber),
	}
}

func AzkiRequest(mobileNumber int, m map[string]bool, wg *sync.WaitGroup) {
	defer wg.Done()

	targetBody := newAzki(mobileNumber)
	jsonByte, err := json.Marshal(targetBody)
	if err != nil {
		fmt.Println(fmt.Errorf(err.Error()))
	}

	resp, err := http.Post("https://www.azki.com/api/vehicleorder/v2/app/auth/check-login-availability/", "application/json", bytes.NewBuffer([]byte(jsonByte)))
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
		m["azki"] = true
		return
	}
	m["azki"] = false

	fmt.Println("azki")
}
