package static

import (
	"time"

	"github.com/Metadiv-Atomic-Engine/atomic-util/static/handler"
	"github.com/Metadiv-Atomic-Engine/atomic-util/static/model/entity"
	"github.com/Metadiv-Atomic-Engine/atomic-util/static/module"
	"github.com/Metadiv-Atomic-Engine/atomic/atomic"
)

func Install() {
	atomic.Engine.InstallModule(module.Static)

	atomic.NewDBMigration(
		&entity.Static{},
	)

	/*
		uploads
	*/
	atomic.NewPostApi(
		"Nj2GXgsgxC3r0MNIWwUwx",
		handler.API_UPLOAD,
		"/upload",
		nil,
		nil,
	)
	atomic.NewPostApi(
		"ydWU95F6uStPFK77JdKRg",
		handler.API_GLOBAL_UPLOAD,
		"/global/upload",
		nil,
		nil,
	)

	/*
		Serving
	*/
	atomic.NewGetApi(
		"ZQUyXXB80dDdcyh561aIo",
		handler.API_ASSET_SERVE,
		"/assets/:uuid",
		nil,
		nil,
	)
	atomic.NewGetApi(
		"rn6vROiRAYsC1PummxpEJ",
		handler.API_PUBLIC_ASSET_SERVE,
		"/public/assets/:uuid",
		&atomic.CacheOpt{
			Duration: time.Minute,
		},
		&atomic.RateLimitOpt{
			Rate:     60 * 2,
			Duration: time.Minute,
		},
	)

	/*
		Download
	*/
	atomic.NewGetApi(
		"P16lyxvcGo2FZwBHQL3l8",
		handler.API_ASSET_DOWNLOAD,
		"/assets/:uuid/download",
		nil,
		nil,
	)
	atomic.NewGetApi(
		"67eONDx0sjFo4vixBg5Mv",
		handler.API_PUBLIC_ASSET_DOWNLOAD,
		"/public/assets/:uuid/download",
		&atomic.CacheOpt{
			Duration: time.Minute,
		},
		&atomic.RateLimitOpt{
			Rate:     60 * 2,
			Duration: time.Minute,
		},
	)

	/*
		Static management
	*/
	atomic.NewGetApi(
		"IyT7XyfHPSIO54rllv1aP",
		handler.API_STATIC_LIST,
		"/static",
		nil,
		nil,
	)
	atomic.NewGetApi(
		"zaP9HJm5Rz3RPYFj7CVpH",
		handler.API_STATIC_GET,
		"/static/:uuid",
		nil,
		nil,
	)
	atomic.NewPutApi(
		"LpdMeMfWf3ZWFsNgCcUsj",
		handler.API_STATIC_UPDATE,
		"/static/:uuid",
		nil,
		nil,
	)
	atomic.NewDeleteApi(
		"hiwfEEFbDe9Ml3TDjQMVK",
		handler.API_STATIC_DELETE,
		"/static/:uuid",
		nil,
		nil,
	)
	atomic.NewPostApi(
		"9GMWhPPRXOptaFJudFThU",
		handler.API_STATIC_PIN,
		"/static/pin",
		nil,
		nil,
	)

	/*
		Jobs
	*/
	atomic.NewCronJob(
		"wabX8MlSPHnIWH82isjW6",
		handler.JOB_REMOVE_UNPINNED,
		"@every 1m",
		true,
	)
}
