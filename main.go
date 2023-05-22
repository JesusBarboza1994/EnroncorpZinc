package main

import (
	"log"
	"fmt"
	"net/http"
	"github.com/JesusBarboza1994/EnroncorpZinc/config"
	"github.com/JesusBarboza1994/EnroncorpZinc/controller"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)



func main() {
	r := chi.NewRouter()

	// Agrega el middleware de recuperación de errores
	r.Use(middleware.Recoverer)
	r.Post("/search", controller.Search)

	if !config.IndexExist() {
		fmt.Println("Index created")
		config.UpZinc()
		config.LoopUsers()
	}
	log.Fatal(http.ListenAndServe(":8000", r))


	
}