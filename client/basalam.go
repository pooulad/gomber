package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
)

type BasalamBody struct {
	Mobile string `json:"mobile"`
}

func newBasalam(mobileNumber int) *BasalamBody {
	return &BasalamBody{
		Mobile: fmt.Sprintf("0%d", mobileNumber),
	}
}

func BasalamRequest(mobileNumber int, m map[string]bool, wg *sync.WaitGroup) {
	defer wg.Done()

	targetBody := newBasalam(mobileNumber)
	jsonByte, err := json.Marshal(targetBody)
	if err != nil {
		fmt.Println(fmt.Errorf(err.Error()))
	}

	resp, err := http.Post("https://auth.basalam.com/otp-request", "application/json", bytes.NewBuffer([]byte(jsonByte)))
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
		m["basalam"] = true
		return
	}
	m["basalam"] = false

	fmt.Println("basalam")
}
