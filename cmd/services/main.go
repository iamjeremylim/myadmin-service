package main

import (
	"database/sql"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/iamjeremylim/myadmin-service/api/stores"
	"github.com/iamjeremylim/myadmin-service/cmd/services/config"
	db "github.com/iamjeremylim/myadmin-service/db/sqlc"
)

func main() {
	// Read configuration
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal(err.Error())
	}

	// Instantiates the database
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	// Instantiates the store service
	queries := db.New(conn)
	storeService := stores.NewService(queries)

	// Register our service handlers to the router
	router := gin.Default()

	corsConfig := cors.Config{
		AllowWildcard: true,
		AllowOrigins:  allowedOrigins(),
		AllowMethods:  []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:  []string{"Authorization", "Content-Type", "X-Signature", "X-Origin", "X-User-Source"},
	}
	router.Use(cors.New(corsConfig))

	storeService.RegisterHandlers(router)

	// Start the server
	log.Println("Hello, running router!")
	router.Run()

}

func allowedOrigins() (origins []string) {
	origins = []string{
		"http://127.0.0.1:5173",
	}

	return
}
