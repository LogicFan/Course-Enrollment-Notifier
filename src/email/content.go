package email

import (
	"time"
)

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
func (content *Content) Create(subject string, body string) {
	if content == nil {
		return
	}

	content.date = time.Now().Format(layout)
	content.subject = subject
	content.body = body

	return
}

func (content *Content) toBytes(from string, to string) []byte {
	if content == nil {
		return nil
	}

	content.from = from
	content.to = to

	retVal := ""
	retVal = retVal + "From: " + content.from + "\n"
	retVal = retVal + "To: " + content.to + "\n"
	retVal = retVal + "Date: " + content.date + "\n"
	retVal = retVal + "Subject: " + content.subject + "\n"
	retVal = retVal + "\n"
	retVal = retVal + content.body + "\n"

	return []byte(retVal)
}
