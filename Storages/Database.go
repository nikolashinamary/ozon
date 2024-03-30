package Storages

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

func DBConnection(log1 *logrus.Logger) *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		log1.Fatalln(err)
		return nil
	}
	host := os.Getenv("POSTGRES_HOST")
	dbname := os.Getenv("POSTGRES_DB")
	username := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	port := os.Getenv("POSTGRES_PORT")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host, username, password, dbname, port)
	_, err = gorm.Open(postgres.Open(dsn))
	if err != nil {
		log1.Warningln("Unable to access database")
		log1.Infoln("Trying to create database")
		dsn1 := fmt.Sprintf("host=%s user=%s password=%s dbname=postgres port=%s sslmode=disable",
			host, username, password, port)
		Db, err := gorm.Open(postgres.Open(dsn1))
		if err != nil {
			log1.Fatalln("Unable to access new database")
		}
		query := fmt.Sprintf("CREATE DATABASE %s", dbname)
		err = Db.Exec(query).Error
		if err != nil {
			log1.Fatalln("Unable to create database")
		}
		dbSQL, err := Db.DB()
		if err != nil {
			log1.Infoln("Unable to close database")
			return nil
		}
		err = dbSQL.Close()
		if err != nil {
			log1.Errorln("Unable to close database")
			return nil
		}
		log1.Infoln("Database successfully created!")
	}
	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		log1.Fatalln("Unable to connect to database")
	}
	return db
}

func DBDisconnection(db *gorm.DB, log1 *logrus.Logger) {
	temp, err := db.DB()
	if err != nil {
		logrus.Infoln("Disconnect error")
		return
	}
	err = temp.Close()
	if err != nil {
		logrus.Infoln("Disconnect error")
		return
	}
}

type Database struct {
	Database *gorm.DB
}

func DatabaseConstr(db *gorm.DB) *Database {
	return &(Database{Database: db})
}

func (database *Database) GetByHash(hash string, log1 *logrus.Logger) string {
	var mod Model
	err := database.Database.Where("short_url=?", hash).First(&mod).Error
	if err != nil {
		log1.Errorln("Unable to get")
		return ""
	}
	return mod.Url

}

func (database *Database) GetByUrl(url string, log1 *logrus.Logger) string {
	var mod Model
	err := database.Database.Where("url=?", url).First(&mod).Error
	if err != nil {
		log1.Errorln("Unable to get")
		return ""
	}
	return mod.ShortUrl

}

func (database *Database) WriteByUrl(hash string, url string, log1 *logrus.Logger) error {
	mod := Model{ShortUrl: hash, Url: url}
	return database.Database.Create(&mod).Error

}

func (database *Database) DeleteByUrl(url string, log1 *logrus.Logger) {
	var obj Model
	database.Database.Where("url=?", url).Delete(&obj)
}

func (database *Database) ContainsByUrl(url string, log1 *logrus.Logger) bool {
	err := database.Database.Where("url=?", url).First(&Model{}).Error
	if err != nil {
		return false
	}
	return true
}

func (database *Database) ContainsByHash(hash string, log1 *logrus.Logger) bool {
	err := database.Database.Where("short_url=?", hash).First(&Model{}).Error
	if err != nil {
		return false
	}
	return true
}
