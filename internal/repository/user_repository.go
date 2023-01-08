package repository

import (
	"context"
	"database/sql"
	"github.com/arvians-id/go-clean-architecture/internal/model"
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
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return UserRepository{
		DB: db,
	}
}

func (repository *UserRepository) FindAll(ctx context.Context) ([]*model.User, error) {
	query := "SELECT * FROM users ORDER BY id DESC"
	rows, err := repository.DB.QueryContext(ctx, query)
	if err != nil {
		log.Println("[UserRepository][FindAll] problem querying to db, err: ", err.Error())
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Println("[UserRepository][FindAll] problem closing db rows, err: ", err.Error())
			return
		}
	}(rows)

	var users []*model.User
	for rows.Next() {
		var user model.User
		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			log.Println("[UserRepository][FindAll] problem with scanning db row, err: ", err.Error())
			return nil, err
		}
		users = append(users, &user)
	}

	return users, nil
}

func (repository *UserRepository) FindByID(ctx context.Context, id int64) (*model.User, error) {
	query := "SELECT * FROM users WHERE id = $1"
	row := repository.DB.QueryRowContext(ctx, query, id)

	var user model.User
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		log.Println("[UserRepository][FindByID] problem with scanning db row, err: ", err.Error())
		return nil, err
	}

	return &user, nil
}

func (repository *UserRepository) Create(ctx context.Context, user *model.User) (*model.User, error) {
	query := "INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING id, created_at, updated_at"
	row := repository.DB.QueryRowContext(ctx, query, user.Name, user.Email, user.Password)

	var id int64
	err := row.Scan(&id, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		log.Println("[UserRepository][Create] problem with scanning db row, err: ", err.Error())
		return nil, err
	}

	user.ID = id

	return user, nil
}

func (repository *UserRepository) Update(ctx context.Context, user *model.User) (*model.User, error) {
	query := "UPDATE users SET name = $1, password = $2, updated_at = $3 WHERE id = $4"
	_, err := repository.DB.ExecContext(ctx, query, user.Name, user.Password, user.UpdatedAt, user.ID)
	if err != nil {
		log.Println("[UserRepository][Update] problem querying to db, err: ", err.Error())
		return nil, err
	}

	return user, nil
}

func (repository *UserRepository) Delete(ctx context.Context, id int64) error {
	query := "DELETE FROM users WHERE id = $1"
	_, err := repository.DB.ExecContext(ctx, query, id)
	if err != nil {
		log.Println("[UserRepository][Delete] problem querying to db, err: ", err.Error())
		return err
	}

	return nil
}
