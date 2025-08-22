package repository

import (
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"user-service/internal/domain"
)

type usersPostgresStore struct {
	db *pgxpool.Pool
}

func NewUsersPostgresStore(db *pgxpool.Pool) domain.UserRepo {
	return &usersPostgresStore{
		db: db,
	}
}

func (store *usersPostgresStore) CreateUser(user *domain.CreateUserDTO) (domain.User, error) {
	panic("not implemented")
}

func (store *usersPostgresStore) GetUser(id uuid.UUID) (domain.User, error) {
	panic("not implemented")
}

func (store *usersPostgresStore) GetUsers() ([]domain.User, error) {
	panic("not implemented")
}

func (store *usersPostgresStore) UpdateUser(id uuid.UUID, updateUser domain.UpdateUserDTO) (domain.User, error) {
	panic("not implemented")
}

func (store *usersPostgresStore) UpdateRoleUser(id uuid.UUID, role domain.UpdateRoleUserDTO) (domain.User, error) {
	panic("not implemented")
}

func (store *usersPostgresStore) RemoveUser(id uuid.UUID) error {
	panic("not implemented")
}
