package models

type UserView struct {
	ID         int    `json:"id" gorm:"primaryKey;autoIncrement"`
	Code       string `json:"code" gorm:"type:varchar(50);not null"`
	Username   string `json:"username" gorm:"type:varchar(30);not null;unique"`
	Name       string `json:"name" gorm:"type:varchar(50);not null"`
	Gender     string `json:"gender" gorm:"type:char(1);null"`
	Address    string `json:"address" gorm:"type:varchar(100);null"`
	Occupation string `json:"occupation" gorm:"type:varchar(50);null"`
	RoleID     int    `json:"role_id" gorm:"type:int;not null"`
	RoleName   string `json:"role_name"`
}