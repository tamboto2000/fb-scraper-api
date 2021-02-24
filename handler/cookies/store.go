package cookies

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/tamboto2000/facebook"
	"github.com/tamboto2000/fb-scraper-api/model"
	"github.com/tamboto2000/random"
)

// Store store cookie
func (c *Cookies) Store(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	cookie := new(model.Cookie)
	if err := json.NewDecoder(r.Body).Decode(cookie); err != nil {
		if err == io.EOF {
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(model.HTTPResponse{
				Code:    400,
				Message: "request body is empty",
			})

			return
		}

		w.WriteHeader(400)
		json.NewEncoder(w).Encode(model.HTTPResponse{
			Code:    400,
			Message: err.Error(),
		})

		return
	}

	r.Body.Close()

	if cookie.Cookie == "" {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(model.HTTPResponse{
			Code:    400,
			Message: "cookie can not be empty",
		})

		return
	}

	cl := facebook.New()
	cl.SetCookieStr(cookie.Cookie)
	if err := cl.Init(); err != nil {
		if err != facebook.ErrInvalidSession {
			w.WriteHeader(500)
			json.NewEncoder(w).Encode(model.HTTPResponse{
				Code:    500,
				Message: err.Error(),
			})

			return
		}

		w.WriteHeader(401)
		json.NewEncoder(w).Encode(model.HTTPResponse{
			Code:    401,
			Message: err.Error(),
		})

		return
	}

	id := random.RandStr(20)
	cookie.ID = id
	cookie.CreatedAt = time.Now().Format("2006-01-02T15:04:05-0700")
	cookie.SetClient(cl)
	c.Cookies.Store(id, cookie)

	w.WriteHeader(200)
	json.NewEncoder(w).Encode(cookie)
}
