package main

import (
	"context"
	"delete-issues/config"
	"github.com/machinebox/graphql"
	"log"
	"strconv"
	"time"
)

func main() {
	config.Load()
	client := graphql.NewClient("https://api.github.com/graphql")
	username := config.Get("username")
	repo := config.Get("repo")
	token := config.Get("token")
	from, _ := strconv.ParseUint(config.Get("issueFrom"), 10, 32)
	till, _ := strconv.ParseUint(config.Get("issueTill"), 10, 32)

	for from < till+1 {
		get_id_req := graphql.NewRequest(`
		query ($username: String!, $repo: String!, $issueNumber: Int!) {
			repository(owner: $username, name: $repo) {
   				 issue(number: $issueNumber) {
   			   		id
				 }
			}
		}
	`)

		delete_issue_req := graphql.NewRequest(`
			mutation($issueID: ID!) {
			  deleteIssue(input: { issueId: $issueID, clientMutationId: "auto-delete-spam-issues" }) {
				repository {
				  id
				}
			  }
			}
		`)

		get_id_req.Var("username", username)
		get_id_req.Var("repo", repo)
		get_id_req.Var("issueNumber", from)
		get_id_req.Header.Set("Authorization", "token "+token)

		var issue_id_resp IssueIDResponse

		if err := client.Run(context.Background(), get_id_req, &issue_id_resp); err != nil {
			log.Fatal(err)
		}

		delete_issue_req.Var("issueID", issue_id_resp.Repository.Issue.Id)
		delete_issue_req.Header.Set("Authorization", "token "+token)

		if er := client.Run(context.Background(), delete_issue_req, nil); er != nil {
			log.Fatal(er)
		}

		println("Deleted Issue #" + strconv.Itoa(int(from)))
		from = from + 1
		time.Sleep(1000 * time.Millisecond)
	}
}
