package router

import "github.com/gin-gonic/gin"

func Initialice() {
	//inicializa una ruta utilizando la configuración por defecto de gin
	router := gin.Default()

	//definiendo una ruta
	inicializeRoutes(router)
	// Run the server
	router.Run(":8080") // listen and serve on 0.0.0.0:8080
}
