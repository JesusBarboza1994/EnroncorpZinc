package main

import (
	"log"
	"fmt"
	"sync"
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
		// Ya no se crea el index, ya que el _bulk lo crea por si solo
		// config.UpZinc()
		var wg sync.WaitGroup
		wg.Add(1)
		go func() {
			defer wg.Done()
			config.LoopUsers("../enron_mail_20110402/maildir")
		}()
		wg.Wait()
		config.UploadTotalInfo()
	}


	// Ruta para el profiling

	fmt.Println("Listening on port 8000")
	log.Fatal(http.ListenAndServe(":8000", handler))

}