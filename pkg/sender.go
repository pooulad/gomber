package pkg

import (
	"fmt"
	// "sync"
)

func SendSms(mobileNumber int, amount int) {
	// wg := &sync.WaitGroup{}
	fmt.Println(amount)
	fmt.Println(mobileNumber)
	fmt.Println("execute program here")
}
