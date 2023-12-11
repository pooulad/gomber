package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
)

type Attributes struct {
	Type        string `json:"type"`
	ShowMsg     string `json:"show_msg"`
	Country     string `json:"country"`
	CountryCode string `json:"country_code"`
	Number      string `json:"number"`
	Flag        string `json:"flag"`
}

type FilimoBody struct {
	Type       string     `json:"type"`
	ID         string     `json:"id"`
	Attributes Attributes `json:"attributes"`
}

func newFilimo(mobileNumber int) *FilimoBody {
	return &FilimoBody{
		Type: "country_code",
		ID:   "",
		Attributes: Attributes{
			Type:        "success",
			ShowMsg:     "yes",
			Country:     "iran",
			CountryCode: "98",
			Number:      fmt.Sprintf("(98)%d", mobileNumber),
			Flag:        "https://www.filimo.com/public/public/images/flags/iran.png",
		},
	}
}

func FilimoRequest(mobileNumber int, m map[string]bool, wg *sync.WaitGroup) {
	targetBody := newFilimo(mobileNumber)
	jsonByte, err := json.Marshal(targetBody)
	if err != nil {
		log.Fatal(err)
	}
	
	client := &http.Client{}
	req, err := http.NewRequest("POST", "https://filimo.com/api/fa/v1/user/Authenticate/country_code", bytes.NewBuffer(jsonByte))
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	if resp.StatusCode == 200 {
		m["filimo"] = true
		return
	}
	m["filimo"] = false

	defer wg.Done()
}
