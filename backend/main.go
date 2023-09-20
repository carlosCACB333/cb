package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("🚀 Starting server on port: " + os.Getenv("PORT"))
	if os.Getenv("GIN_MODE") == "release" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		fmt.Println("Init migrations")
		// InitMigrations()
		gin.SetMode(gin.DebugMode)
	}
	r := SetupRouter()
	r.Run(":" + os.Getenv("PORT"))

}
