package request

type SmtpJobCreate struct {
	To  []string `json:"to"`
	Cc  []string `json:"cc"`
	Bcc []string `json:"bcc"`

	AccountId  uint `json:"account_id"`
	TemplateId uint `json:"template_id"`

	Value  map[string]string `json:"value"`
	Locale string            `json:"locale"`
}
