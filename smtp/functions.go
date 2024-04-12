package smtp

import (
	"github.com/Metadiv-Atomic-Engine/atomic-util/smtp/errs"
	"github.com/Metadiv-Atomic-Engine/atomic-util/smtp/model/entity"
	"github.com/Metadiv-Atomic-Engine/atomic-util/smtp/repo"
	"github.com/Metadiv-Atomic-Engine/atomic-util/smtp/service"
	"gorm.io/gorm"
)

func SendEmail(
	tx *gorm.DB,
	to []string,
	cc []string,
	bcc []string,
	accountID uint,
	templateID uint,
	value map[string]string,
	locale string,
	sendImmediately bool,
	workspaceID ...uint,
) (errCode string) {
	account := repo.SmtpAccountRepo.FindByID(tx, accountID, workspaceID...)
	if account == nil {
		return errs.SMTP_ACCOUNT_NOT_FOUND
	}

	template := repo.SmtpTemplateRepo.FindByID(tx, templateID, workspaceID...)
	if template == nil {
		return errs.SMTP_TEMPLATE_NOT_FOUND
	}

	job := &entity.SmtpJob{}
	job.AccountId = accountID
	job.TemplateId = templateID
	job.Status = job.StatusPending()
	job.SetTo(to)
	job.SetCc(cc)
	job.SetBcc(bcc)
	job.Locale = locale
	job.SetValue(value)
	repo.SmtpJobRepo.Save(tx, job, workspaceID...)
	if sendImmediately {
		service.SmtpService.SendEmail(tx, job)
	}
	return ""
}
