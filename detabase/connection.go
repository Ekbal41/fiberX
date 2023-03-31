package detabase

import (
	"log"

	"github.com/ekbal41/fiberX/v1/utils"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() *gorm.DB {
	// load config file
	utils.LoadConfig()
	// connect to database 
	dcn := viper.GetString("DB_DCS")
    dbConnected, err := gorm.Open(postgres.Open(dcn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	log.Println("Database connected")
	// assign db to global variable
	DB = dbConnected
	return dbConnected
}