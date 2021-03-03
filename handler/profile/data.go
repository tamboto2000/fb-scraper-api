package profile

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/tamboto2000/facebook"
	"github.com/tamboto2000/fb-scraper-api/model"
)

var profileFields = []string{
	"basic",
	"workAndEducation",
	"placesLived",
	"contactAndBasicInfo",
	"familyAndRelationships",
	"details",
	"lifeEvents",
}

// Data fetch profile data
func (p *Profile) Data(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// get cookie id
	cookieID := r.Header.Get("cookie-id")
	if cookieID == "" {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(model.HTTPResponse{
			Code:    400,
			Message: "header cookie-id can not be empty",
		})

		return
	}

	cl, ok := p.cookieHandler.Cookies.Load(cookieID)
	if !ok {
		w.WriteHeader(404)
		json.NewEncoder(w).Encode(model.HTTPResponse{
			Code:    404,
			Message: "cookie not found",
		})

		return
	}

	vars := mux.Vars(r)
	username := vars["username"]
	fieldStr := vars["fields"]

	prof, err := cl.(*model.Cookie).Client().Profile(username)
	if err != nil {
		if err == facebook.ErrInvalidSession {
			w.WriteHeader(401)
			json.NewEncoder(w).Encode(model.HTTPResponse{
				Code:    401,
				Message: "cookie expired or invalid, will automatically removed by server",
			})

			p.cookieHandler.Cookies.Delete(cookieID)

			return
		}

		if err == facebook.ErrUserNotFound {
			w.WriteHeader(404)
			json.NewEncoder(w).Encode(model.HTTPResponse{
				Code:    404,
				Message: "user not found",
			})

			return
		}

		w.WriteHeader(500)
		json.NewEncoder(w).Encode(model.HTTPResponse{
			Code:    500,
			Message: err.Error(),
		})

		return
	}

	var fields []string

	if fieldStr == "all" {
		fields = profileFields
	} else if fieldStr != "basic" {
		fields = strings.Split(fieldStr, ",")
	}

	for _, field := range fields {
		notFound := true
		for _, f := range profileFields {
			if field == f {
				notFound = false
				break
			}
		}

		if notFound {
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(model.HTTPResponse{
				Code:    400,
				Message: "unknown field \"" + field + "\"",
			})

			return
		}
	}

	if err := prof.SyncAbout(); err != nil {
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(model.HTTPResponse{
			Code:    500,
			Message: err.Error(),
		})

		return
	}

	err = syncAboutDetail(prof, fields)
	if err != nil {
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(model.HTTPResponse{
			Code:    500,
			Message: err.Error(),
		})

		return
	}

	w.WriteHeader(200)
	json.NewEncoder(w).Encode(prof)
}

func syncAboutDetail(prof *facebook.Profile, fields []string) error {
	for _, f := range fields {
		if f == "workAndEducation" {
			if err := prof.About.SyncWorkAndEducation(); err != nil {
				return err
			}
		}

		if f == "placesLived" {
			if err := prof.About.SyncPlacesLived(); err != nil {
				return err
			}
		}

		if f == "contactAndBasicInfo" {
			if err := prof.About.SyncContactAndBasicInfo(); err != nil {
				return err
			}
		}

		if f == "familyAndRelationships" {
			if err := prof.About.SyncFamilyAndRelationships(); err != nil {
				return err
			}
		}

		if f == "details" {
			if err := prof.About.SyncDetails(); err != nil {
				return err
			}
		}

		if f == "lifeEvents" {
			if err := prof.About.SyncLifeEvents(); err != nil {
				return err
			}
		}
	}

	return nil
}
