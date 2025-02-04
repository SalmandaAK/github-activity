package view

import (
	"fmt"
	"strings"

	"github.com/SalmandaAK/github-user-activity/internal/model"
)

// eventMessageFunc is a function which formats the event and prints it to terminal.
type eventMessageFunc func(*model.UserEvent) string

var eventMessageMap = map[string]eventMessageFunc{
	"CommitCommentEvent":            commitCommentEventMessage,
	"CreateEvent":                   createEventMessage,
	"DeleteEvent":                   deleteEventMessage,
	"ForkEvent":                     forkEventMessage,
	"GollumEvent":                   gollumEventMessage,
	"IssueCommentEvent":             issueCommentEvent,
	"IssuesEvent":                   issuesEventMessage,
	"MemberEvent":                   memberEventMessage,
	"PublicEvent":                   publicEventMessage,
	"PullRequestEvent":              pullRequestEventMessage,
	"PullRequestReviewEvent":        pullRequestReviewEventMessage,
	"PullRequestReviewCommentEvent": pullRequestReviewCommentEventMessage,
	"PullRequestReviewThreadEvent":  pullRequestReviewThreadEventMessage,
	"PushEvent":                     pushEventMessage,
	"ReleaseEvent":                  releaseEventMessage,
	"SponsorshipEvent":              sponsorshipEvent,
	"WatchEvent":                    watchEventMessage,
}

// Payload.Action: created
func commitCommentEventMessage(ue *model.UserEvent) string {
	return fmt.Sprintf("%s a commit comment in %s %s", ue.Payload.Action, ue.Repo.Name, formatTimePassed(ue.CreatedAt))
}

// Payload.Reftype: branch, tag, or repository
func createEventMessage(ue *model.UserEvent) string {
	return fmt.Sprintf("created a %s in %s %s", ue.Payload.RefType, ue.Repo.Name, formatTimePassed(ue.CreatedAt))
}

// Payload.Reftype: branch or tag
func deleteEventMessage(ue *model.UserEvent) string {
	return fmt.Sprintf("deleted a %s in %s %s", ue.Payload.RefType, ue.Repo.Name, formatTimePassed(ue.CreatedAt))
}

func forkEventMessage(ue *model.UserEvent) string {
	return fmt.Sprintf("created a fork in %s %s", ue.Repo.Name, formatTimePassed(ue.CreatedAt))
}

// Payload.PagesAction: created or edited
func gollumEventMessage(ue *model.UserEvent) string {
	return fmt.Sprintf("%s wiki page in %s %s", ue.Payload.PagesAction, ue.Repo.Name, formatTimePassed(ue.CreatedAt))
}

// Payload.Action: created, edited, or deleted
func issueCommentEvent(ue *model.UserEvent) string {
	return fmt.Sprintf("%s issue comment in issue #%d in %s %s", ue.Payload.Action, ue.Payload.Issue.Number, ue.Repo.Name, formatTimePassed(ue.CreatedAt))
}

// Payload.Action: opened, edited, closed, reopened, assigned, unassigned, labeled, or unlabeled
func issuesEventMessage(ue *model.UserEvent) string {
	return fmt.Sprintf("%s issue #%d in %s %s", ue.Payload.Action, ue.Payload.Issue.Number, ue.Repo.Name, formatTimePassed(ue.CreatedAt))
}

// Payload.Action: created or edited, Payload.Member.Login: user that added
func memberEventMessage(ue *model.UserEvent) string {
	var msg string
	if ue.Payload.Action == "created" {
		msg = fmt.Sprintf("%s %s to be a collaborator in %s %s", ue.Payload.Action, ue.Payload.Member.Login, ue.Repo.Name, formatTimePassed(ue.CreatedAt))
	} else if ue.Payload.Action == "edited" {
		msg = fmt.Sprintf("%s %s collaborator permissions in %s %s", ue.Payload.Action, ue.Payload.Member.Login, ue.Repo.Name, formatTimePassed(ue.CreatedAt))
	}
	return msg
}

func publicEventMessage(ue *model.UserEvent) string {
	return fmt.Sprintf("has made %s become public %s", ue.Repo.Name, formatTimePassed(ue.CreatedAt))
}

// Payload.Action: opened, edited, closed, reopened, assigned, unassigned, review_requested, review_request_removed, labeled, unlabeled, or syncronize
func pullRequestEventMessage(ue *model.UserEvent) string {
	return fmt.Sprintf("%s pull request #%d in %s %s", strings.ReplaceAll(ue.Payload.Action, "_", " "), ue.Payload.PullRequest.Number, ue.Repo.Name, formatTimePassed(ue.CreatedAt))
}

// Payload.Action: created
func pullRequestReviewEventMessage(ue *model.UserEvent) string {
	return fmt.Sprintf("%s a pull request review in pull request #%d in %s %s", ue.Payload.Action, ue.Payload.PullRequest.Number, ue.Repo.Name, formatTimePassed(ue.CreatedAt))
}

// Payload.Action: created
func pullRequestReviewCommentEventMessage(ue *model.UserEvent) string {
	return fmt.Sprintf("%s a pull request review comment in pull request #%d in %s %s", ue.Payload.Action, ue.Payload.PullRequest.Number, ue.Repo.Name, formatTimePassed(ue.CreatedAt))
}

// Payload.Action: resolved or unresolved.
func pullRequestReviewThreadEventMessage(ue *model.UserEvent) string {
	return fmt.Sprintf("marked pull request #%d as %s in %s %s", ue.Payload.PullRequest.Number, ue.Payload.Action, ue.Repo.Name, formatTimePassed(ue.CreatedAt))
}

func pushEventMessage(ue *model.UserEvent) string {
	if ue.Payload.Size == 1 {
		return fmt.Sprintf("pushed %d commit to %s %s", ue.Payload.Size, ue.Repo.Name, formatTimePassed(ue.CreatedAt))
	}
	return fmt.Sprintf("pushed %d commits to %s %s", ue.Payload.Size, ue.Repo.Name, formatTimePassed(ue.CreatedAt))
}

// Payload.Action: published
func releaseEventMessage(ue *model.UserEvent) string {
	return fmt.Sprintf("%s %s %s", ue.Payload.Action, ue.Repo.Name, formatTimePassed(ue.CreatedAt))
}

// Payload.Action: created
func sponsorshipEvent(ue *model.UserEvent) string {
	return fmt.Sprintf("%s sponsorship in %s %s", ue.Payload.Action, ue.Repo.Name, formatTimePassed(ue.CreatedAt))
}

func watchEventMessage(ue *model.UserEvent) string {
	return fmt.Sprintf("starred %s %s", ue.Repo.Name, formatTimePassed(ue.CreatedAt))
}
