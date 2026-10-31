package user

import (
	"errors"
	"fmt"
	"hexagonalArchitecture/business/user"
	"time"

	"gorm.io/gorm"
)

//GormRepository The implementation of user.Repository object
type GormRepository struct {
	DB *gorm.DB
}

type UserTable struct {
	ID          int       `gorm:"id;primaryKey;autoIncrement"`
	Name        string    `gorm:"name"`
	Email       string    `gorm:"email"`
	PhoneNumber string    `gorm:"phone_number"`
	Username    string    `gorm:"username"`
	Password    string    `gorm:"password"`
	CreatedAt   time.Time `gorm:"created_at"`
	UpdatedAt   time.Time `gorm:"updated_at"`
	DeletedAt   time.Time `gorm:"deleted_at"`
}

// ToUser maps a row to the domain model. Stored rows are trusted as-is; validation happens where
// users are created (user.NewUser), so Email and Username convert without re-checking.
func (col *UserTable) ToUser() user.FindUser {
	return user.FindUser{
		ID:          col.ID,
		Name:        col.Name,
		Email:       user.Email(col.Email),
		PhoneNumber: col.PhoneNumber,
		Username:    user.Username(col.Username),
		CreatedAt:   col.CreatedAt,
		UpdatedAt:   col.UpdatedAt,
		DeletedAt:   &col.DeletedAt,
	}
}

//NewGormDBRepository Generate Gorm DB user repository
func NewGormDBRepository(db *gorm.DB) *GormRepository {
	return &GormRepository{
		db,
	}
}

//FindUserByID If data not found will return nil without error
func (r *GormRepository) FindUserByID(id int) (*user.FindUser, error) {
	var userData UserTable

	err := r.DB.First(&userData, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, user.ErrUserNotFound
		}
		return nil, fmt.Errorf("database error: %w", err)
	}

	user := userData.ToUser()

	return &user, nil
}
