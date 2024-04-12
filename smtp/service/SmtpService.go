package service

import (
	"github.com/Metadiv-Atomic-Engine/atomic-util/smtp/mapper"
	"github.com/Metadiv-Atomic-Engine/atomic-util/smtp/model/entity"
	"github.com/Metadiv-Atomic-Engine/atomic-util/smtp/repo"
	"gopkg.in/mail.v2"
	"gorm.io/gorm"
)

var SmtpService = new(smtpService)

type smtpService struct{}

func (s *smtpService) SendEmail(tx *gorm.DB, jobEntity *entity.SmtpJob) {

	job := mapper.SmtpJobMapper.ToDTO(jobEntity)
	if job.Account == nil || job.Template == nil {
		s.handleFailJob(tx, jobEntity)
		return
	}

	m := mail.NewMessage()
	m.SetHeader("From", job.Account.User)
	m.SetHeader("To", job.To...)
	m.SetHeader("Cc", job.Cc...)
	m.SetHeader("Bcc", job.Bcc...)

	var bodyType string
	if job.Template.Type == "html" {
		bodyType = "text/html"
	} else {
		bodyType = "text/plain"
	}

	m.SetHeader("Subject", job.Template.GetSubject(job.Locale, job.Value))
	m.SetBody(bodyType, job.Template.GetContent(job.Locale, job.Value))

	d := mail.NewDialer(job.Account.Host, job.Account.Port, job.Account.User, job.Account.Password)
	if err := d.DialAndSend(m); err != nil {
		s.handleFailJob(tx, jobEntity)
	}
}

func (s *smtpService) handleFailJob(tx *gorm.DB, jobEntity *entity.SmtpJob) {
	jobEntity.TryTimes++
	if jobEntity.TryTimes >= 3 {
		jobEntity.Status = jobEntity.StatusFailed()
	} else {
		jobEntity.Status = jobEntity.StatusRetrying()
	}
	repo.SmtpJobRepo.Save(tx, jobEntity)
}
