package mailer

import (
	"crypto/tls"
	"fmt"
	"strconv"

	mail "github.com/wneessen/go-mail"

	"example.com/config"
	"example.com/models"
)

func SendProjectConfirmation(event models.ProjectCreatedEvent) error {

	message := mail.NewMsg()

	if err := message.From(config.SMTP.Email); err != nil {
		return err
	}

	if err := message.To(event.Email); err != nil {
		return err
	}

	message.Subject("Your ValtQ Discovery Session")

	html := fmt.Sprintf(`
<h1>Welcome to ValtQ 🚀</h1>

<p>Hello <b>%s</b>,</p>

<p>Thank you for submitting your project.</p>

<hr>

<p><b>Reference ID:</b> %s</p>

<p><b>Project:</b> %s</p>

<p><b>Service:</b> %s</p>

<p><b>Budget:</b> %s</p>

<p><b>Description:</b> %s</p>

<br>

<p>
<a href="%s">Join Discovery Meeting</a>
</p>

<br>

<p>Regards,<br>ValtQ Team</p>
`,
		event.ClientName,
		event.ReferenceID,
		event.ProjectName,
		event.Service,
		event.Budget,
		event.Description,
		event.MeetingLink,
	)

	message.SetBodyString(mail.TypeTextHTML, html)

	port, err := strconv.Atoi(config.SMTP.Port)
	if err != nil {
		return err
	}

	client, err := mail.NewClient(
		config.SMTP.Host,
		mail.WithPort(port),
		mail.WithSMTPAuth(mail.SMTPAuthPlain),
		mail.WithUsername(config.SMTP.Email),
		mail.WithPassword(config.SMTP.Password),
		mail.WithTLSPolicy(mail.TLSMandatory),
		mail.WithTLSConfig(&tls.Config{
			ServerName: config.SMTP.Host,
		}),
	)

	if err != nil {
		return err
	}

	return client.DialAndSend(message)
}
