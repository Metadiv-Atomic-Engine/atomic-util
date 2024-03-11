package request

type SettingGet struct {
	/*
		Keys separated by comma.
	*/
	Keys string `form:"keys" json:"keys" binding:"required"`
}
