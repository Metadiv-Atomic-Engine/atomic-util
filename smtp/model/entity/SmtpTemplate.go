package entity

import "github.com/Metadiv-Atomic-Engine/atomic/base"

type SmtpTemplate struct {
	base.Model
	base.ModelWorkspace

	Type string `json:"type"` // text or html
	Name []byte `json:"name"`

	SubjectEn  []byte `json:"subject_en"`
	SubjectZht []byte `json:"subject_zht"`
	SubjectZhs []byte `json:"subject_zhs"`

	ContentEn  []byte `json:"content_en"`
	ContentZht []byte `json:"content_zht"`
	ContentZhs []byte `json:"content_zhs"`
}
