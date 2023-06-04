package main

import (
	"log"
	"fmt"
	"sync"
	"net/http"
	"github.com/JesusBarboza1994/EnroncorpZinc/config"
	"github.com/JesusBarboza1994/EnroncorpZinc/controller"
	"github.com/go-chi/chi"
	"github.com/rs/cors"
	// "bytes"
)



func main() {
	r := chi.NewRouter()

	// Agregar CORS disponibles para el cliente
	corsOptions := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:5500", "http://127.0.0.1:5500", "http://127.0.0.1:5501"},
		AllowedMethods: []string{ "POST"},
	})

	// Agregar el middleware CORS 
	handler := corsOptions.Handler(r)

	r.Post("/search", controller.Search)


	if !config.IndexExist() {
		config.UpZinc()
		var wg sync.WaitGroup
		wg.Add(1)
		go func() {
			defer wg.Done()
			config.LoopUsers("../enron_mail_20110402/maildir")
		}()
		wg.Wait()
		// config.UploadTotalInfo()
	}


	// Ruta para el profiling

	fmt.Println("Listening on port 8000")
	log.Fatal(http.ListenAndServe(":8000", handler))

}

// func main() {
// 	// URL del punto de conexión
// 	url := "http://localhost:4080/api/article/_multi"
// 	username := "admin"
// 	password := "Complexpass#123"
// 	// Cuerpo de la solicitud con documentos JSON separados por líneas
// 	body := `{"Year": 1796, "City": "Atenas", "Sport": "Natación", "Discipline": "Natación", "Athlete": "HAJOS, Alfred", "Country": "HUN", "Gender": "Hombres", "Event": "100M Estilo Libre", "Medal": "Oro", "Season": "verano"}
// 					{"Year": 1996, "City": "Atenas", "Sport": "Natación", "Discipline": "Natación", "Athlete": "HERSCHMANN, Otto", "Country": "AUT", "Gender": "Hombres", "Event": "100M Estilo Libre", "Medal": "Plata", "Season": "verano"}
// 					{"Year": 2096, "City": "Atenas", "Sport": "Natación", "Discipline": "Natación", "Athlete": "CHASAPIS, Spiridon", "Country": "GRE", "Gender": "Hombres", "Event": "100M Estilo Libre para Marineros", "Medal": "Plata", "Season": "verano"}`

// 	// Crear un objeto de tipo bytes.Buffer para almacenar el cuerpo de la solicitud
// 	reqBody := bytes.NewBufferString(body)

// 	// Crear una solicitud POST
// 	req, err := http.NewRequest("POST", url, reqBody)
// 	if err != nil {
// 		fmt.Println("Error al crear la solicitud:", err)
// 		return
// 	}

// 	// Establecer el tipo de contenido del encabezado de la solicitud
// 	req.Header.Set("Content-Type", "application/json")
// 	req.SetBasicAuth(username, password)
	
// 	// Crear un cliente HTTP y enviar la solicitud
// 	client := &http.Client{}
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		fmt.Println("Error al enviar la solicitud:", err)
// 		return
// 	}

// 	// Verificar el código de respuesta
// 	if resp.StatusCode == http.StatusOK {
// 		fmt.Println("Solicitud enviada correctamente.")
// 	} else {
// 		fmt.Println("Error al enviar la solicitud. Código de respuesta:", resp.StatusCode)
// 	}
// }