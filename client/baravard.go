package client

import (
	"fmt"
	"io"
	"net/http"
	"sync"
)

func BaravardRequest(mobileNumber int, m map[string]bool, wg *sync.WaitGroup) {
	defer wg.Done()

	url := fmt.Sprintf("https://baravard.com/Ajax/SendToken?pn=0%d", mobileNumber)

	resp, err := http.Get(url)
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
		m["baravard"] = true
		return
	}
	m["baravard"] = false

	fmt.Println("baravard")
}
