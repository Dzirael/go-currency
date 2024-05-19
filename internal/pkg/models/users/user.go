package users

import (
	"time"

	"github.com/Dzirael/go-curenncy/internal/pkg/models"
)

type User struct {
	models.Model
	Email        string `gorm:"column:email;not null;unique_index:email" json:"email" form:"email"`
	IsSubscribed bool   `gorm:"column:isSubscribed;not null;" json:"isSubscribed" form:"isSubscribed"`
}

func (m *User) BeforeCreate() error {
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	return nil
}

func (m *User) BeforeUpdate() error {
	m.UpdatedAt = time.Now()
	return nil
}
