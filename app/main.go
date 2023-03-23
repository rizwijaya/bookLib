package main

import (
	error "bookLib/pkg/http-error"
	routesV1 "bookLib/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(cors.Default())

	router = routesV1.NewRouter(router)
	router.NoRoute(error.PageNotFound())
	router.NoMethod(error.NoMethod())

	router.Run(":8080")
}
