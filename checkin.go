package main

type Checkin struct {
	ID string `json:"id"`
	CreatedAt int `json:"createdAt"`
	Type string `json:"type"`
	Private bool `json:"private"`
	Visibility string `json:"visibility"`
	TimeZoneOffset int `json:"timeZoneOffset"`
	User struct {
		   ID string `json:"id"`
		   FirstName string `json:"firstName"`
		   Gender string `json:"gender"`
		   Relationship string `json:"relationship"`
		   Photo struct {
			      Prefix string `json:"prefix"`
			      Suffix string `json:"suffix"`
		      } `json:"photo"`
		   IsAnonymous bool `json:"isAnonymous"`
	   } `json:"user"`
	Venue struct {
		   ID string `json:"id"`
		   Name string `json:"name"`
		   Contact struct {
			      Phone string `json:"phone"`
			      FormattedPhone string `json:"formattedPhone"`
			      Facebook string `json:"facebook"`
			      FacebookUsername string `json:"facebookUsername"`
			      FacebookName string `json:"facebookName"`
		      } `json:"contact"`
		   Location struct {
			      Address string `json:"address"`
			      Lat float64 `json:"lat"`
			      Lng float64 `json:"lng"`
			      PostalCode string `json:"postalCode"`
			      Cc string `json:"cc"`
			      City string `json:"city"`
			      State string `json:"state"`
			      Country string `json:"country"`
			      FormattedAddress []string `json:"formattedAddress"`
		      } `json:"location"`
		   Categories []struct {
			   ID string `json:"id"`
			   Name string `json:"name"`
			   PluralName string `json:"pluralName"`
			   ShortName string `json:"shortName"`
			   Icon struct {
				      Prefix string `json:"prefix"`
				      Suffix string `json:"suffix"`
			      } `json:"icon"`
			   Primary bool `json:"primary"`
		   } `json:"categories"`
		   Verified bool `json:"verified"`
		   Stats struct {
			      CheckinsCount int `json:"checkinsCount"`
			      UsersCount int `json:"usersCount"`
			      TipCount int `json:"tipCount"`
		      } `json:"stats"`
		   URL string `json:"url"`
		   BeenHere struct {
			      UnconfirmedCount int `json:"unconfirmedCount"`
			      Marked bool `json:"marked"`
			      LastCheckinExpiredAt int `json:"lastCheckinExpiredAt"`
		      } `json:"beenHere"`
	   } `json:"venue"`
}