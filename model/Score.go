package model

type Score struct {
	ID                    uint    `json:"id"; gorm:"primary_key"`
	StudentID             string  `json:"studentid" binding:"required"`
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
	MonthID               string  `json:"monthid"`
}
