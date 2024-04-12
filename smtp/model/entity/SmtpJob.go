package entity

import (
	"encoding/json"
	"slices"
	"strings"

	"github.com/Metadiv-Atomic-Engine/aes"
	"github.com/Metadiv-Atomic-Engine/atomic/atomic"
)

type SmtpJob struct {
	ID uint `json:"id"`

	To  []byte `json:"to"`
	Cc  []byte `json:"cc"`
	Bcc []byte `json:"bcc"`

	AccountId uint         `json:"account_id"`
	Account   *SmtpAccount `json:"account"`

	TemplateId uint          `json:"template_id"`
	Template   *SmtpTemplate `json:"template"`

	Value  []byte `json:"value"`
	Locale string `json:"locale"`

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

func (j *SmtpJob) VerifyStatus(status string) (ok bool) {
	return slices.Contains([]string{
		j.StatusPending(),
		j.StatusRetrying(),
		j.StatusSuccess(),
		j.StatusFailed(),
	}, status)
}

func (j *SmtpJob) GetTo() []string {
	return strings.Split(j.decrypt(j.To), ",")
}

func (j *SmtpJob) GetCc() []string {
	return strings.Split(j.decrypt(j.Cc), ",")
}

func (j *SmtpJob) GetBcc() []string {
	return strings.Split(j.decrypt(j.Bcc), ",")
}

func (j *SmtpJob) GetValue() map[string]string {
	value := make(map[string]string)
	err := json.Unmarshal([]byte(j.Value), &value)
	if err != nil {
		atomic.Engine.Logger.ERROR(err)
		return nil
	}
	return value
}

func (j *SmtpJob) encrypt(text string) []byte {
	return aes.EncryptTextToBytes(text, atomic.Engine.EnvString(atomic.DB_ENCRYPT_KEY))
}

func (j *SmtpJob) decrypt(bytes []byte) string {
	return aes.DecryptBytesToString(bytes, atomic.Engine.EnvString(atomic.DB_ENCRYPT_KEY))
}

func (j *SmtpJob) SetValue(value map[string]string) {
	bytes, err := json.Marshal(value)
	if err != nil {
		atomic.Engine.Logger.ERROR(err)
		return
	}
	j.Value = j.encrypt(string(bytes))
}

func (j *SmtpJob) SetTo(to []string) {
	j.To = j.encrypt(strings.Join(to, ","))
}

func (j *SmtpJob) SetCc(cc []string) {
	j.Cc = j.encrypt(strings.Join(cc, ","))
}

func (j *SmtpJob) SetBcc(bcc []string) {
	j.Bcc = j.encrypt(strings.Join(bcc, ","))
}
