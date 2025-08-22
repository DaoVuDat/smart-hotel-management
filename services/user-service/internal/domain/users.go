package domain

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID          uuid.UUID
	Email       string
	Name        string
	Phone       string
	Role        string
	Status      string
	DateOfBirth string
	CreatedAt   *time.Time
	UpdatedAt   *time.Time
}

type UserRepo interface {
	CreateUser(user *CreateUserDTO) (User, error)
	GetUser(id uuid.UUID) (User, error)
	GetUsers() ([]User, error)
	UpdateUser(id uuid.UUID, updateUser UpdateUserDTO) (User, error)
	UpdateRoleUser(id uuid.UUID, role UpdateRoleUserDTO) (User, error)
	RemoveUser(id uuid.UUID) error
}

type UserSvc interface {
	CreateUser(user *CreateUserDTO) (UserDTO, error)
	GetUser(id uuid.UUID) (UserDTO, error)
	GetUsers() ([]UserDTO, error)
	UpdateUser(id uuid.UUID, updateUser UpdateUserDTO) (UserDTO, error)
	UpdateRoleUser(id uuid.UUID, role UpdateRoleUserDTO) (UserDTO, error)
	RemoveUser(id uuid.UUID) error
}
