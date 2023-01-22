package database

import (
	"fmt"
	orgDomain "hardtmann/smartlab/api/organization/domain"
	userDomain "hardtmann/smartlab/api/user/domain"
	"log"
	"net/url"

	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

// func GetDB() *gorm.DB {
// 	if db == nil {
// 		database, err := gorm.Open(sqlite.Open("file:"+viper.GetString("DB_LOCATION")+"?cache=shared&mode=rwc"), &gorm.Config{
// 			Logger: logger.New(
// 				log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
// 				logger.Config{
// 					SlowThreshold:             time.Second, // Slow SQL threshold
// 					LogLevel:                  logger.Info, // Log level
// 					IgnoreRecordNotFoundError: false,       // Ignore ErrRecordNotFound error for logger
// 					Colorful:                  true,        // Disable color
// 				},
// 			),
// 		})
// 		if err != nil {
// 			panic(err)
// 		}
// 		err = database.Exec("PRAGMA foreign_keys = ON").Error
// 		if err != nil {
// 			panic(err)
// 		}

// 		db = database
// 	}
// 	return db
// }

func GetDB() *gorm.DB {
	dbHost := viper.GetString(`databaseHost`)
	dbPort := viper.GetString(`databasePort`)
	dbUser := viper.GetString(`databaseUser`)
	dbPass := viper.GetString(`databasePass`)
	dbName := viper.GetString(`databaseName`)

	dbDetailInformation := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	sqlxDb, err := InitializeMySqlDB(dbDetailInformation)
	if err != nil {
		log.Println("failed to initialize DB")
		log.Fatalln(err)
		return nil
	}

	if sqlxDb == nil {
		log.Println("failed,sqlxDb nil")
		log.Fatalln(err)
		return nil
	}

	return GetMySqlDB(sqlxDb)
}

func InitializeMySqlDB(dbDetailInfo string) (dbConnection *sqlx.DB, err error) {
	val := url.Values{}
	val.Add("parseTime", "1")
	dsn := fmt.Sprintf("%s?%s", dbDetailInfo, val.Encode())
	log.Println(dsn)
	dbConnection, err = sqlx.Connect(`mysql`, dsn)
	if err != nil {
		log.Println("failed to initialize my sql")
		log.Fatalln(err)
		return nil, err
	}

	if dbConnection == nil {
		log.Println("failed, dbConnection nil")
		log.Fatalln(err)
		return nil, err
	}
	return
}

func GetMySqlDB(db *sqlx.DB) *gorm.DB {
	gorm, err := gorm.Open(mysql.New(mysql.Config{
		Conn: db.DB,
	}), &gorm.Config{})
	if err != nil {
		log.Println("failed to open gorm connection")
		log.Fatalln(err)
		return nil
	}

	if gorm == nil {
		log.Println("failed, gorm nil")
		log.Fatalln(err)
		return nil
	}
	return gorm
}

func AutoMigrate() {
	GetDB().AutoMigrate(
		&orgDomain.Organization{},
		&userDomain.User{},
	)
	GetDB().Save(&orgDomain.Organization{
		Name: "Hardtmann",
	})
}
