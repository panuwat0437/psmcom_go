package api

import (
	"main/db"
	"main/interceptor"
	"main/model"

	"github.com/gin-gonic/gin"
)

// SetupProductAPI - call this method to setup product route group
func SetupScoreAPI(router *gin.Engine) {
	ScoreAPI := router.Group("/api/v1")
	{
		ScoreAPI.GET("/score/:id", interceptor.JwtVerify, getScoreByID)

	}
}

type ResponseScore struct {
	StudentID             string  `json:"studentid"`
	Name                  string  `json:"name"`
	Year                  string  `json:"year"`
	LineUpCount           int64   `json:"lineupcount"`
	LineUpTotalScore      float64 `json:"lineuptotalscore"`
	MissLineUpCount       int64   `json:"misslineupcount"`
	MissLineUpTotalScore  float64 `json:"misslineuptotalscore"`
	LeaveCount            int64   `json:"leavecount"`
	BehavedWellCount      int64   `json:"behavedwellcount"`
	BehavedWellTotalScore float64 `json:"behavedwelltotalscore"`
	BehavedBadCount       int64   `json:"behavedbadcount"`
	BehavedBadTotalScore  float64 `json:"behavedbadtotalscore"`
	TotalScore            float64 `json:"totalscore"`
	AccumulateScore       float64 `json:"accumulatescore"`
	MonthName             string  `json:"monthname"`
}

func getScoreByID(c *gin.Context) {
	var month = c.Param("id")

	var checkmonth model.Month
	checkmonthresult := db.GetDB().Where("id = ?", month).First(&checkmonth)
	if checkmonthresult.Error != nil {
		c.JSON(404, gin.H{"message": "Data Not Found"})
		return
	}

	var score ResponseScore
	db.GetDB().Raw("SELECT * FROM `scores` LEFT JOIN months ON scores.month_id = months.id WHERE `student_id` = ? AND month_id = ?", c.GetString("jwt_student_id"), month).First(&score)
	c.JSON(200, gin.H{"rows": checkmonthresult.RowsAffected, "data": score})
}
