package controllers

import (
	"strconv"

	"github.com/PedroPaiao/rest-api-go/database"
	"github.com/PedroPaiao/rest-api-go/models"
	"github.com/gin-gonic/gin"
)

func ShowBook(c *gin.Context) {
	id := c.Param("id")

	newId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"DEU RUIM": "ID has to be integer",
		})

		return
	}

	db := database.GetDatabase()

	var book models.Book
	err = db.First(&book, newId).Error

	if err != nil {
		c.JSON(404, gin.H{
			"DEU RUIM": "Cannot find book " + err.Error(),
		})

		return
	}

	c.JSON(200, book)
}

func CreateBook(c *gin.Context) {
	db := database.GetDatabase()

	var book models.Book

	err := c.ShouldBindJSON(&book)
	if err != nil {
		c.JSON(422, gin.H{
			"DEU RUIM": "Unprossessable entity " + err.Error(),
		})

		return
	}

	err = db.Create(&book).Error

	if err != nil {
		c.JSON(422, gin.H{
			"DEU RUIM": "Cannot create a book " + err.Error(),
		})

		return
	}

	c.JSON(200, book)
}

func IndexBooks(c *gin.Context) {
	db := database.GetDatabase()

	var books []models.Book
	err := db.Find(&books).Error

	if err != nil {
		c.JSON(422, gin.H{
			"DEU RUIM": "Cannot load list of books",
		})

		return
	}

	c.JSON(200, books)
}

func UpdateBook(c *gin.Context) {
	id := c.Param("id")

	newId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"DEU RUIM": "ID has to be integer",
		})

		return
	}

	db := database.GetDatabase()

	var book models.Book
	err = db.First(&book, newId).Error

	if err != nil {
		c.JSON(404, gin.H{
			"DEU RUIM": "Cannot find book " + err.Error(),
		})

		return
	}

	err = c.ShouldBindJSON(&book)
	if err != nil {
		c.JSON(422, gin.H{
			"DEU RUIM": "Unprossessable entity " + err.Error(),
		})

		return
	}

	err = db.Save(&book).Error

	if err != nil {
		c.JSON(422, gin.H{
			"DEU RUIM": "Cannot update a book " + err.Error(),
		})

		return
	}

	c.JSON(200, book)
}

func DestroyBook(c *gin.Context) {
	id := c.Param("id")

	newId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"DEU RUIM": "ID has to be integer",
		})

		return
	}

	db := database.GetDatabase()

	var book models.Book
	err = db.First(&book, newId).Error

	if err != nil {
		c.JSON(404, gin.H{
			"DEU RUIM": "Cannot find book " + err.Error(),
		})

		return
	}

	err = db.Delete(&book).Error
	if err != nil {
		c.JSON(422, gin.H{
			"DEU RUIM": "Cannot delete the book!",
		})

		return
	}

	c.JSON(200, &book)
}
