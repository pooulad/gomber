package pkg

import (
	"fmt"
	"sync"
	"time"

	"github.com/pooulad/gomber/client"
)

func SendSms(mobileNumber int, amount int) {
	wg := &sync.WaitGroup{}

	start := time.Now()
	defer func() {
		fmt.Println("Execution Time: ", time.Since(start))
	}()

	clientsStatus := make(map[string]bool)
	clients := []func(int, map[string]bool, *sync.WaitGroup){
		client.DigikalaRequest,
	}

	wg.Add(len(clients))
	for _, fn := range clients {
		fn(mobileNumber, clientsStatus, wg)
		wg.Done()
	}

	wg.Wait()
	fmt.Println(clientsStatus)
}
