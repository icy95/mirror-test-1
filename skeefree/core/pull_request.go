package core

import (
	"time"
	"gopkg.in/go-playground/webhooks.v5/github"
)

type PullRequestStatus string

const (
	PullRequestStatusDetected  PullRequestStatus = "detected"
	PullRequestStatusQueued    PullRequestStatus = "queued"
	PullRequestStatusCancelled PullRequestStatus = "cancelled"
	PullRequestStatusComplete  PullRequestStatus = "complete"
	PullRequestStatusUnknown   PullRequestStatus = "unknown"
)

type PullRequest struct {
	Id                           int64               `db:"id" json:"id"`
	Number                       int                 `db:"pull_request_number" json:"pull_request_number"`
	Title                        string              `db:"title" json:"title"`
	TimeAdded                    time.Time           `db:"added_timestamp" json:"added_timestamp"`
}


func NewPullRequest(payload *github.PullRequestPayload) *PullRequest {
	return &PullRequest{
		Number: int(payload.PullRequest.Number),
		Title: payload.PullRequest.Title,
	}
}