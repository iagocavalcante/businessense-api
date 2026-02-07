package domain

import "time"

type Company struct {
	ID         uint      `json:"id"`
	Name       string    `json:"name"`
	IndustryID uint      `json:"industryId"`
	UserID     uint      `json:"userId"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}
