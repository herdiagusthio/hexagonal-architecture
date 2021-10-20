package migration

import (
	"hexagonalArchitecture/repository/user"

	"gorm.io/gorm"
)

func InitMigrate(db *gorm.DB) {
	db.AutoMigrate(&user.UserTable{})
}
