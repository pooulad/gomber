package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"log"
	"net/http"
	"sync"
)

type FilimoBody struct {
	Mobile string `json:"mobile"`
	Guid   string `json:"guid"`
}

func newFilimo(mobileNumber int) *FilimoBody {
	id := uuid.New()
	return &FilimoBody{
		Mobile: fmt.Sprintf("%d", mobileNumber),
		Guid:   id.String(),
	}
}

func FilimoRequest(mobileNumber int, m map[string]bool, wg *sync.WaitGroup) {
	defer wg.Done()

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

	fmt.Println("filimo")
}
