package model

type MySQLConf struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Database string `json:"database"`
	Username string `json:"username"`
	Password string `json:"password"`
	LogoMode bool   `json:"logo_mode"`
}
