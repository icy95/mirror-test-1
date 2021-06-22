package app

import (
	"fmt"
	"strings"

	"github.com/startupheroes/kubebot/skeefree/core"
	"github.com/startupheroes/kubebot/skeefree/db"
	"gopkg.in/go-playground/webhooks.v5/github"
)



func HandleIssueCommentPayload(prNumber int, comment string) {

	if comment == "!migrate" {
		//check if pr approved and mergeable
		//add pull request with status queued
		//check if any running or queued pr's in db
		// if there is, comment:"queued"
		// if there is not, trigger workflow, set status running
	}

	if strings.Contains(comment, "-- skeema:diff") {
		info := core.ParseSkeemaDiffStatements(comment)
		for _, statement := range info.Statements {
			migration := core.NewMigration(info.SchemaName, statement)
			fmt.Println(migration.PRStatement)
			fmt.Printf("----------------------")
		}
	}
}

func HandlePullRequestPayload(payload *github.PullRequestPayload) {
	pr := core.NewPullRequest(payload)
	db := db.NewBackend()
	db.SubmitPR(pr)
}