package handler

import (
	"github.com/Metadiv-Atomic-Engine/atomic-util/smtp/errs"
	"github.com/Metadiv-Atomic-Engine/atomic-util/smtp/mapper"
	"github.com/Metadiv-Atomic-Engine/atomic-util/smtp/model/dto"
	"github.com/Metadiv-Atomic-Engine/atomic-util/smtp/model/request"
	"github.com/Metadiv-Atomic-Engine/atomic-util/smtp/module"
	"github.com/Metadiv-Atomic-Engine/atomic-util/smtp/repo"
	"github.com/Metadiv-Atomic-Engine/atomic/atomic"
)

var API_SMTP_JOB_CREATE = atomic.NewApiHandler(
	module.Smtp,
	"s1rgKOJgp66jueNqRFKAb",
	"Create smtp job",
	func(ctx *atomic.Context[request.SmtpJobCreate]) {

		account := repo.SmtpAccountRepo.FindByID(ctx.DB, ctx.Request.AccountId, ctx.WorkspaceID())
		if account == nil {
			ctx.Err(errs.SMTP_ACCOUNT_NOT_FOUND)
			return
		}

		template := repo.SmtpTemplateRepo.FindByID(ctx.DB, ctx.Request.TemplateId, ctx.WorkspaceID())
		if template == nil {
			ctx.Err(errs.SMTP_TEMPLATE_NOT_FOUND)
			return
		}

		job := mapper.SmtpJobMapper.FromCreateRequest(ctx.Request)
		job = repo.SmtpJobRepo.Save(ctx.DB, job, ctx.WorkspaceID())

		job.Account = account
		job.Template = template

		ctx.OK(mapper.SmtpJobMapper.ToDTO(job))
	},
	&atomic.TypescriptOpt{
		Models:       []any{dto.SmtpJob{}, request.SmtpJobCreate{}},
		FunctionName: "createSmtpJob",
		Body:         "SmtpJobCreate",
		Response:     "SmtpJob",
	},
)
