package main

import (
	"fmt"

	"github.com/akshith6212/portscanner/portscanner"
)

func main() {
	fmt.Println("Port Scanning")
	results := port.InitialScan("localhost")
	fmt.Println(results)
}
