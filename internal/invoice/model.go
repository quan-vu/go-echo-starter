package invoice

import (
	"time"

	"gorm.io/gorm"
)

type Invoice struct {
	ID           uint           `json:"id" gorm:"primaryKey"`
	CustomerName string         `json:"customer_name"`
	Amount       float64        `json:"amount"`
	DueDate      time.Time      `json:"due_date"`
	Status       string         `json:"status"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`
}
