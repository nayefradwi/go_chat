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
	SentBy      int
	Content     string
	ContentType string
	SentAt      time.Time
	IsRead      bool
}
