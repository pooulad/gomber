package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
)

type Khodro45Body struct {
	Mobile string `json:"mobile"`
}

func newKhodro45(mobileNumber int) *Khodro45Body {
	return &Khodro45Body{
		Mobile: fmt.Sprintf("0%d", mobileNumber),
	}
}

func Khodro45Request(mobileNumber int, m map[string]bool, wg *sync.WaitGroup) {
	defer wg.Done()

	targetBody := newKhodro45(mobileNumber)
	jsonByte, err := json.Marshal(targetBody)
	if err != nil {
		fmt.Println(fmt.Errorf(err.Error()))
	}

	resp, err := http.Post("https://khodro45.com/api/v1/customers/otp/", "application/json", bytes.NewBuffer([]byte(jsonByte)))
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
		m["khodro45"] = true
		return
	}
	m["khodro45"] = false

	fmt.Println("khodro45")
}
