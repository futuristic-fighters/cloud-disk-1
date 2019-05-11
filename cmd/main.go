package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"os"
)

func init() {
	err := godotenv.Load("/home/chen/goProject/src/cloud-storage/cmd/.env")
	if err != nil {
		fmt.Println("load .env error", err.Error())
		os.Exit(-1)
	}
}

func main() {
	var err chan error
	r := gin.Default()

	go func() {
		err <- r.Run(":8080")
	}()

	fmt.Println("err:", <-err)
}
