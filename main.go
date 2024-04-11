package main

import (
	"log"

	"github.com/AlperRehaYAZGAN/fezz/pkg"
)

func main() {
	// log
	log.Println("Started Fezz on port 7071")
	// init app
	if err := pkg.Init(); err != nil {
		panic(err)
	}
}
