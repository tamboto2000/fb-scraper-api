package model

import "github.com/tamboto2000/facebook"

// Profile contains profile data and errors on retrieving data
type Profile struct {
	Errors []string `json:"errors"`
	*facebook.Profile
}
