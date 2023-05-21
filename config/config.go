package config
import(
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"io/ioutil"
)

// Creación del indice base en ZincSearch
func UpZinc(){
	url := "http://localhost:4080/api/index"

	data := map[string]interface{}{
		"name":         "enron_zinc_v02",
		"storage_type": "disk",
		"shard_num":    1,
		"mappings": map[string]interface{}{
			"properties": map[string]interface{}{
				"File": map[string]interface{}{
					"type":          "text",
					"index":         true,
					"store":         true,
					"highlightable": true,
				},
				"Message-ID": map[string]interface{}{
					"type":          "text",
					"index":         true,
					"store":         true,
					"highlightable": true,
				},
				"Date": map[string]interface{}{
					"type":          "text",
					"index":         true,
					"store":         true,
					"highlightable": true,
				},
				"From": map[string]interface{}{
					"type":          "text",
					"index":         true,
					"store":         true,
					"highlightable": true,
				},
				"To": map[string]interface{}{
					"type":          "text",
					"index":         true,
					"store":         true,
					"highlightable": true,
				},
				"Subject": map[string]interface{}{
					"type":          "text",
					"index":         true,
					"store":         true,
					"highlightable": true,
				},
				"Mime-Version": map[string]interface{}{
					"type":          "text",
					"index":         true,
					"store":         true,
					"highlightable": true,
				},
				"Content-Type": map[string]interface{}{
					"type":          "text",
					"index":         true,
					"store":         true,
					"highlightable": true,
				},
				"Content-Transfer-Encoding": map[string]interface{}{
					"type":          "text",
					"index":         true,
					"store":         true,
					"highlightable": true,
				},
				"X-From": map[string]interface{}{
					"type":          "text",
					"index":         true,
					"store":         true,
					"highlightable": true,
				},
				"X-To": map[string]interface{}{
					"type":          "text",
					"index":         true,
					"store":         true,
					"highlightable": true,
				},
				"Xcc": map[string]interface{}{
					"type":          "text",
					"index":         true,
					"store":         true,
					"highlightable": true,
				},
				"Xbcc": map[string]interface{}{
					"type":          "text",
					"index":         true,
					"store":         true,
					"highlightable": true,
				},
				"XFolder": map[string]interface{}{
					"type":          "text",
					"index":         true,
					"store":         true,
					"highlightable": true,
				},
				"XOrigin": map[string]interface{}{
					"type":          "text",
					"index":         true,
					"store":         true,
					"highlightable": true,
				},
				"XFileName": map[string]interface{}{
					"type":          "text",
					"index":         true,
					"store":         true,
					"highlightable": true,
				},
				"Message": map[string]interface{}{
					"type":          "text",
					"index":         true,
					"store":         true,
					"highlightable": true,
				},
			},
		},
	}	
	
	// Convertir los datos a formato JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error al serializar JSON:", err)
		return
	}

	// Crear de solicitud POST
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
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
	respData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error al leer la respuesta:", err)
		return
	}

	// Mostrar la respuesta
	fmt.Println("Respuesta:", string(respData))
}