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

var API_SMTP_TEMPLATE_GET = atomic.NewApiHandler(
	module.Smtp,
	"M1idpaTvdmpayPNFHqaLZ",
	"Get SMTP template",
	func(ctx *atomic.Context[base.RequestIDPath]) {
		template := repo.SmtpTemplateRepo.FindByID(ctx.DB, ctx.Request.ID, ctx.WorkspaceID())
		if template == nil {
			ctx.Err(errs.SMTP_TEMPLATE_NOT_FOUND)
			return
		}
		ctx.OK(mapper.SmtpTemplateMapper.ToDTO(template))
	},
	&atomic.TypescriptOpt{
		Models:       []any{dto.SmtpAccount{}},
		FunctionName: "getSmtpTemplate",
		Paths:        []string{"id"},
		Response:     "SmtpTemplate",
	},
)
