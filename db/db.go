package db

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego-develop-1.x/configs"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"io/ioutil"
	"os"
	"time"
)

var DB *gorm.DB

func ReadConfigFromFile(path string, config interface{}) error {
	dir, _ := os.Getwd()
	fmt.Printf("ReadConfigFromFile:%v/n", dir)

	file, _ := os.Open(path)
	defer file.Close()
	bytes, _ := ioutil.ReadAll(file)
	return json.Unmarshal(bytes, &config)
}

func InitDB() {
	var conf configs.DB
	if err := ReadConfigFromFile("configs/db.json", &conf); err != nil {
		DB, err := gorm.Open(mysql.New(mysql.Config{
			DSN:                       conf.DSN,
			DefaultStringSize:         256,
			DisableDatetimePrecision:  true,
			DontSupportRenameIndex:    true,
			DontSupportRenameColumn:   true,
			SkipInitializeWithVersion: false,
		}), &gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				TablePrefix:   "",
				SingularTable: true,
			},
			Logger: logger.Default.LogMode(logger.Info),
		})
		if err != nil {
			panic(err)
		}
		sqlDB, _ := DB.DB()
		sqlDB.SetMaxIdleConns(10)
		sqlDB.SetMaxOpenConns(100)
		sqlDB.SetConnMaxLifetime(time.Hour)
	}

}
