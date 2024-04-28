package main

import (
	"go-core/config"
	"go-core/databases/models"
	"go-core/services"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	initUser()
	initPermissisons()
}

func initUser() {
	godotenv.Load()
	db := config.GetDatabaseInstance()

	username := os.Getenv("ADMIN_USERNAME")
	password, err := services.HashPassword(os.Getenv("ADMIN_PASSWORD"))
	if err == nil {
		var user models.User
		rs := db.Where("username = ?", username).First(&user)
		if rs.Error != nil {
			user = models.User{
				Username:    username,
				Password:    password,
				FullName:    "Administrator",
				AccountType: config.SUPER_ADMIN,
				Status:      config.USER_VISIBLE,
			}

			rs := db.Create(&user)
			if rs.Error == nil {
				log.Println("Create user successfull !")
			}
		}
	}
}

func initPermissisons() {
	modules := config.GetPermissionsConfig()
	db := config.GetDatabaseInstance()
	for _, module := range modules {
		itemModule := models.Module{Name: module.Name, Description: module.Description}
		result := db.FirstOrCreate(&itemModule, models.Module{Name: module.Name})
		if result.Error != nil {
			log.Println("Error save module: ", result.Error.Error())
			continue
		}

		for _, permission := range module.Permissisons {
			itemPermission := models.Permission{
				Name:        permission.Name,
				Description: permission.Description,
				Module:      &itemModule,
			}
			result := db.FirstOrCreate(&itemPermission, models.Permission{Name: permission.Name})
			if result.Error != nil {
				log.Println("Error save permission: ", permission.Name, result.Error.Error())
			}
		}
	}
}
