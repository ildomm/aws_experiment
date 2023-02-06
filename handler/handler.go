package handler

import (
	"context"
	"errors"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/aws"
	"log"
)

type handler struct {
	AwsConfig *aws.Config
}

func Build(config *aws.Config) (*handler, error) {
	if config == nil {
		return nil, errors.New("AWS configuration cant be nil")
	}

	return &handler{AwsConfig: config}, nil
}

func (hd *handler) Handle(ctx context.Context, event events.S3Event) (events.S3Event, error) {
	response := &events.S3Event{}

	// TODO: Implement parser
	log.Printf("Processing event: %v", event)

	return *response, nil
}