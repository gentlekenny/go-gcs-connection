package main

import (
	"fmt"
	"log"

	proxy "github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	ConnectToGCSMySQLDatabase()
}

func ConnectToGCSMySQLDatabase() *gorm.DB {

	user := "username"   // Username
	passwd := "password" // Password
	instance := "projectId:region:instance"
	// ProjectID : This is the project ID or name. It identifies the Google Cloud project that contains the MySQL instance.
	// Region: This is the region where the MySQL instance is located. In this case, it is the us-east1 region.
	// Instace: This is the name of the MySQL instance within the specified project and region. It uniquely identifies the MySQL instance.
	// Example : lazy-project:us-east-1:mysql-live
	database := "database" // Dataabase name

	cfg := proxy.Cfg(instance, user, passwd)
	cfg.DBName = database
	db, err := proxy.DialCfg(cfg)
	if err != nil {
		fmt.Println("failed to connect database")
		panic("failed to connect database")
	}

	// In case you want to convert that connection to GORM
	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn: db,
	}), &gorm.Config{})

	if err != nil {
		fmt.Println("failed to connect as a gorm database")
		panic("failed to connect as gorm database")
	}

	log.Println("Successfully connected to MySQL database.")

	return gormDB
}
