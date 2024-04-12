package handler

import (
	"github.com/Metadiv-Atomic-Engine/atomic-util/smtp/errs"
	"github.com/Metadiv-Atomic-Engine/atomic-util/smtp/mapper"
	"github.com/Metadiv-Atomic-Engine/atomic-util/smtp/model/dto"
	"github.com/Metadiv-Atomic-Engine/atomic-util/smtp/module"
	"github.com/Metadiv-Atomic-Engine/atomic-util/smtp/repo"
	"github.com/Metadiv-Atomic-Engine/atomic/atomic"
	"github.com/Metadiv-Atomic-Engine/atomic/base"
)

var API_SMTP_JOB_GET = atomic.NewApiHandler(
	module.Smtp,
	"ihumJ3NtiYvmfbBXcQxZ2",
	"Get smtp job",
	func(ctx *atomic.Context[base.RequestIDPath]) {

		job := repo.SmtpJobRepo.FindByID(ctx.DB, ctx.Request.ID, ctx.WorkspaceID())
		if job == nil {
			ctx.Err(errs.SMTP_JOB_NOT_FOUND)
			return
		}
		ctx.OK(mapper.SmtpJobMapper.ToDTO(job))
	},
	&atomic.TypescriptOpt{
		Models:       []any{dto.SmtpJob{}},
		FunctionName: "getSmtpJob",
		Paths:        []string{"id"},
		Response:     "SmtpJob",
	},
)
