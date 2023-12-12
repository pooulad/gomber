package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
)

type DivarBody struct {
	Phone string `json:"phone"`
}

func newDivar(mobileNumber int) *DivarBody {
	return &DivarBody{
		Phone: fmt.Sprintf("0%d", mobileNumber),
	}
}

func DivarRequest(mobileNumber int, m map[string]bool, wg *sync.WaitGroup) {
	defer wg.Done()

	targetBody := newDivar(mobileNumber)
	jsonByte, err := json.Marshal(targetBody)
	if err != nil {
		fmt.Println(fmt.Errorf(err.Error()))
	}

	resp, err := http.Post("https://api.divar.ir/v5/auth/authenticate", "application/json", bytes.NewBuffer([]byte(jsonByte)))
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
		m["divar"] = true
		return
	}
	m["divar"] = false

	fmt.Println("divar")
}
