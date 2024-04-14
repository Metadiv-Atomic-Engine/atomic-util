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

var API_SMTP_ACCOUNT_UPDATE = atomic.NewApiHandler(
	module.Smtp,
	"1IJwEP6jzV5KEa2M7KbZP",
	"Update SMTP Account",
	func(ctx *atomic.Context[request.SmtpAccountUpdate]) {

		account := repo.SmtpAccountRepo.FindByID(ctx.DB, ctx.Request.ID, ctx.WorkspaceID())
		if account == nil {
			ctx.Err(errs.SMTP_ACCOUNT_NOT_FOUND)
			return
		}

		account = mapper.SmtpAccountMapper.FromUpdateRequest(account, ctx.Request)
		account = repo.SmtpAccountRepo.Save(ctx.DB, account, ctx.WorkspaceID())
		ctx.OK(mapper.SmtpAccountMapper.ToDTO(account))
	},
	&atomic.TypescriptOpt{
		Models:       []any{request.SmtpAccountUpdate{}, dto.SmtpAccount{}},
		FunctionName: "updateSmtpAccount",
		Paths:        []string{"id"},
		Body:         "SmtpAccountUpdate",
		Response:     "SmtpAccount",
	},
)
