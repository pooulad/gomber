package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
)

type SheypoorBody struct {
	Username string `json:"username"`
}

func newSheypoor(mobileNumber int) *SheypoorBody {
	return &SheypoorBody{
		Username: fmt.Sprintf("0%d", mobileNumber),
	}
}

func SheypoorRequest(mobileNumber int, m map[string]bool, wg *sync.WaitGroup) {
	defer wg.Done()

	targetBody := newSheypoor(mobileNumber)
	jsonByte, err := json.Marshal(targetBody)
	if err != nil {
		fmt.Println(fmt.Errorf(err.Error()))
	}

	resp, err := http.Post("https://www.sheypoor.com/api/v10.0.0/auth/send", "application/json", bytes.NewBuffer([]byte(jsonByte)))
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
		m["sheypoor"] = true
		return
	}
	m["sheypoor"] = false

	fmt.Println("sheypoor")
}
