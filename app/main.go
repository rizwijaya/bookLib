package main

import (
	database "bookLib/infrastructures/databases"
	routesV1 "bookLib/modules/v1/book/routes"
	error "bookLib/pkg/http-error"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(cors.Default())
	db := database.NewDatabases()

	router = routesV1.NewRouter(router, db)
	router.NoRoute(error.PageNotFound())
	router.NoMethod(error.NoMethod())

	router.Run(":8080")
}
