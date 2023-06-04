package config

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	// "path"
	"path/filepath"
	"strings"

	// "sync"
	"github.com/JesusBarboza1994/EnroncorpZinc/model"
)

var i = 1
var j = 1
func UploadTotalInfo() {
	url := "http://localhost:4080/api/_bulk"
	username := "admin"
	password := "Complexpass#123"
	filePath := "data.ndjson"

	// Lee el contenido del archivo
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Error al abrir el archivo:", err)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	var jsonData bytes.Buffer

	for {
		line, err := reader.ReadBytes('\n')
		if err != nil {
			break
		}
		jsonData.Write(line)
	}

	req, err := http.NewRequest("POST", url, &jsonData)
	if err != nil {
		log.Fatal("Error al crear la solicitud HTTP:", err)
		return
	}

	req.SetBasicAuth(username, password)
	req.Header.Set("Content-Type", "application/json")

	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Error al enviar la solicitud HTTP:", err)
		return
	}
	defer resp.Body.Close()

	// Leer la respuesta
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error al leer la respuesta HTTP:", err)
		return
	}

	// Mostrar la respuesta en el registro
	log.Println(string(respBody))
}
func ExtractInfoItem(filePath, folderName, userFolderName string){
	
	dataFile, err := os.Stat(filePath)
	if err != nil {
		fmt.Println("Error al obtener información del archivo:", err)
		return
	}
	if dataFile.IsDir() {
		files, err := os.ReadDir(filePath)
		if err != nil {
			fmt.Println("Error al leer la carpeta:", err)
			return
		}
		for _, file := range files {
			// Obtiene el nombre del archivo
			fileName := file.Name()
			
			// Ruta completa del archivo
			filePath := filepath.Join(filePath, fileName)

			// Ejecuta recursivamente la función
			ExtractInfoItem(filePath, folderName, userFolderName)	

		}
	}else{
		// if j == 152{
		// 	fmt.Println("%v %v", folderName, userFolderName)
		// }
		// Lee el archivo

		if !isHeavyFile(filePath){

			fileData, err := ioutil.ReadFile(filePath)
			if err != nil {
				fmt.Println("Error al leer el archivo:", err)
				return 
			}
		
			// Convierte el contenido del archivo a una cadena
			fileContent := string(fileData)
		
			// Separa el contenido del archivo por líneas
			lines := strings.Split(fileContent, "\n")
		
			// Crea un mapa para almacenar los campos del archivo
			emailMap := make(map[string]interface{})
		
			// Itera sobre las líneas del archivo y asigna los campos al mapa
			// Revisa que los campos no esten completos para no sobreescribir los valores
			for _, line := range lines {
				if strings.HasPrefix(line, "Message-ID:") && emailMap["Message-ID"] == nil {
					emailMap["Message-ID"] = strings.TrimSpace(strings.TrimPrefix(line, "Message-ID:"))
				} else if strings.HasPrefix(line, "Date:") && emailMap["Date"] == nil {
					emailMap["Date"] = strings.TrimSpace(strings.TrimPrefix(line, "Date:"))
				} else if strings.HasPrefix(line, "From:") && emailMap["From"] == nil {
					emailMap["From"] = strings.TrimSpace(strings.TrimPrefix(line, "From:"))
				} else if strings.HasPrefix(line, "To:") && emailMap["To"] == nil {
					emailMap["To"] = strings.TrimSpace(strings.TrimPrefix(line, "To:"))
				} else if strings.HasPrefix(line, "Subject:") && emailMap["Subject"] == nil {
					emailMap["Subject"] = strings.TrimSpace(strings.TrimPrefix(line, "Subject:"))
				} else if strings.HasPrefix(line, "Mime-Version:") && emailMap["Mime-Version"] == nil {
					emailMap["Mime-Version"] = strings.TrimSpace(strings.TrimPrefix(line, "Mime-Version:"))
				} else if strings.HasPrefix(line, "Content-Type:") && emailMap["Content-Type"] == nil {
					emailMap["Content-Type"] = strings.TrimSpace(strings.TrimPrefix(line, "Content-Type:"))
				} else if strings.HasPrefix(line, "Content-Transfer-Encoding:") && emailMap["Content-Transfer-Encoding"] == nil {
					emailMap["Content-Transfer-Encoding"] = strings.TrimSpace(strings.TrimPrefix(line, "Content-Transfer-Encoding:"))
				} else if strings.HasPrefix(line, "X-From:") && emailMap["X-From"] == nil {
					emailMap["X-From"] = strings.TrimSpace(strings.TrimPrefix(line, "X-From:"))
				} else if strings.HasPrefix(line, "X-To:") && emailMap["X-To"] == nil {
					emailMap["X-To"] = strings.TrimSpace(strings.TrimPrefix(line, "X-To:"))
				} else if strings.HasPrefix(line, "X-cc:") && emailMap["X-cc"] == nil {
					emailMap["X-cc"] = strings.TrimSpace(strings.TrimPrefix(line, "X-cc:"))
				} else if strings.HasPrefix(line, "X-bcc:") && emailMap["X-bcc"] == nil {
					emailMap["X-bcc"] = strings.TrimSpace(strings.TrimPrefix(line, "X-bcc:"))
				} else if strings.HasPrefix(line, "X-Folder:") && emailMap["X-Folder"] == nil {
					emailMap["X-Folder"] = strings.TrimSpace(strings.TrimPrefix(line, "X-Folder:"))
				} else if strings.HasPrefix(line, "X-Origin:") && emailMap["X-Origin"] == nil {
					emailMap["X-Origin"] = strings.TrimSpace(strings.TrimPrefix(line, "X-Origin:"))
				} else if strings.HasPrefix(line, "X-FileName:") && emailMap["X-FileName"] == nil {
					emailMap["X-FileName"] = strings.TrimSpace(strings.TrimPrefix(line, "X-FileName:"))
				} else {
					// Agrega el contenido sin nombre al campo Message
					if _, ok := emailMap["Message"]; !ok {
						emailMap["Message"] = line
					} else {
						emailMap["Message"] = emailMap["Message"].(string) + "\n" + line
					}
				}
			}
		
			// Convierte el mapa a un objeto Email
			email := model.Email{
				File:                    folderName,
				User:                    userFolderName,
				MessageID:               getStringValue(emailMap["Message-ID"]),
				Date:                    getStringValue(emailMap["Date"]),
				From:                    getStringValue(emailMap["From"]),
				To:                      getStringValue(emailMap["To"]),
				Subject:                 getStringValue(emailMap["Subject"]),
				MimeVersion:             getStringValue(emailMap["Mime-Version"]),
				ContentType:             getStringValue(emailMap["Content-Type"]),
				ContentTransferEncoding: getStringValue(emailMap["Content-Transfer-Encoding"]),
				XFrom:                   getStringValue(emailMap["X-From"]),
				XTo:                     getStringValue(emailMap["X-To"]),
				Xcc:                     getStringValue(emailMap["X-cc"]),
				Xbcc:                    getStringValue(emailMap["X-bcc"]),
				XFolder:                 getStringValue(emailMap["X-Folder"]),
				XOrigin:                 getStringValue(emailMap["X-Origin"]),
				XFileName:               getStringValue(emailMap["X-FileName"]),
				Message:                 getStringValue(emailMap["Message"]),
			}
		
			// sendItem(email)
			// m.Lock()
			sendEmailsByBatch(email)
			// m.Unlock()
		}
	}

}

