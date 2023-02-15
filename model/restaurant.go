package model

import "time"

type Restaurant struct {
	Id            string    `json:"id"`
	Cuisine       Cuisine   `json:"cuisine"`
	CostBracket   int       `json:"cost_bracket"`
	Rating        float32   `json:"rating"`
	IsRecommended bool      `json:"is_recommended"`
	OnboardedTime time.Time `json:"onboarded_time"`
}

func (r Restaurant) IsNewlyCreated() bool {
	if r.OnboardedTime.After(time.Now().Add(-2 * 24 * time.Hour)) {
		return true
	}
	return false
}
