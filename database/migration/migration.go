package main

import (
	"go/service1/src/config"
	"go/service1/src/entities"
	"log"

	"github.com/joho/godotenv"
)

type table struct {
	tableName interface{}
}

func registerTable() []*table {
	return []*table{
		{tableName: &entities.StatusUser{}},
		{tableName: &entities.User{}},
		{tableName: &entities.Product{}},
	}
}

func RunMigration() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error when loading .env file")
	}

	config.ConnectDB()

	db := *config.DB

	for _, table := range registerTable() {
		dropTable := db.Migrator().DropTable(table.tableName)

		if dropTable != nil {
			log.Print(err.Error())
		}
	}

	err = db.Migrator().AutoMigrate(&entities.StatusUser{}, &entities.User{}, &entities.Product{})

	if err != nil {
		log.Print(err.Error())
	}

	err = db.Migrator().CreateConstraint(&entities.User{}, "Products")

	if err != nil {
		log.Print(err.Error())
	}

	bool1 := db.Migrator().HasConstraint(&entities.User{}, "Products")
	log.Print("HasConstraint")
	log.Print(bool1)
	bool2 := db.Migrator().HasConstraint(&entities.User{}, "fk_public_users_products")
	log.Print("HasConstraint")
	log.Print(bool2)

	tbStatus := db.Create(&entities.StatusUser{
		Id:     1,
		Status: "admin",
	})

	if tbStatus.Error != nil {
		log.Print(tbStatus.Error.Error())
	}
}

func main() {
	RunMigration()
}
