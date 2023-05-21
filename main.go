package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
	"log"
	"bytes"
	"net/http"
	"github.com/JesusBarboza1994/EnroncorpZinc/model"
	"github.com/JesusBarboza1994/EnroncorpZinc/config"
)



func main() {
	//Se registra la ruta del archivo
	filePath := "./enron_mail_20110402/maildir/allen-p/sent_items/2" 

	// Lee el archivo
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
		MessageID:              emailMap["Message-ID"].(string),
		Date:                   emailMap["Date"].(string),
		From:                   emailMap["From"].(string),
		To:                     emailMap["To"].(string),
		Subject:                emailMap["Subject"].(string),
		MimeVersion:            emailMap["Mime-Version"].(string),
		ContentType:            emailMap["Content-Type"].(string),
		ContentTransferEncoding: emailMap["Content-Transfer-Encoding"].(string),
		XFrom:                  emailMap["X-From"].(string),
		XTo:                    emailMap["X-To"].(string),
		Xcc:                    emailMap["X-cc"].(string),
		Xbcc:                   emailMap["X-bcc"].(string),
		XFolder:                emailMap["X-Folder"].(string),
		XOrigin:                emailMap["X-Origin"].(string),
		XFileName:              emailMap["X-FileName"].(string),
		Message:                emailMap["Message"].(string),
	}

	// Convierte el objeto Email a JSON
	jsonData, err := json.MarshalIndent(email, "", "  ")
	if err != nil {
		fmt.Println("Error al serializar JSON:", err)
		return
	}

	fmt.Printf("%s\n", jsonData)
	config.UpZinc()

	req, err := http.NewRequest("POST", "http://localhost:4080/api/enron_zinc_v01/_doc", bytes.NewBuffer(jsonData))
    if err != nil {
        log.Fatal(err)
    }
	req.SetBasicAuth("admin", "Complexpass#123")
	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
    if err != nil {
        log.Fatal(err)
    }
	defer resp.Body.Close()
}