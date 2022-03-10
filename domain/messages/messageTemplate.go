package messages

import "time"

type MessageTemplate struct {
	Service          string    `json:"service"`
	AlertName        string    `json:"alertName"`
	ProblemStartedAt time.Time `json:"problemStartedAt"`
	ProblemEndAt     time.Time `json:"problemEndAt"`
	Description      string    `json:"description"`
	Status           string    `json:"status"`
}
