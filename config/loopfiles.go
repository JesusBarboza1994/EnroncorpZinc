package config

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
)

func LoopUsers(folderPath string) {
	wg := &sync.WaitGroup{}

	// Obtiene la lista de carpetas de usuarios
	usersFolders, err := os.ReadDir(folderPath)
	if err != nil {
		fmt.Println("Error al leer la carpeta:", err)
		return
	}

	// Recorre las carpetas
	for i, usersFolder := range usersFolders {
		if i < 2 {
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
	wg := &sync.WaitGroup{}

	// Obtiene la lista de carpetas
	folders, err := os.ReadDir(folderPath)
	if err != nil {
		fmt.Println("Error al leer la carpeta:", err)
		return
	}

	// Recorre las carpetas
	for j, folder := range folders {
		if j < 10 {
			fmt.Println(folder.Name())
			folderName := folder.Name()
			wg.Add(1)
			go func(folderPath, folderName, userFolderName string) {
				defer wg.Done()
				LoopFiles(filepath.Join(folderPath, folderName), folderName, userFolderName)
			}(folderPath, folderName, userFolderName)
		}
	}

	wg.Wait()
}

func LoopFiles(folderPath, folderName, userFolderName string) {
	wg := &sync.WaitGroup{}

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

		wg.Add(1)
		// Ejecuta la función ExtractInfoItem como goroutine
		go func(filePath, folderName, userFolderName string) {
			defer wg.Done()
			ExtractInfoItem(filePath, folderName, userFolderName)
		}(filePath, folderName, userFolderName)
	}

	wg.Wait()
}
