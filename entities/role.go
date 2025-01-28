package entities

import "time"

type Role struct {
	ID        int       `json:"id" gorm:"primaryKey;autoIncrement"`
	Code      string    `json:"code" gorm:"type:varchar(50);not null"`
	Name      string    `json:"name" gorm:"type:varchar(50);not null"`
	IsActive  bool      `json:"is_active" gorm:"type:bool;not null;default:true"`
	CreatedBy int       `json:"created_by" gorm:"type:int;not null;default:1"`
	CreatedAt time.Time `json:"created_at" gorm:"type:datetime;not null;default:CURRENT_TIMESTAMP"`
	UpdatedBy int       `json:"updated_by" gorm:"type:int;not null;default:1"`
	UpdatedAt time.Time `json:"updated_at" gorm:"type:datetime;not null;default:CURRENT_TIMESTAMP"`
}