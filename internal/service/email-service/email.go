package emailservice

import (
	"fmt"
	"html/template"
	"log"

	"github.com/wneessen/go-mail"
)

type Mailer struct {
	client *mail.Client
	from   string
}

type MailConfig struct {
	Bcc          *[]string
	Cc           *[]string
	To           string
	RepylTo      *string
	Message      *string
	Importance   *mail.Importance
	Format       mail.ContentType
	Template     *template.Template
	TemplateData any
}

func New(host string, port int, user, pass, from string) (*Mailer, error) {
	c, err := mail.NewClient(host,
		mail.WithPort(port),
		mail.WithSMTPAuth(mail.SMTPAuthPlain),
		mail.WithUsername(user),
		mail.WithPassword(pass),
	)
	if err != nil {
		return nil, err
	}

	return &Mailer{
		client: c,
		from:   from,
	}, nil
}

func (m *Mailer) SendEmail(cfg *MailConfig) error {
	msg := mail.NewMsg()

	if err := msg.From(m.from); err != nil {
		return fmt.Errorf("invalid from address: %w", err)
	}

	if cfg.Importance != nil {
		msg.SetImportance(*cfg.Importance)
	}

	if cfg.Bcc != nil {
		for _, bcc := range *cfg.Bcc {
			if err := msg.AddBcc(bcc); err != nil {
				log.Fatalf("failed to set BCC address: %s", err)
			}
		}

	}

	if cfg.Cc != nil {
		for _, cc := range *cfg.Cc {
			if err := msg.AddCc(cc); err != nil {
				log.Fatalf("failed to set BCC address: %s", err)
			}
		}

	}

	if cfg.RepylTo != nil {
		msg.SetAddrHeader("Reply-To", "replies@example.com")
	}

	if cfg.Template != nil {
		if err := msg.SetBodyHTMLTemplate(cfg.Template, cfg.TemplateData); err != nil {
			log.Fatalf("failed to set body: %s", err)
		}

	}

	if cfg.Message != nil {
		msg.SetBodyString(cfg.Format, *cfg.Message)
	}
	if err := msg.To(cfg.To); err != nil {
		log.Fatalf("failed to send to address: %s", err)
	}

	return m.client.DialAndSend(msg)
}


