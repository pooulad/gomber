package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
)

type PindoBody struct {
	Phone string `json:"phone"`
}

func newPindo(mobileNumber int) *PindoBody {
	return &PindoBody{
		Phone: fmt.Sprintf("0%d", mobileNumber),
	}
}

func PindoRequest(mobileNumber int, m map[string]bool, wg *sync.WaitGroup) {
	defer wg.Done()

	targetBody := newPindo(mobileNumber)
	jsonByte, err := json.Marshal(targetBody)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.Post("https://api.pindo.ir/v1/user/login-register/", "application/json", bytes.NewBuffer([]byte(jsonByte)))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b))

	if resp.StatusCode == 200 {
		m["pindo"] = true
		return
	}
	m["pindo"] = false

	fmt.Println("pindo")
}
