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

var API_SMTP_JOB_UPDATE = atomic.NewApiHandler(
	module.Smtp,
	"evgEXkWc6to2OV9JEsrE0",
	"Update smtp job",
	func(ctx *atomic.Context[request.SmtpJobUpdate]) {

		job := repo.SmtpJobRepo.FindByID(ctx.DB, ctx.Request.ID, ctx.WorkspaceID())
		if job == nil {
			ctx.Err(errs.SMTP_JOB_NOT_FOUND)
			return
		}

		if !job.VerifyStatus(ctx.Request.Status) {
			ctx.Err(errs.SMTP_JOB_INVALID_STATUS)
			return
		}

		job = mapper.SmtpJobMapper.FromUpdateRequest(job, ctx.Request)

		job = repo.SmtpJobRepo.Save(ctx.DB, job, ctx.WorkspaceID())
		ctx.OK(mapper.SmtpJobMapper.ToDTO(job))
	},
	&atomic.TypescriptOpt{
		Models:       []any{request.SmtpJobUpdate{}, dto.SmtpJob{}},
		FunctionName: "updateSmtpJob",
		Paths:        []string{"id"},
		Body:         "SmtpJobUpdate",
		Response:     "SmtpJob",
	},
)
