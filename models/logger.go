package models

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type auditWriter struct {
	auditDb *gorm.DB
}

func (w *auditWriter) Printf(format string, args ...interface{}) {
	if len(args) == 4 {
		w.auditDb.Exec("INSERT INTO audit_log(location, logTime, rowsChanged, sqlQuery) VALUES (@location, @logTime, @rowsChanged, @sqlQuery)",
			sql.Named("location", args[0]),
			sql.Named("logTime", args[1]),
			sql.Named("rowsChanged", args[2]),
			sql.Named("sqlQuery", args[3]),
		)
	} else {
		joined := []string{}
		for _, item := range args {
			joined = append(joined, fmt.Sprintf("%v", item), " | ")
		}
		message := fmt.Sprintf("No. args: %d, Error: %s", len(args), joined)
		w.auditDb.Exec("INSERT INTO audit_log(logError) VALUES (@logError)",
			sql.Named("logError", message))
	}
}

func OpenAuditLogger() (*logger.Interface, error) {

	filePath := "logs.db"

	if _, err := os.Stat(filePath); errors.Is(err, os.ErrNotExist) {

		_, err := os.Create(filePath)
		if err != nil {
			log.Println("sqlite Create error", err)
			return nil, err
		}

		log.Println("Audit log file created")
	} else {

		log.Println("Audit log file exists")
	}

	db, err := gorm.Open(sqlite.Open(filePath), &gorm.Config{})
	if err != nil {
		log.Println("sqlite Open error", err)
		return nil, err
	}

	err = db.Exec(`
	CREATE TABLE IF NOT EXISTS "audit_log" (
		"sysTime" 			DATETIME DEFAULT CURRENT_TIMESTAMP, 
		"logTime" 			VARCHAR(64), 
		"location"			VARCHAR(512), 
		"line" 				VARCHAR(64), 
		"rowsChanged" 		VARCHAR(64), 
		"error" 			VARCHAR(512), 
		"sqlQuery" 			TEXT,
		"logError" 			VARCHAR(512)
	)
	`).Error
	if err != nil {
		log.Fatal(err)
	}

	// Migrate the schema
	// err = db.AutoMigrate(&Entry{})
	// if err != nil {
	// 	log.Println("gorm.AutoMigrate - OpenSqliteLogger error", err)
	// 	return nil, err
	// }

	logWriter := auditWriter{
		auditDb: db,
	}

	logInterface := logger.New(
		&logWriter,
		logger.Config{
			SlowThreshold:        time.Second, // Slow SQL threshold
			LogLevel:             logger.Info, // Log level
			Colorful:             false,       // Disable color
			ParameterizedQueries: false,       // Include parameters into query
		},
	)

	return &logInterface, nil
}
