package pkg

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/pooulad/gomber/util"
)

func Execute() {
	banner, err := os.ReadFile("./program/banner.txt")
	if err != nil {
		log.Fatal(err)
	}
	usage, err := os.ReadFile("./program/usage.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprint(os.Stdout, util.Colorize(util.ColorGreen, string(banner))+"\n")
	fmt.Fprint(os.Stdout, string(usage)+"\n")

	scanner := bufio.NewScanner(os.Stdin)
	inputs := []int{}
	var count int
	fmt.Print("Enter mobile target: ")
loop:
	for scanner.Scan() {
		if count == 0 {
			fmt.Print("Enter number of requests: ")
		}

		text := scanner.Text()
		if text == "" {
			fmt.Println(fmt.Errorf("please insert somthing"))
			break
		}

		input, err := strconv.Atoi(text)
		if err != nil {
			if count == 0 {
				fmt.Println(fmt.Errorf("please insert mobile number to start"))
			} else {
				fmt.Println(fmt.Errorf("please insert number of requests to start"))
			}
			break
		}

		inputs = append(inputs, input)

		if count >= 1 {
			break
		}
		count++
		goto loop
	}

	if len(inputs) < 2 {
		log.Fatal("Please insert the inputs correctly")
	}
	SendSms(inputs[0], inputs[1])

	if err := scanner.Err(); err != nil {
		log.Fatal("Some problem happend in scanner")
	}
}
