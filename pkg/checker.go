package pkg

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func Execute() {
	banner, err := os.ReadFile("./program/banner.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprint(os.Stdout, string(banner)+"\n")

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Usage : ")
	fmt.Println("Enter mobile target: [target number]")
	fmt.Println("Enter number of requests: [number of requests]")

	fmt.Println("Sample : ")
	fmt.Println("Enter mobile target: 09191234567")
	fmt.Println("Enter number of requests: 20")

	fmt.Println("Type start and enter")

	inputs := []int{}
loop:
	for scanner.Scan() {
		fmt.Println(len(inputs))
		if len(inputs) >= 2 {
			break
		}
		if len(inputs) == 0 {
			fmt.Print("Enter mobile target: ")
		} else {
			fmt.Print("Enter number of requests: ")
		}
		text := scanner.Text()
		if text == "" {
			fmt.Println(fmt.Errorf("please insert somthing"))
			break
		}

		input, err := strconv.Atoi(text)
		if err != nil {
			if len(inputs) == 0 {
				fmt.Println(fmt.Errorf("please insert mobile number to start"))
			} else {
				fmt.Println(fmt.Errorf("please insert number of requests to start"))
			}
			break
		}

		inputs = append(inputs, input)
		goto loop
	}

	fmt.Println(inputs)
	// SendSms(mobileNumber, amount)

	if err := scanner.Err(); err != nil {
		log.Fatal("Some problem happend in scanner")
	}
}
