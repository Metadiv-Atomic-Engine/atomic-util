package handler

import (
	"github.com/Metadiv-Atomic-Engine/atomic-util/smtp/mapper"
	"github.com/Metadiv-Atomic-Engine/atomic-util/smtp/model/dto"
	"github.com/Metadiv-Atomic-Engine/atomic-util/smtp/model/request"
	"github.com/Metadiv-Atomic-Engine/atomic-util/smtp/module"
	"github.com/Metadiv-Atomic-Engine/atomic-util/smtp/repo"
	"github.com/Metadiv-Atomic-Engine/atomic/atomic"
)

var API_SMTP_TEMPLATE_CREATE = atomic.NewApiHandler(
	module.Smtp,
	"IAuB8LS8uBhyGy1LhzjQM",
	"Create SMTP Template",
	func(ctx *atomic.Context[request.SmtpTemplateCreate]) {
		template := mapper.SmtpTemplateMapper.FromCreateRequest(ctx.Request)
		template = repo.SmtpTemplateRepo.Save(ctx.DB, template, ctx.WorkspaceID())
		ctx.OK(mapper.SmtpTemplateMapper.ToDTO(template))
	},
	&atomic.TypescriptOpt{
		Models:       []any{request.SmtpTemplateCreate{}, dto.SmtpTemplate{}},
		FunctionName: "createSmtpTemplate",
		Body:         "SmtpTemplateCreate",
		Response:     "SmtpTemplate",
	},
)
