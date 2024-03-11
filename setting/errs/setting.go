package errs

import "github.com/Metadiv-Atomic-Engine/atomic/atomic"

var SETTING_KEY_NOT_FOUND = atomic.NewError(
	"66kOACaZHj8AUTg6gLP2O",
	"Setting key not found",
	"設定鍵未找到",
	"设置键未找到",
)

var WORKSPACE_HAS_NO_SETTINGS = atomic.NewError(
	"Pu93YPvmvxyD8eWLI55u1",
	"Workspace has no settings",
	"工作區沒有設定",
	"工作区没有设置",
)
