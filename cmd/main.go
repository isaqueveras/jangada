package main

import (
	"log"

	"github.com/isaqueveras/jangada/internal/cli"
)

func main() {
	if err := cli.New(); err != nil {
		log.Println(err.Error())
	}
}
