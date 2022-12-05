package main

import (
	"fmt"
	"net/http"

	"github.com/bugslayer_332/crudAPI/internal/routes"
	"github.com/bugslayer_332/crudAPI/internal/service"
	"github.com/gorilla/handlers"
	"github.com/supertokens/supertokens-golang/supertokens"
)

func main() {
	supertokens.Init(service.SuperTokensConfig)

	fmt.Println("Welcome to my website")

	routes.Routes()
	http.ListenAndServe(":8080", handlers.CORS(
		handlers.AllowedHeaders(append([]string{"Content-Type"},
			supertokens.GetAllCORSHeaders()...)),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}),
		handlers.AllowedOrigins([]string{"http://localhost:3000"}),
		handlers.AllowCredentials(),
	)(supertokens.Middleware(routes.Routes())))

}
