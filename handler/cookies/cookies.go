package cookies

import "sync"

// Cookies contains handler functions for store and delete cookies
type Cookies struct {
	Cookies *sync.Map
}

// NewCookies initiate Cookies handler
func NewCookies() *Cookies {
	return &Cookies{Cookies: new(sync.Map)}
}
