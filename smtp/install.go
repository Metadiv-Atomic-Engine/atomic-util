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

	/* smtp template */
	atomic.NewGetApi(
		"8Pa54U03B1LPpK8gKjCcu",
		handler.API_SMTP_TEMPLATE_LIST,
		"/smtp/template",
		nil,
		nil,
	)

	atomic.NewGetApi(
		"h6oJhU9cgViiB4vZR0pBl",
		handler.API_SMTP_TEMPLATE_GET,
		"/smtp/template/:id",
		nil,
		nil,
	)

	atomic.NewPostApi(
		"jSyrN4SX1IiKJokjzLQLi",
		handler.API_SMTP_TEMPLATE_CREATE,
		"/smtp/template",
		nil,
		nil,
	)

	atomic.NewPutApi(
		"GppKnbMHN6xOWwROKGjmq",
		handler.API_SMTP_TEMPLATE_UPDATE,
		"/smtp/template/:id",
		nil,
		nil,
	)

	atomic.NewDeleteApi(
		"TUV8CWMmu5fYBbwUMY7Hc",
		handler.API_SMTP_TEMPLATE_DELETE,
		"/smtp/template/:id",
		nil,
		nil,
	)

	/* smtp job */
	atomic.NewGetApi(
		"tUfD3bmE5CmX3dYr1RHt6",
		handler.API_SMTP_JOB_LIST,
		"/smtp/job",
		nil,
		nil,
	)

	atomic.NewGetApi(
		"tIbLQDfmCNJgHCRIu8EU1",
		handler.API_SMTP_JOB_GET,
		"/smtp/job/:id",
		nil,
		nil,
	)

	atomic.NewPostApi(
		"sjO5kwJ8kbOd89mNuXIwM",
		handler.API_SMTP_JOB_CREATE,
		"/smtp/job",
		nil,
		nil,
	)

	atomic.NewPutApi(
		"T4uj67UX5j5JF1gRIARPM",
		handler.API_SMTP_JOB_UPDATE,
		"/smtp/job/:id",
		nil,
		nil,
	)

	atomic.NewDeleteApi(
		"AvvtDRLDFm6htqP9JQ3ri",
		handler.API_SMTP_JOB_DELETE,
		"/smtp/job/:id",
		nil,
		nil,
	)
}
