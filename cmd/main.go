package main

import (
	"fmt"
	"github.com/ildomm/aws_experiment/util"
	"log"
	"os"
)

const Application = "aws_experiment"

func main() {

	// Configure log standards
	log.Default().SetOutput(os.Stdout)
	log.Default().SetFlags(log.Llongfile)
	log.SetPrefix(fmt.Sprintf("%s:", Application))

	// Configure timezone
	err := util.SetGlobalTimezoneAsUTC()
	if err != nil {
		log.Fatal(err)
	}

	// DEBUG
	log.Print("Running")
	// DEBUG
}
