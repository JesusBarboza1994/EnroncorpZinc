package config

import(
	"os"
	"fmt"
	"path/filepath"
	"sync"
)

func LoopUsers(){

	wg := &sync.WaitGroup{}
	// Declara la ruta base principal
	folderPath := "../enron_mail_20110402/maildir"

	// Obtiene la lista de carpetas de usuarios
	usersFolders, err := os.ReadDir(folderPath)
	if err != nil {
		fmt.Println("Error al leer la carpeta:", err)
		return
	}
	
	// Recorre las carpetas
	for i, usersFolder := range usersFolders {
		if i<3{
			wg.Add(1)
			fmt.Println(usersFolder.Name())
			userFolderName := usersFolder.Name()
			go func(folderPath, userFolderName string) {
				defer wg.Done()
				LoopFoldersOfAUser(folderPath, userFolderName)
			}(filepath.Join(folderPath, userFolderName), userFolderName)
		}
	}
	wg.Wait()
}
func LoopFoldersOfAUser(folderPath, userFolderName string) {

	// Obtiene la lista de carpetas
	folders, err := os.ReadDir(folderPath)
	if err != nil {
		fmt.Println("Error al leer la carpeta:", err)
		return
	}

	// Recorre las carpetas
	for _, folder := range folders {
		fmt.Println(folder.Name())
		folderName := folder.Name()
		LoopFiles(filepath.Join(folderPath, folderName), folderName, userFolderName)
	}
}
func LoopFiles(folderPath, folderName, userFolderName string){
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
			ExtractInfoItem(filePath, folderName, userFolderName)
	}
}