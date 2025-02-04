package model

type Payload struct {
	Action      string      `json:"action"`
	RefType     string      `json:"ref_type"`
	PagesAction string      `json:"pages[][action]"`
	Size        int         `json:"size"`
	PullRequest PullRequest `json:"pull_request"`
	Member      Member      `json:"member"`
	Issue       Issue       `json:"issue"`
}
