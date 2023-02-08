package server

import (
	auth "ASIDK_Dispenser_Service/Auth"
	"github.com/gin-gonic/gin"
	"log"
)

func Start() {
	log.Printf("Server started")

	r := gin.Default()

	authMiddleware, err := auth.Auth()
	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}
	errInit := authMiddleware.MiddlewareInit()

	if errInit != nil {
		log.Fatal("authMiddleware.MiddlewareInit Error:" + errInit.Error())
	}

	TldZone := r.Group("/tld")
	TldZone.Use(authMiddleware.MiddlewareFunc())
	{
		{
			TldZone.GET("/:zoneid/pass/:pass")
			TldZone.GET("")
		}
	}
}
