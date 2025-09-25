package main

import (
	"TAPI/handler"
	"TAPI/repository"
	"TAPI/service"
	"log"

	"database/sql"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pl"
)

func main() {
	//loading the values from .env
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}
	// DB authentication string
	DBauth := fmt.Sprintf("%s://%s:%s@%s/%s?sslmode=%s",
		"postgres",
		os.Getenv("DBuser"),
		os.Getenv("DBpassword"),
		os.Getenv("DBconn"),
		os.Getenv("DBname"),
		os.Getenv("DBsslmode"),
	)
	// the database connection
	sqlconn, err := sql.Open("postgres", DBauth)
	if err != nil {
		log.Fatalf("Failed to open database connection: %s", err)
	}
	defer sqlconn.Close()

	err = sqlconn.Ping()
	if err != nil {
		log.Fatalf("Failed to ping database: %s", err)
	}
	fmt.Println("Successfully connected to Database")

	// Dependency injection phase
	repo := repository.NewPGModelRepo(sqlconn)
	svc := service.NewRepoContractInstance(repo)
	hdlr := handler.NewServiceContractInstance(svc)

	ginCxt := gin.Default()

	ginCxt.POST("/register", hdlr.HandleCreateDeviceRequest)
	ginCxt.POST("/update", hdlr.HandleUpdateMeshRequest)
	ginCxt.POST("/retrieve", hdlr.HandleDeviceRetrieval)

	log.Println("Starting server on :8080")
	if err := ginCxt.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %s", err)
	}
}
