package cookies

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tamboto2000/fb-scraper-api/model"
)

// Load load cookies by id
func (c *Cookies) Load(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := mux.Vars(r)["id"]
	cookie, ok := c.Cookies.Load(id)
	if !ok {
		w.WriteHeader(404)
		json.NewEncoder(w).Encode(model.HTTPResponse{
			Code:    404,
			Message: "cookie not found",
		})

		return
	}

	w.WriteHeader(200)
	json.NewEncoder(w).Encode(cookie)
}
