package main

import (
	"github.com/jchavannes/iiproject/app/cmd"
	"log"
)

func main() {
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}
