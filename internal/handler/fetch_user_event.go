package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/SalmandaAK/github-user-activity/internal/model"
	"github.com/SalmandaAK/github-user-activity/internal/view"
)

var (
	errUserNotFound       = errors.New("user not found")
	errForbidden          = errors.New("access forbidden")
	errServiceUnavailable = errors.New("service unavailable")
)

func FetchGithubUserEvent(username string) {
	resp, err := http.Get(fmt.Sprintf("http://api.github.com/users/%s/events", username))
	if err != nil {
		fmt.Printf("Error fetching data: %v", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == 404 {
		fmt.Printf("Error: %v: %v.\n", errUserNotFound, username)
		return
	}

	if resp.StatusCode == 403 {
		fmt.Printf("Error: %v.\n", errForbidden)
		return
	}

	if resp.StatusCode == 503 {
		fmt.Printf("Error: %v.\n", errServiceUnavailable)
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response: %v", err)
		return
	}

	var events []*model.UserEvent
	err = json.Unmarshal(body, &events)
	if err != nil {
		fmt.Printf("Error unmarshaling JSON: %v", err)
		return
	}

	view.DisplayUserEventMessage(events, username)
}
