package service

import (
	"github.com/google/uuid"
	"user-service/internal/domain"
)

type userSvc struct {
	repo domain.UserRepo
}

func NewUserSvc(repo domain.UserRepo) domain.UserSvc {
	return &userSvc{
		repo: repo,
	}
}

func (svc *userSvc) CreateUser(user *domain.CreateUserDTO) (domain.UserDTO, error) {
	panic("not implemented")
}

func (svc *userSvc) GetUser(id uuid.UUID) (domain.UserDTO, error) {
	panic("not implemented")
}

func (svc *userSvc) GetUsers() ([]domain.UserDTO, error) {
	panic("not implemented")
}

func (svc *userSvc) UpdateUser(id uuid.UUID, updateUser domain.UpdateUserDTO) (domain.UserDTO, error) {
	panic("not implemented")
}

func (svc *userSvc) UpdateRoleUser(id uuid.UUID, role domain.UpdateRoleUserDTO) (domain.UserDTO, error) {
	panic("not implemented")
}

func (svc *userSvc) RemoveUser(id uuid.UUID) error {
	panic("not implemented")
}
