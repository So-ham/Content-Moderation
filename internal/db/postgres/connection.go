package postgres

import (
	"os"
	"time"

	"github.com/So-ham/Content-Moderation/internal/entities"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Connect establishes a connection to PostgreSQL using environment variables
// and returns a configured GORM DB instance
func Connect() *gorm.DB {
	connectionString := os.Getenv("DATABASE_URL")
	if connectionString == "" {
		panic("DATABASE_URL environment variable is not set")
	}

	// Open connection to database
	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("failed to connect database: " + err.Error())
	}

	// Get underlying SQL database to configure connection pool
	sqlDB, err := db.DB()
	if err != nil {
		panic("failed to get database instance: " + err.Error())
	}

	// Configure connection pool settings
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	// Verify connection
	if err := sqlDB.Ping(); err != nil {
		panic("failed to ping database: " + err.Error())
	}

	// AutoMigrate all entities
	if err := db.AutoMigrate(
		&entities.User{},
		&entities.Post{},
		&entities.Review{},
		&entities.Moderation{},
		&entities.Comment{},
	); err != nil {
		panic("failed to auto-migrate database: " + err.Error())
	}

	return db
}
