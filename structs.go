package main

type IssueIDResponse struct {
	Repository Repository `json:"repository"`
}

type Repository struct {
	Issue Issue `json:"issue"`
}

type Issue struct {
	Id string `json:"id"`
}
