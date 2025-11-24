package database

import (
	"fmt"
	"igaming-platform/internal/config"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

var db *gorm.DB

func SetDB(connection *gorm.DB) {
	db = connection
}

func GetDB() *gorm.DB {
	return db
}

func connect(env *config.Env) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true",
		env.DB.Username,
		env.DB.Password,
		env.DB.Host,
		env.DB.Port,
		env.DB.Database,
	)

	connection, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}

	sqlDB, err := connection.DB()
	if err != nil {
		log.Fatal("Failed to get underlying SQL DB:", err)
	}

	err = sqlDB.Ping()
	if err != nil {
		log.Fatal("Database connection failed:", err)
	}

	log.Println("Successfully connected to the database!")

	SetDB(connection)
}

func migrateUp(env *config.Env) {
	dsn := fmt.Sprintf(
		"mysql://%s:%s@tcp(%s:%s)/%s",
		env.DB.Username,
		env.DB.Password,
		env.DB.Host,
		env.DB.Port,
		env.DB.Database,
	)

	migration, err := migrate.New(
		"file://./internal/database/migrations",
		dsn,
	)

	if err != nil {
		log.Fatalf("Failed to create migrate instance: %v", err)
	}

	if err := migration.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Failed to run migrations: %v", err)
	}
}

func Initialize(env *config.Env) {
	connect(env)

	migrateUp(env)
}

func Close() {
	db := GetDB()

	if db == nil {
		return
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Println("Error getting underlying SQL DB:", err)
		return
	}

	if err := sqlDB.Close(); err != nil {
		log.Println("Error closing database connection:", err)
	} else {
		log.Println("Database connection closed.")
	}
}
