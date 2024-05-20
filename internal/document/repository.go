package document

import (
	"fmt"
	"worker-report-document-linker/internal/provider"

	"github.com/ProovGroup/lib-env/database"
)

type Repository struct {
	db       database.Database
	Document *Document
}

func NewRepository(doc *Document) *Repository {
	return &Repository{
		db:       provider.GetDB(),
		Document: doc,
	}
}

func (r *Repository) Save(proovCode string, link string) error {
	var count int
	r.db.QueryRow(`SELECT COUNT(1) FROM rptdocuments WHERE proov_code = $1 AND url = $2`, proovCode, link).Scan(&count)
	if count > 0 {
		fmt.Println("[WARN] This document has already been linked to this proov_code", proovCode, link)
		return nil
	}
	_, err := r.db.Exec(
		`INSERT INTO rptdocuments (name, type, proov_code, url) VALUES ($1, $2, $3, $4)`,
		r.Document.Name, r.Document.Type, proovCode, link,
	)
	if err != nil {
		return err
	}
	fmt.Println("[INFO] Document linked successfully")
	return nil
}
