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

type SnappFoodBody struct {
	CellPhone string `json:"cellphone"`
}

func newSnappFood(mobileNumber int) *SnappFoodBody {
	return &SnappFoodBody{
		CellPhone: fmt.Sprintf("0%d", mobileNumber),
	}
}

func SnappFoodRequest(mobileNumber int, m map[string]bool, wg *sync.WaitGroup) {
	defer wg.Done()

	targetBody := newSnappFood(mobileNumber)
	jsonByte, err := json.Marshal(targetBody)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(jsonByte))

	resp, err := http.Post("https://snappfood.ir/mobile/v2/user/loginMobileWithNoPass?lat=35.774&long=51.418&optionalClient=WEBSITE&client=WEBSITE&deviceType=WEBSITE&appVersion=8.1.1&UDID=2e10eb3c-5df3-4a6c-ba16-398448e4fd76&locale=fa", "application/json", bytes.NewBuffer([]byte(jsonByte)))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b))

	if resp.StatusCode == 200 {
		m["snappfood"] = true
		return
	}
	m["snappfood"] = false

	fmt.Println("snappfood")
}
