package smtp

import (
	"github.com/Metadiv-Atomic-Engine/atomic-util/smtp/handler"
	"github.com/Metadiv-Atomic-Engine/atomic-util/smtp/module"
	"github.com/Metadiv-Atomic-Engine/atomic/atomic"
)

func Install() {
	atomic.Engine.InstallModule(module.Smtp)

	/* smtp account */
	atomic.NewGetApi(
		"eerpcRVewubckVEvSqMng",
		handler.API_SMTP_ACCOUNT_LIST,
		"/smtp/account",
		nil,
		nil,
	)

	atomic.NewGetApi(
		"RrOgXH6kAE26CQuwcOCmH",
		handler.API_SMTP_ACCOUNT_GET,
		"/smtp/account/:id",
		nil,
		nil,
	)

	atomic.NewPostApi(
		"0ZjEqJHYoTTWGWgWNgwPB",
		handler.API_SMTP_ACCOUNT_CREATE,
		"/smtp/account",
		nil,
		nil,
	)

	atomic.NewPutApi(
		"xUw3enZ4y8We64nn18JIV",
		handler.API_SMTP_ACCOUNT_UPDATE,
		"/smtp/account/:id",
		nil,
		nil,
	)

	atomic.NewDeleteApi(
		"Xl4r1qAGaijierDQopNNY",
		handler.API_SMTP_ACCOUNT_DELETE,
		"/smtp/account/:id",
		nil,
		nil,
	)
}
