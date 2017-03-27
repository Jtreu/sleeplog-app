package infrastructure

import mailgun "gopkg.in/mailgun/mailgun-go.v1"

type MailHandler struct {
	mg      mailgun.Mailgun
	options *MailHandlerOptions
}

type MailHandlerOptions struct {
	Domain        string
	PrivateApiKey string
	PublicApiKey  string
}

func NewMailHandler(options MailHandlerOptions) *MailHandler {
	mailHandler := &MailHandler{
		options: &options,
	}
	mg := mailgun.NewMailgun(options.Domain, options.PrivateApiKey, options.PublicApiKey)
	mailHandler.mg = mg

	return mailHandler
}

func (mailer *MailHandler) Send(from, to, subject, content string) error {
	message := mailer.mg.NewMessage(from, subject, "", to)
	message.SetHtml(content)
	_, _, err := mailer.mg.Send(message)
	return err
}
