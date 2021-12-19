package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"
)

func GetConn() *sql.DB {
	if appName := os.Getenv("APP_NAME"); appName == ""{
		log.Fatal("APP_NAME ENVIRONMENT CANNOT BE NULL")
	}
	user := os.Getenv("DB_USER")
	if user == "" {
		log.Fatal("DB_USER ENVIRONMENT CANNOT BE NULL")
	}
	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		log.Fatal("DB_NAME ENVIRONMENT CANNOT BE NULL")
	}
	dsn := fmt.Sprintf("%s:@/%s?parseTime=true",user,dbName)
	conn, err := sql.Open("mysql",dsn)
	if err != nil {
		fmt.Println("Connection Failed!!!")
		panic(err)
	}
	conn.SetMaxIdleConns(10) //min 10 connection
	conn.SetMaxOpenConns(100) //max 100 connection
	conn.SetConnMaxIdleTime(5 * time.Minute) //if iin 5 minutes nothing happen? db will close the connection
	conn.SetConnMaxLifetime(60 * time.Minute) //after 60 minutes, the connection will create new connection
	return conn
}
