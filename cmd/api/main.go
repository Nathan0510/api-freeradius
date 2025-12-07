package main

import (
    "os"
    "fmt"
    "log"
    "api-freeradius/db"
    "api-freeradius/internal/api/controllers"
    "github.com/joho/godotenv"
    "github.com/gin-gonic/gin"
)

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

    r.GET("/ping", controllers.Ping)
    r.POST("/user", controllers.User)
    r.POST("/nas", controllers.Nas)


    r.Run()

}
