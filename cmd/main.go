package main

import (
	"cloud-disk/svc/user"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"os"
)

func init() {
	err := godotenv.Load("/home/chen/goProject/src/cloud-disk/cmd/.env")
	if err != nil {
		panic(err.Error())
	}

	if os.Getenv("MYSQL_USER") == "" {
		panic("error: env MYSQL_USER no set")
	}

	if os.Getenv("MYSQL_PASSWORD") == "" {
		panic("error: env MYSQL_PASSWORD no set")
	}

	if os.Getenv("MYSQL_HOST") == "" {
		panic("error: env MYSQL_HOST no set")
	}

	if os.Getenv("MYSQL_DATABASE") == "" {
		panic("error: env MYSQL_DATABASE no set")
	}
}

func main() {
	var db *sql.DB
	{
		var err error
		var link = os.Getenv("MYSQL_USER") + ":" + os.Getenv("MYSQL_PASSWORD") + "@tcp(" + os.Getenv("MYSQL_HOST") + ")/" + os.Getenv("MYSQL_DATABASE")
		db, err = sql.Open("mysql", link)

		if err != nil {
			panic(err.Error())
		}
	}

	var err chan error
	var r = gin.Default()
	{
		user.MakeSvc(db, r)
	}

	go func() {
		err <- r.Run(":8080")
	}()

	fmt.Println("err:", <-err)
}
