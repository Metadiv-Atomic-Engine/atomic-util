package handler

import (
	"github.com/Metadiv-Atomic-Engine/atomic-util/smtp/mapper"
	"github.com/Metadiv-Atomic-Engine/atomic-util/smtp/model/dto"
	"github.com/Metadiv-Atomic-Engine/atomic-util/smtp/module"
	"github.com/Metadiv-Atomic-Engine/atomic-util/smtp/repo"
	"github.com/Metadiv-Atomic-Engine/atomic/atomic"
	"github.com/Metadiv-Atomic-Engine/atomic/base"
)

var API_SMTP_TEMPLATE_LIST = atomic.NewApiHandler(
	module.Smtp,
	"IF9XhfjLm3yP2HaTGbMQf",
	"Delete SMTP template",
	func(ctx *atomic.Context[base.RequestListing]) {
		templates, page := repo.SmtpTemplateRepo.FindAllComplex(ctx.DB, ctx.Request.BuildDecryptedSimilarClause(
			"name",
			"subject_en",
			"subject_zht",
			"subject_zhs",
			"content_en",
			"content_zht",
			"content_zhs",
		), ctx.Page, ctx.Sort, "", ctx.WorkspaceID())
		ctx.OK(mapper.SmtpTemplateMapper.ToDTOs(templates), page)
	},
	&atomic.TypescriptOpt{
		Models:       []any{dto.SmtpTemplate{}},
		FunctionName: "listSmtpTemplate",
		Forms:        []string{"page", "size", "by", "asc", "keyword"},
		Response:     "SmtpTemplate[]",
	},
)
