package dto

type SmtpAccount struct {
	ID uint `json:"id"`

	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
}
