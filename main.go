package main

import (
	"fmt"
	"log"
)

func main() {
	app := &Application{}

	err := app.loadCookies(app)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(app)
}
