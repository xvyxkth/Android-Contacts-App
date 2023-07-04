package main

import (
	"github.com/AvyakthKrishnaKumar/go-crud/initializers"
	"github.com/AvyakthKrishnaKumar/go-crud/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
