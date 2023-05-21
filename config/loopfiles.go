package config

import(
	"os"
	"fmt"
	"path/filepath"
)


func LoopFoldersOfAUser() {
	// Declara la ruta base para un usuario
	folderPath := "./enron_mail_20110402/maildir/allen-p"
	j := 1
	// Obtiene la lista de carpetas
	folders, err := os.ReadDir(folderPath)
	if err != nil {
		fmt.Println("Error al leer la carpeta:", err)
		return
	}

	// Recorre las carpetas
	for _, folder := range folders {
		fmt.Printf("j es %v \n", j)
		j = j +  1
		fmt.Println(folder.Name())
		folderName := folder.Name()
		LoopFiles(filepath.Join(folderPath, folderName), folderName)
	}
}
func LoopFiles(folderPath, folderName string){
	// Lee el contenido de la carpeta
	files, err := os.ReadDir(folderPath)
	if err != nil {
		fmt.Println("Error al leer la carpeta:", err)
		return
	}

	// Recorre los archivos de la carpeta
	for _, file := range files {
		// Verifica si es un archivo regular
		
			// Obtiene el nombre del archivo
			fileName := file.Name()
		fmt.Printf("el nombre es %v \n", fileName)
			// Ruta completa del archivo
			filePath := filepath.Join(folderPath, fileName)

			// Extraer información del archivo
			ExtractInfoItem(filePath, folderName)
	}
}