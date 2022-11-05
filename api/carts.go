package api

import (
	"a21hc3NpZ25tZW50/model"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

func (api *API) AddCart(w http.ResponseWriter, r *http.Request) {
	// Get username context to struct model.Cart.
	//username := "" // TODO: replace this
	cookie, _ := r.Cookie("session_token")
	if cookie == nil {
		w.WriteHeader(http.StatusUnauthorized)
		_ = json.NewEncoder(w).Encode(model.ErrorResponse{Error: "http: named cookie not present"})
		return
	}

	r.ParseForm()

	// Check r.Form with key product, if not found then return response code 400 and message "Request Product Not Found".
	// TODO: answer here

	var list []model.Product
	for _, formList := range r.Form {
		for _, v := range formList {
			item := strings.Split(v, ",")
			p, _ := strconv.ParseFloat(item[2], 64)
			q, _ := strconv.ParseFloat(item[3], 64)
			total := p * q
			list = append(list, model.Product{
				Id:       item[0],
				Name:     item[1],
				Price:    item[2],
				Quantity: item[3],
				Total:    total,
			})
		}
	}

	if len(list) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(model.ErrorResponse{Error: "Request Product Not Found"})
	}

	// Add data field Name, Cart and TotalPrice with struct model.Cart.
	session, _ := api.sessionsRepo.TokenExist(cookie.Value)
	cart := model.Cart{
		Name: session.Username,
		Cart: list,
	} // TODO: replace this

	for _, product := range cart.Cart {
		cart.TotalPrice += product.Total
	}

	_, err := api.cartsRepo.CartUserExist(cart.Name)
	if err != nil {
		api.cartsRepo.AddCart(cart)
	} else {
		api.cartsRepo.UpdateCart(cart)
	}
	api.dashboardView(w, r)

}
