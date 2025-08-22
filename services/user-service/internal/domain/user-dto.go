package domain

import "github.com/google/uuid"

type UserDTO struct {
	ID          uuid.UUID `json:"id"`
	Email       string    `json:"email"`
	Name        string    `json:"name"`
	Phone       string    `json:"phone"`
	Role        string    `json:"role"`
	DateOfBirth string    `json:"dateOfBirth"`
}

type CreateUserDTO struct {
	Email       string `json:"email"`
	Name        string `json:"name"`
	Phone       string `json:"phone"`
	Role        string `json:"role"`
	DateOfBirth string `json:"dateOfBirth"`
}

type UpdateUserDTO struct {
	Email       *string `json:"email"`
	Name        *string `json:"name"`
	Phone       *string `json:"phone"`
	DateOfBirth *string `json:"dateOfBirth"`
}

type UpdateRoleUserDTO struct {
	Role string `json:"role"`
}
