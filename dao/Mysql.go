package dao

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"metro/settings"
	"time"
)

var db *gorm.DB

//func InitMysql(config *settings.MysqlConfig) (error, *gorm.DB) {
//	Db, err := gorm.Open(mysql.Open("root:123456@tcp(localhost:3306)/metro?parseTime=true&loc=Asia%2FShanghai"), &gorm.Config{})
//	if err != nil {
//		log.Fatal(err)
//		return err, nil
//	}
//	sqlDb, _ := Db.DB()
//	// SetMaxOpenConns 设置打开数据库连接的最大数量。
//	sqlDb.SetMaxOpenConns(5)
//	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
//	sqlDb.SetMaxIdleConns(2)
//	sqlDb.SetConnMaxIdleTime(time.Minute * 300)
//	defer sqlDb.Close()
//	return nil, Db
//}

func NewConnection() *gorm.DB {
	var dbUri string
	dbUri = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true",
		settings.Conf.MysqlConfig.User,
		settings.Conf.MysqlConfig.Password,
		settings.Conf.MysqlConfig.Host,
		settings.Conf.MysqlConfig.Port,
		settings.Conf.MysqlConfig.Database)

	conn, err := gorm.Open(mysql.Open(dbUri), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}
	return conn
}

func Setup() error{
	db = NewConnection()
	sqlDb, err := db.DB()
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDb.SetMaxOpenConns(30)
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDb.SetMaxIdleConns(10)
	sqlDb.SetConnMaxLifetime(time.Second * 300) //设置连接空闲超时
	//defer sqlDb.Close()
	return err
}

func GetDB() *gorm.DB {
	sqlDb, err := db.DB()
	if err != nil {
		sqlDb.Close()
		db = NewConnection()
	}
	return db
}