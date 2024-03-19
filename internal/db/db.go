package db

import (
	"worker-report-document-linker/internal/document"

	"github.com/ProovGroup/env"
)

func SaveDocument(e *env.Env, proovCode string, document *document.Document, link string) error {
	_, err := e.Exec(
		`INSERT INTO rptdocuments (name, type, proov_code, url) VALUES ($1, $2, $3, $4)`, 
		document.Name, document.Type, proovCode, link,
	)
	if err != nil {
		return err
	}
	return nil	
}
