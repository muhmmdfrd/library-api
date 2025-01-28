package users

import (
	"library-api/entities"
	common "library-api/models"
	"library-api/models/requests"
	"library-api/modules/users/models"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRepo interface {
	GetByAuth(request requests.AuthRequest) (models.UserView, error)
	GetPaged(filter common.TableFilter) (common.PaginationData, error)
	Create(add models.UserAdd) (models.UserView, error)
}

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *userRepo {
	return &userRepo{db: db}
}


func (r *userRepo) GetByAuth(request requests.AuthRequest) (models.UserView, error) {
	var user entities.User
	var view models.UserView

	err := r.getUser().Where("username = ?", request.Username).First(&user).Error
	if err != nil {
		return view, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		return view, err
	}

	view = models.UserView{ID: user.ID, Code: user.Code, Username: user.Username, Name: user.Name, Gender: user.Gender, Address: user.Address, Occupation: user.Occupation, RoleID: user.RoleID, RoleName: user.Role.Name}

	return view, nil
}

func (r *userRepo) GetPaged(filter common.TableFilter) (common.PaginationData, error) {
	var entities []entities.User
	var view []models.UserView
	var result common.PaginationData
	var count int64

	r.getUser().Count(&count)

	offset := (filter.Index - 1) * filter.Size
	fields := "`user`.`id`,`user`.`code`,`user`.`username`,`user`.`name`,`user`.`gender`,`user`.`address`,`user`.`occupation`,`user`.`role_id`,`role`.`name`"
	join := "JOIN role ON role.id = `user`.role_id"
	query := r.getUser().Select(fields).Joins(join).Limit(filter.Size).Offset(offset).Find(&view)

	if len(filter.Keyword) > 0 {
		keyword := "%" + filter.Keyword + "%"
		query = query.Where("user.name LIKE ?", keyword).Or("user.username LIKE ?", keyword).Find(&view)
		query.Count(&count)
	}

	err := query.Find(&entities).Error
	if err != nil {
		return result, err
	}

	result = common.PaginationData{
		Size:     filter.Size,
		Filtered: len(view),
		Data:     view,
		Total:    int(count),
	}

	return result, nil
}

func (r *userRepo) Create(add models.UserAdd) (models.UserView, error) {
	var view models.UserView

	err := r.db.Table("user").Create(&add).Error
	if err != nil {
		return view, err
	}

	r.getUser().Where(&models.UserView{ID: add.ID}).Model(&view)

	return view, nil
}

func (r *userRepo) getUser() (*gorm.DB) {
	return r.db.Table("user").Where("`user`.is_active")
}