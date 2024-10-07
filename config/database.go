package config

import (
    "log"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "os"
    "github.com/joho/godotenv"
)

var DB *gorm.DB

func ConnectDB() {

    errorDb := godotenv.Load()
    if errorDb != nil {
        log.Fatal("Error loading .env file")
    }

    dbUser := os.Getenv("DB_USER")
    dbPassword := os.Getenv("DB_PASSWORD")
    dbName := os.Getenv("DB_NAME")
    dbHost := os.Getenv("DB_HOST")
    dbPort := os.Getenv("DB_PORT")

    dsn := dbUser + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"

    var err error
    DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to database: ", err)
    }
    log.Println("Database connected successfully")
}
