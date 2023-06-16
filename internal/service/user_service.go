package service

import (
	"context"
	"github.com/arvians-id/go-clean-architecture/internal/http/presenter/request"
	"github.com/arvians-id/go-clean-architecture/internal/model"
	"github.com/arvians-id/go-clean-architecture/internal/repository"
	"github.com/arvians-id/go-clean-architecture/util"
)

type UserServiceContract interface {
	FindAll(ctx context.Context) ([]*model.User, error)
	FindByID(ctx context.Context, id int64) (*model.User, error)
	Create(ctx context.Context, request *request.CreateUserRequest) (*model.User, error)
	Update(ctx context.Context, request *request.UpdateUserRequest) (*model.User, error)
	Delete(ctx context.Context, id int64) error
}

type UserService struct {
	UserRepository repository.UserRepositoryContract
}

func NewUserService(userRepository repository.UserRepositoryContract) UserServiceContract {
	return &UserService{
		UserRepository: userRepository,
	}
}

func (service *UserService) FindAll(ctx context.Context) ([]*model.User, error) {
	return service.UserRepository.FindAll(ctx)
}

func (service *UserService) FindByID(ctx context.Context, id int64) (*model.User, error) {
	return service.UserRepository.FindByID(ctx, id)
}

func (service *UserService) Create(ctx context.Context, request *request.CreateUserRequest) (*model.User, error) {
	passwordHashed, err := util.HashPassword(request.Password)
	if err != nil {
		return nil, err
	}

	var user model.User
	user.Name = request.Name
	user.Email = request.Email
	user.Password = passwordHashed

	return service.UserRepository.Create(ctx, &user)
}

func (service *UserService) Update(ctx context.Context, request *request.UpdateUserRequest) (*model.User, error) {
	checkUser, err := service.FindByID(ctx, request.ID)
	if err != nil {
		return nil, err
	}

	newPassword := checkUser.Password
	if request.Password != "" {
		passwordHashed, err := util.HashPassword(request.Password)
		if err != nil {
			return nil, err
		}
		newPassword = passwordHashed
	}

	checkUser.ID = request.ID
	checkUser.Name = request.Name
	checkUser.Password = newPassword

	return service.UserRepository.Update(ctx, checkUser)
}

func (service *UserService) Delete(ctx context.Context, id int64) error {
	_, err := service.FindByID(ctx, id)
	if err != nil {
		return err
	}

	return service.UserRepository.Delete(ctx, id)
}
