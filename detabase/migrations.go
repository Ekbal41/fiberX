package detabase

import (
	models "github.com/ekbal41/fiberX/detabase/model"
)

// this function is called in init function in main.go
func InitMigration() {
	// Migrate the schema
	DB.AutoMigrate(&models.User{})
}
