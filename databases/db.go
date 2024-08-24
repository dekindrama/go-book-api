package databases

import (
	"log"

	"github.com/dekindrama/go-book-api/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Mysql *gorm.DB

func NewMysql() {
	//* init instance
	var err error
	Mysql, err = gorm.Open(mysql.Open("root:@tcp(127.0.0.1:3306)/testing?parseTime=true"), &gorm.Config{})
	if err != nil {
		log.Fatalf("could not connect to the database: %v", err)
	}

	//* Perform auto-migration
	err = Mysql.AutoMigrate(&models.BookModel{})
	if err != nil {
		panic("failed to migrate database")
	}
}
