package main

import (
	"net/http"
	"net/url"
	"encoding/json"
	"errors"
)

type CheckinDispatcher struct {
	Checkin Checkin
	Publisher Publisher
}

func NewCheckinDispatcher(request *http.Request, publisher Publisher) (*CheckinDispatcher, error) {
	if err := request.ParseForm(); err != nil {
		return nil, err
	}
	row := request.PostFormValue("checkin")
	if row == "" {
		return nil, errors.New("failed form chickin is empty")
	}
	j, err := url.QueryUnescape(row)
	if err != nil {
		return nil, errors.New("failed query unescape " + row)
	}
	var checkin Checkin
	err = json.Unmarshal([]byte(j), &checkin)
	if err != nil {
		return nil, errors.New(err.Error() + ", json " + j)
	}
	dispatcher := CheckinDispatcher{checkin, publisher}
	return &dispatcher, nil
}

func (dispatcher *CheckinDispatcher) Dispatch() {
	if dispatcher.Checkin.Venue.Name != "合同会社コベリン" {
		return
	}
	dispatcher.Publisher.PublishCheckin("おは")
}
