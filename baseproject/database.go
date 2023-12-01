package baseproject

import (
	"baseapp/config"
	"baseapp/models"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/glebarez/sqlite"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const dlm = ""
const inMemoryDatabaseFile = "storage.db"

func ConnectToDatabase(c *config.Config, logWriter io.Writer) (db *gorm.DB, err error) {
    switch c.Database {
    case "mysql":
        db, err = ConnectToDatabaseMySQL(c, logWriter)
    case "sqlite":
    default:
        db, err = ConnectToDatabaseSQLite(c, logWriter)
    }
    return
}

func ConnectToDatabaseMySQL(c *config.Config, logWriter io.Writer) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", c.DatabaseUsername, c.DatabasePassword, c.DatabaseHost, c.DatabasePort, c.DatabaseName)

	newLogger := logger.New(
		log.New(logWriter, dlm, log.LstdFlags),     // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Warn, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      false,       // Include params in the SQL log
			Colorful:                  false,       // Disable color
		},
	)

	db, err := gorm.Open(
		mysql.Open(dsn),
		&gorm.Config{
			Logger: newLogger,
		},
	)
	if err != nil {
		log.Println("gorm.Open mysql error", err)
		panic(err)
	}

	// Migrate the schema
	err = db.AutoMigrate(&models.User{}, &models.SystemRole{}, &models.Token{}, &models.Session{})
	if err != nil {
		log.Println("gorm.AutoMigrate error", err)
		panic(err)
	}

	log.Println("Successfully connected! Database is up ...")

	return db, err
}

func ConnectToDatabaseSQLite(c *config.Config, logWriter io.Writer) (*gorm.DB, error) {

	newLogger := logger.New(
		log.New(logWriter, dlm, log.LstdFlags),     // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Warn, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      false,       // Include params in the SQL log
			Colorful:                  false,       // Disable color
		},
	)

	if _, err := os.Stat(inMemoryDatabaseFile); errors.Is(err, os.ErrNotExist) {
		_, err := os.Create(inMemoryDatabaseFile)
		if err != nil {
			log.Println("sqlite Create error", err)
			return nil, err
		}
		log.Println("Storage db file created")
	} else {
		log.Println("Storage db file exists")
	}

	db, err := gorm.Open(
		sqlite.Open(inMemoryDatabaseFile),
		&gorm.Config{
			Logger: newLogger,
		},
	)
	if err != nil {
		log.Println("gorm.Open sqlite error", err)
		panic(err)
	}

	// Migrate the schema
	err = db.AutoMigrate(&models.User{}, &models.SystemRole{}, &models.Token{}, &models.Session{})
	if err != nil {
		log.Println("gorm.AutoMigrate error", err)
		panic(err)
	}

	log.Println("Successfully connected! Database is up ...")

	return db, err
}
