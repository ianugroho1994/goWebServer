package app

import (
	"fmt"
	"goWebServer/shared/logger"
	"net/url"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	fmt.Println("start server initialization")

	echoInstance := InitializeEcho()
	dbConnection, error := InitializeDatabase(GetDbInformation())

	if error != nil {
		logger.Log.Fatal(error)
	}

	gormConnection, error := InitializeGormConnection(dbConnection)

	if error != nil {
		logger.Log.Fatal(error)
	}

	timeout := time.Duration(viper.GetInt("context.timeout")) * time.Second

	server := InitializeWebServerDependency(gormConnection, echoInstance, timeout)
	server.EchoInstance.Logger.Fatal(server.EchoInstance.Start(viper.GetString("server.address")))
}

func init() {
	viper.SetConfigFile(`./config/config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	logger.Log = logger.NewLogger("")
	logger.Log.SetLogLevel(logrus.DebugLevel)
}

func InitializeEcho() *echo.Echo {
	echoInstance := echo.New()

	skipUptimeCheckFunc := func(c echo.Context) bool {
		return strings.Contains(c.Request().UserAgent(), "GoogleStackdriverMonitoring-UptimeChecks")
	}

	echoInstance.Use(middleware.CORS())
	echoInstance.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Skipper: skipUptimeCheckFunc,
	}))
	echoInstance.Use(middleware.BodyDumpWithConfig(middleware.BodyDumpConfig{
		Skipper: skipUptimeCheckFunc,
		Handler: func(c echo.Context, reqBody, resBody []byte) {
			logger.Log.Info(string(reqBody))
			logger.Log.Info(string(resBody))
		},
	}))
	return echoInstance
}

func GetDbInformation() string {
	dbHost := viper.GetString(`database.host`)
	dbPort := viper.GetString(`database.port`)
	dbUser := viper.GetString(`database.user`)
	dbPass := viper.GetString(`database.pass`)
	dbName := viper.GetString(`database.name`)

	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
}

func InitializeDatabase(dbDetailInfo string) (dbConn *sqlx.DB, err error) {
	val := url.Values{}
	val.Add("parseTime", "1")
	dsn := fmt.Sprintf("%s?%s", dbDetailInfo, val.Encode())
	dbConn, err = sqlx.Connect(`mysql`, dsn)
	if err != nil {
		logger.Log.Fatal(err)
		return nil, err
	}
	return
}

func InitializeGormConnection(db *sqlx.DB) (dbConn *gorm.DB, err error) {
	dbConn, err = gorm.Open(mysql.New(mysql.Config{
		Conn: db.DB,
	}), &gorm.Config{})
	if err != nil {
		logger.Log.Fatal(err)
		return nil, err
	}
	return
}