// var batch = ""
var client = &http.Client{}
var batch bytes.Buffer
func sendEmailsByBatch(email model.Email) {
	url := "http://localhost:4080/api/enron_zinc_v03/_multi"
	username := "admin"
	password := "Complexpass#123"

	encoder := json.NewEncoder(&batch)
	err := encoder.Encode(email)
	if err != nil {
		fmt.Println("Error al serializar JSON:", err)
		return
	}

	if i == 20000 {
		fmt.Printf("batch: %v \n", j)

		// Crear una solicitud POST
		req, err := http.NewRequest("POST", url, bytes.NewReader(batch.Bytes()))
		if err != nil {
			fmt.Println("Error al crear la solicitud:", err)
			return
		}

		// Establecer el tipo de contenido del encabezado de la solicitud
		req.Header.Set("Content-Type", "application/json")
		req.SetBasicAuth(username, password)

		// Crear un cliente HTTP y enviar la solicitud
		
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("Error al enviar la solicitud:", err)
			return
		}

		// Verificar el código de respuesta
		if resp.StatusCode == http.StatusOK {
			fmt.Println("Solicitud enviada correctamente.")
		} else {
			fmt.Println("Error al enviar la solicitud. Código de respuesta:", resp.StatusCode)
			return
		}

		i = 0
		batch.Reset()
		j = j + 1
	}

	i = i + 1
}

func isHeavyFile(filePath string) bool{
	// Obtener información del archivo
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		fmt.Println("Error al obtener información del archivo:", err)
		return true
	}
	// Verificar el tamaño del archivo
	fileSize := fileInfo.Size()
	if fileSize > 1048576 { // Tamaño mayor a 1 MB
		fmt.Println("El archivo es demasiado grande.")
		return true
	}
	return false
	
}

func getStringValue(value interface{}) string {
	if value == nil {
		return "-"
	}

	if str, ok := value.(string); ok {
		return str
	}

	return ""
}

// var file *os.File
func sendItem(email model.Email) {
	// Convierte el objeto Email a JSON en una sola línea
	jsonData, err := json.Marshal(email)
	if err != nil {
		fmt.Println("Error al serializar JSON:", err)
		return
	}

	// Reemplaza los dos puntos ":" con ": " en el JSON
	jsonData = bytes.ReplaceAll(jsonData, []byte(`:`), []byte(`: `))
	// Reemplaza los dos puntos "," con ", " en el JSON
	jsonData = bytes.ReplaceAll(jsonData, []byte(`,`), []byte(`, `))
	jsonData = append(jsonData, '\n')

	// Abre el archivo en modo append

	file, err := os.OpenFile("data.ndjson", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal("Error al abrir el archivo:", err)
		return
	}
	defer file.Close()

	// Verifica si el jsonData no es nulo o vacío antes de escribir la línea del índice
	if len(jsonData) > 0 {
		// Escribe la línea del índice en el archivo NDJSON
		indexLine := `{ "index" : { "_index" : "enron_zinc_v03" } }` + "\n"
		entry := indexLine + string(jsonData)
		// Escribe el objeto JSON en el archivo NDJSON
		_, err = file.WriteString(entry)
		if err != nil {
			log.Fatal("Error al escribir en el archivo:", err)
			return
		}
	}
}





		// fmt.Printf("%s\n", jsonData)
		
	
		// req, err := http.NewRequest("POST","http://localhost:4080/api/enron_zinc_v03/_doc", bytes.NewBuffer(jsonData))
		// 	if err != nil {
		// 			log.Fatal(err)
		// 	}
		// req.SetBasicAuth("admin", "Complexpass#123")
		// req.Header.Set("Content-Type", "application/json")
		// resp, err := http.DefaultClient.Do(req)
		// 	if err != nil {
		// 			log.Fatal(err)
		// 	}
		// defer resp.Body.Close()
