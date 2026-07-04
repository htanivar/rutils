# Email Package Documentation

## Overview

The `email` package provides high-level functionality for sending emails via SMTP. It supports plain text, HTML alternatives, multiple recipients (To, CC, BCC), attachments, and connection-reuse for batch sends. Crucially, it features modern Go support for `context.Context` cancellation and timeout control.

## API Reference

### Configuration

#### Config
Configures the SMTP client.
```go
type Config struct {
	Host     string // SMTP server host (e.g. "smtp.gmail.com")
	Port     int    // SMTP server port (e.g. 587)
	Username string // Username/email for auth
	Password string // Password or App Password
	From     string // Sender email address (From header)
}
```

---

### Client

#### NewClient
Creates a new email client config.
```go
func NewClient(cfg *Config) *Client
```

#### NewClientWithDialer
Creates a client using a mockable or custom `dialer` interface (useful for unit tests).
```go
func NewClientWithDialer(from string, d dialer) *Client
```

#### Send
Sends a single email message using a background context.
```go
func (c *Client) Send(msg *Message) error
```

#### SendWithContext
Sends a single email message with context cancellation and timeout support.
```go
func (c *Client) SendWithContext(ctx context.Context, msg *Message) error
```

#### SendBatch
Sends multiple messages in a batch. Reuses the SMTP connection.
```go
func (c *Client) SendBatch(msgs []*Message) error
```

#### SendBatchWithContext
Sends multiple messages in a batch with context cancellation support.
```go
func (c *Client) SendBatchWithContext(ctx context.Context, msgs []*Message) error
```

---

### Message

#### Message Struct
Defines the content and recipients of an email.
```go
type Message struct {
	To          []string // Primary recipients
	CC          []string // Carbon-copy recipients
	BCC         []string // Blind carbon-copy recipients
	Subject     string   // Subject line
	Body        string   // Plain text body alternative
	HTMLBody    string   // HTML body content
	Attachments []string // File paths of attachments to include
}
```

---

## Code Examples

### Sending an Email with Timeout
```go
package main

import (
	"context"
	"log"
	"time"

	"github.com/htanivar/rutils/email"
)

func main() {
	client := email.NewClient(&email.Config{
		Host:     "smtp.mailtrap.io",
		Port:     2525,
		Username: "my_username",
		Password: "my_password",
		From:     "sender@example.com",
	})

	msg := &email.Message{
		To:       []string{"recipient@example.com"},
		Subject:  "Hello!",
		Body:     "This is the plain text fallback.",
		HTMLBody: "<h1>Hello!</h1><p>This is the HTML body.</p>",
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := client.SendWithContext(ctx, msg); err != nil {
		log.Fatalf("failed to send: %v", err)
	}
}
```

### Sending a Batch of Emails
```go
package main

import (
	"context"
	"log"

	"github.com/htanivar/rutils/email"
)

func main() {
	client := email.NewClient(&email.Config{
		Host:     "smtp.mailtrap.io",
		Port:     2525,
		Username: "my_username",
		Password: "my_password",
		From:     "sender@example.com",
	})

	msgs := []*email.Message{
		{
			To:      []string{"user1@example.com"},
			Subject: "Digest 1",
			Body:    "Digest details for user 1",
		},
		{
			To:      []string{"user2@example.com"},
			Subject: "Digest 2",
			Body:    "Digest details for user 2",
		},
	}

	if err := client.SendBatchWithContext(context.Background(), msgs); err != nil {
		log.Fatalf("failed to send batch: %v", err)
	}
}
```
