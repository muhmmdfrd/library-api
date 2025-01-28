package main

import (
	"library-api/modules/auth"
	"library-api/modules/users"
	"library-api/utils"
	"log"
	"os"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
    godotenv.Load()

    dbHost := os.Getenv("DB_HOST")
    dbUser := os.Getenv("DB_USER")
    dbPassword := os.Getenv("DB_PASSWORD")
    dbName := os.Getenv("DB_NAME")
    dbPort := os.Getenv("DB_PORT")

    dsn := dbUser + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"

    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
        QueryFields: true,
        PropagateUnscoped: true,
    })
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }

    // redis
    utils.InitRedis()

    // repositories
    userRepo := users.NewUserRepo(db)

    // handlers
    authHandler := auth.NewAuthHandler(userRepo)
    userHandler := users.NewUserHandler(userRepo)

    r := gin.Default()

    // gzip compression
    r.Use(gzip.Gzip(gzip.BestSpeed))

    auth.SetupAuthRoutes(r, authHandler)
    users.SetupUserRoutes(r, userHandler)

    r.Run(":8080")
}