package email

import (
	"errors"
	"net/smtp"
	"time"
)

// Email an object that can send email
type Email struct {
	server string
	auth   smtp.Auth
}

// Content contains information about content
type Content struct {
	from    string
	to      string
	date    string
	subject string
	body    string
}

const layout = "Mon, 2 Jan 2006 15:04:05 MST"

// Create construct a Content object who is missing from field
func (content *Content) Create(to string, subject string, body string) error {
	if content == nil {
		return errors.New("content should not be nil")
	}

	content.to = to
	content.date = time.Now().Format(layout)
	content.subject = subject
	content.body = body

	return nil
}

func (content *Content) toBytes(from string) []byte {
	if content == nil {
		return nil
	}

	content.from = from

	retVal := ""
	retVal = retVal + "From: " + content.from + "\n"
	retVal = retVal + "To: " + content.to + "\n"
	retVal = retVal + "Date: " + content.date + "\n"
	retVal = retVal + "Subject: " + content.subject + "\n"
	retVal = retVal + "\n"
	retVal = retVal + content.body + "\n"

	return []byte(retVal)
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
func (e *Email) Send(to string, content Content) error {
	var from = "notifier.uwaterloo@gmail.com"

	return smtp.SendMail(e.server, e.auth, from, []string{to}, content.toBytes(from))
}
