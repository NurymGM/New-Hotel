package main

import (
	"github.com/NurymGM/New-Hotel/initializers"
	"github.com/NurymGM/New-Hotel/models"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
