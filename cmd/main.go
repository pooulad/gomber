package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/pooulad/gomber/pkg"
)

func main() {
	banner, err := os.ReadFile("./program/banner.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprint(os.Stdout, string(banner)+"\n")

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter mobile target: ")

	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			fmt.Println(fmt.Errorf("please insert mobile number to start"))
			break
		}

		mobileNumber, err := strconv.Atoi(text)
		if err != nil {
			fmt.Println(fmt.Errorf("please insert mobile number to start"))
			break
		}
		pkg.Execute(mobileNumber)

	}

	if err := scanner.Err(); err != nil {
		log.Fatal("Some problem happend in scanner")
	}
}
