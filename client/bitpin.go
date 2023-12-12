package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
)

type BitpinBody struct {
	Phone string `json:"phone"`
}

func newBitpin(mobileNumber int) *BitpinBody {
	return &BitpinBody{
		Phone: fmt.Sprintf("0%d", mobileNumber),
	}
}

func BitpinRequest(mobileNumber int, m map[string]bool, wg *sync.WaitGroup) {
	defer wg.Done()

	targetBody := newBitpin(mobileNumber)
	jsonByte, err := json.Marshal(targetBody)
	if err != nil {
		fmt.Println(fmt.Errorf(err.Error()))
	}

	resp, err := http.Post("https://api.bitpin.org/v2/usr/signin/", "application/json", bytes.NewBuffer([]byte(jsonByte)))
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
		m["bitpin"] = true
		return
	}
	m["bitpin"] = false

	fmt.Println("bitpin")
}
