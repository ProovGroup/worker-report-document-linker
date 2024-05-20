package webhooksNotifier

import (
	"fmt"
	"worker-report-document-linker/internal/provider"

	"github.com/ProovGroup/lib-core-models-report"
	webhook "github.com/ProovGroup/lib-core-models-webhook"
	env "github.com/ProovGroup/lib-env"
)

const QUEUE_NAME = "webhooks-notifier"

type MessageEvent struct {
	ProovCode string            `json:"proov_code"`
	Owner     int               `json:"owner"`
	State     string            `json:"state"`
	Event     webhook.EventType `json:"event"`
	WebHookID int               `json:"webhook_id,omitempty"`
}

func Send(proovCode string) error {
	e := provider.GetEnv()
	wh, isExist := e.GetQueue(env.NewQueueSelector(QUEUE_NAME))
	if isExist == false {
		return fmt.Errorf("queue not found (%s)\n", QUEUE_NAME)
	}

	db := provider.GetDB()
	r, err := report.GetReport(db, proovCode)
	if err != nil {
		return fmt.Errorf("report.GetReport(db, proovCode): %s", err)
	}

	msgEvent := MessageEvent{
		ProovCode: r.ProovCode,
		Owner:     r.Owner,
		State:     r.State,
		Event:     webhook.MatrixCreated,
	}

	message := wh.NewMessage()
	message.WriteJSON(msgEvent)

	return message.Send()
}
