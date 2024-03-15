package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	FistName  string         `json:"first_name" validate:"required"`
	LastName  string         `json:"last_name" validate:"required"`
	Email     string         `json:"email" validate:"required,email"`
	Password  string         `json:"password,omitempty" validate:"required"`
	RoleID    uint           `json:"role_id" validate:"required"`
	Role      *Role          `json:"role"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" faker:"-"`
}

func (u *User) HashPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.Password = string(hashedPassword)
	return nil
}

func (u *User) ComparePassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}

func (u *User) ClearPassword() {
	u.Password = ""
}

type UserWithToken struct {
	User  *User  `json:"user"`
	Token string `json:"token"`
}
