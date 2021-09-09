package model

type EmailConf struct {
	Username string `json:"user"`
	Password string `json:"pass"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	Title    string `json:"title"`
	Sender   string `json:"sender"`
	Body     string `json:"body"`
}
