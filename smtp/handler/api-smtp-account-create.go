package handler

import (
	"github.com/Metadiv-Atomic-Engine/atomic-util/smtp/mapper"
	"github.com/Metadiv-Atomic-Engine/atomic-util/smtp/model/dto"
	"github.com/Metadiv-Atomic-Engine/atomic-util/smtp/model/request"
	"github.com/Metadiv-Atomic-Engine/atomic-util/smtp/module"
	"github.com/Metadiv-Atomic-Engine/atomic-util/smtp/repo"
	"github.com/Metadiv-Atomic-Engine/atomic/atomic"
)

var API_SMTP_ACCOUNT_CREATE = atomic.NewApiHandler(
	module.Smtp,
	"NH1vn0NpYp9mZBXaAnXoY",
	"Create SMTP Account",
	func(ctx *atomic.Context[request.SmtpAccountCreate]) {
		account := mapper.SmtpAccountMapper.FromCreateRequest(ctx.Request)
		account = repo.SmtpAccountRepo.Save(ctx.DB, account, ctx.WorkspaceID())
		ctx.OK(mapper.SmtpAccountMapper.ToDTO(account))
	},
	&atomic.TypescriptOpt{
		Models:       []any{dto.SmtpAccount{}, request.SmtpAccountCreate{}},
		FunctionName: "createSmtpAccount",
		Body:         "SmtpAccountCreate",
		Response:     "SmtpAccount",
	},
)
