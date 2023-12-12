package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
)

type IdpayBody struct {
	Username string `json:"username"`
}

func newIdpay(mobileNumber int) *IdpayBody {
	return &IdpayBody{
		Username: fmt.Sprintf("0%d", mobileNumber),
	}
}

func IdpayRequest(mobileNumber int, m map[string]bool, wg *sync.WaitGroup) {
	defer wg.Done()

	targetBody := newIdpay(mobileNumber)
	jsonByte, err := json.Marshal(targetBody)
	if err != nil {
		fmt.Println(fmt.Errorf(err.Error()))
	}

	resp, err := http.Post("https://panel.idpay.ir/api/v1/user/authenticate", "application/json", bytes.NewBuffer([]byte(jsonByte)))
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
		m["idpay"] = true
		return
	}
	m["idpay"] = false

	fmt.Println("idpay")
}
