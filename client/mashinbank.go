package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
)

type MashinbankBody struct {
	MobileNumber string `json:"mobileNumber"`
}

func newMashinbank(mobileNumber int) *MashinbankBody {
	return &MashinbankBody{
		MobileNumber: fmt.Sprintf("0%d", mobileNumber),
	}
}

func MashinbankRequest(mobileNumber int, m map[string]bool, wg *sync.WaitGroup) {
	defer wg.Done()

	targetBody := newMashinbank(mobileNumber)
	jsonByte, err := json.Marshal(targetBody)
	if err != nil {
		fmt.Println(fmt.Errorf(err.Error()))
	}

	resp, err := http.Post("https://mashinbank.com/api2/users/check", "application/json", bytes.NewBuffer([]byte(jsonByte)))
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
		m["mashinbank"] = true
		return
	}
	m["mashinbank"] = false

	fmt.Println("mashinbank")
}
