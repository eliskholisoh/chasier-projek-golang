package api

import (
	"a21hc3NpZ25tZW50/model"
	"encoding/json"
	"net/http"
)

func (api *API) ImgProfileView(w http.ResponseWriter, r *http.Request) {
	// View with response image `img-avatar.png` from path `assets/images`
	// TODO: answer here
	cookie, _ := r.Cookie("session_token")
	if cookie == nil {
		w.WriteHeader(http.StatusUnauthorized)
		_ = json.NewEncoder(w).Encode(model.ErrorResponse{Error: "http: named cookie not present"})
		return
	}
}

func (api *API) ImgProfileUpdate(w http.ResponseWriter, r *http.Request) {
	// Update image `img-avatar.png` from path `assets/images`
	// TODO: answer here
	cookie, _ := r.Cookie("session_token")
	if cookie == nil {
		w.WriteHeader(http.StatusUnauthorized)
		_ = json.NewEncoder(w).Encode(model.ErrorResponse{Error: "http: named cookie not present"})
		return
	}

	creds := model.Credentials{
		Username: r.PostFormValue("username"),
		Password: r.PostFormValue("password"),
	} // TODO: replace this

	// Handle request if creds is empty send response code 400, and message "Username or Password empty"
	// TODO: answer here
	if creds.Username == "" && creds.Password == "" {
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(model.ErrorResponse{Error: "Error: internal server error"})
		return
	}

	api.dashboardView(w, r)
}
