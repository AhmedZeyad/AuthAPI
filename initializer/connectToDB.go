package initializer

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnecteToDb() {

}

var DB *gorm.DB

func ConnectDB() {

	var err error
	dsn := os.Getenv("DSN")
	fmt.Print(dsn)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("filed to conecte to DB")
	}
}
