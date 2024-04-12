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

var API_SMTP_TEMPLATE_UPDATE = atomic.NewApiHandler(
	module.Smtp,
	"7XcotpLWcvCJWU6SJ631Q",
	"Update SMTP Template",
	func(ctx *atomic.Context[request.SmtpTemplateUpdate]) {
		template := repo.SmtpTemplateRepo.FindByID(ctx.DB, ctx.Request.ID, ctx.WorkspaceID())
		if template == nil {
			ctx.Err(errs.SMTP_TEMPLATE_NOT_FOUND)
			return
		}
		template = mapper.SmtpTemplateMapper.FromUpdateRequest(template, ctx.Request)
		template = repo.SmtpTemplateRepo.Save(ctx.DB, template, ctx.WorkspaceID())
		ctx.OK(mapper.SmtpTemplateMapper.ToDTO(template))
	},
	&atomic.TypescriptOpt{
		Models:       []any{request.SmtpTemplateUpdate{}, dto.SmtpTemplate{}},
		FunctionName: "updateSmtpTemplate",
		Body:         "SmtpTemplateUpdate",
		Response:     "SmtpTemplate",
	},
)
