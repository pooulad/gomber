package client

import (
	"fmt"
	"io"
	"net/http"
	"sync"
)

func TorobRequest(mobileNumber int, m map[string]bool, wg *sync.WaitGroup) {
	defer wg.Done()

	url := fmt.Sprintf("https://api.torob.com/v4/user/phone/send-pin/?phone_number=0%d", mobileNumber)

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
		m["torob"] = true
		return
	}
	m["torob"] = false

	fmt.Println("torob")
}
