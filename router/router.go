package router

import "github.com/gin-gonic/gin"

func Initialice() {
	//inicializa una ruta utilizando la configuración por defecto de gin
	router := gin.Default()

	//definiendo una ruta
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	// Estamos corriendo a nuestra api
	router.Run(":8080") // listen and serve on 0.0.0.0:8080
}
