package main

import (
    "os"
    "fmt"
    "log"
    "api-freeradius/db"
    "github.com/joho/godotenv"
    "github.com/gin-gonic/gin"
    _ "api-freeradius/docs"
    "github.com/swaggo/files"
    "github.com/swaggo/gin-swagger"
    "api-freeradius/internal/api/routes"
)

// @title Radius API
// @version 1.0
// @description API for managing FreeRADIUS users and NAS devices
// @BasePath /

func main() {

    godotenv.Load()
    host := os.Getenv("DB_HOST")
    user := os.Getenv("DB_USER")
    password := os.Getenv("DB_PASSWORD")
    dbname := os.Getenv("DB_NAME")
    port := os.Getenv("DB_PORT")
    sslmode := os.Getenv("DB_SSLMODE")

    dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
        host, user, password, dbname, port, sslmode,
    )

    if err := db.Connect(dsn); err != nil {
        log.Fatal("erreur connexion db:", err)
        return
    }

    r := gin.Default()
    url := ginSwagger.URL("/swagger/doc.json")
    r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

    routes.SetupRouter(r)

    r.Run()

}
