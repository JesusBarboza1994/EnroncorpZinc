package main

import (
	"log"
	"fmt"
	"net/http"
	"github.com/JesusBarboza1994/EnroncorpZinc/config"
	"github.com/JesusBarboza1994/EnroncorpZinc/controller"
	"github.com/go-chi/chi"
	"github.com/rs/cors"
)

func main() {
	r := chi.NewRouter()

	// Agregar CORS disponibles para el cliente
	corsOptions := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:5500","http://localhost:5501", "http://127.0.0.1:5500", "http://127.0.0.1:5501"},
		AllowedMethods: []string{ "POST"},
	})

	// Agregar el middleware CORS 
	handler := corsOptions.Handler(r)

	// Ruta principal de la API
	r.Post("/search", controller.Search)

	//Se ejecuta la carga de datos solo si el index no existe
	if !config.IndexExist() {
		config.LoopUsers("../enron_mail_20110402/maildir")
		fmt.Println("Ultimo batch")
		config.SendBatch()
	}

	// Levantar el servidor
	fmt.Println("Listening on port 8000")
	log.Fatal(http.ListenAndServe(":8000", handler))

}
