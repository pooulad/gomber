package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
)

type Bit24Body struct {
	Mobile      string `json:"mobile"`
	CountryCode string `json:"country_code"`
}

func newBit24(mobileNumber int) *Bit24Body {
	return &Bit24Body{
		Mobile:      fmt.Sprintf("0%d", mobileNumber),
		CountryCode: "98",
	}
}

func Bit24Request(mobileNumber int, m map[string]bool, wg *sync.WaitGroup) {
	defer wg.Done()

	targetBody := newBit24(mobileNumber)
	jsonByte, err := json.Marshal(targetBody)
	if err != nil {
		fmt.Println(fmt.Errorf(err.Error()))
	}

	resp, err := http.Post("https://bit24.cash/auth/bit24/api/v3/auth/check-mobile", "application/json", bytes.NewBuffer([]byte(jsonByte)))
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
		m["bit24"] = true
		return
	}
	m["bit24"] = false

	fmt.Println("bit24")
}
