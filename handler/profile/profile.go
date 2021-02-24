package profile

import "github.com/tamboto2000/fb-scraper-api/handler/cookies"

// Profile contains handler functions for retrieving profile data
type Profile struct {
	cookieHandler *cookies.Cookies
}

// NewProfile init Profile handler
func NewProfile(cookieHandler *cookies.Cookies) *Profile {
	return &Profile{cookieHandler: cookieHandler}
}
