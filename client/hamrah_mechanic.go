package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
)

type HamrahMechanicBody struct {
	PhoneNumber string `json:"PhoneNumber"`
}

func newHamrahMechanic(mobileNumber int) *HamrahMechanicBody {
	return &HamrahMechanicBody{
		PhoneNumber: fmt.Sprintf("0%d", mobileNumber),
	}
}

func HamrahMechanicRequest(mobileNumber int, m map[string]bool, wg *sync.WaitGroup) {
	defer wg.Done()

	targetBody := newHamrahMechanic(mobileNumber)
	jsonByte, err := json.Marshal(targetBody)
	if err != nil {
		fmt.Println(fmt.Errorf(err.Error()))
	}

	resp, err := http.Post("https://www.hamrah-mechanic.com/api/v1/membership/otp", "application/json", bytes.NewBuffer([]byte(jsonByte)))
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
		m["hamrah-mechanic"] = true
		return 
	}
	m["hamrah-mechanic"] = false

	fmt.Println("hamrah-mechanic")
}
