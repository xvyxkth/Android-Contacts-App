package controllers

import (
	"github.com/AvyakthKrishnaKumar/go-crud/initializers"
	"github.com/AvyakthKrishnaKumar/go-crud/models"
	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {
	// Get Data off request body

	var body struct {
		Name         string
		Phone_Number int
	}

	c.Bind(&body)

	// Create a Contact
	post := models.Post{Name: body.Name, Phone_Number: body.Phone_Number}

	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	// Return It
	c.JSON(200, gin.H{
		"posts": post,
	})
}

func PostsIndex(c *gin.Context) {
	// Get the Contacts
	var posts []models.Post
	initializers.DB.Find(&posts)

	// Respond with them
	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func PostsShow(c *gin.Context) {
	// Get ID from URL
	id := c.Param("id")

	// Find the contacts
	var post models.Post
	initializers.DB.First(&post, id)

	// Respond with them
	c.JSON(200, gin.H{
		"posts": post,
	})
}

func PostsUpdate(c *gin.Context) {
	// Get ID from the URL
	id := c.Param("id")

	//Get the data off the request body
	var body struct {
		Name         string
		Phone_Number int
	}

	c.Bind(&body)

	// Find the contact we're updating
	var post models.Post
	initializers.DB.First(&post, id)

	// Update it
	initializers.DB.Model(&post).Updates(models.Post{
		Name:         body.Name,
		Phone_Number: body.Phone_Number,
	})

	// Respond with it
	c.JSON(200, gin.H{
		"contact": post,
	})
}

func PostsDelete(c *gin.Context) {
	// Get the ID of the URL
	id := c.Param("id")

	// Delete the Contact
	initializers.DB.Delete(&models.Post{}, id)

	// Respond
	c.Status(200)

}
