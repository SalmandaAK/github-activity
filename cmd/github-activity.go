package main

import (
	"fmt"
	"os"

	"github.com/SalmandaAK/github-user-activity/internal/handler"
)

func main() {
	username, err := handler.ParseInput(os.Args)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	handler.FetchGithubUserEvent(username)
}
