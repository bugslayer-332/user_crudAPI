package routes

import (
	"github.com/bugslayer_332/crudAPI/internal/service"
	"github.com/gorilla/mux"
)

func Routes() *mux.Router {
	r := mux.NewRouter()
	// r.HandleFunc("/", welcome).Methods("POST")
	r.HandleFunc("/create", service.CreateUser).Methods("POST")
	r.HandleFunc("/read", service.GetAllUser).Methods("GET")
	r.HandleFunc("/read/{id}", service.GetUser).Methods("GET", "OPTIONS")

	r.HandleFunc("/update/{id}", service.UpdateUser).Methods("PUT")
	r.HandleFunc("/delete/{id}", service.DeleteUser).Methods("DELETE")

	return r

}
