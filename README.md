# EnroncorpZinc
Este es un proyecto que busca indexar una base datos usando el motor de búsqueda de ZincSearch.
Para ejecutar el proyecto deberás tener la base de datos. La puedes conseguir en la url: http://www.cs.cmu.edu/~enron/enron_mail_20110402.tgz(423MB)
Una vez descomprimida deberás poner la carpeta al lado del proyecto (no dentro). Esto es porque las rutas están preparadas de esa manera.
Luego, al ejecutar el comando `go run .` se deberá ejecutar el proyecto.
- Al inicio empezará a crear y editar el archivo data.ndjson con los datos de la base de datos en el formato que acepta zincsearch (es el mismo que Elastic search).
- Una vez finalice de recorrer todas las carpetas, empezará a subir el archivo a la base de datos. RECUERDA, tienes que haber levantado zincsearch en tu host local previamente (utiliza el puerto 4080 por defecto).
- Cargada toda la información, el puerto 8000 estará disponible para consulta. La información se irá cargando de a pocos una vez mostrado el log del puerto.
- Finalmente, ejecuta el archivo index.html en el puerto 5500 o 5501 (puertos que por defecto usa LiveServer). Puedes usar otros; sin embargo, debes habilitar los cors para el puerto que desees (se encuentran en el archivo main.go).
Con ello ya puedes entrar y consultar de manera eficiente la base de datos de correos.

CONSIDERACIONES:
Para que se ejecute el proyecto deberás tener instalado ZincSearch así como una versión a Go.
