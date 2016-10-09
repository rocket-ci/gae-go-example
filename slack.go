package main

import "github.com/nlopes/slack"

type Slack struct {
	Token string
}

func NewSlack(token string) *Slack {
	s := &Slack{token}
	return s
}

func (publisher *Slack) PublishCheckin(message string) {
	app := slack.New(publisher.Token)
	params := slack.NewPostMessageParameters()
	params.Username = "numa08"
	params.AsUser = true
	println("post message")
	app.PostMessage("timesheets", message, params)
	println("posted")
}

