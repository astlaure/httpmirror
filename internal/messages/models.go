package messages

import "time"

type (
	ProxyRequest struct {
		ID        uint
		Service   string
		Tracking  string
		CreatedAt time.Time `db:"created_at"`
		UpdatedAt time.Time `db:"updated_at"`
	}

	ProxyMessage struct {
		ID        uint
		RequestID uint `db:"request_id"`
		Type      string
		Status    uint
		Path      string
		Protocol  string
		Headers   string
		Body      string
		CreatedAt time.Time `db:"created_at"`
		UpdatedAt time.Time `db:"updated_at"`
	}
)
