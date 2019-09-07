package main

import (
	"log"

	sis "github.com/zhuharev/simple-image-server"
)

func main() {
	server, err := sis.New()
	if err != nil {
		log.Fatal(err)
	}
	server.Run()
}
