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

    dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
        os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"), os.Getenv("DB_SSLMODE"),
    )

    if err := db.Connect(dsn); err != nil {
        log.Fatal("Error connexion db:", err)
        return
    }

    r := gin.Default()
    url := ginSwagger.URL("/swagger/doc.json")
    r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

    routes.SetupRouter(r)

    r.Run()

}
