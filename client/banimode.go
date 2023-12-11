package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
)

type BanimodeBody struct {
	Phone string `json:"phone"`
}

func newBanimode(mobileNumber int) *BanimodeBody {
	return &BanimodeBody{
		Phone: fmt.Sprintf("%d", mobileNumber),
	}
}

func BanimodeRequest(mobileNumber int, m map[string]bool, wg *sync.WaitGroup) {
	defer wg.Done()

	targetBody := newBanimode(mobileNumber)
	jsonByte, err := json.Marshal(targetBody)
	if err != nil {
		log.Fatal(err)
	}

	client := &http.Client{}
	req, err := http.NewRequest("POST", "https://mobapi.banimode.com/api/v2/auth/request", bytes.NewBuffer(jsonByte))
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	if resp.StatusCode == 200 {
		m["banimode"] = true
		return
	}
	m["banimode"] = false

	fmt.Println("banimode")
}
