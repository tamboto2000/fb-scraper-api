package main

import (
	"github.com/tamboto2000/fb-scraper-api/handler/cookies"
	"github.com/tamboto2000/fb-scraper-api/handler/profile"
)

var cookiesHandler = cookies.NewCookies()
var profileHandler = profile.NewProfile(cookiesHandler)
