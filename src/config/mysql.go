package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

var DB *gorm.DB

type DBConfig struct {
	Host     string
	Port     int
	User     string
	DBName   string
	Password string
}

func BuildDBConfig() *DBConfig {
	port, _ := strconv.Atoi(GetInv("MYSQL_PORT"))
	dbConfig := DBConfig{
		Host:     GetInv("MYSQL_HOST"),
		Port:     port,
		User:     GetInv("MYSQL_USER"),
		DBName:   GetInv("MYSQL_DBNAME"),
		Password: GetInv("MYSQL_PASSWORD"),
	}

	return &dbConfig
}

func DbURL(dbConfig *DBConfig) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
	)
}

func GetInv(key string) string {
	err := godotenv.Load(".env")

	if err != nil {
		panic(err)
	}

	return os.Getenv(key)
}
