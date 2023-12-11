package client

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
)

func Gishe7Request(mobileNumber int, m map[string]bool, wg *sync.WaitGroup) {
	defer wg.Done()

	resp, err := http.Post(fmt.Sprintf("https://gisheh7.ir:8080/api/user/send-token?mobile=0%d", mobileNumber), "application/json", nil)
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
		m["gishe7"] = true
		return
	}
	m["gishe7"] = false

	fmt.Println("gishe7")
}
