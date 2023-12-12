package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
)

type KarnamehBody struct {
	Phone string `json:"phone"`
}

func newKarnameh(mobileNumber int) *KarnamehBody {
	return &KarnamehBody{
		Phone: fmt.Sprintf("0%d", mobileNumber),
	}
}

func KarnamehRequest(mobileNumber int, m map[string]bool, wg *sync.WaitGroup) {
	defer wg.Done()

	targetBody := newKarnameh(mobileNumber)
	jsonByte, err := json.Marshal(targetBody)
	if err != nil {
		fmt.Println(fmt.Errorf(err.Error()))
	}

	resp, err := http.Post("https://api.karnameh.com/v1/carinspection/auth/authenticate", "application/json", bytes.NewBuffer([]byte(jsonByte)))
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
		m["karnameh"] = true
		return
	}
	m["karnameh"] = false

	fmt.Println("karnameh")
}
