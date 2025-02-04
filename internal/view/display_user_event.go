package view

import (
	"fmt"
	"strings"

	"github.com/SalmandaAK/github-user-activity/internal/model"
)

func DisplayUserEventMessage(events []*model.UserEvent, username string) {
	// Print events to os.Stdout
	// "past 90 days" because only events created within the past 90 days can be fetched.
	if len(events) == 0 {
		fmt.Printf("%s has no activity for the past 90 days\n", username)
		return
	}

	if len(events) == 1 {
		fmt.Printf("Found \x1b[1;33m%d\x1B[0m recent activity for user \x1b[1;33m%s\x1B[0m.\n", len(events), username)
	} else {
		fmt.Printf("Found \x1b[1;33m%d\x1B[0m recent activities for user \x1b[1;33m%s\x1B[0m.\n", len(events), username)
	}
	for _, event := range events {
		msgFunc, ok := eventMessageMap[event.Type]
		if ok {
			msg := msgFunc(event)
			msg = strings.Replace(msg, string(msg[0]), strings.ToUpper(string(msg[0])), 1)
			fmt.Printf("- %v\n", msg)
		}
	}
}
