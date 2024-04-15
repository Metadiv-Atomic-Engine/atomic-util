package handler

import (
	"github.com/Metadiv-Atomic-Engine/atomic-util/smtp/mapper"
	"github.com/Metadiv-Atomic-Engine/atomic-util/smtp/model/dto"
	"github.com/Metadiv-Atomic-Engine/atomic-util/smtp/module"
	"github.com/Metadiv-Atomic-Engine/atomic-util/smtp/repo"
	"github.com/Metadiv-Atomic-Engine/atomic/atomic"
	"github.com/Metadiv-Atomic-Engine/atomic/base"
)

var API_SMTP_ACCOUNT_LIST = atomic.NewApiHandler(
	module.Smtp,
	"ljaoyjA5O38dbxo0HQG8y",
	"SMTP account list",
	func(ctx *atomic.Context[base.RequestListing]) {
		accounts, page := repo.SmtpAccountRepo.FindAllComplex(ctx.DB, ctx.Request.BuildDecryptedSimilarClause(
			"host",
			"port",
			"user",
			"password",
		), ctx.Page, ctx.Sort, "", ctx.WorkspaceID())
		ctx.OK(mapper.SmtpAccountMapper.ToDTOs(accounts), page)
	},
	&atomic.TypescriptOpt{
		Models:       []any{dto.SmtpAccount{}},
		FunctionName: "listSmtpAccount",
		Forms:        []string{"page", "size", "by", "asc", "keyword"},
		Response:     "SmtpAccount[]",
	},
)
