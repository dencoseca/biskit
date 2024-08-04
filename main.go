package main

import (
	"fmt"
	"log"
	"os"
)

const cookiesFile = "./.idea/httpRequests/http-client.cookies"

func main() {
	cookies, err := os.ReadFile(cookiesFile)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(cookies))
}
