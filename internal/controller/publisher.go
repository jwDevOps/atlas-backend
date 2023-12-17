package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jwDevOps/atlas-backend/internal/database"
	"github.com/jwDevOps/atlas-backend/internal/models"
)

func GetAllPublishers(c *gin.Context, m *database.DbManager) {
	publishers, err := m.QueryPublishers()

	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"message": "cannot fetch publishers from database",
		})
		return
	}
	c.JSON(200, gin.H{
		"publishers": publishers,
	})
}

func AddPublishers(c *gin.Context, m *database.DbManager) {
	var pub models.Publisher

	if err := c.BindJSON(&pub); err != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"message": "cannot parse post data",
		})
		return
	}

	if _, err := m.InsertPublisher(pub.FirstName, pub.LastName); err != nil {
		c.AbortWithStatusJSON(500, gin.H{
			"message": "cannot add publisher to database",
		})
		return
	}

	c.Status(200)
}

func DeletePublisher(c *gin.Context, m *database.DbManager) {
	type URI struct {
		Id int `json:"id" uri:"id"`
	}

	uri := URI{}
	if err := c.BindUri(&uri); err != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"message": "cannot parse url parameter",
		})
		return
	}

	if _, err := m.DeletePublisher(uri.Id); err != nil {
		c.AbortWithStatusJSON(500, gin.H{
			"message": "Verkündiger konnte nicht gelöscht werden",
		})
		return
	}
	c.Status(200)
}
