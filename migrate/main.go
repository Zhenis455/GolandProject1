package migrate

import (
	"log"

	"github.com/Zhenis455/Goland1/initializer"
	"github.com/Zhenis455/Goland1/model"
)

func RunMigrations() {
	err := initializer.DB.AutoMigrate(&model.User{})
	if err != nil {
		log.Fatal("Migration failed:", err)
	}
	log.Println("Migrations applied")
}
