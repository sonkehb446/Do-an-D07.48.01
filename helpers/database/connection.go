package database

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"Hybrid-blog/helpers/config"
)

var (
	once sync.Once

	dbMaster *gorm.DB
)

type dbName string

const connectDBNameMaster = dbName("master")

type dbInfo struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
}

func DB() *gorm.DB {
	once.Do(func() {
		dbMaster = newConnection(connectDBNameMaster)
	})
	return dbMaster
}

func newConnection(dn dbName) *gorm.DB {
	d, err := dbConnectionInfo(dn)
	if err != nil {
		log.Fatalf("err: %s", err.Error())
		os.Exit(1)
	}

	var connectionString string

	switch config.GetEnv() {
	case config.EnvLocal:
		connectionString = fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s connect_timeout=9",
			d.Host,
			d.User,
			d.Name,
			d.Password)
	case config.EnvTesting, config.EnvProduction:
		connectionString = fmt.Sprintf("host=%s user=%s dbname=%s sslmode=require password=%s connect_timeout=3",
			d.Host,
			d.User,
			d.Name,
			d.Password)
	default:
		fmt.Println("undefined environment")
		os.Exit(1)
	}

	db, err := gorm.Open("postgres", connectionString)
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
	env := config.GetEnvValue()
	db.DB()
	db.DB().SetConnMaxLifetime(time.Hour * 24)
	db.DB().SetMaxIdleConns(env.Database.MaxIdleConnections)
	db.DB().SetMaxOpenConns(env.Database.MaxOpenConnections)

	if config.GetEnv() != config.EnvProduction && config.GetEnv() != config.EnvTesting {
		db.LogMode(true)
	}
	return db
}

func dbConnectionInfo(dn dbName) (dbInfo, error) {
	env := config.GetEnvValue()
	switch dn {
	case connectDBNameMaster:
		return dbInfo{
			Host:     env.Database.Host,
			Port:     env.Database.Port,
			User:     env.Database.User,
			Name:     env.Database.Name,
			Password: env.Database.Password,
		}, nil
	default:
		return dbInfo{}, fmt.Errorf("[dbConnectionInfo]local invalid dbName: %s", dn)
	}
}
