package config

import (
	"github.com/Metadiv-Atomic-Engine/atomic-util/static/module"
	"github.com/Metadiv-Atomic-Engine/atomic/atomic"
)

var (
	STATIC_FOLDER_NAME = atomic.NewEnvString(
		module.Static,
		"STATIC_FOLDER_NAME",
		"Static folder name",
		"static_assets",
	)
	STATIC_REVERSE_TIME_IN_MINUTES = atomic.NewEnvInt(
		module.Static,
		"STATIC_REVERSE_TIME_IN_MINUTES",
		"Static reverse time in minutes",
		60,
	)
)
