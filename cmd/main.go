package main

import (
	"cloud-disk/logger"
	"cloud-disk/svc/user"
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"os"
)

var (
	LogPath       string
	MySQLUser     string
	MySQLPassword string
	MySQLHost     string
	MySQLDatabase string
	RuntimePath   string
	LogPathName   string
	LogExpire     string
	ListenPort    string
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err.Error())
	}

	if MySQLUser = os.Getenv("MYSQL_USER"); MySQLUser == "" {
		panic("error: env MYSQL_USER no set")
	}

	if MySQLPassword = os.Getenv("MYSQL_PASSWORD"); MySQLPassword == "" {
		panic("error: env MYSQL_PASSWORD no set")
	}

	if MySQLHost = os.Getenv("MYSQL_HOST"); MySQLHost == "" {
		panic("error: env MYSQL_HOST no set")
	}

	if MySQLDatabase = os.Getenv("MYSQL_DATABASE"); MySQLDatabase == "" {
		panic("error: env MYSQL_DATABASE no set")
	}

	if RuntimePath = os.Getenv("RUNTIME_PATH"); RuntimePath == "" {
		panic("error: env RUNTIME_PATH no set")
	}

	if LogPathName = os.Getenv("LOG_PATH_NAME"); LogPathName == "" {
		panic("error: env LOG_PATH_NAME no set")
	}

	if LogExpire = os.Getenv("LOG_EXPIRE"); LogExpire == "" {
		panic("error: env LOG_EXPIRE no set")
	}

	if ListenPort = os.Getenv("LISTEN_PORT"); ListenPort == "" {
		panic("error: env LOG_EXPIRE no set")
	}

	LogPath = RuntimePath + "/" + LogPathName
}

func main() {
	var db *sql.DB
	{
		var err error
		db, err = sql.Open(
			"mysql",
			MySQLUser+":"+MySQLPassword+"@tcp("+MySQLHost+")/"+MySQLDatabase,
		)

		if err != nil {
			panic(err.Error())
		}
	}

	var loggerDebug = true
	var ginMode = os.Getenv("GIN_RUN_MODE")
	gin.SetMode(ginMode)

	//gin的运行模式为release的时候
	//才记录日志在文件
	//否则打印在控制台
	if gin.Mode() == gin.ReleaseMode {
		gin.DisableConsoleColor()

		//gin mode == debug or test
		loggerDebug = false
	}

	var l = logger.New(RuntimePath, LogPath, "gin-", LogExpire, loggerDebug)
	var r = gin.New()
	r.Use(l.GinLogger(), gin.Recovery())
	{
		user.MakeSvc(db, r, l)
	}

	var err chan error

	go func() {
		err <- r.Run(ListenPort)
	}()

	l.Write("err:", (<-err).Error())
}
