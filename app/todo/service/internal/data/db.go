package data

import (
	"fmt"
	"github.com/ljxsteam/coinside-backend-kratos/app/todo/service/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"sync"
)

var _db *gorm.DB
var once sync.Once

// NewDB 对应db的为singalton模式
func NewDB(config *config.Config) *gorm.DB {
	once.Do(func() {
		host := config.GetString("data.mariadb.host")
		port := config.GetString("data.mariadb.port")
		username := config.GetString("data.mariadb.username")
		password := config.GetString("data.mariadb.password")
		dbName := config.GetString("data.mariadb.db_name")

		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, dbName)
		var err error
		_db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

		if err != nil {
			log.Fatalf("Fatal error open database:%s %s \n", dsn, err)
		} else {
			log.Printf("Opened database:%s %s \n", dsn, err)
		}
	})

	return _db
}
