package lambda

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"worker-report-document-linker/internal/document"
	webhooksNotifier "worker-report-document-linker/internal/sqs/webhooks-notifier"

	webhook "github.com/ProovGroup/lib-core-models-webhook"
	"github.com/aws/aws-lambda-go/events"
)

type MessageEvent struct {
	ProovCode string            `json:"proov_code"`
	Owner     int               `json:"owner"`
	State     string            `json:"state"`
	Event     webhook.EventType `json:"event"`
	WebHookID int               `json:"webhook_id,omitempty"`
}

func Handler(ctx context.Context, event events.SQSEvent) error {
	for _, sqsRecord := range event.Records {
		s3Event := events.S3Event{}

		err := json.Unmarshal([]byte(sqsRecord.Body), &s3Event)
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

			docToLink := document.Document{
				Name: file[0],
				Type: file[1],
				Path: document.Path{
					Region: s3Region,
					Bucket: s3Bucket,
					Key:    s3Key,
				},
			}

			fmt.Println("[INFO] ProovCode:", proovCode)
			fmt.Println("[INFO] Document:", docToLink)

			// Get permalink
			link := docToLink.GetPermalink()
			if link == "" {
				return fmt.Errorf("[ERROR] link is empty")
			}

			// Save document
			err := document.NewRepository(&docToLink).Save(proovCode, link)
			if err != nil {
				fmt.Println("[ERROR] NewRepository(&document).Save:", err)
				return err
			}

			fmt.Println("[INFO] Document linked successfully")

			// Send an event to webhooks-notifier only if the document linked is a matrix report
			if strings.HasSuffix(s3Key, "matrix.pdf") {
				webhooksNotifier.Send(proovCode)
				fmt.Println("[INFO] Webhook message sent to webhooks-notifier queue")
			}
		}
	}

	return nil
}
