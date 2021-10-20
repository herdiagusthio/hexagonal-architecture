package user

import (
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

func (col *UserTable) ToUser() user.FindUser {
	var user user.FindUser

	user.ID = col.ID
	user.Name = col.Name
	user.Email = col.Email
	user.PhoneNumber = col.PhoneNumber
	user.Username = col.Username
	user.CreatedAt = col.CreatedAt
	user.UpdatedAt = col.UpdatedAt
	user.DeletedAt = &col.DeletedAt

	return user
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
		return nil, err
	}

	user := userData.ToUser()

	return &user, nil
}
