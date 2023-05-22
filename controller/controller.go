package controller

import(
	"fmt"
	"encoding/json"
	"net/http"
	"path/filepath"
	"bytes"
	"log"
	"io/ioutil"
	"github.com/JesusBarboza1994/EnroncorpZinc/model"
)
func Search(w http.ResponseWriter, r *http.Request){
	
	filepath.Join(model.SearchUrl, "_search")

	data := map[string]interface{}{
		"search-type": "text",
		"query": map[string]interface{}{
			"term":  "sent_items",
			"field": "_all",
		},
		"sort_fields":  []string{"-@timestamp"},
		"from":         0,
		"max_results":  20,
		"_source":      []string{},
	}
	// Convertir los datos a formato JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error al serializar JSON:", err)
		return
	}

	// Crear de solicitud POST
	req, err := http.NewRequest("POST", "http://localhost:4080/api/enron_zinc_v03/_search", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error al crear la solicitud:", err)
		return
	}

	// Establecer las cabeceras de la solicitud
	req.Header.Set("Content-Type", "application/json")

	// Agregar las credenciales de autenticación Basic
	req.SetBasicAuth("admin", "Complexpass#123")

	// Realizar la solicitud HTTP
	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error al enviar la solicitud:", err)
		return
	}
	defer resp.Body.Close()
	// Leer la respuesta
	log.Println(resp.StatusCode)
	respData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error al leer la respuesta:", err)
		return
	}

	// Establecer las cabeceras de la respuesta
	w.Header().Set("Content-Type", "application/json")

	// Escribir los datos de respuesta en el http.ResponseWriter
	w.Write(respData)
}