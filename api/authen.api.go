package api

import (
	"main/db"
	"main/interceptor"
	"main/model"

	"github.com/gin-gonic/gin"
)

func SetupAuthenAPI(router *gin.Engine) {
	authenAPI := router.Group("/api/v1")
	{
		authenAPI.POST("/login", login)
	}
}

func login(c *gin.Context) {
	var score model.Score

	if c.ShouldBind(&score) == nil {
		var queryScore model.Score
		if err := db.GetDB().First(&queryScore, "student_id = ?", score.StudentID).Error; err != nil {
			c.JSON(401, gin.H{"result": "UnAuthorization"})
		} else {
			token := interceptor.JwtSign(queryScore)
			c.JSON(200, gin.H{"result": "Authorization", "token": token})
		}

	} else {
		c.JSON(400, gin.H{"status": "Required Data Not Found"})
	}
}
