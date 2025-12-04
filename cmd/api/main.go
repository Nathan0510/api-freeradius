package main

import (
    "os"
    "fmt"
    "log"
    "api-freeradius/db"
    "api-freeradius/internal/app"
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

    if err := app.CreateUserRadcheck("toto", "toto"); err != nil {
        log.Fatal("Erreur ajout radcheck user:", err)
    }

    if err := app.CreateUserRadreply("toto", "Framed-IP-Address", "192.168.1.50"); err != nil {
        log.Fatal("Erreur ajout radreply user:", err)
    }

    if err := app.CreateUserRadreply("toto", "Mikrotik-Group", "Profile-Internet"); err != nil {
        log.Fatal("Erreur ajout radreply user:", err)
    }

//    db.DB.Model(&models.Radreply{}).Where(&models.Radreply{Username: "toto", Attribute: "Mikrotik-Group"}).Update("Value", "Profile-Customer1")

    if err := app.UpdateUserRadreply("toto", "Mikrotik-Group", "Profile-Customer1"); err != nil {
        log.Println("Erreur update Radreply :", err)
    }


    r := gin.Default()

    r.GET("/ping", controllers.ping)

    r.Run()

}
