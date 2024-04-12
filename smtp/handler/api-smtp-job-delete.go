package handler

import (
	"github.com/Metadiv-Atomic-Engine/atomic-util/smtp/errs"
	"github.com/Metadiv-Atomic-Engine/atomic-util/smtp/module"
	"github.com/Metadiv-Atomic-Engine/atomic-util/smtp/repo"
	"github.com/Metadiv-Atomic-Engine/atomic/atomic"
	"github.com/Metadiv-Atomic-Engine/atomic/base"
)

var API_SMTP_JOB_DELETE = atomic.NewApiHandler(
	module.Smtp,
	"HRJw0awEZ5Iqj56yTnna3",
	"Delete smtp job",
	func(ctx *atomic.Context[base.RequestIDPath]) {

		job := repo.SmtpJobRepo.FindByID(ctx.DB, ctx.Request.ID, ctx.WorkspaceID())
		if job == nil {
			ctx.Err(errs.SMTP_JOB_NOT_FOUND)
			return
		}

		if !repo.SmtpJobRepo.Delete(ctx.DB, job) {
			ctx.Err(errs.SMTP_JOB_NOT_FOUND)
			return
		}

		ctx.OK(nil)
	},
	&atomic.TypescriptOpt{
		FunctionName: "deleteSmtpJob",
		Paths:        []string{"id"},
	},
)
