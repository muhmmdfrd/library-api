package users

import (
	"crypto/md5"
	"fmt"
	"library-api/entities"
	"library-api/models/requests"
	"library-api/modules/users/models"
	"time"

	"gorm.io/gorm"
)

type UserRepo interface {
	Auth(request requests.AuthRequest) (models.UserView, error)
	Get() ([]models.UserView, error)
	CreateTemp() (bool, error)
}

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *userRepo {
	return &userRepo{db: db}
}

func (r *userRepo) Auth(request requests.AuthRequest) (models.UserView, error) {
	var user entities.User
	var view models.UserView

	err := r.db.Table("user").Model(&user).Where("username = ? AND password = ?", request.Username, request.Password).First(&view).Error
	if err != nil {
		return view, err
	}

	return view, nil
}

func (r *userRepo) Get() ([]models.UserView, error) {
	var users []entities.User
	var view []models.UserView

	err := r.db.Table("user").Model(&users).Find(&view).Error
	if err != nil {
		return nil, err
	}
	return view, nil
}

func (r *userRepo) CreateTemp() (bool, error) {
	var users []models.UserAdd

	for i := 103; i <= 5000; i++ {
		user := models.UserAdd{
			Code:       generateUUID(), // Fungsi untuk menghasilkan UUID
			Username:   fmt.Sprintf("user%d", i),
			Password:   fmt.Sprintf("%x", md5Hash(fmt.Sprintf("password%d", i))), // Hash MD5
			Name:       fmt.Sprintf("Name %d", i),
			Gender:     getGender(i), // Gender bergantian antara 'M' dan 'F'
			Address:    fmt.Sprintf("Address %d", i),
			Occupation: fmt.Sprintf("Occupation %d", i),
			RoleID:     1,
		}
		users = append(users, user)	
	}

	result := r.db.Table("user").Create(&users)
	if result.Error != nil {
		return false, result.Error
	}

	return true, nil
}

func generateUUID() string {
	return fmt.Sprintf("%d", time.Now().UnixNano())
}

func md5Hash(text string) []byte {
	hash := md5.Sum([]byte(text))
	return hash[:]
}

func getGender(i int) string {
	if i % 2 == 0 {
		return "F"
	}
	return "M"
}