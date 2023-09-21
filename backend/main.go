package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {

	if os.Getenv("GIN_MODE") == "release" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		fmt.Println("Init migrations")
		// InitMigrations()
		gin.SetMode(gin.DebugMode)
	}
	r := SetupRouter()
	fmt.Println("🚀 Starting server on port: " + os.Getenv("PORT"))
	r.Run(":" + os.Getenv("PORT"))

}
