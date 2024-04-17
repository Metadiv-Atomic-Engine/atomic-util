package handler

import (
	"github.com/Metadiv-Atomic-Engine/atomic-util/static/model/request"
	"github.com/Metadiv-Atomic-Engine/atomic-util/static/module"
	"github.com/Metadiv-Atomic-Engine/atomic-util/static/repo"
	"github.com/Metadiv-Atomic-Engine/atomic/atomic"
	"github.com/Metadiv-Atomic-Engine/sql"
)

var API_STATIC_PIN = atomic.NewApiHandler(
	module.Static,
	"36XD4m9pFkNrChaTn9Lar",
	"Pin static",
	func(ctx *atomic.Context[request.StaticPin]) {
		statics := repo.StaticRepo.FindAll(
			ctx.DB, sql.In("uuid", ctx.Request.UUIDs), "", ctx.WorkspaceID())
		for i := range statics {
			statics[i].Pinned = true
		}
		if len(statics) > 0 {
			if repo.StaticRepo.SaveAll(ctx.DB, statics, ctx.WorkspaceID()) == nil {
				ctx.InternalServerErr("Failed to pin statics")
				return
			}
		}
		ctx.OK(nil)
	},
	&atomic.TypescriptOpt{
		FunctionName: "pinStatic",
		Models:       []any{request.StaticPin{}},
		Body:         "StaticPin",
	},
)
