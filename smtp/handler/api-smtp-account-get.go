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

var API_SMTP_ACCOUNT_GET = atomic.NewApiHandler(
	module.Smtp,
	"8SNFa5aZB0ljLx7SSPk6N",
	"Get SMTP account",
	func(ctx *atomic.Context[base.RequestIDPath]) {

		account := repo.SmtpAccountRepo.FindByID(ctx.DB, ctx.Request.ID, ctx.WorkspaceID())
		if account == nil {
			ctx.Err(errs.SMTP_ACCOUNT_NOT_FOUND)
			return
		}
		ctx.OK(mapper.SmtpAccountMapper.ToDTO(account))
	},
	&atomic.TypescriptOpt{
		Models:       []any{dto.SmtpAccount{}},
		FunctionName: "getSmtpAccount",
		Paths:        []string{"id"},
		Response:     "SmtpAccount",
	},
)
