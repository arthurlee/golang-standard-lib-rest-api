package routes

import (
	"net/http"
	"../controllers"
)

func CreateRoutes(mux *http.ServeMux, uc *controllers.UserController)  {
	//mux.HandleFunc("/register", uc.Register)
}
