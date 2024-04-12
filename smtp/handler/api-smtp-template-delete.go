package handler

import (
	"github.com/Metadiv-Atomic-Engine/atomic-util/smtp/module"
	"github.com/Metadiv-Atomic-Engine/atomic/atomic"
	"github.com/Metadiv-Atomic-Engine/atomic/base"
)

var API_SMTP_TEMPLATE_DELETE = atomic.NewApiHandler(
	module.Smtp,
	"6YsHyWg0stBUGYZ65XTFf",
	"Delete SMTP template",
	func(ctx *atomic.Context[base.RequestIDPath]) {

	},
	&atomic.TypescriptOpt{
		FunctionName: "deleteSmtpTemplate",
		Paths:        []string{"id"},
	},
)
