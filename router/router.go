package router

import "github.com/gin-gonic/gin"

func Initialice() {
	//TODO: inicializa una ruta utilizando la configuración por defecto de gin
	router := gin.Default()

	//TODO: definiendo una ruta
	inicializeRoutes(router)
	//TODO: Run the server
	router.Run(":8080") //TODO: listen and serve on 0.0.0.0:8080
}
