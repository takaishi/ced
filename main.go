package main

import (
	"log"
	"os"
)

func main() {
	if err := RunCLI(os.Args[1:]); err != nil {
		log.Fatal(err)
	}
}
