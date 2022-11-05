package api

import (
	"net/http"
)

func (api *API) ImgProfileView(w http.ResponseWriter, r *http.Request) {
	// View with response image `img-avatar.png` from path `assets/images`
	// TODO: answer here
}

func (api *API) ImgProfileUpdate(w http.ResponseWriter, r *http.Request) {
	// Update image `img-avatar.png` from path `assets/images`
	// TODO: answer here

	api.dashboardView(w, r)
}
