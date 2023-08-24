package gormfunc

import (
	"encoding/json"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

type dbConfig struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	DBName   string `json:"dbname"`
}

func (User) TableName() string {
	return "user"
}
func loadConfig() (dbConfig, error) {
	var config dbConfig
	file, err := os.Open("gormfunc/config.json")
	if err != nil {
		return config, err
	}
	defer file.Close()

	if err := json.NewDecoder(file).Decode(&config); err != nil {
		return config, err
	}

	return config, nil
}

func getDSN(config dbConfig) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		config.User, config.Password, config.Host, config.Port, config.DBName)
}

func ConnectDB() (*gorm.DB, error) {
	config, err := loadConfig()
	if err != nil {
		return nil, err
	}
	dsn := getDSN(config)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("connect to database successfully")
	return db, nil
}
