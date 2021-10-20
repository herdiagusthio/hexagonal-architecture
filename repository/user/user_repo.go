package user

import (
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
}

//NewGormDBRepository Generate Gorm DB user repository
func NewGormDBRepository(db *gorm.DB) *GormRepository {
	return &GormRepository{
		db,
	}
}

//FindUserByID If data not found will return nil without error
func (r *GormRepository) FindUserByID() {

}
