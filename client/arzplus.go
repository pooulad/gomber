package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
)

type ArzplusBody struct {
	Phone string `json:"phone"`
}

func newArzplus(mobileNumber int) *ArzplusBody {
	return &ArzplusBody{
		Phone: fmt.Sprintf("0%d", mobileNumber),
	}
}

func ArzplusRequest(mobileNumber int, m map[string]bool, wg *sync.WaitGroup) {
	defer wg.Done()

	targetBody := newArzplus(mobileNumber)
	jsonByte, err := json.Marshal(targetBody)
	if err != nil {
		fmt.Println(fmt.Errorf(err.Error()))
	}

	resp, err := http.Post("https://api.arzplus.net/api/v1/accounts/signup/init/", "application/json", bytes.NewBuffer([]byte(jsonByte)))
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
		m["arzplus"] = true
		return
	}
	m["arzplus"] = false

	fmt.Println("arzplus")
}
