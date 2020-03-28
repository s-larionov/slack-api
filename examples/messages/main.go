package main

import (
	"fmt"

	log "github.com/sirupsen/logrus"

	"slack"
	"slack/events/base"
	"slack/events/message"
)

func main() {
	api := slack_api.NewEventsAPI(":9090", "/slack")

	events := api.Subscribe(base.TypeMessage)
	go func() {
		for event := range events {
			msg, ok := event.Event.(*message.Message)
			if !ok {
				log.Error("unable to cast event to *message.Message")
				continue
			}
			fmt.Println(msg.Text)
		}
	}()

	_ = api.ListenAndServe()
}
