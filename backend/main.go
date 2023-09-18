package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("🚀 Starting server...")
	if os.Getenv("debug") == "release" {
		InitMigrations()
	}
	r := SetupRouter()
	r.Run(":" + os.Getenv("PORT"))

}
