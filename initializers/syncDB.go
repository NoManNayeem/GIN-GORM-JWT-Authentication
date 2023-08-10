package initializers

import "example/go-auth/models"

func SyncDB() {
	DB.AutoMigrate(&models.User{})
}
