package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
)

type DigiPayBody struct {
	CellNumber string `json:"cellNumber"`
}

func newDigiPay(mobileNumber int) *DigiPayBody {
	return &DigiPayBody{
		CellNumber: fmt.Sprintf("0%d", mobileNumber),
	}
}

func DigiPayRequest(mobileNumber int, m map[string]bool, wg *sync.WaitGroup) {
	defer wg.Done()

	targetBody := newDigiPay(mobileNumber)
	jsonByte, err := json.Marshal(targetBody)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.Post("https://www.mydigipay.com/digipay/api/users/send-sms", "application/json", bytes.NewBuffer([]byte(jsonByte)))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b))

	if resp.StatusCode == 200 {
		m["digipay"] = true
		return
	}
	m["digipay"] = false

	fmt.Println("digipay")
}
