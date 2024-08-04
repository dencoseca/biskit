package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const cookiesFile = "./.idea/httpRequests/http-client.cookies"

type Cookie struct {
	Domain string
	Path   string
	Name   string
	Value  string
	Date   string
}

type Application struct {
	cookies []Cookie
}

func (a *Application) String() string {
	var result string
	for i, cookie := range a.cookies {
		result += fmt.Sprintf("Cookie %d: %+v\n", i+1, cookie)
	}
	return result
}

func main() {
	app := &Application{}

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

	for scanner.Scan() {
		for scanner.Scan() {
			line := scanner.Text()
			parts := strings.Fields(line)

			app.cookies = append(app.cookies, Cookie{
				Domain: parts[0],
				Path:   parts[1],
				Name:   parts[2],
				Value:  parts[3],
				Date:   strings.Join(parts[4:], " "),
			})
		}
	}

	if scanner.Err() != nil {
		log.Fatal(err)
	}

	fmt.Println(app)
}
