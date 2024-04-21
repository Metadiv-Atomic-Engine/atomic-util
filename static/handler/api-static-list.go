package handler

import (
	"strings"

	"github.com/Metadiv-Atomic-Engine/atomic-util/static/model/entity"
	"github.com/Metadiv-Atomic-Engine/atomic-util/static/model/request"
	"github.com/Metadiv-Atomic-Engine/atomic-util/static/module"
	"github.com/Metadiv-Atomic-Engine/atomic-util/static/repo"
	"github.com/Metadiv-Atomic-Engine/atomic/atomic"
	"github.com/Metadiv-Atomic-Engine/sql"
)

var API_STATIC_LIST = atomic.NewApiHandler(
	module.Static,
	"UI7cPHeyg4tg2T3yf0WEP",
	"List static",
	func(ctx *atomic.Context[request.StaticList]) {
		cls := make([]*sql.Clause, 0)
		cls = append(cls, sql.Eq("pinned", true))
		if ctx.Request.IncludeUnpinned {
			cls = append(cls, sql.Eq("pinned", false))
		}
		cls = append(cls, ctx.Request.BuildSimilarClause(
			"uuid",
			"filename",
			"file_type",
		))
		if ctx.Request.UUIDs != "" {
			cls = append(cls, sql.In("uuid", strings.Split(ctx.Request.UUIDs, ",")))
		}
		statics, page := repo.StaticRepo.FindAllComplex(ctx.DB,
			sql.And(cls...), ctx.Page, ctx.Sort, "", ctx.WorkspaceID())
		if statics == nil {
			ctx.InternalServerErr()
			return
		}

		ctx.OK(statics, page)
	},
	&atomic.TypescriptOpt{
		Models:       []any{entity.Static{}},
		FunctionName: "listStaticInfo",
		Forms:        []string{"page", "size", "by", "asc", "keyword", "uuids", "include_unpinned"},
		Response:     "Static[]",
	},
)
