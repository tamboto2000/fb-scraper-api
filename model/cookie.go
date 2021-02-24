package model

import "github.com/tamboto2000/facebook"

// Cookie for storing or get a cookie
type Cookie struct {
	ID        string `json:"id"`
	Cookie    string `json:"cookie"`
	Label     string `json:"label,omitempty"`
	CreatedAt string `json:"createdAt"`
	fbClient  *facebook.Facebook
}

// Client get client for requesting Facebook APIs
func (c *Cookie) Client() *facebook.Facebook {
	return c.fbClient
}

// SetClient set facebook client for requesting Facebook APIs
func (c *Cookie) SetClient(cl *facebook.Facebook) {
	c.fbClient = cl
}
