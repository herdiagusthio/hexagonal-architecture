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

	userController "hexagonalArchitecture/api/v1/user"
	userService "hexagonalArchitecture/business/user"
	userRepository "hexagonalArchitecture/repository/user"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	echo "github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func newDatabaseConnection(confg *config.AppConfig) *gorm.DB {
	configDB := map[string]string{
		"DB_Username": os.Getenv("GOHEXAGONAL_DB_USERNAME"),
		"DB_Password": os.Getenv("GOHEXAGONAL_DB_PASSWOR"),
		"DB_Port":     os.Getenv("GOHEXAGONAL_DB_PORT"),
		"DB_Host":     os.Getenv("GOHEXAGONAL_DB_HOST"),
		"DB_Name":     os.Getenv("GOHEXAGONAL_DB_NAME"),
	}

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		configDB["DB_Username"],
		configDB["DB_Passwor"],
		configDB["DB_Host"],
		configDB["DB_Port"],
		configDB["DB_Name"])

	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	migration.InitMigrate(db)

	return db
}

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

	//create echo htp
	e := echo.New()

	//register API path and handler
	api.RegisterPath(e, userController)

	// run servr
	go func() {
		address := fmt.Sprintf("localhost:%d", config.AppPort)

		if err := e.Start(address); err != nil {
			log.Info("shutting down the server")
		}
	}()

	// Wait for interrupt signalto gracefully shutdown the server with
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	// a timeout of 10 seconds to shutdown the server
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}
}
