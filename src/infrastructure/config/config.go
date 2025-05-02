package config

import (
	"be/src/domain/models"
	"fmt"
	"log"
	"sync"

	"github.com/joho/godotenv"
	"github.com/sakirsensoy/genv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type appConfig struct {
	Port       int
	Debug      bool
	DBPort     int
	DBHost     string
	DBUser     string
	DBPassword string
	DBName     string
	JWTSecret  string
}

var (
	App  *appConfig
	DB   *gorm.DB
	once sync.Once
)

// LoadConfig initializes the application configuration and database connection
func LoadConfig() error {
	var err error
	once.Do(func() {
		err = godotenv.Load()
		if err != nil {
			log.Println("Error loading .env file:", err)
		}

		App = &appConfig{
			Port:       genv.Key("PORT").Default(8080).Int(),
			Debug:      genv.Key("DEBUG").Default(false).Bool(),
			DBPort:     genv.Key("DB_PORT").Default(5432).Int(),
			DBHost:     genv.Key("DB_HOST").Default("localhost").String(),
			DBUser:     genv.Key("DB_USER").Default("gorm").String(),
			DBName:     genv.Key("DB_NAME").Default("gorm").String(),
			DBPassword: genv.Key("DB_PASSWORD").String(),
			JWTSecret:  genv.Key("JWT_SECRET").String(),
		}

		// Validate critical environment variables
		if App.DBPassword == "" {
			err = fmt.Errorf("DB_PASSWORD is required")
			return
		}

		err = initDB()
		if err != nil {
			log.Println("Failed to connect to database:", err)
		}
	})

	return err
}

// initDB initializes the database connection and runs migrations
func initDB() error {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai",
		App.DBHost, App.DBUser, App.DBPassword, App.DBName, App.DBPort)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("could not open database connection: %v", err)
	}

	// Apply migrations
	if err := db.AutoMigrate(
		&models.User{},
		&models.Project{},
		&models.Task{},
		&models.Note{},
		&models.Timeline{},
		&models.Comment{},
	); err != nil {
		return fmt.Errorf("could not migrate database: %v", err)
	}

	DB = db
	return nil
}
