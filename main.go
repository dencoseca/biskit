package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const cookiesFile = "./.idea/httpRequests/http-client.cookies"

func main() {
	file, err := os.Open(cookiesFile)
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	scanner := bufio.NewScanner(file)
	scanner.Scan()

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if scanner.Err() != nil {
		log.Fatal(err)
	}
}
