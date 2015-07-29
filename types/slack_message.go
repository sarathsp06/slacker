package types

import "time"

//SlackMessage formatof the SlackMessage outgoing webhook
type SlackMessage struct {
	Text      string
	User      string
	Channel   string
	Timestamp *time.Time
}
