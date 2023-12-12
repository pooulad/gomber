package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
)

type GanjeBody struct {
	PhoneNumber string `json:"phone_number"`
}

func newGanje(mobileNumber int) *GanjeBody {
	return &GanjeBody{
		PhoneNumber: fmt.Sprintf("+98%d", mobileNumber),
	}
}

func GanjeRequest(mobileNumber int, m map[string]bool, wg *sync.WaitGroup) {
	defer wg.Done()

	targetBody := newGanje(mobileNumber)
	jsonByte, err := json.Marshal(targetBody)
	if err != nil {
		fmt.Println(fmt.Errorf(err.Error()))
	}

	resp, err := http.Post("https://api.ganje.net/v1/account/end-user-before-login/", "application/json", bytes.NewBuffer([]byte(jsonByte)))
	if err != nil {
		fmt.Println(fmt.Errorf(err.Error()))
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(fmt.Errorf(err.Error()))
	}
	fmt.Println(string(b))

	if resp.StatusCode == 200 || resp.StatusCode == 201 {
		m["ganje"] = true
		return
	}
	m["ganje"] = false

	fmt.Println("ganje")
}
