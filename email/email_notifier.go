package email

import (
	"fmt"

	"gopkg.in/mail.v2"
)

type dialer interface {
	DialAndSend(msgs ...*mail.Message) error
	Dial() (mail.SendCloser, error)
}

type gomailDialer struct{ d *mail.Dialer }

func (g gomailDialer) DialAndSend(msgs ...*mail.Message) error { return g.d.DialAndSend(msgs...) }
func (g gomailDialer) Dial() (mail.SendCloser, error)          { return g.d.Dial() }

type Config struct {
	Host     string
	Port     int
	Username string
	Password string
	From     string
}

type Client struct {
	dialer dialer
	from   string
}

func NewClient(cfg *Config) *Client {
	d := mail.NewDialer(cfg.Host, cfg.Port, cfg.Username, cfg.Password)
	return &Client{
		dialer: gomailDialer{d: d},
		from:   cfg.From,
	}
}

// helper for tests / advanced use
func NewClientWithDialer(from string, d dialer) *Client {
	return &Client{from: from, dialer: d}
}

type Message struct {
	To          []string
	CC          []string
	BCC         []string
	Subject     string
	Body        string
	HTMLBody    string
	Attachments []string // file paths
}

func (c *Client) Send(msg *Message) error {
	m := mail.NewMessage()
	m.SetHeader("From", c.from)
	m.SetHeader("To", msg.To...)

	if len(msg.CC) > 0 {
		m.SetHeader("Cc", msg.CC...)
	}
	if len(msg.BCC) > 0 {
		m.SetHeader("Bcc", msg.BCC...)
	}

	m.SetHeader("Subject", msg.Subject)

	if msg.HTMLBody != "" {
		m.SetBody("text/html", msg.HTMLBody)
		if msg.Body != "" {
			m.AddAlternative("text/plain", msg.Body)
		}
	} else {
		m.SetBody("text/plain", msg.Body)
	}

	for _, path := range msg.Attachments {
		m.Attach(path)
	}

	if err := c.dialer.DialAndSend(m); err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}
	return nil
}

func (c *Client) SendBatch(msgs []*Message) error {
	sender, err := c.dialer.Dial()
	if err != nil {
		return fmt.Errorf("failed to dial: %w", err)
	}
	defer sender.Close()

	for _, msg := range msgs {
		m := mail.NewMessage()
		m.SetHeader("From", c.from)
		m.SetHeader("To", msg.To...)
		m.SetHeader("Subject", msg.Subject)

		if msg.HTMLBody != "" {
			m.SetBody("text/html", msg.HTMLBody)
		} else {
			m.SetBody("text/plain", msg.Body)
		}

		if err := mail.Send(sender, m); err != nil {
			return fmt.Errorf("failed to send batch email: %w", err)
		}
	}
	return nil
}
