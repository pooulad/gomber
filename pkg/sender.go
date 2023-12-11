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
		client.BanimodeRequest,
		client.SnappFoodRequest,
		client.DigiPayRequest,
	}

	wg.Add(len(clients))
	for _, fn := range clients {
		time.Sleep(time.Second * 1)
		fn(mobileNumber, clientsStatus, wg)
	}

	wg.Wait()
	fmt.Println(clientsStatus)
}
