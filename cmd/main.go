package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/ildomm/aws_experiment/handler"
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

	// Load AWS configuration
	ctx := context.Background()
	configuration, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// Start the Lambda handler
	builder, err := handler.Build(&configuration)
	if err != nil {
		log.Fatal(err)
	}

	log.Print("Starting handler")
	lambda.Start(builder.Handle)
}
