package main

import (
	"log"

	cli "github.com/isaqueveras/jangada/internal"
)

func main() {
	if err := cli.New(); err != nil {
		log.Println(err.Error())
	}
}
