package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
)

type PoolenoBody struct {
	Mobile string `json:"mobile"`
}

func newPooleno(mobileNumber int) *PoolenoBody {
	return &PoolenoBody{
		Mobile: fmt.Sprintf("0%d", mobileNumber),
	}
}

func PoolenoRequest(mobileNumber int, m map[string]bool, wg *sync.WaitGroup) {
	defer wg.Done()

	targetBody := newPooleno(mobileNumber)
	jsonByte, err := json.Marshal(targetBody)
	if err != nil {
		fmt.Println(fmt.Errorf(err.Error()))
	}

	resp, err := http.Post("https://api.pooleno.ir/v1/auth/check-mobile", "application/json", bytes.NewBuffer([]byte(jsonByte)))
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
		m["pooleno"] = true
		return
	}
	m["pooleno"] = false

	fmt.Println("pooleno")
}
