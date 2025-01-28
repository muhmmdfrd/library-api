package entities

import "time"

type BookRent struct {
	ID         int       `json:"id" gorm:"primaryKey;autoIncrement"`
	Code       string    `json:"code" gorm:"type:varchar(50);not null"`
	StatusID   int       `json:"status_id" gorm:"type:int;not null"`
	UserID     int       `json:"user_id" gorm:"type:int;not null"`
	BookID     int       `json:"book_id" gorm:"type:int;not null"`
	RentDate   time.Time `json:"rent_date" gorm:"type:datetime;not null"`
	RentDay    int       `json:"rent_day" gorm:"type:int;not null"`
	ReturnDate time.Time `json:"return_date" gorm:"type:datetime;not null"`
	IsActive   bool      `json:"is_active" gorm:"type:boolean;default:true"`
	CreatedBy  int       `json:"created_by" gorm:"type:int;not null;default:1"`
	CreatedAt  time.Time `json:"created_at" gorm:"type:datetime;not null;default:CURRENT_TIMESTAMP"`
	UpdatedBy  int       `json:"updated_by" gorm:"type:int;not null;default:1"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"type:datetime;not null;default:CURRENT_TIMESTAMP"`

	User User `gorm:"foreignKey:UserID"`
	Book Book `gorm:"foreignKey:BookID"`
}