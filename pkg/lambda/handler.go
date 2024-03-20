package lambda

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"worker-report-document-linker/internal/db"
	"worker-report-document-linker/internal/document"

	"github.com/ProovGroup/env"
	"github.com/aws/aws-lambda-go/events"
)

func Handler(ctx context.Context, event events.SQSEvent) error {
	for i, sqsRecord := range event.Records {
		e, err := env.GetEnvSqsArnSSM(event.Records[i].EventSourceARN, os.Getenv("AWS_REGION"), env.BDDWrite)
		if err != nil {
			fmt.Println("[ERROR] env.GetEnvSqsArnSSM:", err)
			return err
		}

		s3Event := events.S3Event{}

		err = json.Unmarshal([]byte(sqsRecord.Body), &s3Event)
		if err != nil {
			fmt.Println("[Error] Unmarshal record event", err)
			return err
		}

		for _, s3Record := range s3Event.Records {
			s3Region := s3Record.AWSRegion
			s3Bucket := s3Record.S3.Bucket.Name
			s3Key := s3Record.S3.Object.Key

			parts := strings.Split(s3Key, "/")
			if len(parts) < 2 {
				return fmt.Errorf("[ERROR] Invalid key %s. The key must be in this format [...]/{proov_code}/{file}.{ext}.\n", s3Key)
			}

			proovCode := parts[len(parts)-2]
			file := strings.Split(parts[len(parts)-1], ".")

			document := document.Document {
				Name: file[0],
				Type: file[1],
				Path: document.Path {
					Region: s3Region,
					Bucket: s3Bucket,
					Key:    s3Key,
				},
			}

			fmt.Println("[INFO] ProovCode:", proovCode)
			fmt.Println("[INFO] Document:", document)

			// Get permalink
			link := document.GetPermalink()
			if link == "" {
				return fmt.Errorf("[ERROR] link is empty")
			}

			// Save document
			err = db.SaveDocument(&e, proovCode, &document, link)
			if err != nil {
				fmt.Println("[ERROR] db.SaveDocument:", err)
				return err
			}

			fmt.Println("[INFO] Document linked successfully")
		}
	}

	return nil
}


