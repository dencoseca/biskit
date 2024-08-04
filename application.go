package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

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

const cookiesFile = "./.idea/httpRequests/http-client.cookies"

func (a *Application) loadCookies(app *Application) (err error) {
	file, err := os.Open(cookiesFile)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		closeErr := file.Close()
		if closeErr != nil {
			err = closeErr
		}
	}(file)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		for scanner.Scan() {
			parts := strings.Fields(scanner.Text())

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
		return err
	}

	return nil
}
