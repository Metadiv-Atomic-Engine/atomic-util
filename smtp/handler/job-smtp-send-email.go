package handler

import (
	"github.com/Metadiv-Atomic-Engine/atomic-util/smtp/model/entity"
	"github.com/Metadiv-Atomic-Engine/atomic-util/smtp/module"
	"github.com/Metadiv-Atomic-Engine/atomic-util/smtp/repo"
	"github.com/Metadiv-Atomic-Engine/atomic-util/smtp/service"
	"github.com/Metadiv-Atomic-Engine/atomic/atomic"
	"github.com/Metadiv-Atomic-Engine/sql"
)

var JOB_SMTP_SEND_EMAIL = atomic.NewJobHandler(
	module.Smtp,
	"5P6D9KgjUuCQr21f3mVKb",
	"Send email",
	func() {
		j := &entity.SmtpJob{}

		jobs := repo.SmtpJobRepo.FindAll(
			atomic.Engine.DB.Preload("Account").Preload("Template"),
			sql.Or(
				sql.Eq("status", j.StatusPending()),
				sql.Eq("status", j.StatusRetrying()),
			), "")

		for i := range jobs {
			service.SmtpService.SendEmail(atomic.Engine.DB, &jobs[i])
		}
	},
)
