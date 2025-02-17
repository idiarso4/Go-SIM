package main

import (
    "fmt"
    "log"

    "academic-system/internal/config"
    "academic-system/internal/routes"
    "github.com/gin-gonic/gin"
    "github.com/joho/godotenv"
)

func main() {
    if err := godotenv.Load(); err != nil {
        log.Fatal("Error loading .env file")
    }

    cfg := config.LoadConfig()
    config.InitDB()

    // Initialize Gin router
    r := gin.Default()
    
    // Setup routes
    routes.SetupRoutes(r)

    // Start server
    addr := fmt.Sprintf(":%d", cfg.ServerPort)
    log.Printf("Server starting on %s in %s mode\n", addr, cfg.Environment)
    r.Run(addr)
}