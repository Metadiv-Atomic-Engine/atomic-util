package handler

import (
	"github.com/Metadiv-Atomic-Engine/atomic-util/smtp/errs"
	"github.com/Metadiv-Atomic-Engine/atomic-util/smtp/module"
	"github.com/Metadiv-Atomic-Engine/atomic-util/smtp/repo"
	"github.com/Metadiv-Atomic-Engine/atomic/atomic"
	"github.com/Metadiv-Atomic-Engine/atomic/base"
)

var API_SMTP_TEMPLATE_DELETE = atomic.NewApiHandler(
	module.Smtp,
	"6YsHyWg0stBUGYZ65XTFf",
	"Delete SMTP template",
	func(ctx *atomic.Context[base.RequestIDPath]) {
		template := repo.SmtpTemplateRepo.FindByID(ctx.DB, ctx.Request.ID, ctx.WorkspaceID())
		if template == nil {
			ctx.Err(errs.SMTP_TEMPLATE_NOT_FOUND)
			return
		}

		if !repo.SmtpTemplateRepo.Delete(ctx.DB, template) {
			ctx.InternalServerErr("Failed to delete template")
			return
		}
		ctx.OK(nil)
	},
	&atomic.TypescriptOpt{
		FunctionName: "deleteSmtpTemplate",
		Paths:        []string{"id"},
	},
)
