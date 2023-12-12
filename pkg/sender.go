package pkg

import (
	"fmt"
	"sync"
	"time"

	"github.com/pooulad/gomber/client"
)

func SendSms(mobileNumber int, delay int) {
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
		client.Gishe7Request,
		client.PindoRequest,
		client.DivarRequest,
		client.BasalamRequest,
	}

	wg.Add(len(clients))
	for _, fn := range clients {
		time.Sleep(time.Second * time.Duration(delay))
		fn(mobileNumber, clientsStatus, wg)
	}

	wg.Wait()
	fmt.Println(clientsStatus)
}
