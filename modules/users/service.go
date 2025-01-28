package users

import (
	common "library-api/models"
	"library-api/models/requests"
	"library-api/modules/users/models"

	"github.com/samborkent/uuidv7"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Auth(request requests.AuthRequest) (models.UserView, error)
	GetPagedUser(filter common.TableFilter) (common.PaginationData, error)
	CreateUser(add models.UserAdd) (models.UserView, error)
}

type userService struct {
	userRepo *userRepo
}

func NewUserService(userRepo *userRepo) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (s *userService) Auth(request requests.AuthRequest) (models.UserView, error) {
	return s.userRepo.GetByAuth(request)
}

func (s *userService) GetPagedUser(filter common.TableFilter) (common.PaginationData, error) {
	return s.userRepo.GetPaged(filter)
}

func (s *userService) CreateUser(add models.UserAdd) (models.UserView, error) {
	var view models.UserView

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(add.Password), bcrypt.DefaultCost)
	if err != nil {
		return view, err
	}

	add.Password = string(hashedPassword)
	add.Code = uuidv7.New().String()

	return s.userRepo.Create(add)
}