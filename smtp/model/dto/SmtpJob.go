package dto

type SmtpJob struct {
	ID uint `json:"id"`

	To  []string `json:"to"`
	Cc  []string `json:"cc"`
	Bcc []string `json:"bcc"`

	AccountId uint         `json:"account_id"`
	Account   *SmtpAccount `json:"account"`

	TemplateId uint          `json:"template_id"`
	Template   *SmtpTemplate `json:"template"`

	Value  map[string]string `json:"value"`
	Locale string            `json:"locale"`

	TryTimes int    `json:"try_times"`
	Status   string `json:"status"` // pending, retrying, success, failed
}

func (j *SmtpJob) StatusPending() string {
	return "pending"
}

func (j *SmtpJob) StatusRetrying() string {
	return "retrying"
}

func (j *SmtpJob) StatusSuccess() string {
	return "success"
}

func (j *SmtpJob) StatusFailed() string {
	return "failed"
}
