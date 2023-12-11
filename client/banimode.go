package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
)

type BanimodeBody struct {
	Phone string `json:"phone"`
}

func newBanimode(mobileNumber int) *BanimodeBody {
	return &BanimodeBody{
		Phone: fmt.Sprintf("0%d", mobileNumber),
	}
}

func BanimodeRequest(mobileNumber int, m map[string]bool, wg *sync.WaitGroup) {
	defer wg.Done()

	targetBody := newBanimode(mobileNumber)
	jsonByte, err := json.Marshal(targetBody)
	if err != nil {
		fmt.Println(fmt.Errorf(err.Error()))
	}

	resp, err := http.Post("https://mobapi.banimode.com/api/v2/auth/request", "application/json", bytes.NewBuffer([]byte(jsonByte)))
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
		m["banimode"] = true
		return
	}
	m["banimode"] = false

	fmt.Println("banimode")
}
