package handler

import (
	"github.com/Metadiv-Atomic-Engine/atomic-util/smtp/mapper"
	"github.com/Metadiv-Atomic-Engine/atomic-util/smtp/model/dto"
	"github.com/Metadiv-Atomic-Engine/atomic-util/smtp/module"
	"github.com/Metadiv-Atomic-Engine/atomic-util/smtp/repo"
	"github.com/Metadiv-Atomic-Engine/atomic/atomic"
	"github.com/Metadiv-Atomic-Engine/atomic/base"
	"github.com/Metadiv-Atomic-Engine/sql"
)

var API_SMTP_JOB_LIST = atomic.NewApiHandler(
	module.Smtp,
	"OWqmRDqr8AYO032EK7i4Z",
	"List smtp jobs",
	func(ctx *atomic.Context[base.RequestListing]) {
		jobs, page := repo.SmtpJobRepo.FindAllComplex(ctx.DB, sql.Or(
			ctx.Request.BuildSimilarClause(
				"status",
			),
			ctx.Request.BuildDecryptedSimilarClause(
				"value",
			),
		), ctx.Page, ctx.Sort, "", ctx.WorkspaceID())
		ctx.OK(mapper.SmtpJobMapper.ToDTOs(jobs), page)
	},
	&atomic.TypescriptOpt{
		Models:       []any{dto.SmtpJob{}},
		FunctionName: "listSmtpJobs",
		Forms:        []string{"page", "size", "by", "asc", "status", "keyword"},
		Response:     "SmtpJob[]",
	},
)
