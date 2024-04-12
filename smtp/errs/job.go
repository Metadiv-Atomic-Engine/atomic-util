package errs

import "github.com/Metadiv-Atomic-Engine/atomic/atomic"

var SMTP_JOB_NOT_FOUND = atomic.NewError(
	"7NSqanBlej5JvsziPLzGr",
	"SMTP job not found",
	"找不到SMTP工作",
	"找不到SMTP工作",
)

var SMTP_JOB_INVALID_STATUS = atomic.NewError(
	"c0mxOUZIRnO7NzE9oVXtL",
	"Invalid status",
	"無效的狀態",
	"无效的状态",
)
