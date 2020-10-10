package main

import (
	"GoBasics/Endpoints"
	"fmt"
)

func main() {
	fmt.Println("Listening...")
	Endpoints.HandleRequests()
}