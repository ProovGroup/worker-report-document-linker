package document

import "worker-report-document-linker/internal/permalink"

type Document struct {
	Name string `json:"name"`
	Type string `json:"type"`
	Path Path   `json:"path"`
}

type Path struct {
	Region string `json:"region"`
	Bucket string `json:"bucket"`
	Key    string `json:"key"`
}

func (d Document) GetPermalink() string {
	return permalink.GetPermalink(d.Path.Region, d.Path.Bucket, d.Path.Key)
}
