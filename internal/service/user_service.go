package service

import (
	"context"
	"github.com/arvians-id/go-clean-architecture/internal/http/presenter/request"
	"github.com/arvians-id/go-clean-architecture/internal/http/presenter/response"
	"github.com/arvians-id/go-clean-architecture/internal/model"
	"github.com/arvians-id/go-clean-architecture/internal/repository"
	"github.com/arvians-id/go-clean-architecture/util"
)

type UserServiceContract interface {
	FindAll(ctx context.Context) ([]*response.UserResponse, error)
	FindByID(ctx context.Context, id int64) (*response.UserResponse, error)
	Create(ctx context.Context, request *request.CreateUserRequest) (*response.UserResponse, error)
	Update(ctx context.Context, request *request.UpdateUserRequest) (*response.UserResponse, error)
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

func (service *UserService) FindAll(ctx context.Context) ([]*response.UserResponse, error) {
	users, err := service.UserRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	var userResponses []*response.UserResponse
	for _, user := range users {
		userResponses = append(userResponses, &response.UserResponse{
			ID:        user.ID,
			Name:      user.Name,
			Email:     user.Email,
			CreatedAt: user.CreatedAt.String(),
			UpdatedAt: user.UpdatedAt.String(),
		})
	}

	return userResponses, nil
}

func (service *UserService) FindByID(ctx context.Context, id int64) (*response.UserResponse, error) {
	user, err := service.UserRepository.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return &response.UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt.String(),
		UpdatedAt: user.UpdatedAt.String(),
	}, nil
}

func (service *UserService) Create(ctx context.Context, request *request.CreateUserRequest) (*response.UserResponse, error) {
	passwordHashed, err := util.HashPassword(request.Password)
	if err != nil {
		return nil, err
	}

	user, err := service.UserRepository.Create(ctx, &model.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: passwordHashed,
	})
	if err != nil {
		return nil, err
	}

	return &response.UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt.String(),
		UpdatedAt: user.UpdatedAt.String(),
	}, nil
}

func (service *UserService) Update(ctx context.Context, request *request.UpdateUserRequest) (*response.UserResponse, error) {
	checkUser, err := service.UserRepository.FindByID(ctx, request.ID)
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

	user, err := service.UserRepository.Update(ctx, checkUser)
	if err != nil {
		return nil, err
	}

	return &response.UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt.String(),
		UpdatedAt: user.UpdatedAt.String(),
	}, nil
}

func (service *UserService) Delete(ctx context.Context, id int64) error {
	_, err := service.UserRepository.FindByID(ctx, id)
	if err != nil {
		return err
	}

	return service.UserRepository.Delete(ctx, id)
}
