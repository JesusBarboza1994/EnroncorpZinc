package config

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
)

func LoopUsers(folderPath string){

	// wg := &sync.WaitGroup{}

	// Obtiene la lista de carpetas de usuarios
	usersFolders, err := os.ReadDir(folderPath)
	if err != nil {
		fmt.Println("Error al leer la carpeta:", err)
		return
	}

	// Recorre las carpetas
	for _, usersFolder := range usersFolders {
			// wg.Add(1)
			userFolderName := usersFolder.Name()
			// go func(folderPath, userFolderName string) {
				// defer wg.Done()
				LoopFoldersOfAUser(filepath.Join(folderPath, userFolderName), userFolderName)
			// }
	}
	// wg.Wait()
}

func LoopFoldersOfAUser(folderPath, userFolderName string) {
	// wg := &sync.WaitGroup{}
	// Obtiene la lista de carpetas
	folders, err := os.ReadDir(folderPath)
	if err != nil {
		fmt.Println("Error al leer la carpeta:", err)
		return
	}

	// Recorre las carpetas
	for _, folder := range folders {
			folderName := folder.Name()
			// wg.Add(1)
				// go func(folderPath, folderName, userFolderName string) {
				// 	defer wg.Done()
					LoopFiles(filepath.Join(folderPath, folderName), folderName, userFolderName)
				// }(filepath.Join(folderPath, folderName), folderName, userFolderName)
	}
	// wg.Wait()
}

func LoopFiles(folderPath, folderName, userFolderName string){
	var wg sync.WaitGroup
	var m sync.Mutex
	// Lee el contenido de la carpeta
	files, err := os.ReadDir(folderPath)
	if err != nil {
		fmt.Println("Error al leer la carpeta:", err)
		return
	}

	// Recorre los archivos de la carpeta
	for _, file := range files {
		// Obtiene el nombre del archivo
		fileName := file.Name()
		
		// Ruta completa del archivo
		filePath := filepath.Join(folderPath, fileName)
		// Ejecuta la función ExtractInfoItem como goroutine
		wg.Add(1)
		go func(){
			defer wg.Done()
			ExtractInfoItem(filePath, folderName, userFolderName, &m, &wg)	
		}()
	}
	wg.Wait()
	
}