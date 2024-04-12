package dto

import "strings"

type SmtpTemplate struct {
	ID uint `json:"id"`

	SubjectEn  string `json:"subject_en"`
	SubjectZht string `json:"subject_zht"`
	SubjectZhs string `json:"subject_zhs"`

	ContentEn  string `json:"content_en"`
	ContentZht string `json:"content_zht"`
	ContentZhs string `json:"content_zhs"`
}

func (t *SmtpTemplate) GetSubject(locale string, value map[string]string) string {
	subject := ""
	switch locale {
	case "en":
		subject = t.SubjectEn
	case "zht":
		subject = t.SubjectZht
	case "zhs":
		subject = t.SubjectZhs
	}
	for k, v := range value {
		subject = strings.ReplaceAll(subject, "{{"+k+"}}", v)
	}
	return subject
}

func (t *SmtpTemplate) GetContent(locale string, value map[string]string) string {
	content := ""
	switch locale {
	case "en":
		content = t.ContentEn
	case "zht":
		content = t.ContentZht
	case "zhs":
		content = t.ContentZhs
	}
	for k, v := range value {
		content = strings.ReplaceAll(content, "{{"+k+"}}", v)
	}
	return content
}
