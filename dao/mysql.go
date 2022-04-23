package dao

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"metro/settings"
)

var db *sqlx.DB

func Init(cfg *settings.MysqlConfig) (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=Local", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Database)
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		return
	}
	db.SetMaxIdleConns(200)
	db.SetMaxIdleConns(50)
	return
}

func Close() {
	db.Close()
}

//func NewConnection() *gorm.DB {
//	var dbUri string
//	dbUri = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true",
//		settings.Conf.MysqlConfig.User,
//		settings.Conf.MysqlConfig.Password,
//		settings.Conf.MysqlConfig.Host,
//		settings.Conf.MysqlConfig.Port,
//		settings.Conf.MysqlConfig.Database)
//
//	conn, err := gorm.Open(mysql.Open(dbUri), &gorm.Config{})
//
//	if err != nil {
//		log.Fatal(err)
//	}
//	return conn
//}
//
//func Setup() error{
//	db = NewConnection()
//	sqlDb, err := db.DB()
//	// SetMaxOpenConns 设置打开数据库连接的最大数量。
//	sqlDb.SetMaxOpenConns(30)
//	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
//	sqlDb.SetMaxIdleConns(10)
//	sqlDb.SetConnMaxLifetime(time.Second * 300) //设置连接空闲超时
//	//defer sqlDb.Close()
//	return err
//}
//
//func GetDB() *gorm.DB {
//	sqlDb, err := db.DB()
//	if err != nil {
//		sqlDb.Close()
//		db = NewConnection()
//	}
//	return db
//}
