package email

import (
	"net/smtp"
	"strconv"
)

// Config a structure of configuration
type Config struct {
	Email  string
	Passwd string
	Host   string
	Port   int
}

func (config Config) toAuth() smtp.Auth {
	return smtp.PlainAuth(
		"",
		config.Email,
		config.Passwd,
		config.Host,
	)
}

func (config Config) toServerAddr() string {
	return config.Host + ":" + strconv.FormatInt(int64(config.Port), 10)
}

func (config Config) toEmailAddr() string {
	return config.Email
}
