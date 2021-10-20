package main

import (
	"context"
	"fmt"
	"hexagonalArchitecture/api"
	"hexagonalArchitecture/config"
	"os"
	"os/signal"
	"time"

	"hexagonalArchitecture/repository/migration"

	"hexagonalArchitecture/api"
	userController "hexagonalArchitechture/api/v1/use"
	userService "hexagonalArchitecture/business/user"
	userRepository "hexagonalArchitecture/repository/user"

	"gorm.io/drive/mysql"
	"gorm.io/gorm"

	echo "github.com/labstack/echo/v"
	"github.com/labstack/gommon/log"


func newDatabaseConnection(confg *config.AppConfig) *gorm.DB {
	configDB := map[string]string{
		"DB_Username": os.Getenv("GOHEXAGONAL_DB_USERNAME"),
		"DB_Password": os.Getenv("GOHEXAGONAL_DB_PASSWOR"),
		"DB_Port":     os.Getenv("GOHEXAGONAL_DB_PORT"),
		"DB_Host":     os.Getenv("GOHEXAGONAL_DB_HOST"),
		"DB_Name":     os.Getenv("GOHEXAGONAL_DB_NAME"),
}

	connectionString := fmt.Srintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		configDB["DB_Username"],
		configDB["DB_Passwor"],
		configDB["DB_Host"],
		configDB["DB_Port"],
	configDB["DB_Name"])

	db, err := gormOpen(mysql.Open(connectionString), &gorm.Config{})
	if err != nl {
		anic(err)
}

migration.InitMigrate(db)

	eturn db


func main() {
	//load config if available o set to default
config := config.GetConfig()

	//initialize database connection based on givn config
dbConnection := newDatabaseConnection(config)

	//initiate user repository
userRepo := userRepository.NewGormDBRepository(dbConnection)

	//initiate user service
userService := userService.NewService(userRepo)

	//initiate user controller
userController := userController.NewController(userService)

	//initiate pet repository
// petRepo := petRepository.NewGormDBRepository(dbConnection)

	//initiate pet service
// petService := petService.NewService(petRepo)

	//initiate pet controller
// petController := petController.NewController(petService)

	//initiate auth service
// authService := authService.NewService(userService)

	//initiate auth controller
// authController := authController.NewController(authService)

	//create echo htp
e := echo.New()

	//register API path and handler
api.RegisterPath(e, userController)

	// run servr
	go func() {
	address := fmt.Sprintf("localhost:%d", config.AppPort)

		if err := e.Start(address); err != ni {
			og.Info("shutting down the server")
		}
}()

	// Wait for interrupt signalto gracefully shutdown the server with
	quit := make(chan os.Signal)
	signalNotify(quit, os.Interrupt)
<-quit

	// a timeout of 10 seconds to shutdown the server
	ctx, cancel :=context.WithTimeout(context.Background(), 10*time.Second)
defer cancel()

	if err := e.Shudown(ctx); err != nil {
		og.Fatal(err)
	
}
