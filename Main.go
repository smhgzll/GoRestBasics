package main

import (
	"GoRestBasics/Endpoints"
	"fmt"
)

func main() {
	fmt.Println("Listening...")
	Endpoints.HandleRequests()
}
