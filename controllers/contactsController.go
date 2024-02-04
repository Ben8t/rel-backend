package controllers

import (
	"rel/initializers"
	"rel/models"

	"github.com/gin-gonic/gin"
)

func CreateContact(c *gin.Context) {
	var body struct {
		Name     string
		LastName string
		Email    string
		Phone    string
		LinkedIn string
	}

	c.Bind(&body)
	contact := models.Contact{Name: body.Name, LastName: body.LastName, Email: body.Email, Phone: body.Phone, LinkedIn: body.LinkedIn}
	result := initializers.DB.Create(&contact)

	if result.Error != nil {
		c.Status(400)
		return
	}
	c.JSON(200, gin.H{
		"contact": contact,
	})
}

func ListContacts(c *gin.Context) {
	var contacts []models.Contact
	initializers.DB.Find(&contacts)

	c.JSON(200, gin.H{
		"contacts": contacts,
	})
}

func GetContact(c *gin.Context) {
	id := c.Param("id")

	var contact models.Contact
	initializers.DB.First(&contact, id)

	c.JSON(200, gin.H{
		"contact": contact,
	})
}

func UpdateContact(c *gin.Context) {
	id := c.Param("id")

	var body struct {
		Name     string
		LastName string
		Email    string
		Phone    string
		LinkedIn string
	}

	var contact models.Contact
	initializers.DB.First(&contact, id)

	c.Bind(&body)
	initializers.DB.Model(&contact).Updates(models.Contact{Name: body.Name, LastName: body.LastName, Email: body.Email, Phone: body.Phone, LinkedIn: body.LinkedIn})

	c.JSON(200, gin.H{
		"contact": contact,
	})
}

func DeleteContact(c *gin.Context) {
	id := c.Param("id")
	initializers.DB.Delete(&models.Contact{}, id)
	c.Status(200)
}