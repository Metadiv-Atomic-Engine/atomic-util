package request

type SmtpTemplateCreate struct {
	Type string `json:"type"`
	Name string `json:"name"`

	SubjectEn  string `json:"subject_en"`
	SubjectZht string `json:"subject_zht"`
	SubjectZhs string `json:"subject_zhs"`

	ContentEn  string `json:"content_en"`
	ContentZht string `json:"content_zht"`
	ContentZhs string `json:"content_zhs"`
}
