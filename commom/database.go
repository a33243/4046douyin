package commom

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/url"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	host := viper.GetString("mysql.host")
	port := viper.GetInt("mysql.port")
	database := viper.GetString("mysql.dbname")
	username := viper.GetString("mysql.username")
	password := viper.GetString("mysql.password")
	charset := viper.GetString("mysql.charset")
	loc := viper.GetString("datasource.loc")
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true&loc=%s",
		username,
		password,
		host,
		port,
		database,
		charset,
		url.QueryEscape(loc), // 对参数进行转义，防止解析错误
	) // root:789528@tcp(localhost:3306)/douyin?charset=utf-8&parse=true&loc=shanghai
	db, err := gorm.Open(mysql.Open(args), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	DB = db
	return db
}
