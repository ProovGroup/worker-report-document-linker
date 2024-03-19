package sqs

import "worker-report-document-linker/internal/document"

type Message struct {
	ProovCode string            `json:"proov_code"`
	Document  document.Document `json:"document"`
}

