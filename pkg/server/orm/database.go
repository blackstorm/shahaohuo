package orm

import (
	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"shahaohuo.com/shahaohuo/pkg/config"
	"time"
)

var database *gorm.DB

func Init(mConfig config.MySQLConfig) {

	config := mysql.Config{
		User:                 mConfig.UserName,
		Passwd:               mConfig.Password,
		Addr:                 mConfig.Addr(),
		DBName:               mConfig.Database,
		Net:                  "tcp",
		ParseTime:            true,
		AllowNativePasswords: true,
		Loc:                  time.UTC,
	}

	if db, err := gorm.Open("mysql", config.FormatDSN()); err == nil {
		db.LogMode(mConfig.ShowSQL)
		db.SingularTable(true)
		db.DB().SetMaxOpenConns(10)
		db.DB().SetMaxIdleConns(1)
		if err = db.DB().Ping(); err == nil {
			if err = db.AutoMigrate(&User{}, &Haohuo{}, &Favorite{}, &Image{}, &Comment{}, &Tag{}, &HaohuoTag{}).Error; err == nil {
				database = db
			}
		} else {
			panic(err)
		}
	} else {
		panic(err)
	}

}
