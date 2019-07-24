package email

import (
	"encoding/json"
	"io/ioutil"
	"net/smtp"
	"os"
)

// Email an object that can send email
type Email struct {
	configure Config
}

// InitByConfig initialize email by class Config
func (email *Email) InitByConfig(config Config) {
	if email != nil {
		email.configure = config
	}
}

// InitByFile initialze email by configuration file at path
func (email *Email) InitByFile(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}

	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	json.Unmarshal(bytes, &email.configure)

	return nil
}

// Send send the email
func (email *Email) Send(to string, content Content) error {
	server := email.configure.toServerAddr()
	auth := email.configure.toAuth()
	from := email.configure.toEmailAddr()

	return smtp.SendMail(
		server,
		auth,
		from,
		[]string{to},
		content.toBytes(from, to))
}
