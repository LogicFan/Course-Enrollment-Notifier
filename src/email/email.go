package email

import "net/smtp"

// Email an object that can send email
type Email struct {
	server string
	auth   smtp.Auth
}

var a = smtp.PlainAuth(
	"",
	"notifier.uwaterloo@gmail.com",
	"Logic990316",
	"smtp.gmail.com",
)

// TestEmail test purpose only
var TestEmail = Email{server: "smtp.gmail.com:587", auth: a}

// InitByConfig initialze email by configuration file at path
func (*Email) InitByConfig(path string) error {
	return nil
}

// Send send the email
func (e *Email) Send(from string, to []string, msg []byte) error {
	return smtp.SendMail(e.server, e.auth, from, to, msg)
}
