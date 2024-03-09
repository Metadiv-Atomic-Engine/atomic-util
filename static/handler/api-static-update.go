package handler

import (
	"github.com/Metadiv-Atomic-Engine/atomic-util/static/errs"
	"github.com/Metadiv-Atomic-Engine/atomic-util/static/mapper"
	"github.com/Metadiv-Atomic-Engine/atomic-util/static/model/entity"
	"github.com/Metadiv-Atomic-Engine/atomic-util/static/model/request"
	"github.com/Metadiv-Atomic-Engine/atomic-util/static/module"
	"github.com/Metadiv-Atomic-Engine/atomic-util/static/repo"
	"github.com/Metadiv-Atomic-Engine/atomic/atomic"
)

var API_STATIC_UPDATE = atomic.NewApiHandler(
	module.Static,
	"UrERcgGni6aToA1PHFBIl",
	"Update static info",
	func(ctx *atomic.Context[request.StaticInfoUpdate]) {
		static := repo.StaticRepo.FindByUUID(ctx.DB, ctx.Request.UUID, ctx.WorkspaceID())
		if static == nil {
			ctx.Err(errs.STATIC_NOT_FOUND)
			return
		}

		static = mapper.StaticMapper.FromUpdateRequest(ctx.Request, static)
		static = repo.StaticRepo.Save(ctx.DB, static, ctx.WorkspaceID())
		if static == nil {
			ctx.InternalServerErr()
			return
		}

		ctx.OK(static)
	},
	&atomic.TypescriptOpt{
		FunctionName: "updateStaticInfo",
		Models:       []any{request.StaticInfoUpdate{}, entity.Static{}},
		Paths:        []string{"uuid"},
		Body:         "StaticInfoUpdate",
		Response:     "Static",
	},
)
