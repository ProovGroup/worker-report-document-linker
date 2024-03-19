package lambda

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"worker-report-document-linker/internal/db"
	"worker-report-document-linker/internal/sqs"

	"github.com/ProovGroup/env"
	"github.com/aws/aws-lambda-go/events"
)

func Handler(ctx context.Context, event events.SQSEvent) error {
	for i := range event.Records {
		e, err := env.GetEnvSqsArnSSM(event.Records[i].EventSourceARN, os.Getenv("AWS_REGION"), env.BDDWrite)
		if err != nil {
			fmt.Println("[ERROR] env.GetEnvSqsArnSSM:", err)
			return err
		}

		var message sqs.Message
		if err = json.Unmarshal([]byte(event.Records[i].Body), &message); err != nil {
			fmt.Println("[ERROR] json.Unmarshal:", err)
			return err
		}

		fmt.Println("[INFO] ProovCode:", message.ProovCode)
		fmt.Println("[INFO] Document:", message.Document)

		// Get permalink
		link := message.Document.GetPermalink()
		if link == "" {
			return fmt.Errorf("[ERROR] link is empty")
		}

		// Save document
		err = db.SaveDocument(&e, message.ProovCode, &message.Document, link)
		if err != nil {
			fmt.Println("[ERROR] db.SaveDocument:", err)
			return err
		}

		fmt.Println("[INFO] Document linked successfully")
	}

	return nil
}


