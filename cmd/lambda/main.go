package main

import (
	l "worker-report-document-linker/pkg/lambda"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(l.Handler)
}
