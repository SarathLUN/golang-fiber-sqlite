package initializers

import (
	"github.com/SarathLUN/golang-fiber-sqlite/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

var DB *gorm.DB

// ConnectDB connects to the SQLite database
func ConnectDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// implement DB.Logger
	DB.Logger = logger.Default.LogMode(logger.Info)
	log.Println("Running Migration")
	err = DB.AutoMigrate(&models.Note{})
	if err != nil {
		log.Fatalln("Migration Failed: ", err.Error())
	}
	log.Println("🚀 DB connected and Migration Completed")
}
