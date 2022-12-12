package routers

import (
	"privy-test/controllers"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/cakes", controllers.GetAllCakes).Methods("GET", "OPTIONS")
	router.HandleFunc("/cakes/{id}", controllers.GetCakeById).Methods("GET", "OPTIONS")
	router.HandleFunc("/cakes", controllers.CreateCake).Methods("POST", "OPTIONS")
	router.HandleFunc("/cakes/{id}", controllers.UpdateCakeById).Methods("PATCH", "OPTIONS")
	router.HandleFunc("/cakes/{id}", controllers.DeleteCakeById).Methods("DELETE", "OPTIONS")

	return router
}
