package repository

import (
	"context"
	"github.com/arvians-id/go-clean-architecture/internal/model"
	"gorm.io/gorm"
	"log"
)

type UserRepositoryContract interface {
	FindAll(ctx context.Context) ([]*model.User, error)
	FindByID(ctx context.Context, id int64) (*model.User, error)
	Create(ctx context.Context, user *model.User) (*model.User, error)
	Update(ctx context.Context, user *model.User) (*model.User, error)
	Delete(ctx context.Context, id int64) error
}

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepositoryContract {
	return &UserRepository{
		DB: db,
	}
}

func (repository *UserRepository) FindAll(ctx context.Context) ([]*model.User, error) {
	var users []*model.User
	err := repository.DB.WithContext(ctx).Find(&users).Error
	if err != nil {
		log.Println("[UserRepository][FindAll] problem with scanning db row, err: ", err.Error())
		return nil, err
	}

	return users, nil
}

func (repository *UserRepository) FindByID(ctx context.Context, id int64) (*model.User, error) {
	var user model.User
	err := repository.DB.WithContext(ctx).First(&user, id).Error
	if err != nil {
		log.Println("[UserRepository][FindByID] problem with scanning db row, err: ", err.Error())
		return nil, err
	}

	return &user, nil
}

func (repository *UserRepository) Create(ctx context.Context, user *model.User) (*model.User, error) {
	err := repository.DB.WithContext(ctx).Create(user).Error
	if err != nil {
		log.Println("[UserRepository][Create] problem with scanning db row, err: ", err.Error())
		return nil, err
	}

	return user, nil
}

func (repository *UserRepository) Update(ctx context.Context, user *model.User) (*model.User, error) {
	err := repository.DB.WithContext(ctx).Select("name", "password").Updates(&user).Error
	if err != nil {
		log.Println("[UserRepository][Update] problem querying to db, err: ", err.Error())
		return nil, err
	}

	return user, nil
}

func (repository *UserRepository) Delete(ctx context.Context, id int64) error {
	var user model.User
	err := repository.DB.WithContext(ctx).Delete(&user, id).Error
	if err != nil {
		log.Println("[UserRepository][Delete] problem querying to db, err: ", err.Error())
		return err
	}

	return nil
}
