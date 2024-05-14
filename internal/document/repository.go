package document

import (
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
	_, err := r.db.Exec(
		`INSERT INTO rptdocuments (name, type, proov_code, url) VALUES ($1, $2, $3, $4)`,
		r.Document.Name, r.Document.Type, proovCode, link,
	)
	return err
}
