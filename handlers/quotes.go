package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetQuotes(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Quotes will be fetched here",
	})
}

func SavePreferences(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Preferences will be saved here",
	})
}

func GetPreferences(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Preferences will be fetched here",
	})
}
