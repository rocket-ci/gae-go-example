package main

import (
	"testing"
	"strings"
	"net/url"
	"net/http/httptest"
)

func TestNewCheckinDispatcher(t *testing.T) {
	publisher := NewSlack("TOKEN")
	values := url.Values{}
	values.Set("checkin", `{"id":"57f881f6498ee39652524105","createdAt":1475903990,"type":"checkin","private":true,"visibility":"private","timeZoneOffset":540,"user":{"id":"298030","firstName":"numanuma08","gender":"male","relationship":"self","photo":{"prefix":"https:\/\/irs3.4sqi.net\/img\/user\/","suffix":"\/298030-QHG3PODZXNT3GHML.jpg"},"isAnonymous":false},"venue":{"id":"4b0587a1f964a520989d22e3","name":"サンシャインシティ","contact":{"phone":"0338983331","formattedPhone":"03-3898-3331","facebook":"163995166958353","facebookUsername":"sunshinecity.spot","facebookName":"Sunshine City"},"location":{"address":"東池袋3-1-1","lat":35.72953903034779,"lng":139.71790797507603,"postalCode":"170-8630","cc":"JP","city":"豊島区","state":"東京都","country":"日本","formattedAddress":["東池袋3-1-1","豊島区, 東京都","170-8630"]},"categories":[{"id":"4bf58dd8d48988d1fd941735","name":"ショッピングモール","pluralName":"ショッピングモール","shortName":"モール","icon":{"prefix":"https:\/\/ss3.4sqi.net\/img\/categories_v2\/shops\/mall_","suffix":".png"},"primary":true}],"verified":false,"stats":{"checkinsCount":53686,"usersCount":24021,"tipCount":23},"url":"http:\/\/www.sunshinecity.co.jp","beenHere":{"unconfirmedCount":0,"marked":false,"lastCheckinExpiredAt":0}}}`)
	request := httptest.NewRequest("POST", "http://localhost:8080", strings.NewReader(values.Encode()))
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	dispatcher, err := NewCheckinDispatcher(request, publisher)
	noError(t, err)
	if dispatcher == nil{
		t.Fail()
	}
}

type MockPublisher struct {
	Published bool
}

func (mock *MockPublisher) PublishCheckin(message string) {
	mock.Published = true
}

func TestCheckinDispatcher_Should_Dispatch(t *testing.T) {
	publisher := &MockPublisher{false}
	values := url.Values{}
	values.Set("checkin", `{"id":"57f881f6498ee39652524105","createdAt":1475903990,"type":"checkin","private":true,"visibility":"private","timeZoneOffset":540,"user":{"id":"298030","firstName":"numanuma08","gender":"male","relationship":"self","photo":{"prefix":"https:\/\/irs3.4sqi.net\/img\/user\/","suffix":"\/298030-QHG3PODZXNT3GHML.jpg"},"isAnonymous":false},"venue":{"id":"4b0587a1f964a520989d22e3","name":"合同会社コベリン","contact":{"phone":"0338983331","formattedPhone":"03-3898-3331","facebook":"163995166958353","facebookUsername":"sunshinecity.spot","facebookName":"Sunshine City"},"location":{"address":"東池袋3-1-1","lat":35.72953903034779,"lng":139.71790797507603,"postalCode":"170-8630","cc":"JP","city":"豊島区","state":"東京都","country":"日本","formattedAddress":["東池袋3-1-1","豊島区, 東京都","170-8630"]},"categories":[{"id":"4bf58dd8d48988d1fd941735","name":"ショッピングモール","pluralName":"ショッピングモール","shortName":"モール","icon":{"prefix":"https:\/\/ss3.4sqi.net\/img\/categories_v2\/shops\/mall_","suffix":".png"},"primary":true}],"verified":false,"stats":{"checkinsCount":53686,"usersCount":24021,"tipCount":23},"url":"http:\/\/www.sunshinecity.co.jp","beenHere":{"unconfirmedCount":0,"marked":false,"lastCheckinExpiredAt":0}}}`)
	request := httptest.NewRequest("POST", "http://localhost:8080", strings.NewReader(values.Encode()))
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	dispatcher, err := NewCheckinDispatcher(request, publisher)
	noError(t, err)
	dispatcher.Dispatch()
	if !publisher.Published {
		t.Fail()
	}
}

func TestCheckinDispatcher_Should_Not_Dispatch(t *testing.T) {
	publisher := &MockPublisher{false}
	values := url.Values{}
	values.Set("checkin", `{"id":"57f881f6498ee39652524105","createdAt":1475903990,"type":"checkin","private":true,"visibility":"private","timeZoneOffset":540,"user":{"id":"298030","firstName":"numanuma08","gender":"male","relationship":"self","photo":{"prefix":"https:\/\/irs3.4sqi.net\/img\/user\/","suffix":"\/298030-QHG3PODZXNT3GHML.jpg"},"isAnonymous":false},"venue":{"id":"4b0587a1f964a520989d22e3","name":"べべべべ","contact":{"phone":"0338983331","formattedPhone":"03-3898-3331","facebook":"163995166958353","facebookUsername":"sunshinecity.spot","facebookName":"Sunshine City"},"location":{"address":"東池袋3-1-1","lat":35.72953903034779,"lng":139.71790797507603,"postalCode":"170-8630","cc":"JP","city":"豊島区","state":"東京都","country":"日本","formattedAddress":["東池袋3-1-1","豊島区, 東京都","170-8630"]},"categories":[{"id":"4bf58dd8d48988d1fd941735","name":"ショッピングモール","pluralName":"ショッピングモール","shortName":"モール","icon":{"prefix":"https:\/\/ss3.4sqi.net\/img\/categories_v2\/shops\/mall_","suffix":".png"},"primary":true}],"verified":false,"stats":{"checkinsCount":53686,"usersCount":24021,"tipCount":23},"url":"http:\/\/www.sunshinecity.co.jp","beenHere":{"unconfirmedCount":0,"marked":false,"lastCheckinExpiredAt":0}}}`)
	request := httptest.NewRequest("POST", "http://localhost:8080", strings.NewReader(values.Encode()))
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	dispatcher, err := NewCheckinDispatcher(request, publisher)
	noError(t, err)
	dispatcher.Dispatch()
	if publisher.Published {
		t.Fail()
	}

}