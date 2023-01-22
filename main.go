package main

import (
	"hardtmann/smartlab/api"
	"log"
	"os"
	"strconv"

	"github.com/spf13/viper"
)

func main() {
	log.Println("start app")

	db_user := os.Getenv("DB_USER")
	db_pass := os.Getenv("DB_PASS")
	db_host := os.Getenv("DB_HOST")
	db_port := os.Getenv("DB_PORT")
	apps_port := os.Getenv("PORT")

	viper.SetDefault("AES_PKEY_32BIT", "12345678901234567890123456789013")
	viper.SetDefault(`databaseHost`, db_host)
	viper.SetDefault(`databasePort`, db_port)
	viper.SetDefault(`databaseUser`, db_user)
	viper.SetDefault(`databasePass`, db_pass)
	viper.SetDefault(`databaseName`, "smartlab_db")

	viper.SetConfigFile("config/config.env")
	viper.ReadInConfig()
	log.Println("finish viper")
	//database.AutoMigrate()
	port, _ := strconv.Atoi(apps_port)
	api.NewServer().Start(port)
}
