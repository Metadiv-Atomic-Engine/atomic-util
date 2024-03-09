package handler

import (
	"github.com/Metadiv-Atomic-Engine/atomic-util/static/errs"
	"github.com/Metadiv-Atomic-Engine/atomic-util/static/model/entity"
	"github.com/Metadiv-Atomic-Engine/atomic-util/static/module"
	"github.com/Metadiv-Atomic-Engine/atomic-util/static/repo"
	"github.com/Metadiv-Atomic-Engine/atomic/atomic"
	"github.com/Metadiv-Atomic-Engine/atomic/base"
	"github.com/Metadiv-Atomic-Engine/sql"
)

var API_STATIC_GET = atomic.NewApiHandler(
	module.Static,
	"oL00LtsNb66OpFQQyqojU",
	"Get static info by id",
	func(ctx *atomic.Context[base.RequestUUIDPath]) {
		static := repo.StaticRepo.FindOne(ctx.DB, sql.Eq("uuid", ctx.Request.UUID), ctx.WorkspaceID())
		if static == nil {
			ctx.Err(errs.STATIC_NOT_FOUND)
			return
		}
		ctx.OK(static)
	},
	&atomic.TypescriptOpt{
		Models:       []any{entity.Static{}},
		FunctionName: "getStaticInfo",
		Paths:        []string{"uuid"},
		Response:     "Static",
	},
)
