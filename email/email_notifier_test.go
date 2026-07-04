package email

import (
	"context"
	"errors"
	"io"
	"os"
	"testing"

	"gopkg.in/mail.v2"
)


type mockSendCloser struct {
	sendErr error
	sent    int
}

func (m *mockSendCloser) Send(from string, to []string, msg io.WriterTo) error {
	if m.sendErr != nil {
		return m.sendErr
	}
	// Force message serialization (catches some formatting/attachment issues)
	if _, err := msg.WriteTo(io.Discard); err != nil {
		return err
	}
	m.sent++
	return nil
}

func (m *mockSendCloser) Close() error { return nil }

type mockDialer struct {
	dialAndSendErr error
	dialErr        error
	sender         *mockSendCloser
	dialAndSendN   int
	dialN          int
}

func (m *mockDialer) DialAndSend(msgs ...*mail.Message) error {
	m.dialAndSendN++
	if m.dialAndSendErr != nil {
		return m.dialAndSendErr
	}
	// optional: serialize here too
	for _, mm := range msgs {
		if _, err := mm.WriteTo(io.Discard); err != nil {
			return err
		}
	}
	return nil
}

func (m *mockDialer) Dial() (mail.SendCloser, error) {
	m.dialN++
	if m.dialErr != nil {
		return nil, m.dialErr
	}
	if m.sender == nil {
		m.sender = &mockSendCloser{}
	}
	return m.sender, nil
}

func TestClient_Send(t *testing.T) {
	tmp, err := os.CreateTemp("", "attach-*.txt")
	if err != nil {
		t.Fatal(err)
	}
	_ = tmp.Close()
	t.Cleanup(func() { _ = os.Remove(tmp.Name()) })

	tests := []struct {
		name           string
		msg            *Message
		dialAndSendErr error
		wantErr        bool
	}{
		{
			name: "plain text",
			msg: &Message{
				To:      []string{"a@b.com"},
				Subject: "s",
				Body:    "body",
			},
		},
		{
			name: "html only",
			msg: &Message{
				To:       []string{"a@b.com"},
				Subject:  "s",
				HTMLBody: "<b>x</b>",
			},
		},
		{
			name: "html + alt text",
			msg: &Message{
				To:       []string{"a@b.com"},
				Subject:  "s",
				Body:     "plain",
				HTMLBody: "<b>x</b>",
			},
		},
		{
			name: "cc+bcc",
			msg: &Message{
				To:      []string{"a@b.com"},
				CC:      []string{"c@d.com"},
				BCC:     []string{"e@f.com"},
				Subject: "s",
				Body:    "body",
			},
		},
		{
			name: "with attachment",
			msg: &Message{
				To:          []string{"a@b.com"},
				Subject:     "s",
				Body:        "body",
				Attachments: []string{tmp.Name()},
			},
		},
		{
			name: "dialAndSend error bubbles up",
			msg: &Message{
				To:      []string{"a@b.com"},
				Subject: "s",
				Body:    "body",
			},
			dialAndSendErr: errors.New("boom"),
			wantErr:        true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			md := &mockDialer{dialAndSendErr: tt.dialAndSendErr}
			c := NewClientWithDialer("from@x.com", md)

			err := c.Send(tt.msg)
			if (err != nil) != tt.wantErr {
				t.Fatalf("err=%v wantErr=%v", err, tt.wantErr)
			}
			if md.dialAndSendN != 1 {
				t.Fatalf("DialAndSend called %d times, want 1", md.dialAndSendN)
			}
		})
	}
}

func TestClient_SendBatch(t *testing.T) {
	tests := []struct {
		name     string
		msgs     []*Message
		dialErr  error
		sendErr  error
		wantErr  bool
		wantSent int
	}{
		{
			name: "multiple",
			msgs: []*Message{
				{To: []string{"a@b.com"}, Subject: "1", Body: "x"},
				{To: []string{"c@d.com"}, Subject: "2", HTMLBody: "<p>y</p>"},
			},
			wantSent: 2,
		},
		{
			name: "empty batch still dials (current behavior)",
			msgs: []*Message{},
		},
		{
			name:    "dial error",
			msgs:    []*Message{{To: []string{"a@b.com"}, Subject: "1", Body: "x"}},
			dialErr: errors.New("dial fail"),
			wantErr: true,
		},
		{
			name:    "send error",
			msgs:    []*Message{{To: []string{"a@b.com"}, Subject: "1", Body: "x"}},
			sendErr: errors.New("send fail"),
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			md := &mockDialer{
				dialErr: tt.dialErr,
				sender:  &mockSendCloser{sendErr: tt.sendErr},
			}
			c := NewClientWithDialer("from@x.com", md)

			err := c.SendBatch(tt.msgs)
			if (err != nil) != tt.wantErr {
				t.Fatalf("err=%v wantErr=%v", err, tt.wantErr)
			}

			if tt.wantErr {
				return
			}

			if md.dialN != 1 {
				t.Fatalf("Dial called %d times, want 1", md.dialN)
			}
			if md.sender.sent != tt.wantSent {
				t.Fatalf("sent=%d want=%d", md.sender.sent, tt.wantSent)
			}
		})
	}
}

func TestClient_SendWithContext_Cancellation(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel() // Cancel immediately

	md := &mockDialer{}
	c := NewClientWithDialer("from@x.com", md)

	msg := &Message{
		To:      []string{"recipient@example.com"},
		Subject: "Test",
		Body:    "body",
	}

	err := c.SendWithContext(ctx, msg)
	if !errors.Is(err, context.Canceled) {
		t.Fatalf("expected context.Canceled error, got: %v", err)
	}
}

func TestClient_SendBatchWithContext_Cancellation(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel() // Cancel immediately

	md := &mockDialer{}
	c := NewClientWithDialer("from@x.com", md)

	msgs := []*Message{
		{To: []string{"recipient@example.com"}, Subject: "Test", Body: "body"},
	}

	err := c.SendBatchWithContext(ctx, msgs)
	if !errors.Is(err, context.Canceled) {
		t.Fatalf("expected context.Canceled error, got: %v", err)
	}
}

