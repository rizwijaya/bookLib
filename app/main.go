package main

import (
	database "bookLib/infrastructures/databases"
	routesV1 "bookLib/modules/v1/book/routes"
	error "bookLib/pkg/http-error"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

//	@title			Book Library API
//	@description	This is a sample server for a book library.
//	@version		1.0.0
//	@termsOfService	http://swagger.io/terms/
//	@contact.name	Swagger API Team
//	@contact.email	admin@rizwijaya.com
//	@licence.name	MIT
//	@licence.url	http://opensource.org/licenses/MIT
//	@host			localhost:8080
//	@BasePath		/

func main() {
	// config, err := config.New()
	// if err != nil {
	// 	panic(err)
	// }

	router := gin.Default()
	router.Use(cors.Default())
	db := database.NewDatabases()

	router = routesV1.NewRouter(router, db)
	router.NoRoute(error.PageNotFound())
	router.NoMethod(error.NoMethod())
	port := os.Getenv("PORT")
	router.Run(":" + port)
}
