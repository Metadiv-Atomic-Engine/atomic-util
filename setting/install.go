package setting

import (
	"time"

	"github.com/Metadiv-Atomic-Engine/atomic-util/setting/handler"
	"github.com/Metadiv-Atomic-Engine/atomic-util/setting/model/entity"
	"github.com/Metadiv-Atomic-Engine/atomic-util/setting/module"
	"github.com/Metadiv-Atomic-Engine/atomic/atomic"
)

/*
Install setting module to support system settings.
*/
func Install() {
	atomic.Engine.InstallModule(module.Setting)

	atomic.NewDBMigration(
		entity.SystemSetting{},
	)

	/*
		Workspace user apis
	*/
	atomic.NewGetApi(
		"yEHOQv5LvvUY6PSE2BXgC",
		handler.API_SETTING_WORKSPACE_GET,
		"/setting",
		nil,
		nil,
	)
	atomic.NewPutApi(
		"k9ElVWUstVF1D5uNTjZ8o",
		handler.API_SETTING_WORKSPACE_UPDATE,
		"/setting",
		nil,
		nil,
	)

	/*
		Global apis
	*/
	atomic.NewGetApi(
		"bY7EZosrQSEsNCskdvSgi",
		handler.API_SETTING_GLOBAL_GET,
		"/sys/setting",
		nil,
		nil,
	)
	atomic.NewPutApi(
		"UQtRYI2ud4ObTNNxqH1vX",
		handler.API_SETTING_GLOBAL_UPDATE,
		"/sys/setting",
		nil,
		nil,
	)

	/*
		Public apis
	*/
	atomic.NewGetApi(
		"Pft8tL28CI8ESmNV7dng1",
		handler.API_SETTING_PUBLIC_GET,
		"/public/setting",
		nil,
		&atomic.RateLimitOpt{
			Rate:     60,
			Duration: time.Minute,
		},
	)
	atomic.NewGetApi(
		"3MlJ7H6YgaVOGBeQ3j1m2",
		handler.API_SETTING_PUBLIC_GLOBAL_GET,
		"/public/global/setting",
		nil,
		&atomic.RateLimitOpt{
			Rate:     60,
			Duration: time.Minute,
		},
	)
}
