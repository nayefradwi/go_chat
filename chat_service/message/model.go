package message

import "time"

const (
	TEXT  = iota
	IMAGE = iota
	VIDEO = iota
	FILE  = iota
	URL   = iota
)

type Message struct {
	// todo: missing id
	// todo: missing chat id (index)
	SentBy      int       `json:"sentBy"`
	Content     string    `json:"content"`
	ContentType string    `json:"contentType"`
	SentAt      time.Time `json:"sentAt"`
	IsRead      bool      `json:"isRead"`
	ReadAt      time.Time `json:"readAt"`
}
