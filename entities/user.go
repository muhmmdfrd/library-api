package entities

import "time"

type User struct {
	ID         int       `json:"id" gorm:"primaryKey;autoIncrement"`
	Code       string    `json:"code" gorm:"type:varchar(50);not null"`
	Username   string    `json:"username" gorm:"type:varchar(30);not null;unique"`
	Password   string    `json:"password" gorm:"type:varchar(50);not null"`
	Name       string    `json:"name" gorm:"type:varchar(50);not null"`
	Gender     string    `json:"gender" gorm:"type:char(1);null"`
	Address    string    `json:"address" gorm:"type:varchar(100);null"`
	Occupation string    `json:"occupation" gorm:"type:varchar(50);null"`
	RoleID     int       `json:"role_id" gorm:"type:int;not null"`
	IsActive   bool      `json:"is_active" gorm:"type:bool;default:true"`
	CreatedBy  int       `json:"created_by" gorm:"type:int;not null;default:1"`
	CreatedAt  time.Time `json:"created_at" gorm:"type:datetime;not null;default:CURRENT_TIMESTAMP"`
	UpdatedBy  int       `json:"updated_by" gorm:"type:int;not null;default:1"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"type:datetime;not null;default:CURRENT_TIMESTAMP"`
	DeletedAt time.Time  `json:"deleted_at" gorm:"type:datetime;null"`

	Role Role `gorm:"foreignKey:RoleID"`
}