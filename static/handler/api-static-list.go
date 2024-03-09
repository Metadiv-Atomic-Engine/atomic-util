package handler

import (
	"github.com/Metadiv-Atomic-Engine/atomic-util/static/model/entity"
	"github.com/Metadiv-Atomic-Engine/atomic-util/static/module"
	"github.com/Metadiv-Atomic-Engine/atomic-util/static/repo"
	"github.com/Metadiv-Atomic-Engine/atomic/atomic"
	"github.com/Metadiv-Atomic-Engine/atomic/base"
)

var API_STATIC_LIST = atomic.NewApiHandler(
	module.Static,
	"UI7cPHeyg4tg2T3yf0WEP",
	"List static",
	func(ctx *atomic.Context[base.RequestListing]) {
		statics, page := repo.StaticRepo.FindAllComplex(ctx.DB, ctx.Request.BuildSimilarClause(
			"uuid",
			"filename",
			"file_type",
		), ctx.Page, ctx.Sort, "", ctx.WorkspaceID())
		if statics == nil {
			ctx.InternalServerErr()
			return
		}

		ctx.OK(statics, page)
	},
	&atomic.TypescriptOpt{
		Models:       []any{entity.Static{}},
		FunctionName: "listStaticInfo",
		Forms:        []string{"page", "size", "by", "asc", "keyword"},
		Response:     "Static[]",
	},
)
