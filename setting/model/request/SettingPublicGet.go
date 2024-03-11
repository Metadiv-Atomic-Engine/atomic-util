package request

type SettingPublicGet struct {
	Workspace uint `form:"workspace" json:"workspace" binding:"required"`
	SettingGet
}
