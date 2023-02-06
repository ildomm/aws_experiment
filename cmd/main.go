package main

import (
	"fmt"
	"log"
)

const Application = "aws_experiment"

func main() {
	log.SetPrefix(fmt.Sprintf("%s:", Application))


	// DEBUG
	log.Print("Running")
	// DEBUG

}
