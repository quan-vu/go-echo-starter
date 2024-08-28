package models

import (
	"time"

	"gorm.io/gorm"
)

type DateField struct {
	time.Time
}

func (ct *DateField) UnmarshalJSON(b []byte) error {
	var err error
	// Try parsing the date in the format "2006-01-02"
	ct.Time, err = time.Parse(`"2006-01-02"`, string(b))
	if err != nil {
		// If it fails, try parsing the date in the format "2006-01-02 15:04:05 -0700 UTC"
		ct.Time, err = time.Parse(`"2006-01-02 15:04:05 -0700 UTC"`, string(b))
	}
	return err
}

type Invoice struct {
	ID           uint           `json:"id" gorm:"primaryKey"`
	CustomerName string         `json:"customer_name"`
	Amount       float64        `json:"amount"`
	DueDate      time.Time      `json:"due_date" gorm:"datetime"`
	Status       string         `json:"status"`
	CreatedAt    time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt    time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-" `
}
