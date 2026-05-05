# Email Package Documentation

## Overview

The email package provides functionality for sending emails via SMTP with support for attachments, HTML content, and various email features. It follows a simple client-message pattern for sending emails.

## Key Components

### Email Client

The `Client` type handles the SMTP connection and email sending logic.

#### NewClient
```go
c := email.NewClient(host, port, username, password)
```

Creates a new email client with the specified SMTP server credentials.

**Parameters:**
- `host`: SMTP server hostname (e.g., "smtp.gmail.com")
- `port`: SMTP server port (e.g., 587)
- `username`: SMTP username
- `password`: SMTP password

#### Send
```go
err := client.Send(msg)
```

Sends an email message through the SMTP server.

**Parameters:**
- `msg`: Pointer to an email.Message struct

**Returns:**
- `error`: nil if successful, error if failed to send

### Email Message

The `Message` struct defines the content and recipients of an email.

```go
type Message struct {
	From       string
	To         []string
	CC         []string
	BCC        []string
	ReplyTo    []string
	Subject    string
	Body       string
	HTMLBody   string
	Attachments []File
}
```

**Fields:**
- `From`: Sender email address
- `To`: Recipient email addresses (slice)
- `CC`: Carbon copy recipients (slice)
- `BCC`: Blind carbon copy recipients (slice)
- `ReplyTo`: Reply-to addresses (slice)
- `Subject`: Email subject line
- `Body`: Plain text body content
- `HTMLBody`: HTML body content
- `Attachments`: Slice of File structs for file attachments

### File Attachment

The `File` struct represents a file attachment.

```go
type File struct {
	Name string
	Data []byte
}
```

**Fields:**
- `Name`: Name of the file as it will appear in the email
- `Data`: Raw byte data of the file

### Message Creation Functions

Utility functions for creating common message types:

#### NewPlainTextMessage
```go
msg := email.NewPlainTextMessage(to, subject, body)
```

Creates a new message with plain text content.

**Parameters:**
- `to`: Recipient email address
- `subject`: Email subject
- `body`: Plain text body content

**Returns:**
- `*Message`: Pointer to a new Message struct

#### NewHTMLMessage
```go
msg := email.NewHTMLMessage(to, subject, html)
```

Creates a new message with HTML content.

**Parameters:**
- `to`: Recipient email address
- `subject`: Email subject
- `html`: HTML body content

**Returns:**
- `*Message`: Pointer to a new Message struct

## Usage Examples

### Sending a Plain Text Email

```go
// Create email client
c := email.NewClient("smtp.gmail.com", 587, "user@gmail.com", "password")

// Create plain text message
msg := email.NewPlainTextMessage("recipient@example.com", "Test Subject", "Hello World!")

// Send the email
if err := c.Send(msg); err != nil {
	log.Fatal(err)
}
```

### Sending an HTML Email

```go
// Create email client
c := email.NewClient("smtp.gmail.com", 587, "user@gmail.com", "password")

// Create HTML message
htmlContent := "<h1>Hello World!</h1><p>This is an HTML email!</p>"
msg := email.NewHTMLMessage("recipient@example.com", "Test HTML Email", htmlContent)

// Send the email
if err := c.Send(msg); err != nil {
	log.Fatal(err)
}
```

### Sending an Email with Attachment

```go
// Read file data
fileData, err := os.ReadFile("/path/to/document.pdf")
if err != nil {
	log.Fatal(err)
}

// Create email client
c := email.NewClient("smtp.gmail.com", 587, "user@gmail.com", "password")

// Create message
msg := email.NewPlainTextMessage("recipient@example.com", "Document Attached", "Please find the document attached.")

// Add attachment
msg.Attachments = []email.File{
	{
		Name: "document.pdf",
		Data: fileData,
	},
}

// Send the email
if err := c.Send(msg); err != nil {
	log.Fatal(err)
}
```

### Sending an Email with Multiple Recipients

```go
// Create email client
c := email.NewClient("smtp.gmail.com", 587, "user@gmail.com", "password")

// Create message with multiple recipients
msg := &email.Message{
	To:    []string{"recipient1@example.com", "recipient2@example.com"},
	CC:    []string{"cc@example.com"},
	BCC:   []string{"bcc@example.com"},
	ReplyTo: []string{"replyto@example.com"},
	Subject: "Multiple Recipients",
	Body:    "This email has multiple recipients.",
}

// Send the email
if err := c.Send(msg); err != nil {
	log.Fatal(err)
}
```

## Testing

The email package includes comprehensive tests that mock the SMTP client for testing without actually sending emails. Tests cover:
- Plain text messages
- HTML messages
- Email with attachments
- Various error conditions
- Mock SMTP server responses

Run tests with:
```bash
make test-email
```

## Error Handling

The email package handles various error conditions:
- Network connectivity issues
- Authentication failures
- Invalid email addresses
- Message formatting issues

Errors are returned from the Send method and should be checked by the caller.

## Dependencies

The email package uses the following external dependency:
- `gopkg.in/mail.v2`: For SMTP email sending functionality

This dependency is specified in the go.mod file and will be automatically downloaded when the package is imported.

## Security Considerations

- SMTP credentials should be stored securely (e.g., environment variables, secret management systems)
- Consider using app-specific passwords for services like Gmail
- Validate recipient email addresses to prevent email injection attacks
- Sanitize HTML content to prevent XSS attacks in HTML emails

## Performance

- The client reuses SMTP connections when possible
- Large attachments may impact performance due to memory usage
- Consider chunking large files or using file streaming for very large attachments

## Future Improvements

- Add support for OAuth2 authentication
- Implement connection pooling for high-volume email sending
- Add support for email templates
- Implement rate limiting to prevent spam
- Add DKIM/SPF support for improved email deliverability
