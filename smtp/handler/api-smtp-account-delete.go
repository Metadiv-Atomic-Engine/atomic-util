package handler

import (
	"github.com/Metadiv-Atomic-Engine/atomic-util/smtp/errs"
	"github.com/Metadiv-Atomic-Engine/atomic-util/smtp/module"
	"github.com/Metadiv-Atomic-Engine/atomic-util/smtp/repo"
	"github.com/Metadiv-Atomic-Engine/atomic/atomic"
	"github.com/Metadiv-Atomic-Engine/atomic/base"
)

var API_SMTP_ACCOUNT_DELETE = atomic.NewApiHandler(
	module.Smtp,
	"2erub49cusiU307AoMpT3",
	"Delete SMTP account",
	func(ctx *atomic.Context[base.RequestIDPath]) {

		account := repo.SmtpAccountRepo.FindByID(ctx.DB, ctx.Request.ID, ctx.WorkspaceID())
		if account == nil {
			ctx.Err(errs.SMTP_ACCOUNT_NOT_FOUND)
			return
		}

		if !repo.SmtpAccountRepo.Delete(ctx.DB, account) {
			ctx.Err(errs.SMTP_ACCOUNT_NOT_FOUND)
			return
		}

		ctx.OK(nil)
	},
	&atomic.TypescriptOpt{
		FunctionName: "deleteSmtpAccount",
		Paths:        []string{"id"},
	},
)
