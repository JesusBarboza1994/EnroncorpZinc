package main

import (
	"log"
	"fmt"
	"net/http"
	"github.com/JesusBarboza1994/EnroncorpZinc/config"
	"github.com/JesusBarboza1994/EnroncorpZinc/controller"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/rs/cors"
)



func main() {
	r := chi.NewRouter()

	// Agrega el middleware de recuperación de errores
	r.Use(middleware.Recoverer)

	// Agregar CORS disponibles para el cliente
	corsOptions := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:5500", "http://127.0.0.1:5500", "http://127.0.0.1:5501"},
		AllowedMethods: []string{ "POST"},
	})

	// Agregar el middleware CORS 
	handler := corsOptions.Handler(r)

	r.Post("/search", controller.Search)

	if !config.IndexExist() {
		fmt.Println("Index created")
		config.UpZinc()
		config.LoopUsers()
	}
	fmt.Println("Listening on port 8000")
	log.Fatal(http.ListenAndServe(":8000", handler))

}